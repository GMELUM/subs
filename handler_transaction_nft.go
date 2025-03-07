package main

import (
	"encoding/json"
	"subs/shared/models"

	"github.com/elum-utils/tonsub"
)

// HandlerTransactionNFT processes a received NFT transaction.
// It checks if the NFT collection is whitelisted, prepares the data, and adds it to a queue.
func HandlerTransactionNFT(t *tonsub.RootNFT) {
	// Check if the collection is in the whitelist.
	if _, exists := collections[t.Body.CollectionAddress]; !exists {
		return // Exit if the collection is not whitelisted.
	}

	// Prepare NFT data to add to the queue by marshaling it to JSON.
	bytes, err := json.Marshal(models.EncodeData{
		Type: "received_nft", // Set the type of data.
		Data: t.Body,         // Include the NFT body data.
	})
	if err != nil {
		// Exit if there is an error during marshaling.
		return
	}

	// Add the marshaled data to the queue.
	err = i.Add(bytes)
	if err != nil {
		// Handle error if adding to the queue fails.
		return
	}
}
