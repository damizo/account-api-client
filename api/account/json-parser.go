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

const ErrorMessageWildcard = "error_message"

func ParseFrom(data AccountData) *bytes.Buffer {
	accountData := CreateAccountCommand{Data: data}
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

	if strings.Contains(responseInString, ErrorMessageWildcard) {
		errorResponse := Error{}
		err = json2.Unmarshal(bodyBytes, &errorResponse)
		return AccountCreatedResponse{}, errorResponse
	}
	accountCreatedResponse := AccountCreatedResponse{}
	err = json2.Unmarshal(bodyBytes, &accountCreatedResponse)
	return accountCreatedResponse, Error{}
}

func ParseToFetchQueryResponse(data io.Reader) (FetchAccountQuery, Error) {
	bodyBytes, err := ioutil.ReadAll(data)
	if err != nil {
		log.Fatal(err)
	}
	responseInString := string(bodyBytes)

	if strings.Contains(responseInString, ErrorMessageWildcard) {
		errorResponse := Error{}
		err = json2.Unmarshal(bodyBytes, &errorResponse)
		return FetchAccountQuery{}, errorResponse
	}
	fetchAccountQuery := FetchAccountQuery{}
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
	return response.StatusCode == 200 || response.StatusCode == 201 || response.StatusCode == 204
}
