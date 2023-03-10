package checkers

import (
	"pizza-site-backend/models"
)

func (u *models.User) IsUserEmpty() bool {
	userMap := structs.Map(u)
	for _, i := range userMap {
		if i != "" {
			return false
		}
	}
	return true
}
