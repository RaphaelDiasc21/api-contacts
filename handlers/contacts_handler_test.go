package handlers_test

import (
	"api-contacts/resources"
	// "fmt"
	// "net/http"
	// "net/http/httptest"
	// "strings"
	// "testing"

	// "github.com/gorilla/mux"
	"github.com/stretchr/testify/mock"
)

type MockRepository struct {
	mock.Mock
}

func (mock *MockRepository) InsertContacts(contactResources *resources.ContactResources) error {
	mock.Called()
	return nil
}

// func TestPostContactsSuccess(t *testing.T) {
// 	body := strings.NewReader("{\"contacts\":[{\"_id\":0,\"name\":\"TESTE\",\"cellphone\":\"5519888888888\"}]}")
// 	req := httptest.NewRequest(http.MethodPost,"http://localhost.com",body)
// 	w := httptest.NewRecorder()

// 	mockRepository := new(MockRepository)

// 	mockRepository.On("InsertContacts").Return(nil)

// 	req = mux.SetURLVars(req,map[string]string{"client":"macapa"})
// 	CreateContactHandler(w,req)

// 	fmt.Println(w.Result().StatusCode)
// }