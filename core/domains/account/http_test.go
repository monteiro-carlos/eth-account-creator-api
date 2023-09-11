package account_test

import (
	"encoding/json"
	"eth-account-creator-api/core/domains/account"
	"eth-account-creator-api/core/domains/account/models"
	"eth-account-creator-api/internal/log"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

var (
	publicKeyTestHttp = "0x044bca2f51047d1c9cd24a054ced4366e61496b872dae72ee6ebfccfaa7e3aa513980c8d146ef57173c2e321e6e0ae5a7720d23a7f6d178dcf4d803730118b7c3a"
	addressTestHttp   = &models.Address{
		Address: "0x6971c9D18A384d11a0060622f9adB8e1C985CCca",
	}
)

func routesSetup() *gin.Engine {
	gin.SetMode(gin.TestMode)
	routes := gin.Default()

	return routes
}

func setupHttp() (*account.Handler, *gin.Engine, error) {
	logger, err := log.NewLogger()
	if err != nil {
		return nil, nil, err
	}

	service, err := account.NewAccountService(logger)
	if err != nil {
		return nil, nil, err
	}

	handler := account.NewHandler(service, logger)

	routes := routesSetup()

	return handler, routes, nil
}

func TestCreateNewAccount(t *testing.T) {
	tests := []struct {
		name     string
		wantCode int
	}{
		{
			name:     "Creating a new Ethereum account",
			wantCode: http.StatusOK,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			hdr, router, err := setupHttp()
			if err != nil {
				assert.FailNow(t, err.Error())
			}

			router.GET("/account", hdr.CreateNewAccount)
			req, _ := http.NewRequest("GET", "/account", nil)

			res := httptest.NewRecorder()
			router.ServeHTTP(res, req)

			if eq := assert.Equal(t, test.wantCode, res.Code); !eq {
				return
			}

			actual := new(models.Account)
			if err := json.Unmarshal(res.Body.Bytes(), &actual); err != nil {
				assert.FailNow(t, err.Error())
			}
			assert.NotNil(t, actual)

			assert.IsType(t, &models.Account{}, actual)
		})
	}
}

func TestGetAccountFromPubKey(t *testing.T) {
	tests := []struct {
		name       string
		wantCode   int
		inputParam string
		resp       *models.Address
	}{
		{
			name:       "Getting Ethereum address from pub key",
			wantCode:   http.StatusOK,
			inputParam: publicKeyTestHttp,
			resp:       addressTestHttp,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			hdr, router, err := setupHttp()
			if err != nil {
				assert.FailNow(t, err.Error())
			}

			router.GET("/account/:publicKey", hdr.GetAccountFromPubKey)
			req, _ := http.NewRequest("GET", "/account/"+test.inputParam, nil)

			res := httptest.NewRecorder()
			router.ServeHTTP(res, req)

			if eq := assert.Equal(t, test.wantCode, res.Code); !eq {
				return
			}

			actual := new(models.Address)
			if err := json.Unmarshal(res.Body.Bytes(), &actual); err != nil {
				assert.FailNow(t, err.Error())
			}
			assert.NotNil(t, actual)

			want := test.resp

			assert.Equal(t, reflect.TypeOf(want), reflect.TypeOf(want))
			assert.Equal(t, test.resp, actual)
		})
	}
}
