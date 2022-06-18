*Example CLI Command to transfer eth*

in terminal run => go run main.go
Go to => http://localhost:8080/api/v1/eth/latest-block on browser

in terminal =>
curl -d '{"privKey":"PRIVATEKEY", "to":"TO ADDRESS", "amount":5000000000000000000}' -H "Content-Type: application/json" -X POST http://localhost:8080/api/v1/eth/send-eth



# API CALLS
    CONTRACT ABIS AND ADDRESSES

    BRIDGE LOGIC
        => 
