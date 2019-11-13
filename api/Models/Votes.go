package Models

import (
	"fmt"

	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/taingk/goxit/api/Config"
	"gopkg.in/go-playground/validator.v9"
)

func GetAllVote(b *[]Vote) (err error) {
	if err = Config.DB.Where("is_deleted = 0").Find(b).Error; err != nil {
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
	b.IsDeleted = 0
	if err = Config.DB.Create(b).Error; err != nil {
		return err
	}
	return nil
}

func GetOneVote(b *Vote, uuid string) (err error) {
	if err := Config.DB.Where("uuid = ? AND is_deleted = 0", uuid).First(b).Error; err != nil {
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
	b.IsDeleted = 1
	Config.DB.Where("uuid = ?", uuid).Save(b)
	return nil
}
