package handler

import (
	"log"
	"net/http"
	"sport_bookie_server/src/middleware"
	"time"
	"sport_bookie_server/src/db"
	"sport_bookie_server/src/model"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

func generatorToken(userID string) (string, time.Time, error) {
	expire := middleware.AuthMiddleware.TimeFunc().UTC().Add(middleware.AuthMiddleware.Timeout)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"ID":       userID,
		"exp":      expire.Unix(),
		"orig_iat": middleware.AuthMiddleware.TimeFunc().Unix(),
	})
	tokenString, err := token.SignedString([]byte(middleware.AuthMiddleware.Key))
	if err != nil {
		return "", time.Time{}, err
	}
	return tokenString, expire, nil
}

// RegisterHandler handle register
func RegisterHandler(c *gin.Context) {

	var register model.Login
	err := c.ShouldBind(&register)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusOK, gin.H{"code": 3, "token": "", "expire": ""})
		return
	}

	username := register.Username
	password := register.Password

	hash, err := bcrypt.GenerateFromPassword([]byte(password), 5)
	if err != nil {
		log.Println(err)
	}
	hashPassword := string(hash)
	newUser := model.User{
		Username:     username,
		Password:     hashPassword,
		InitialCredit: 10000,
		CreatedAt:    time.Now(),
		LastOnlineAt: time.Now(),
	}
	res, err := db.Users.InsertOne(c, newUser)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusOK, gin.H{"code": 2, "token": "", "expire": ""})
		return
	}
	id := res.InsertedID.(primitive.ObjectID)
	token, expire, _ := generatorToken(id.Hex())
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusOK, gin.H{"code": 3, "token": "", "expire": ""})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": 1, "token": token, "expire": expire})
	return
}

// code
// 	0: '',
// 	1: '', // success
// 	2: 'Username taken!',
// 	3: 'Invalid Input!',
// 	4: 'Server Error!',
// 	5: 'Username invalid!',
// 	6: 'Username at least 4 characters!',
// 	7: 'Password at least 4 characters!'
