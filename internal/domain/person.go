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

func (personBody Person) VerifyStruct() (bool, string) {
	if personBody.Name == "" {
		err := "name required"
		return false, err
	} else if personBody.Birthday == "" {
		err := "birthday required"
		return false, err
	} else if personBody.Phone == 0 {
		err := "phone required"
		return false, err
	}

	return true, ""
}
