package Models

import (
	"fmt"

	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/taingk/goxit/api/Config"
)

func GetAllVote(b *[]Vote) (err error) {
	if err = Config.DB.Find(b).Error; err != nil {
		return err
	}
	return nil
}

func AddNewVote(b *Vote) (err error) {
	if err = Config.DB.Create(b).Error; err != nil {
		return err
	}
	return nil
}

func GetOneVote(b *Vote, id string) (err error) {
	if err := Config.DB.Where("uuid = ?", id).First(b).Error; err != nil {
		return err
	}
	return nil
}

func PutOneVote(b *Vote, id string) (err error) {
	fmt.Println(b)
	Config.DB.Save(b)
	return nil
}

func DeleteVote(b *Vote, id string) (err error) {
	Config.DB.Where("uuid = ?", id).Delete(b)
	return nil
}
