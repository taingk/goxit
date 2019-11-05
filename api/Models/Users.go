package Models

import (
	"fmt"

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
	var encryptedPassword = Helpers.encryptPassword(b.Password)
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
