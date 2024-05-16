package entities

type Provider struct {
	WalletAddress string
	Region        string
	Tokens        []string
	Attributes    Attributes
	IsActive      bool
	Domain        string
}
