package Models

import (
	"fmt"

	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/taingk/goxit/api/Config"
	"golang.org/x/crypto/bcrypt"
)

func GetAllUser(b *[]User) (err error) {
	if err = Config.DB.Find(b).Error; err != nil {
		return err
	}
	return nil
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func AddNewUser(b *User) (err error) {

	password := b.Password
	hash, _ := HashPassword(password) // ignore error for the sake of simplicity

	fmt.Println("Password:", password)
	fmt.Println("Hash:", hash)

	if err = Config.DB.Create(b).Error; err != nil {
		return err
	}

	return nil
}

func GetOneUser(b *User, id string) (err error) {
	if err := Config.DB.Where("uuid = ?", id).First(b).Error; err != nil {
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
	Config.DB.Where("uuid = ?", id).Delete(b)
	return nil
}
