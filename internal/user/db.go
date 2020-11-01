package user

import (
	"github.com/vmware/vending/external/db"
	"golang.org/x/crypto/bcrypt"
	"log"
)

type Users []User

func (u *User) FetchById() error{
	if err := db.Db.First(&u, u.ID).Error; err != nil {
		return err
	}
	return nil
}

func (u *User) FetchByUsername() error{
	if err := db.Db.Where("username = ?", u.Username).First(&u).Error; err != nil {
		return err
	}
	return nil
}

func (u *User) DeleteByUsername() error{
	if err := db.Db.Where("username = ?", u.Username).Delete(&User{}).Error; err != nil {
		return err
	}
	return nil
}

func (us *Users) GetAllUsers() error{
	var users []User
	if err := db.Db.Find(&us).Error; err != nil {
		return err
	}
	copy(*us, users)
	return nil
}

func (u *User) Save() error {
	hash, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.MinCost)
	if err != nil {
		log.Println(err)
	}
	u.Password = string(hash)
	if err = db.Db.Create(&u).Error; err != nil {
		return err
	}
	return nil
}

func (u *User) Authenticate() bool {
	du := &User{
		Username: u.Username,
	}
	if err := du.FetchByUsername(); err != nil {
		return false
	}
	if err := bcrypt.CompareHashAndPassword([]byte(du.Password), []byte(u.Password)); err != nil {
		return false
	}
	return true
}
