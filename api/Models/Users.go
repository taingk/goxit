package Models

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/taingk/goxit/api/Config"
	"github.com/taingk/goxit/api/Helpers"
	"gopkg.in/go-playground/validator.v9"
)

func GetAllUser(b *[]User) (err error) {
	if err = Config.DB.Where("is_deleted = 0").Find(b).Error; err != nil {
		return err
	}
	return nil
}

func AddNewUser(b *User) (err error) {
	var validate *validator.Validate
	validate = validator.New()
	if err = validate.Struct(b); err != nil {
		return err
	}

	var encryptedPassword = Helpers.EncryptPassword(b.Password)
	b.Password = encryptedPassword
	b.IsDeleted = 0
	if err = Config.DB.Create(b).Error; err != nil {
		return err
	}
	return nil
}

func GetOneUser(b *User, uuid string) (err error) {
	if err := Config.DB.Where("uuid = ? AND is_deleted = 0", uuid).First(b).Error; err != nil {
		return err
	}
	return nil
}

func PutOneUser(b *User, uuid string) (err error) {
	var validate *validator.Validate
	validate = validator.New()
	if err = validate.Struct(b); err != nil {
		return err
	}

	fmt.Println(b)
	Config.DB.Save(b)
	return nil
}

func DeleteUser(b *User, uuid string) (err error) {
	b.IsDeleted = 1
	Config.DB.Where("uuid = ?", uuid).Save(b)
	return nil
}

func GetUserByEmailPassword(b *User) (err error) {
	if err := Config.DB.Where("email = ? AND password = ? AND is_deleted = 0", b.Email, Helpers.EncryptPassword(b.Password)).First(b).Error; err != nil {
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
