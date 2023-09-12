package models

type Account struct {
	PrivateKey string `json:"privateKey"`
	PublicKey  string `json:"publicKey"`
	Address    string `json:"address"`
}

type Address struct {
	Address string `json:"address"`
}

type Transaction struct {
	Amount      float64 `json:"amount"`
	PrivateKey  string  `json:"privateKey"`
	AddressTo   string  `json:"addressTo"`
	AddressFrom string  `json:"addressFrom"`
	GasLimit    uint64  `json:"gasLimit"`
	GasPrice    uint64  `json:"gasPrice"`
}
