package Models

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/taingk/goxit/api/Config"
	"github.com/taingk/goxit/api/Helpers"
)

func GetAllUser(b *[]User) (err error) {
	if err = Config.DB.Find(b).Error; err != nil {
		return err
	}
	return nil
}

func AddNewUser(b *User) (err error) {
	var encryptedPassword = Helpers.EncryptPassword(b.Password)
	b.Password = encryptedPassword
	if err = Config.DB.Create(b).Error; err != nil {
		return err
	}
	return nil
}

func GetOneUser(b *User, uuid string) (err error) {
	if err := Config.DB.Where("uuid = ?", uuid).First(b).Error; err != nil {
		return err
	}
	return nil
}

func PutOneUser(b *User, uuid string) (err error) {
	fmt.Println(b)
	Config.DB.Save(b)
	return nil
}

func DeleteUser(b *User, uuid string) (err error) {
	Config.DB.Where("uuid = ?", uuid).Delete(b)
	return nil
}

func GetUserByEmailPassword(b *User) (err error) {
	if err := Config.DB.Where("email = ? AND password = ?", b.Email, Helpers.EncryptPassword(b.Password)).First(b).Error; err != nil {
		return err
	}
	return nil
}

func GetLoggedUser(b *User) (err error) {
	if err := Config.DB.Where("uuid = ? AND access_level = ?", b.UUID, b.AccessLevel).First(b).Error; err != nil {
		return err
	}
	return nil
}

func Authorize(c *gin.Context) bool {
	accessLevel := c.Request.Header.Get("AccessLevel")
	uuid := c.Request.Header.Get("Authorization")
	var user User
	user.UUID = uuid
	parsedAccessLevel, _ := strconv.ParseInt(accessLevel, 10, 64)
	user.AccessLevel = int(parsedAccessLevel)
	err := GetLoggedUser(&user)

	if parsedAccessLevel == 1 && err == nil {
		return true
	} else {
		return false
	}
}

func AuthorizeSimpleUser(c *gin.Context) bool {
	accessLevel := c.Request.Header.Get("AccessLevel")
	uuid := c.Request.Header.Get("Authorization")
	var user User
	user.UUID = uuid
	parsedAccessLevel, _ := strconv.ParseInt(accessLevel, 10, 64)
	user.AccessLevel = int(parsedAccessLevel)
	err := GetLoggedUser(&user)

	if parsedAccessLevel == 0 && err == nil {
		return true
	} else {
		return false
	}
}
