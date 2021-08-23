package account

import "time"

type AccountServerSettings struct {
	url string
}

type AccountAttributes struct {
	AccountClassification   *string  `json:"account_classification,omitempty"`
	AccountMatchingOptOut   *bool    `json:"account_matching_opt_out,omitempty"`
	AccountNumber           string   `json:"account_number,omitempty"`
	AlternativeNames        []string `json:"alternative_names,omitempty"`
	BankID                  string   `json:"bank_id,omitempty"`
	BankIDCode              string   `json:"bank_id_code,omitempty"`
	BaseCurrency            string   `json:"base_currency,omitempty"`
	Bic                     string   `json:"bic,omitempty"`
	Country                 *string  `json:"country,omitempty"`
	Iban                    string   `json:"iban,omitempty"`
	JointAccount            *bool    `json:"joint_account,omitempty"`
	Name                    []string `json:"name,omitempty"`
	SecondaryIdentification string   `json:"secondary_identification,omitempty"`
	Status                  *string  `json:"status,omitempty"`
	Switched                *bool    `json:"switched,omitempty"`
}

type AccountData struct {
	Attributes     AccountAttributes `json:"attributes,omitempty"`
	ID             string            `json:"id,omitempty"`
	OrganisationID string            `json:"organisation_id,omitempty"`
	Type           string            `json:"type,omitempty"`
	Version        *int64            `json:"version,omitempty"`
}

type Link struct {
	Self string `json:"self"`
}

type Error struct {
	ErrorMessage string `json:"error_message"`
}

type AccountCreatedResponse struct {
	Data AccountData `json:"data"`
	Link Link           `json:"links"`
}

type AccountCreated struct {
	Attributes     AccountAttributes `json:"attributes,omitempty"`
	ID             string            `json:"id,omitempty"`
	OrganisationID string            `json:"organisation_id,omitempty"`
	Type           string            `json:"type,omitempty"`
	Version        *int64            `json:"version,omitempty"`
	CreatedOn      time.Time         `json:"created_on"`
	ModifiedOn     time.Time         `json:"modified_on"`
}

type AccountDeletedResponse struct {
	ID string
}

type CreateAccountRequest struct {
	Data AccountData `json:"data"`
}
type FetchAccountResponse struct {
	Data AccountData `json:"data"`
	Link Link        `json:"links"`
}