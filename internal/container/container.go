package container

import (
	"eth-account-creator-api/core/domains/account"
	"eth-account-creator-api/core/domains/health"
	"eth-account-creator-api/internal/log"
)

type components struct {
	Log *log.Logger
}

type Services struct {
	Account account.ServiceI
	Health  health.ServiceI
}

type Dependency struct {
	Components components
	Services   Services
}

func New() (*Dependency, error) {
	cmp, err := setupComponents()
	if err != nil {
		return nil, err
	}

	accountService, err := account.NewAccountService(
		cmp.Log,
	)
	if err != nil {
		return nil, err
	}

	healthService, err := health.NewService(
		cmp.Log,
	)
	if err != nil {
		return nil, err
	}

	srv := Services{
		Account: accountService,
		Health:  healthService,
	}

	dep := Dependency{
		Components: *cmp,
		Services:   srv,
	}

	return &dep, err
}

func setupComponents() (*components, error) {
	logger, err := log.NewLogger()
	if err != nil {
		return nil, err
	}

	return &components{
		Log: logger,
	}, nil
}
