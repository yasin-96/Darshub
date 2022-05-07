package UserServiceTest

import (
	"testing"

	"dev.azure.com/learn-website-orga/_git/learn-website/backend/src/UserService/data"
	"dev.azure.com/learn-website-orga/_git/learn-website/backend/src/util"
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
