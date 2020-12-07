package models

import (
	"errors"
	"github.com/jinzhu/gorm"
	
)

type Beer struct{
	ID        uint32    `gorm:"primary_key;auto_increment" json:"id"`
	Name      string    `gorm:"size:255;not null;" json:"name"`
	Brewery   string    `gorm:"size:255;not null;" json:"brewery"`
	Country   string    `gorm:"size:100;not null;" json:"country"`

}

func (b *Beer) FindBeers(db *gorm.DB) (*[]Beer, error) {
	var err error
	beers := []Beer{}
	err = db.Debug().Model(&Beer{}).Limit(100).Find(&beers).Error
	if err != nil {
		return &[]Beer{}, err
	}
	return &beers, err
}

func (b *Beer) FindBeerByID(db *gorm.DB, uid uint32) (*Beer,error){
	var err error
	err = db.Debug().Model(Beer{}).Where("id = ?", uid).Take(&b).Error
	if err != nil{
		return &Beer{}, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return &Beer{}, errors.New("Beer Not Found")
	}
	return b, err

}