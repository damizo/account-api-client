package account

import (
	"bytes"
	json2 "encoding/json"
)

func ParseFrom(data AccountData) *bytes.Buffer{
	accountData := CreateAccountCommand{Data: data}
	json, _ := json2.Marshal(accountData)
	reader := bytes.NewBuffer(json)
	return reader
}

func to(data CreateAccountCommand) *CreateAccountCommand {
	return nil
}