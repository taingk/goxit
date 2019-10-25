package Models

import (
	"fmt"

	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/taingk/goxit/api/Config"
)

func GetAllUser(b *[]User) (err error) {
	if err = Config.DB.Find(b).Error; err != nil {
		return err
	}
	return nil
}

func AddNewUser(b *User) (err error) {
	if err = Config.DB.Create(b).Error; err != nil {
		return err
	}
	return nil
}

func GetOneUser(b *User, id string) (err error) {
	if err := Config.DB.Where("UUID = ?", id).First(b).Error; err != nil {
		return err
	}
	return nil
}

func PutOneUser(b *User, id string) (err error) {
	fmt.Println(b)
	Config.DB.Save(b)
	return nil
}

func DeleteUser(b *User, id string) (err error) {
	Config.DB.Where("UUID = ?", id).Delete(b)
	return nil
}
