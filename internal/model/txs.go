package model

import "time"

type Txss struct {
	Data      string    `json:"data"`
	Logs      []Logs    `json:"logs"`
	GasWanted string    `json:"gas_wanted"`
	GasUsed   string    `json:"gas_used"`
	Timestamp time.Time `json:"timestamp"`
	Height    string    `json:"height"`
	RawLog    string    `json:"raw_log"`
	Tx        Tx        `json:"tx"`
	Txhash    string    `json:"txhash"`
}

type Attributes struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type Events struct {
	Type       string       `json:"type"`
	Attributes []Attributes `json:"attributes"`
}

type Logs struct {
	Events []Events `json:"events"`
}

type Value1 struct {
	Sender    string `json:"sender"`
	Recipient string `json:"recipient"`
	ID        string `json:"id"`
	DenomID   string `json:"denom_id"`
	Name      string `json:"name"`
	URI       string `json:"uri"`
	Data      string `json:"data"`
}

type Msg struct {
	Type  string `json:"type"`
	Value Value1 `json:"value"`
}

type Amount struct {
	Denom  string `json:"denom"`
	Amount string `json:"amount"`
}

type Fee struct {
	Amount []Amount `json:"amount"`
	Gas    string   `json:"gas"`
}

type Value struct {
	Msg           []Msg         `json:"msg"`
	Fee           Fee           `json:"fee"`
	Signatures    []interface{} `json:"signatures"`
	Memo          string        `json:"memo"`
	TimeoutHeight string        `json:"timeout_height"`
}

type Tx struct {
	Type  string `json:"type"`
	Value Value  `json:"value"`
}
