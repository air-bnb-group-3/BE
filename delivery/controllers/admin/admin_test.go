package admin

import (
	"app_airbnb/delivery/controllers/auth"
	"app_airbnb/delivery/middlewares"
	"app_airbnb/entities"
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

type MockAuthLib struct{}

func (m *MockAuthLib) Login(email, password string) (entities.User, error) {
	if email != "test@admin.com" && password != "test" {
		return entities.User{}, errors.New("record not found")
	}
	return entities.User{Model: gorm.Model{ID: 1}, Email: email, Password: password}, nil
}

type MockAuthAdminLib struct{}

func (m *MockAuthAdminLib) Login(email, password string) (entities.User, error) {
	if email != "admin" && password != "admin" {
		return entities.User{}, errors.New("record not found")
	}
	return entities.User{Model: gorm.Model{ID: 1}, Email: email, Password: password}, nil
}

type MockUserLib struct{}

func (m *MockUserLib) Register(newUser entities.User) (entities.User, error) {
	if newUser.Name != "test" && newUser.Email != "test" && newUser.Password != "test" {
		return entities.User{}, errors.New("record not found")
	}
	return entities.User{}, nil

}

func (m *MockUserLib) GetById(userId int) (entities.User, error) {
	return entities.User{}, nil
}

func (m *MockUserLib) Update(userId int, newUser entities.User) (entities.User, error) {
	return entities.User{}, nil
}

func (m *MockUserLib) Delete(Userid int) error {
	return nil
}

func (m *MockUserLib) GetAll() ([]entities.User, error) {
	return []entities.User{}, nil
}

type MockFalseLib struct{}

func (m *MockFalseLib) Register(newUser entities.User) (entities.User, error) {
	if newUser.Name != "test" && newUser.Email != "test" && newUser.Password != "test" {
		return entities.User{}, errors.New("record not found")
	}
	return entities.User{}, errors.New("invalid input")

}

func (m *MockFalseLib) GetById(userId int) (entities.User, error) {
	return entities.User{}, errors.New("False Object")
}

func (m *MockFalseLib) Update(userId int, newUser entities.User) (entities.User, error) {
	return entities.User{}, errors.New("False Object")
}

func (m *MockFalseLib) Delete(Userid int) error {
	return errors.New("False Object")
}

func (m *MockFalseLib) GetAll() ([]entities.User, error) {
	return []entities.User{}, errors.New("False Object")
}

func TestCreate(t *testing.T) {

	t.Run("Failed to Create", func(t *testing.T) {
		e := echo.New()

		reqBody, _ := json.Marshal(AdminCreateRequestFormat{
			Name:     "",
			Email:    "",
			Password: "",
		})

		req := httptest.NewRequest(http.MethodPost, "/users/register", bytes.NewBuffer(reqBody))
		res := httptest.NewRecorder()
		req.Header.Set("Content-Type", "application/json")
		// req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", nil))

		context := e.NewContext(req, res)
		context.SetPath("/users")

		userController := New(&MockFalseLib{})
		userController.Register()(context)

		response := ResponseFormat{}

		json.Unmarshal([]byte(res.Body.Bytes()), &response)

		assert.Equal(t, 400, response.Code)
		assert.Equal(t, "There is some problem from input", response.Message)

	})

	t.Run("Failed to Access", func(t *testing.T) {
		e := echo.New()

		reqBody, _ := json.Marshal(map[string]string{
			"email":    "anonim",
			"password": "anonim",
		})

		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(reqBody))
		res := httptest.NewRecorder()

		req.Header.Set("Content-Type", "application/json")
		// req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtToken))

		context := e.NewContext(req, res)
		context.SetPath("/users")

		userController := New(&MockFalseLib{})
		userController.Register()(context)

		response := ResponseFormat{}

		json.Unmarshal([]byte(res.Body.Bytes()), &response)

		assert.Equal(t, 500, response.Code)
		assert.Equal(t, "There is some error on server", response.Message)

	})

	t.Run("Success to Create", func(t *testing.T) {
		e := echo.New()

		reqBody, _ := json.Marshal(map[string]string{
			"email":    "test",
			"password": "test",
			"name":     "test",
		})

		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(reqBody))
		res := httptest.NewRecorder()

		req.Header.Set("Content-Type", "application/json")
		// req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtToken))

		context := e.NewContext(req, res)
		context.SetPath("/users")

		userController := New(&MockUserLib{})
		userController.Register()(context)

		response := ResponseFormat{}

		json.Unmarshal([]byte(res.Body.Bytes()), &response)

		assert.Equal(t, 201, response.Code)
		assert.Equal(t, "Success Create Admin", response.Message)

	})
}

func TestGetAll(t *testing.T) {

	var jwtToken string

	t.Run("Success Login", func(t *testing.T) {
		e := echo.New()

		reqBody, _ := json.Marshal(map[string]string{
			"email":    "test",
			"password": "test",
		})

		req := httptest.NewRequest(http.MethodPost, "/admin/login", bytes.NewBuffer(reqBody))
		res := httptest.NewRecorder()
		req.Header.Set("Content-Type", "application/json")
		context := e.NewContext(req, res)

		authController := auth.New(&MockAuthLib{})
		authController.Login()(context)

		response := auth.LoginRespFormat{}
		json.Unmarshal([]byte(res.Body.Bytes()), &response)

		jwtToken = response.Data["token"].(string)

		assert.Equal(t, response.Message, "berhasil masuk, mendapatkan token baru")
		assert.NotNil(t, response.Data["token"])
	})

	t.Run("Success to Create", func(t *testing.T) {
		e := echo.New()

		req := httptest.NewRequest(http.MethodGet, "/", bytes.NewBuffer(nil))
		res := httptest.NewRecorder()

		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtToken))

		context := e.NewContext(req, res)
		context.SetPath("/users")

		userController := New(&MockUserLib{})
		if err := middlewares.JwtMiddleware()(userController.GetAll())(context); err != nil {
			return
		}

		response := ResponseFormat{}

		json.Unmarshal([]byte(res.Body.Bytes()), &response)
		assert.Equal(t, 200, response.Code)
		assert.Equal(t, "Success Get All User", response.Message)
	})
}
