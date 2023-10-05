package handler

import (
	"beautify/pkg/domain"
	mock "beautify/pkg/mock/userUsecaseMock"
	"beautify/pkg/utils/request"
	"beautify/pkg/utils/response"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"

	"github.com/jinzhu/copier"
	"github.com/stretchr/testify/assert"
)

func TestUserSignup(t *testing.T) {

	// ctx := context.Background()
	var usr response.UserSignUp
	testCase := map[string]struct {
		signupData    request.SignupUserData
		buildStub     func(useCaseMock *mock.MockUserService, signupData request.SignupUserData)
		checkResponse func(t *testing.T, responseRecorder *httptest.ResponseRecorder)
	}{
		"Valid Signup": {
			signupData: request.SignupUserData{
				UserName:  "JerinMathai",
				FirstName: "Jerin",
				LastName:  "Mathai",
				Age:       25,
				Phone:     "9446371883",
				Email:     "jerin@gmail.com",
				Password:  "password",
			},
			buildStub: func(useCaseMock *mock.MockUserService, signupData request.SignupUserData) {
				// copying signupData to domain.user for pass to Mock usecase
				var user domain.Users
				if err := copier.Copy(&user, signupData); err != nil {
					fmt.Println("Copy failed")
				}
				useCaseMock.EXPECT().SignUp(gomock.Any(), user).Times(1).Return(usr, nil)
			},
			checkResponse: func(t *testing.T, responseRecorder *httptest.ResponseRecorder) {
				assert.Equal(t, http.StatusOK, responseRecorder.Code)

			},
		},
		"Without email": {
			signupData: request.SignupUserData{
				UserName:  "JerinMathai",
				FirstName: "Jerin",
				LastName:  "Mathai",
				Age:       25,
				Phone:     "9446371883",
				Email:     "",
				Password:  "password",
			},
			buildStub: func(useCaseMock *mock.MockUserService, signupData request.SignupUserData) {
				// not expecting calls to usecase
			},
			checkResponse: func(t *testing.T, responseRecorder *httptest.ResponseRecorder) {
				responseStruct, err := getResponseStructFromResponseBody(responseRecorder.Body)
				assert.Nil(t, err)
				expectedMessage := "Missing or invalid entry"
				assert.Equal(t, expectedMessage, responseStruct.Message)
				assert.Equal(t, http.StatusBadRequest, responseRecorder.Code)

			},
		},
	}

	for testName, test := range testCase {
		test := test
		t.Run(testName, func(t *testing.T) {
			t.Parallel()
			ctrl := gomock.NewController(t)
			mockUseCase := mock.NewMockUserService(ctrl)
			test.buildStub(mockUseCase, test.signupData)

			AuthHandler := NewUserHandler(mockUseCase)

			server := gin.Default()
			server.POST("/signup", AuthHandler.UserSignup)

			jsonData, err := json.Marshal(test.signupData)
			assert.NoError(t, err)
			body := bytes.NewBuffer(jsonData)

			mockRequest, err := http.NewRequest(http.MethodPost, "/signup", body)
			assert.NoError(t, err)
			responseRecorder := httptest.NewRecorder()

			server.ServeHTTP(responseRecorder, mockRequest)
			test.checkResponse(t, responseRecorder)

		})

	}
}

// convert / un marshal response body to response.Response struct
func getResponseStructFromResponseBody(responseBody *bytes.Buffer) (responseStruct response.Response, err error) {
	data, err := io.ReadAll(responseBody)
	json.Unmarshal(data, &responseStruct)
	return
}
