package account

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/testcontainers/testcontainers-go"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"
)

var settings = AccountServerSettings{url: "http://0.0.0.0:8087"}
var accountClient = NewAccountClient(settings)
var id = uuid.New().String()

func TestMain(m *testing.M) {
	abs, _ := filepath.Abs("../../docker-compose.yml")
	composeFilePaths := []string{abs}
	identifier := strings.ToLower(uuid.New().String())

	var compose = testcontainers.NewLocalDockerCompose(composeFilePaths, identifier)
	compose.WithCommand([]string{"up", "-d", "--force-recreate"}).
		Invoke()
	time.Sleep(5 * time.Second)
	exitVal := m.Run()

	compose.WithCommand([]string{"down"}).
		Invoke()
	os.Exit(exitVal)
}

func Test_Should_Create_Account(t *testing.T) {
	var country = "GB"
	names := []string{"Sam", "Holder"}

	accountData := buildAccountData(names, country, id)
	actualAccount, _ := accountClient.CreateAccount(accountData)
	version := int64(0)

	expectedCreatedAccount := AccountCreatedResponse{Data: AccountData{
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

	actualAccount, _ := accountClient.FetchAccount(id)
	version := int64(0)
	expectedFetchAccountQuery := FetchAccountResponse{Data: AccountData{
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

func Test_Should_Not_Fetch_Account_When_Does_Not_Exist(t *testing.T) {
	randomId := uuid.New().String()
	_, e := accountClient.FetchAccount(randomId)
	assert.Equal(t, e.ErrorMessage, fmt.Sprintf("record %s does not exist", randomId))
}

func Test_Should_Not_Delete_Account_When_Does_Not_Exist(t *testing.T) {
	randomId := uuid.New().String()
	accountClient.DeleteAccount(randomId)
}

func Test_Should_Delete_Account(t *testing.T) {
	accountClient.DeleteAccount(id)
	_, e := accountClient.FetchAccount(id)
	assert.Equal(t, e.ErrorMessage, fmt.Sprintf("record %s does not exist", id))
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
