package models

type Account struct {
	PrivateKey string `json:"privateKey"`
	PublicKey  string `json:"publicKey"`
	Address    string `json:"address"`
}

type Address struct {
	Address string `json:"address"`
}
