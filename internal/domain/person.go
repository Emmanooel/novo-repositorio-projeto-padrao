package domain

import (
	"time"

	"gorm.io/gorm"
)

type Person struct {
	gorm.Model
	ID       int
	Login    Login  `gorm:"embedded"`
	Name     string `gorm:"unique" json:"name"`
	Birthday string `gorm:"type:date" json:"birthday"`
	Phone    int    `json:"phone"`
	Score    int    `json:"score"`
}

type Login struct {
	Email    string `gorm:"unique" json:"email"`
	Password string `json:"password"`
}

type Ranking struct {
	Name  string `json:"name"`
	Score int    `json:"score"`
}

type CustomDateType struct {
	time.Time
}

func (personBody *Person) VerifyPerson() (bool, string) {
	if personBody.Name == "" {
		return false, "name required"
	} else if personBody.Birthday == "" {
		return false, "birthday required"
	} else if personBody.Phone == 0 {
		return false, "phone required"
	}

	return true, ""
}

func (loginBody *Login) VerifyLogin() (bool, string) {
	if loginBody.Email == "" {
		return false, "email required"
	} else if loginBody.Password == "" {
		return false, "password required"
	}

	return true, ""
}
