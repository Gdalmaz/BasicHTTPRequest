package helpers

import (
	"auth/database"
	"auth/models"
)

func MailControl(mail string) (bool, error) {
	var user models.User

	err := database.DB.Db.Where("mail=?", mail).First(&user).Error
	if err != nil {
		return false, err
	}
	return true, nil
}
