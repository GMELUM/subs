package models

import "encoding/json"

type DecodeData struct {
	Type string          `json:"type"`
	Data json.RawMessage `json:"data"`
}

type EncodeData struct {
	Type string `json:"type"`
	Data any    `json:"data"`
}

// ReceivedNFT represents the transaction data for receiving an NFT [callback > server]
type ReceivedNFT struct {
	Type       string         `json:"type"`       // Event type
	Address    string         `json:"address"`    // NFT address
	Owner      string         `json:"owner"`      // NFT owner's address
	Collection string         `json:"collection"` // NFT collection address
	Message    string         `json:"message"`    // Transaction comment
	MetaData   map[string]any `json:"meta_data"`  // Additional metadata
	TxHash     string         `json:"tx_hash"`    // Transaction hash
}

// SendNFT represents the transaction data for sending an NFT [server > callback]
type SendNFT struct {
	Address string `json:"address" binding:"required"`  // Address of the NFT, required field.
	Owner   string `json:"owner" binding:"required"`    // Owner of the NFT, required field.
	Message string `json:"message" binding:"omitempty"` // Optional message associated with the NFT send.
	TxHash  string `json:"tx_hash"`                     // Transaction hash
}

// SuccessfulNFT represents the transaction data for a successful NFT sending [callback > server]
type SuccessfulNFT struct {
	Type    string `json:"type"`    // Event type
	Address string `json:"address"` // NFT address
	TxHash  string `json:"tx_hash"` // Transaction hash
}
