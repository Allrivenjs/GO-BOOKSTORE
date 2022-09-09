package utils

import (
	"errors"
	"github.com/gin-gonic/gin"
)

func MatchUserTypeToId(c *gin.Context, userId string) (err error) {
	userType := c.GetString("user_type")
	id := c.GetString("id")
	err = nil
	if userType == "USER" && id != userId {
		err = errors.New("unauthorized to access this resource")
		return err
	}
	err = CheckUserType(c, userType)
	return err
}

func CheckUserType(c *gin.Context, role string) (err error) {
	userType := c.GetString("user_type")
	err = nil
	if userType != role {
		err = errors.New("unauthorized to access this resource")
		return err
	}
	return err
}
