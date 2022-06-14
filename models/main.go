package models

// Models Data structs = Will be used to through api

// Block
type Block struct {
	BlockNumber      int64         `json:"blockNumber"`
	Timestamp        uint64        `json:"timestamp"`
	Difficulty       uint64        `json:"difficultyr"`
	Hash             string        `json:"Hash"`
	TransactionCount int           `json:"transactionsCount"`
	Transactions     []Transaction `json:"transactions"`
}

// Transaction
type Transaction struct {
	Hash     string `json:"Hash"`
	Value    string `json:"Value"`
	Gas      uint64 `json:"Gas"`
	GasPrice uint64 `json:"gasPrice"`
	Nonce    uint64 `json:"nonce"`
	To       string `json:"to"`
	Pending  bool   `json:"pending"`
}

// TransferETHRequest\
type TransferETHRequest struct {
	PrivKey string `json:"privkey"`
	To      string `json:"to"`
	Amount  int64  `json:"amount"`
}

// HashResponse data structure
type HashResponse struct {
	Hash string `json:"Hash"`
}

// BalanceResponse data structure
type BalanceResponse struct {
	Address string `json:"address"`
	Balance string `json:"balance"`
	Symbol  string `json:"symbol"`
	Units   string `json:"uints"`
}

// Error data structure
type Error struct {
	Code    uint64 `json:"code"`
	Message string `json:"message"`
}
