package UserServiceTest

import (
	"testing"

	"darshub.dev/src/UserService/data"
	"darshub.dev/src/util"
)

// func Test_FromJSON_WithNillShouldReturnError(t *testing.T) {
// 	err := util.FromJSON(nil, nil)

// 	if err == nil {
// 		t.Errorf("Response code is not bad requst")
// 	}
// }

// func Test_FromJSON_ForUser_WithNillAsBody_ShouldReturnError(t *testing.T) {
// 	//newuser := &data.User{}

// 	//err := util.FromJSON(newuser, nil)
// 	// err = nil

// 	// if err == nil {
// 	// 	t.Errorf("Response code is not bad requst")
// 	// }
// }

func Test_ToJSON_User_WithNillAsBody_ShouldReturnError(t *testing.T) {
	newuser := &data.User{}

	err := util.ToJSON(newuser, nil)
	err = nil

	if err == nil {
		t.Errorf("Response code is not bad requst")
	}
}

func Test_PasswordDecript(t *testing.T) {
	newuser := &data.User{}

	err := util.ToJSON(newuser, nil)
	err = nil

	if err == nil {
		t.Errorf("Response code is not bad requst")
	}
}
