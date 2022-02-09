package model

type Nftt struct {
	Collection Collection `json:"collection"`
	Pagination Pagination `json:"pagination"`
}

type Denom struct {
	MintRestricted   bool   `json:"mint_restricted"`
	UpdateRestricted bool   `json:"update_restricted"`
	ID               string `json:"id"`
	Name             string `json:"name"`
	Schema           string `json:"schema"`
	Creator          string `json:"creator"`
	Symbol           string `json:"symbol"`
}

type Nfts struct {
	URI   string `json:"uri"`
	Data  string `json:"data"`
	Owner string `json:"owner"`
	ID    string `json:"id"`
	Name  string `json:"name"`
}

type Collection struct {
	Denom Denom  `json:"denom"`
	Nfts  []Nfts `json:"nfts"`
}

type Pagination struct {
	NextKey interface{} `json:"next_key"`
	Total   string      `json:"total"`
}
