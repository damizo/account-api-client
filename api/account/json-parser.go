package account

import (
	"bytes"
	json2 "encoding/json"
	"io"
	"io/ioutil"
	"log"
	"strings"
)

const ErrorMessageWildcard = "error_message"

func ParseFrom(data AccountData) *bytes.Buffer{
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
