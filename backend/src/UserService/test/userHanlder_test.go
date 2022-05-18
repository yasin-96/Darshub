package UserServiceTest

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	userHandler "dev.azure.com/learn-website-orga/_git/learn-website/src/UserService/handlers"
)

func Test_Handler_CreatedNewUser_WithEmptyBodyShouldNotWork(t *testing.T) {
	req := httptest.NewRequest(http.MethodPost, "/user", nil)
	w := httptest.NewRecorder()
	userHandler.RegisterUser(w, req)

	res := w.Result()
	statusCode := res.StatusCode
	fmt.Println(statusCode)
	fmt.Println(res)
	if statusCode != http.StatusBadRequest {
		t.Errorf("Response code is not bad requst")
	}
}

func Test_Handler_CreatedNewUser_WithNoInputShouldNotWork(t *testing.T) {
	user := strings.NewReader(``)

	req := httptest.NewRequest(http.MethodPost, "/user", user)
	w := httptest.NewRecorder()
	userHandler.RegisterUser(w, req)

	res := w.Result()
	defer res.Body.Close()
	_, err := ioutil.ReadAll(res.Body)

	if err != nil {
		t.Errorf("expected error to be nil got %v", err)
		t.Errorf("Invalid Json")
	}

	if res.StatusCode != http.StatusBadRequest {
		t.Errorf("Invalid json")
	}
}

func Test_Handler_CreatedNewUser_WithInvalidJsonSHhouldNotWork(t *testing.T) {
	user := strings.NewReader(`
		"password":   "",
		"first_Name": "",
		"last_Name":  "",
		"bio":        "",
		"birthday":   "",
	`)

	req := httptest.NewRequest(http.MethodPost, "/user", user)
	w := httptest.NewRecorder()
	userHandler.RegisterUser(w, req)

	res := w.Result()
	defer res.Body.Close()
	_, err := ioutil.ReadAll(res.Body)

	if err != nil {
		t.Errorf("expected error to be nil got %v", err)
		t.Errorf("Invalid Json")
	}
	fmt.Println(res.StatusCode)
	if res.StatusCode != http.StatusBadRequest {
		t.Errorf("Invalid Json")
	}
}

func Test_Handler_CreatedNewUser_WithEmptyInputShouldNotWork(t *testing.T) {
	user := strings.NewReader(`
	{
	"password" :"",
	"first_name" :"",
	"last_name" :"",
	"birthday" :"",
	"avatar" :"",
	"email" :"",
	"telNr" :"",
	"company" :"",
	"occupation" :"",
	"school" :"",
	"subject" :"",
	"country" :"",
	"isActive" :"",
	"bio" :"",
	"role" :""
	}
	`)

	req := httptest.NewRequest(http.MethodPost, "/user", user)
	w := httptest.NewRecorder()
	userHandler.RegisterUser(w, req)

	res := w.Result()
	defer res.Body.Close()
	_, err := ioutil.ReadAll(res.Body)

	if err != nil {
		t.Errorf("expected error to be nil got %v", err)
		t.Errorf("Invalid Json")
	}
	fmt.Println(res.StatusCode)
	if res.StatusCode != http.StatusBadRequest {
		t.Errorf("Invalid Json")
	}
}

func Test_Handler_CreatedNewUser_WithMinimumProps(t *testing.T) {
	user := strings.NewReader(`
	{
	"password" :"1234",
	"first_name" :"aa",
	"last_name" :"bb",
	"birthday" : "",
	"avatar" :"",
	"email" :"aa.bb@cc.com",
	"telNr" :"000-132-13",
	"company" :"",
	"occupation" :"",
	"school" :"",
	"subject" :"",
	"country" :"",
	"isActive" :"",
	"bio" :"male",
	"role" :""
	}
	`)

	req := httptest.NewRequest(http.MethodPost, "/user", user)
	w := httptest.NewRecorder()
	userHandler.RegisterUser(w, req)

	res := w.Result()
	defer res.Body.Close()
	_, err := ioutil.ReadAll(res.Body)

	if err != nil {
		t.Errorf("expected error to be nil got %v", err)
		t.Errorf("Invalid Json")
	}
	fmt.Println(res.StatusCode)
	if res.StatusCode != http.StatusBadRequest {
		t.Errorf("Invalid Json")
	}
}

// func Test_Handler_CreatedNewUser_WithMinimumProps(t *testing.T) {
// 	user := strings.NewReader(`
// 	{
// 	"password" :"1234",
// 	"first_name" :"aa",
// 	"last_name" :"bb",
// 	"birthday" : "",
// 	"avatar" :"",
// 	"email" :"aa.bb@cc.com",
// 	"telNr" :"000-132-13",
// 	"company" :"",
// 	"occupation" :"",
// 	"school" :"",
// 	"subject" :"",
// 	"country" :"",
// 	"isActive" :"",
// 	"bio" :"male",
// 	"role" :""
// 	}
// 	`)

// 	req := httptest.NewRequest(http.MethodPost, "/user", user)
// 	w := httptest.NewRecorder()
// 	userHandler.RegisterUser(w, req)

// 	res := w.Result()
// 	defer res.Body.Close()
// 	_, err := ioutil.ReadAll(res.Body)

// 	if err != nil {
// 		t.Errorf("expected error to be nil got %v", err)
// 		t.Errorf("Invalid Json")
// 	}
// 	fmt.Println(res.StatusCode)
// 	if res.StatusCode != http.StatusBadRequest {
// 		t.Errorf("Invalid Json")
// 	}
// }

func Test_Handler_LogUser_WithEmptyBodyShouldNotWork(t *testing.T) {
	req := httptest.NewRequest(http.MethodPost, "/session", nil)
	w := httptest.NewRecorder()
	userHandler.Login(w, req)

	res := w.Result()
	statusCode := res.StatusCode
	fmt.Println(res)
	// fmt.Println(statusCode)
	if statusCode != http.StatusBadRequest {
		t.Errorf("Response code is not bad requst")
	}
}

func Test_Handler_LogUser_WithInValidData(t *testing.T) {
	login := strings.NewReader(`{"email": "", "password": ""}`)

	req := httptest.NewRequest(http.MethodPost, "/session", login)
	w := httptest.NewRecorder()
	userHandler.Login(w, req)

	res := w.Result()
	statusCode := res.StatusCode

	if statusCode != http.StatusBadRequest {
		t.Errorf("User should not have access")
	}
}

func Test_Handler_LogUser_WithMissingPasswordShouldNotWork(t *testing.T) {
	login := strings.NewReader(`{"email": "test@test.com", "password": ""}`)

	req := httptest.NewRequest(http.MethodPost, "/session", login)
	w := httptest.NewRecorder()
	userHandler.Login(w, req)

	res := w.Result()
	statusCode := res.StatusCode

	if statusCode != http.StatusBadRequest {
		t.Errorf("Missing Information")
	}
}

func Test_Handler_LogUser_WithMissingEmailShouldNotWork(t *testing.T) {
	login := strings.NewReader(`{"email": "", "password": "132456"}`)

	req := httptest.NewRequest(http.MethodPost, "/session", login)
	w := httptest.NewRecorder()
	userHandler.Login(w, req)

	res := w.Result()
	statusCode := res.StatusCode

	if statusCode != http.StatusBadRequest {
		t.Errorf("Missing Information")
	}
}

func Test_Handler_LogUser_WithInValidJson(t *testing.T) {
	login := strings.NewReader(`{"email": "", "pass": ""}`)

	req := httptest.NewRequest(http.MethodPost, "/session", login)
	w := httptest.NewRecorder()
	userHandler.Login(w, req)

	res := w.Result()
	statusCode := res.StatusCode

	if statusCode != http.StatusBadRequest {
		t.Errorf("User should have email and password to pass")
	}
}

func Test_Handler_LogUser_WithMissingPropShouldNotWork(t *testing.T) {
	login := strings.NewReader(`{"email": "", }`)

	req := httptest.NewRequest(http.MethodPost, "/session", login)
	w := httptest.NewRecorder()
	userHandler.Login(w, req)

	res := w.Result()
	statusCode := res.StatusCode

	if statusCode != http.StatusInternalServerError {
		t.Errorf("User should have email and password to password as prop")
	}
}

func Test_Handler_LogUser_WithValidData(t *testing.T) {
	login := strings.NewReader(`{"email": "test@test.com", "password": "tt"}`)

	req := httptest.NewRequest(http.MethodPost, "/session", login)
	w := httptest.NewRecorder()
	userHandler.Login(w, req)

	res := w.Result()
	statusCode := res.StatusCode

	if statusCode != http.StatusOK {
		t.Errorf("User should could not logged in.")
	}
}
