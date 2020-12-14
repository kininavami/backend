package product

import (
	"errors"
	"github.com/vmware/vending/external/db"
)

type Products []Product

func (p *Product) FetchById() error{
	if err := db.Db.First(&p, p.ID).Error; err != nil {
		return err
	}
	return nil
}

func (p *Product) FetchByName() error{
	if err := db.Db.Where("name = ?", p.Name).First(&p).Error; err != nil {
		return err
	}
	return nil
}

func (p *Product) DeleteByName() error{
	if err := p.FetchByName(); err != nil {
		return err
	}
	if p.Name == "" {
		return errors.New("record not found")
	}
	if err := db.Db.Where("name = ?", p.Name).Delete(&Product{}).Error; err != nil {
		return err
	}
	return nil
}

func (ps *Products) GetAllProducts() error{
	var products []Product
	if err := db.Db.Find(&ps).Error; err != nil {
		return err
	}
	copy(*ps, products)
	return nil
}

func (p *Product) Save() error {
	if err := db.Db.Create(&p).Error; err != nil {
		return err
	}
	return nil
}
