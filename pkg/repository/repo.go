package repository

import "github.com/jinzhu/gorm"

type Repository struct{}

func (r Repository) GetUserActionAndSubjectByEmail(db *gorm.DB, email string) ([]string, []string, error) {
	var result Result
	err := db.Table("users").Select("actions, subjects").Joins("JOIN roles ON users.role_id = roles.id").Where("users.email = ?", email).Scan(&result).Error

	if err != nil {
		return nil, nil, err
	}

	return result.Actions, result.Subjects, nil
}

type Result struct {
	Actions  []string
	Subjects []string
}
