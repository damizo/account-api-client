package account

import (
	"fmt"
	"net/http"
)

type AccountClient struct {
	settings AccountServerSettings
	client   http.Client
}

const ApiSuffix = "/v1/organisation/accounts"

func NewAccountClient(settings AccountServerSettings) *AccountClient {
	client := http.Client{
		Transport:     nil,
		CheckRedirect: nil,
		Jar:           nil,
		Timeout:       0,
	}
	return &AccountClient{
		settings: settings,
		client:   client,
	}
}

func (a AccountClient) CreateAccount(accountData AccountData) (AccountCreatedResponse, Error) {
	data := ParseFrom(accountData)
	request := NewRequestExecutionBuilder(a.client).
		withUrl(a.settings.url).
		withUrlSuffix(ApiSuffix).
		withMethod(http.MethodPost).
		withBody(data).
		withHeader("Content-Type", "application/vnd.api+json").
		build().
		handle()
	return ParseToAccountCreatedResponse(request.Body)
}


func (a AccountClient) FetchAccount(id string) (FetchAccountQuery, Error) {
	request := NewRequestExecutionBuilder(a.client).
		withUrl(a.settings.url).
		withUrlSuffix(fmt.Sprintf("%s%s", ApiSuffix, "/{id}")).
		withMethod(http.MethodGet).
		withParam("id", id).
		build().
		handle()
	return ParseToFetchQueryResponse(request.Body)
}

func (a AccountClient) DeleteAccount(id string) {
	NewRequestExecutionBuilder(a.client).
		withUrl(a.settings.url).
		withUrlSuffix(fmt.Sprintf("%s%s", ApiSuffix, "/{id}?version=0")).
		withMethod(http.MethodDelete).
		withParam("id", id).
		build().
		handle()
}


