package models

import (
	"errors"
	"math"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm/clause"
)

type User struct {
	Id                uint64  `gorm:"primaryKey" json:"id"`
	Name              *string `gorm:"index" form:"name" json:"name"`
	Username          *string `gorm:"index,unique" form:"username" json:"username"`
	EncryptedPassword *string `json:"-"`
	CreatedAt         *string `gorm:"->" json:"created_at"`
	UpdatedAt         *string `gorm:"->" json:"updated_at"`
}

func GetPagedUser(currentPage int) (p *PaginateContent, err error) {
	var count int64
	var offset int
	var totalPage float64
	// var users []User{}
	users := []User{}

	countRes := Db.Table("users").Count(&count)

	if countRes.Error != nil {
		return nil, countRes.Error
	}

	ttlPage := float64(count) / 10

	totalPage = math.Ceil(ttlPage)

	offset = (currentPage - 1) * 10

	qRes := Db.Table("users").Limit(10).Offset(offset).Order("created_at desc").Find(&users)

	if qRes.Error != nil {
		return nil, qRes.Error
	}

	var pages []int
	for i := 1; i <= int(totalPage); i++ {
		pages = append(pages, i)
	}

	pc := &PaginateContent{
		Count:       count,
		CurrentPage: currentPage,
		TotalPage:   int(totalPage),
		Pages:       pages,
		Contents:    users,
	}

	return pc, nil

}

func (u *User) SetPassword(password, password_confirmation string) error {
	if password == password_confirmation {
		pb := []byte(password)

		hp, err := bcrypt.GenerateFromPassword(pb, bcrypt.DefaultCost)

		if err != nil {
			return err
		}

		sp := string(hp)

		u.EncryptedPassword = &sp

		return nil
	}

	return errors.New("password didn't match")
}

func (u *User) Insert() error {
	id, err := Sf.NextID()

	if err != nil {
		return err
	}

	u.Id = id

	qRes := Db.Table("users").Clauses(clause.Returning{}).Select("Id", "Name", "Username", "EncryptedPassword", "CreatedAt", "UpdatedAt").Create(u)

	if qRes.Error != nil {
		u.Id = 0
		return qRes.Error
	}

	return nil
}

func (u *User) Update() error {

	qRes := Db.Table("users").Save(u)

	if qRes.Error != nil {
		return qRes.Error
	}
	return nil
}

func DeleteUser(id uint64) (*User, error) {
	u := &User{}
	qRes := Db.Table("users").Clauses(clause.Returning{}).Select("Id", "Name", "Username", "EncryptedPassword", "CreatedAt", "UpdatedAt").Delete(u, id)

	if qRes.Error != nil {
		return nil, qRes.Error
	}

	return u, nil
}

func GetUserById(id uint64) (*User, error) {
	u := &User{}
	qRes := Db.Table("users").First(u, id)

	if qRes.Error != nil {
		return nil, qRes.Error
	}

	return u, nil
}

func (u *User) Copy() *User {

	return &User{
		Id:                u.Id,
		Name:              getStrPtr(*u.Name),
		Username:          getStrPtr(*u.Username),
		EncryptedPassword: getStrPtr(*u.EncryptedPassword),
		CreatedAt:         getStrPtr(*u.CreatedAt),
		UpdatedAt:         getStrPtr(*u.UpdatedAt),
	}
}

func getStrPtr(s string) *string {
	return &s
}
