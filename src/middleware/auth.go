package middleware

import (
	"fmt"
	"log"
	"time"

	"sport_bookie_server/src/db"
	"sport_bookie_server/src/jwt"
	"sport_bookie_server/src/model"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/crypto/bcrypt"
)

const identityKey = "ID"
const secretKey = "somesecretkey"

// AuthUser struct
type AuthUser struct {
	ID string `json:"_id" bson:"_id,omitempty"`
}

func getAuthMiddleware() (*jwt.GinJWTMiddleware, error) {
	authMiddleware, err := jwt.New(&jwt.GinJWTMiddleware{
		Realm:       "user_login",
		Key:         []byte(secretKey),
		Timeout:     3600 * time.Hour,
		MaxRefresh:  3600 * time.Hour,
		IdentityKey: identityKey,
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(*AuthUser); ok {
				return jwt.MapClaims{
					identityKey: v.ID,
				}
			}
			return jwt.MapClaims{}
		},
		IdentityHandler: func(c *gin.Context) interface{} {
			claims := jwt.ExtractClaims(c)
			return &AuthUser{
				ID: claims[identityKey].(string),
			}
		},
		Authenticator: func(c *gin.Context) (interface{}, error) {
			var login model.Login
			var err error
			if err = c.ShouldBind(&login); err != nil {
				return nil, jwt.ErrMissingLoginValues
			}
			username := login.Username
			password := login.Password
			filter := bson.M{"username": username}
			var user model.User
			err = db.Users.FindOne(c, filter).Decode(&user)
			if err != nil {
				log.Println(err)
				return nil, jwt.ErrFailedAuthentication
			}
			err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
			if err != nil {
				log.Println(err)
				return nil, jwt.ErrFailedAuthentication
			}
			return &AuthUser{
				ID: user.ID.Hex(),
			}, nil
		},
		Unauthorized: func(c *gin.Context, code int, message string) {
			c.JSON(code, gin.H{
				"code":    code,
				"message": message,
			})
		},
		TokenLookup:   "header: Authorization, query: token", // cookie: jwt",
		TokenHeadName: "Bearer",
		TimeFunc:      time.Now,
	})
	if err != nil {
		log.Fatal("JWT Error: " + err.Error())
		return nil, fmt.Errorf("auth: Fail to setup jwt auth middleware, err: (%v)", err)
	}
	return authMiddleware, nil
}

// AuthMiddleware ...
var AuthMiddleware *jwt.GinJWTMiddleware

func init() {
	jwt, err := getAuthMiddleware()
	if err != nil {
		log.Fatal("JWT Error:" + err.Error())
	}
	AuthMiddleware = jwt
}
