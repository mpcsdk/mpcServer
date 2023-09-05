package model

type AnalzyTx struct {
	Address string
	Txs     []*AnalzyTxData
}
type AnalzyTxData struct {
	Target     string
	MethodId   string
	MethodName string
	Sig        string
	Data       string
	Args       map[string]interface{}
}

// //
type SignTx struct {
	ChainId uint64        `json:"chainId,omitempty"`
	Address string        `json:"address,omitempty"`
	Number  uint64        `json:"number,omitempty"`
	Txs     []*SignTxData `json:"txs,omitempty"`
	TxHash  string        `json:"txHash,omitempty"`
}
type SignTxData struct {
	Target string `json:"target,omitempty"`
	Data   string `json:"data,omitempty"`
}
