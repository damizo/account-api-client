package account

import (
	"bytes"
	json2 "encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

const ErrorMessageFieldName = "error_message"

func ParseFrom(data AccountData) *bytes.Buffer {
	accountData := CreateAccountRequest{Data: data}
	json, _ := json2.Marshal(accountData)
	reader := bytes.NewBuffer(json)
	return reader
}

func ParseToAccountCreatedResponse(data io.Reader) (AccountCreatedResponse, Error) {
	bodyBytes, err := ioutil.ReadAll(data)
	if err != nil {
		log.Fatal(err)
	}
	responseInString := string(bodyBytes)
	if strings.Contains(responseInString, ErrorMessageFieldName) {
		errorResponse := Error{}
		err = json2.Unmarshal(bodyBytes, &errorResponse)
		return AccountCreatedResponse{}, errorResponse
	}
	accountCreatedResponse := AccountCreatedResponse{}
	err = json2.Unmarshal(bodyBytes, &accountCreatedResponse)
	return accountCreatedResponse, Error{}
}

func ParseToFetchQueryResponse(data io.Reader) (FetchAccountResponse, Error) {
	bodyBytes, err := ioutil.ReadAll(data)
	if err != nil {
		log.Fatal(err)
	}
	responseInString := string(bodyBytes)
	if strings.Contains(responseInString, ErrorMessageFieldName) {
		errorResponse := Error{}
		err = json2.Unmarshal(bodyBytes, &errorResponse)
		return FetchAccountResponse{}, errorResponse
	}
	fetchAccountQuery := FetchAccountResponse{}
	err = json2.Unmarshal(bodyBytes, &fetchAccountQuery)
	return fetchAccountQuery, Error{}
}

func ParseToAccountDeletedResponse(id string, response *http.Response) (AccountDeletedResponse, Error) {
	if isSuccess(response) {
		return AccountDeletedResponse{ID: id}, Error{}
	} else {
		return AccountDeletedResponse{}, Error{ErrorMessage: fmt.Sprintf("delete account for id %s ended with failure, status code: %d", id, response.StatusCode)}
	}
}

func isSuccess(response *http.Response) bool {
	HttpCodeOk := 200
	HttpCodeCreated := 201
	HttpCodeNoContent := 204
	return response.StatusCode == HttpCodeOk || response.StatusCode == HttpCodeCreated || response.StatusCode == HttpCodeNoContent
}
