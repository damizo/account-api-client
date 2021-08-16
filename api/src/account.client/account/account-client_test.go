package account

import (
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/testcontainers/testcontainers-go"
	"path/filepath"
	"strings"
	"testing"
	"time"
)

var settings = AccountServerSettings{url: "http://0.0.0.0:8087"}
var accountClient = NewAccountClient(settings)
var id = uuid.New().String()
var compose = testcontainers.NewLocalDockerCompose(nil, "")

func TestRunDockerContainers(t *testing.T) {
	abs, _ := filepath.Abs("../../../../docker-compose.yml")
	composeFilePaths := []string{abs}

	identifier := strings.ToLower(uuid.New().String())
	compose := testcontainers.NewLocalDockerCompose(composeFilePaths, identifier)
	compose.
		WithCommand([]string{"up", "-d"}).
		Invoke()
}

func Test_Should_Create_Account(t *testing.T) {
	time.Sleep(3 * time.Second)
	var country = "GB"
	names := []string{"Sam", "Holder"}

	accountData := buildAccountData(names, country, id)
	actualAccount := accountClient.CreateAccount(accountData)
	version := int64(0)

	expectedCreatedAccount := AccountCreatedResponse{Data: AccountCreated{
		Attributes: AccountAttributes{
			BankID:       "400302",
			BankIDCode:   "GBDSC",
			BaseCurrency: "GBP",
			Bic:          "NWBKGB42",
			Country:      &country,
			Name:         names,
		},
		ID:             id,
		OrganisationID: "eb0bd6f5-c3f5-44b2-b677-acd23cdde73c",
		Type:           "accounts",
		Version:        &version,
		CreatedOn:      time.Time{},
		ModifiedOn:     time.Time{},
	},
		Link: Link{Self: "/v1/organisation/accounts/" + id},
	}

	assert.Equal(t, expectedCreatedAccount.Link, actualAccount.Link)
	assert.Equal(t, expectedCreatedAccount.Data.Attributes, actualAccount.Data.Attributes)
	assert.Equal(t, expectedCreatedAccount.Data.Type, actualAccount.Data.Type)
	assert.Equal(t, expectedCreatedAccount.Data.ID, actualAccount.Data.ID)
}

func Test_Should_Fetch_Account(t *testing.T) {
	var country = "GB"
	names := []string{"Sam", "Holder"}

	actualAccount := accountClient.FetchAccount(id)
	version := int64(0)
	expectedFetchAccountQuery := FetchAccountQuery{Data: AccountData{
		Attributes: AccountAttributes{
			BankID:       "400302",
			BankIDCode:   "GBDSC",
			BaseCurrency: "GBP",
			Bic:          "NWBKGB42",
			Country:      &country,
			Name:         names,
		},
		ID:             id,
		Version:        &version,
		OrganisationID: "eb0bd6f5-c3f5-44b2-b677-acd23cdde73c",
		Type:           "accounts",
	}, Link: Link{Self: "/v1/organisation/accounts/" + id}}
	assert.Equal(t, expectedFetchAccountQuery, actualAccount)
}

func Test_Should_Delete_Account(t *testing.T) {
	actualResponse := accountClient.DeleteAccount(id)
	assert.Equal(t, actualResponse.StatusCode, 204)
	fetchAccountResult := accountClient.FetchAccount(id)
	assert.Emptyf(t, fetchAccountResult, "")
}

func buildAccountData(names []string, country string, id string) AccountData {
	return AccountData{
		Attributes: AccountAttributes{
			BankID:       "400302",
			BankIDCode:   "GBDSC",
			BaseCurrency: "GBP",
			Bic:          "NWBKGB42",
			Country:      &country,
			Name:         names,
		},
		ID:             id,
		OrganisationID: "eb0bd6f5-c3f5-44b2-b677-acd23cdde73c",
		Type:           "accounts",
	}
}
