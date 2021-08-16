package account

import (
	json2 "encoding/json"
	 "./common"
	 "./model"
	"fmt"
	"io/ioutil"
	"log"
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
		Timeout:       5000,
	}
	return &AccountClient{
		settings: settings,
		client:   client,
	}
}

func (a AccountClient) CreateAccount(accountData model.AccountData) model.AccountCreatedResponse {
	data := common.ParseFrom(accountData)
	request, _ := http.NewRequest(http.MethodPost, fmt.Sprintf("%s%s", a.settings.url, ApiSuffix), data)
	request.Header.Add("Content-Type", "application/vnd.api+json")
	response, err := a.client.Do(request)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	bodyBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	accountCreated := model.AccountCreatedResponse{}
	err = json2.Unmarshal(bodyBytes, &accountCreated)
	return accountCreated
}

func (a AccountClient) FetchAccount(id string) model.FetchAccountQuery {
	getRequest, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s%s/%s", a.settings.url, ApiSuffix, id), nil)
	if err != nil {
		log.Fatal(err)
	}
	response, err := a.client.Do(getRequest)
	defer response.Body.Close()

	bodyBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	data := model.FetchAccountQuery{}
	err = json2.Unmarshal(bodyBytes, &data)
	return data
}

func (a AccountClient) DeleteAccount(id string) *http.Response {
	deleteRequest, err := http.NewRequest(http.MethodDelete, fmt.Sprintf("%s%s/%s?version=0", a.settings.url, ApiSuffix, id), nil)
	fmt.Println(deleteRequest.URL)
	response, err := a.client.Do(deleteRequest)
	if err != nil {
		log.Fatal(err)
	}
	return response
}

func (a AccountClient) healthCheck() {
	get, err := http.Get(fmt.Sprintf("%s/%s", a.settings.url, "v1/health"))
	if nil != err {
		fmt.Println(err)
	}
	fmt.Println(get)
}

type AccountServerSettings struct {
	url string
}

func main() {
}
