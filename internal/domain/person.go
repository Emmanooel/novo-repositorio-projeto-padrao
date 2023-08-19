package domain

type Person struct {
	Login    Login  `gorm:"embedded"`
	Name     string `json:"name"`
	Birthday string `json:"birthday"`
	Phone    int    `json:"phone"`
	Score    int    `json:"score"`
}

type Login struct {
	Email    string `json:"email"`
	Password string `json:"password"`
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
