package main

import (
	"encoding/json"
	"fmt"
	"time"
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

type TradeContract struct {
	contractapi.Contract
}

// Trade Structure
type Trade struct {
	TradeID   string  `json:"trade_id"`
	Symbol    string  `json:"symbol"`
	Quantity  int     `json:"quantity"`
	Price     float64 `json:"price"`
	Timestamp string  `json:"timestamp"`
	Status    string  `json:"status"`
}

// Payment Structure
type Payment struct {
	PaymentID string  `json:"payment_id"`
	Sender    string  `json:"sender"`
	Receiver  string  `json:"receiver"`
	Amount    float64 `json:"amount"`
	Timestamp string  `json:"timestamp"`
	Status    string  `json:"status"`
}

// Settlement Structure
type Settlement struct {
	SettlementID string  `json:"settlement_id"`
	TradeID      string  `json:"trade_id"`
	PaymentID    string  `json:"payment_id"`
	Amount       float64 `json:"amount"`
	Timestamp    string  `json:"timestamp"`
	Status       string  `json:"status"`
}

// Initialize Ledger
func (t *TradeContract) InitLedger(ctx contractapi.TransactionContextInterface) error {
	trades := []Trade{
		{TradeID: "T001", Symbol: "AAPL", Quantity: 100, Price: 150.0, Timestamp: time.Now().String(), Status: "OPEN"},
		{TradeID: "T002", Symbol: "GOOGL", Quantity: 50, Price: 2800.0, Timestamp: time.Now().String(), Status: "CLOSED"},
	}

	for _, trade := range trades {
		tradeJSON, err := json.Marshal(trade)
		if err != nil {
			return err
		}
		err = ctx.GetStub().PutState(trade.TradeID, tradeJSON)
		if err != nil {
			return fmt.Errorf("failed to put trade: %s", trade.TradeID)
		}
	}
	return nil
}

// Create Trade
func (t *TradeContract) CreateTrade(ctx contractapi.TransactionContextInterface, tradeID, symbol string, quantity int, price float64) error {
	trade := Trade{
		TradeID:   tradeID,
		Symbol:    symbol,
		Quantity:  quantity,
		Price:     price,
		Timestamp: time.Now().String(),
		Status:    "OPEN",
	}
	tradeJSON, err := json.Marshal(trade)
	if err != nil {
		return err
	}
	return ctx.GetStub().PutState(tradeID, tradeJSON)
}

// Record Payment
func (t *TradeContract) CreatePayment(ctx contractapi.TransactionContextInterface, paymentID, sender, receiver string, amount float64) error {
	payment := Payment{
		PaymentID: paymentID,
		Sender:    sender,
		Receiver:  receiver,
		Amount:    amount,
		Timestamp: time.Now().String(),
		Status:    "PENDING",
	}
	paymentJSON, err := json.Marshal(payment)
	if err != nil {
		return err
	}
	return ctx.GetStub().PutState(paymentID, paymentJSON)
}

// Settlement Process
func (t *TradeContract) CreateSettlement(ctx contractapi.TransactionContextInterface, settlementID, tradeID, paymentID string, amount float64) error {
	settlement := Settlement{
		SettlementID: settlementID,
		TradeID:      tradeID,
		PaymentID:    paymentID,
		Amount:       amount,
		Timestamp:    time.Now().String(),
		Status:       "COMPLETED",
	}
	settlementJSON, err := json.Marshal(settlement)
	if err != nil {
		return err
	}
	return ctx.GetStub().PutState(settlementID, settlementJSON)
}

func main() {
	chaincode, err := contractapi.NewChaincode(new(TradeContract))
	if err != nil {
		fmt.Printf("Error creating trade chaincode: %s", err.Error())
		return
	}
	if err := chaincode.Start(); err != nil {
		fmt.Printf("Error starting trade chaincode: %s", err.Error())
	}
}
