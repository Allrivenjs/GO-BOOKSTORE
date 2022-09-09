package controllers

import (
	"encoding/json"
	"github.com/Allrivenjs/GO-BOOKSTORE/pkg/models"
	"github.com/Allrivenjs/GO-BOOKSTORE/pkg/utils"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"log"
	"net/http"
	"strconv"
)

var User = models.User{}
var validate = validator.New()

func HashPassword() {

}

func VerifyPassword() {

}

func Signup() gin.HandlerFunc {
	return func(c *gin.Context) {
		var user models.User
		if err := c.BindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		validationErr := validate.Struct(user)
		if validationErr != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": validationErr.Error()})
			return
		}
		_, _, count := models.FindEmailCount(&user)
		if count > 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "el email ya existe"})
			return
		}
		_, _, countPhone := models.FindPhoneCount(&user)
		if countPhone > 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "el telefono ya existe"})
			return
		}
		token, refreshToken, _ := utils.GenerateAllTokens(*user.Email, *user.FirstName, *user.LastName, *user.UserType, *&user.ID)
		user.Token = &token
		user.RefreshToken = &refreshToken
		b := user.CreateUser()
		res, _ := json.Marshal(b)
		c.JSON(http.StatusOK, res)

	}
}

func Login() {

}

func GetUsers() gin.HandlerFunc {
	return func(context *gin.Context) {

	}
}

func GetUser() gin.HandlerFunc {
	return func(context *gin.Context) {
		userId := context.Param("user_id")
		if err := utils.MatchUserTypeToId(context, userId); err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
		i, err := strconv.ParseInt(userId, 0, 0)
		if err != nil {
			log.Fatal(err.Error())
		}
		UserDetails, _ := models.GetUserById(i)
		res, _ := json.Marshal(UserDetails)
		context.JSON(http.StatusOK, res)
	}
}
