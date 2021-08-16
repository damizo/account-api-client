package common

import (
	model "../model"
	"bytes"
	json2 "encoding/json"
)

func ParseFrom(data model.AccountData) *bytes.Buffer{
	accountData := model.CreateAccountCommand{Data: data}
	json, _ := json2.Marshal(accountData)
	reader := bytes.NewBuffer(json)
	return reader
}

func to(data model.CreateAccountCommand) *model.CreateAccountCommand {
	return nil
}