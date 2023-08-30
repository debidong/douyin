package auth

import (
	"context"
	"douyin/models"
	"douyin/utils"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.PostForm("token")
		//username := c.PostForm("username")

		if token == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		} else {
			ok, _, err := handleToken(token)
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

// handleToken handle token for authMiddleware
// return the validity of the token, username, and error
func handleToken(token string) (bool, string, error) {
	ctx := context.Background()
	parsedToken, err := jwt.Parse(
		token,
		func(token *jwt.Token) (interface{}, error) {
			return []byte(utils.JwtSecretKey), nil
		})
	fmt.Println(parsedToken)
	var username string
	if claims, ok := parsedToken.Claims.(jwt.MapClaims); ok && parsedToken.Valid {
		if ok {
			username = claims["sub"].(string)
		}
	}
	tokenStored, err := utils.Client.Get(ctx, username).Result()
	if err != nil {
		return false, "", err
	}
	if tokenStored == token {
		return true, "", nil
	} else {
		return false, "", nil
	}

}

// SetToken sets token during login or registration.
func SetToken(username string) (string, error) {
	ctx := context.Background()
	newToken := jwt.New(jwt.SigningMethodHS256)
	claims := newToken.Claims.(jwt.MapClaims)

	//claims["exp"] = time.Now().Add(time.Hour * 1).Unix()
	claims["sub"] = username
	secret := []byte(utils.JwtSecretKey)
	tokenString, err := newToken.SignedString(secret)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	pong, err := utils.Client.Ping(ctx).Result()
	if err != nil {
		fmt.Println("Error pinging Redis:", err)
		return "", err
	}
	fmt.Println("Connected to Redis:", pong)

	err = utils.Client.Set(ctx, username, tokenString, 0).Err()
	if err != nil {
		fmt.Println("Error setting key:", err)
		return "", err
	}

	return tokenString, nil
}

func GetUserFromToken(token string) (models.User, error) {
	ok, username, err := handleToken(token)
	if err != nil {
		return models.User{}, err
	} else if !ok {
		return models.User{}, nil
	} else {
		user := models.User{Username: username}

		// error retrieving user, user doesn't exist
		if err := utils.DB.First(&user).Error; err != nil {
			return models.User{}, err
		} else {
			return user, nil
		}
	}
}
