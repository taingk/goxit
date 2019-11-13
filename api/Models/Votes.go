package Models

import (
	"fmt"

	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/taingk/goxit/api/Config"
	"gopkg.in/go-playground/validator.v9"
)

func GetAllVote(b *[]Vote) (err error) {
	if err = Config.DB.Find(b).Error; err != nil {
		return err
	}
	return nil
}

func AddNewVote(b *Vote) (err error) {
	var validate *validator.Validate
	validate = validator.New()
	if err = validate.Struct(b); err != nil {
		return err
	}

	if err = Config.DB.Create(b).Error; err != nil {
		return err
	}
	return nil
}

func GetOneVote(b *Vote, uuid string) (err error) {
	if err := Config.DB.Where("uuid = ?", uuid).First(b).Error; err != nil {
		return err
	}
	return nil
}

func PutOneVote(b *Vote, uuid string) (err error) {
	var validate *validator.Validate
	validate = validator.New()
	if err = validate.Struct(b); err != nil {
		return err
	}

	fmt.Println(b)
	Config.DB.Save(b)
	return nil
}

func DeleteVote(b *Vote, uuid string) (err error) {
	Config.DB.Where("uuid = ?", uuid).Delete(b)
	return nil
}
