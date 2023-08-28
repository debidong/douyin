package auth

import (
	"context"
	"douyin/utils"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.PostForm("token")
		username := c.PostForm("username")

		if token == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		} else {
			ok, _, err := handleToken(token, username)
			if err != nil {
				c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
				c.Abort()
				fmt.Println(err)
				return
			}
			if ok {
				c.Next()
			} else {
				c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
				c.Abort()
			}

		}
	}
}

// handleToken sets or parses token for a specific user
func handleToken(token string, username string) (bool, string, error) {
	ctx := context.Background()
	// if token is empty string, create a new token
	if token == "" {
		newToken := jwt.New(jwt.SigningMethodHS256)
		claims := newToken.Claims.(jwt.MapClaims)

		//claims["exp"] = time.Now().Add(time.Hour * 1).Unix()
		claims["sub"] = username
		secret := []byte(utils.JwtSecretKey)
		tokenString, err := newToken.SignedString(secret)
		if err != nil {
			fmt.Println(err)
			return false, "", err
		}
		pong, err := utils.Client.Ping(ctx).Result()
		if err != nil {
			fmt.Println("Error pinging Redis:", err)
			return false, "", err
		}
		fmt.Println("Connected to Redis:", pong)

		err = utils.Client.Set(ctx, username, tokenString, 0).Err()
		if err != nil {
			fmt.Println("Error setting key:", err)
			return false, "", err
		}
		return true, tokenString, nil
		// else, parse token
	} else {
		tokenStored, err := utils.Client.Get(ctx, username).Result()
		if err != nil {
			return false, "", fmt.Errorf("error getting key: %w", err)
		} else {
			parsedToken, err := jwt.Parse(
				tokenStored,
				func(token *jwt.Token) (interface{}, error) {
					return utils.JwtSecretKey, nil
				})

			if err != nil {
				return false, "", err
			}

			if claims, ok := parsedToken.Claims.(jwt.MapClaims); ok && parsedToken.Valid {
				if token == tokenStored {
					fmt.Println("User logged in: ", claims["sub"])
					return true, token, nil
				} else {
					// token in request != token in db
					return false, "", nil
				}
			} else {
				return false, "", nil
			}
		}
	}
}

func SetToken(username string) (string, error) {
	ok, token, err := handleToken("", username)
	if err != nil {
		fmt.Println(err)
		return "", err
	} else if !ok {
		return "", nil
	} else {
		return token, nil
	}

}
