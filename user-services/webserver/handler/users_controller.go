package handler

import (
	"net/http"

	us "github.com/asepnur/simple-grpc/user-services/modules/users"
)

// SelectUserController function
func (params *UserParams) SelectUserController() (users []User, code int, err error) {
	var tmpUser []us.User
	q := params.Q
	if q != "" {
		tmpUser, err = us.GetMultipleByFilter(q)
	} else {
		tmpUser, err = us.GetMultipleUser()
	}
	if err != nil {
		code = http.StatusInternalServerError
		return
	}
	for _, el := range tmpUser {
		ct := el.CreateTime.Format(layoutTime)
		if ct == emptyTime {
			ct = "-"
		}
		users = append(users, User{
			UserID:     el.UserID,
			UserEmail:  el.UserEmail,
			FullName:   el.FullName,
			MSISDN:     el.MSISDN,
			CreateTime: ct,
		})
	}
	code = http.StatusOK
	return
}
