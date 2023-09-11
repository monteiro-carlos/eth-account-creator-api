package account_test

import (
	"eth-account-creator-api/core/domains/account"
	"eth-account-creator-api/core/domains/account/models"
	"eth-account-creator-api/internal/log"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	publicKeyTest = "0x044bca2f51047d1c9cd24a054ced4366e61496b872dae72ee6ebfccfaa7e3aa513980c8d146ef57173c2e321e6e0ae5a7720d23a7f6d178dcf4d803730118b7c3a"
	addressTest   = "0x6971c9D18A384d11a0060622f9adB8e1C985CCca"
)

func setup() (*account.Account, error) {
	logger, err := log.NewLogger()
	if err != nil {
		return nil, err
	}

	service, err := account.NewAccountService(logger)
	if err != nil {
		return nil, err
	}

	return service, nil
}

func TestCreateAddress(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "creating Ethereum account",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			srv, err := setup()
			if err != nil {
				t.Error(err)
			}

			account, err := srv.CreateAddress()
			if err != nil {
				t.Error(err)
			}

			assert.IsType(t, models.Account{}, account)
		})
	}
}

func TestFetchAddressFromPubKey(t *testing.T) {
	tests := []struct {
		name  string
		input string
		resp  string
	}{
		{
			name:  "fetching Ethereum address from pub key",
			input: publicKeyTest,
			resp:  addressTest,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			srv, err := setup()
			if err != nil {
				assert.FailNow(t, err.Error())
			}

			address, err := srv.FetchAddressFromPubKey(test.input)
			if err != nil {
				assert.FailNow(t, err.Error())
			}

			assert.Equal(t, test.resp, address.Address)
		})
	}
}
