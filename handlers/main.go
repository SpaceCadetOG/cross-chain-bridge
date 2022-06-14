package handler

import (
	"cross-chain-bridge/models"
	"cross-chain-bridge/modules"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gorilla/mux"
)

type ClientHandler struct {
	*ethclient.Client
}

func (client ClientHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// get parms from url request
	vars := mux.Vars(r)
	module := vars["module"]

	// get Query parms and url request
	address := r.URL.Query().Get("address")
	hash := r.URL.Query().Get("hash")

	// Set The Response Writer
	w.Header().Set("Content-Type", "application/json")

	// Handle each request using the module parameter:
	switch module {
	case "latest-block":
		_block := modules.GetLatestBlock(*client.Client)
		json.NewEncoder(w).Encode(_block)

	case "":
		if hash == " " {
			json.NewEncoder(w).Encode(&models.Error{
				Code:    400,
				Message: "Malformed request",
			})
			return
		}
		txhash := common.HexToHash(hash)
		_tx := modules.GetTxByHash(*client.Client, txhash)
		if _tx != nil {
			json.NewEncoder(w).Encode(_tx)
			return
		}
		json.NewEncoder(w).Encode(&models.Error{
			Code:    400,
			Message: "Tx Not Found",
		})
	case "send-eth":
		decoder := json.NewDecoder(r.Body)
		var t models.TransferETHRequest
		err := decoder.Decode(&t)
		if err != nil {
			fmt.Println(err)
			json.NewEncoder(w).Encode(&models.Error{
				Code:    400,
				Message: "Malformed request",
			})
			return
		}
		_hash, err := modules.TransferETH(*client.Client, t.PrivKey, t.To, t.Amount)
		if err != nil {
			fmt.Println(err)
			json.NewEncoder(w).Encode(&models.Error{
				Code:    500,
				Message: "Internal server error",
			})
			return
		}
		json.NewEncoder(w).Encode(&models.HashResponse{
			Hash: _hash,
		})

	case "get-balance":
		if address == "" {
			json.NewEncoder(w).Encode(&models.Error{
				Code:    400,
				Message: "Malformed request",
			})
			return
		}
		
		balance, err  := modules.GetAddressBalance(*client.Client, address)
		if err != nil {
			fmt.Println(err)
			json.NewEncoder(w).Encode(&models.Error{
				Code:    500,
				Message: "Internal server error",
			})
			return
		}
		json.NewEncoder(w).Encode(&models.BalanceResponse{
			Address: address,
			Balance: balance,
			Symbol: "Ether",
			Units: "Wei",
		})
	}


}


