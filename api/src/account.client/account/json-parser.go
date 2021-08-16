package account

import (
	"bytes"
	json2 "encoding/json"
	"io"
	"io/ioutil"
	"log"
)

func ParseFrom(data AccountData) *bytes.Buffer{
	accountData := CreateAccountCommand{Data: data}
	json, _ := json2.Marshal(accountData)
	reader := bytes.NewBuffer(json)
	return reader
}

func ParseTo(data io.Reader) AccountCreatedResponse {
	bodyBytes, err := ioutil.ReadAll(data)
	if err != nil {
		log.Fatal(err)
	}
	accountCreated := AccountCreatedResponse{}
	err = json2.Unmarshal(bodyBytes, &accountCreated)
	return accountCreated
}