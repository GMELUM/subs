package main

import (
	"encoding/json"
	"errors"
	"subs/config"
	"subs/shared/models"
	"subs/utils"
	"time"

	"github.com/elum-utils/queue"
	"github.com/elum-utils/tonsub"
)

// HandlerQueueInner processes items retrieved from the queue.
// It handles different types of queue data by delegating to specific handlers.
func HandlerQueueInner(item queue.Item, delay func(t time.Duration)) {
	// Deserialize the data from the queue item.
	var data = models.DecodeData{}
	err := json.Unmarshal(item.Data, &data)
	if err != nil {
		return // Exit if data cannot be deserialized.
	}

	// Determine the type of data and handle accordingly.
	switch data.Type {
	case "success_send_nft":
		err = successSendNFT(data)
		if err != nil {
			delay(time.Second * 2)
			return // Exit if handling "success_send_nft" fails.
		}
	case "received_nft":
		err = receivedNFT(data)
		if err != nil {
			delay(time.Second * 2)
			return // Exit if handling "received_nft" fails.
		}
	}

	// Delete the item from the queue after successful handling.
	i.Delete(item.ID)
}

// successSendNFT handles the "success_send_nft" queue data type.
// It sends a callback for a successful NFT send operation.
func successSendNFT(data models.DecodeData) error {
	nft := tonsub.NFT{}
	err := json.Unmarshal(data.Data, &nft)
	if err != nil {
		return err // Return error if deserialization fails.
	}

	// Execute callback REST request with the success NFT details.
	result, err := utils.CallbackREST(config.CallbackURL, models.SuccessfulNFT{
		Type:    "success_send_nft",
		Address: nft.Address,
		TxHash:  nft.TxHash,
	})

	if err != nil {
		return err // Return error if callback fails.
	}

	if result {
		return nil // Return nil if callback indicates success.
	}

	return errors.New("unsuccess") // Return a generic error for unsuccessful callback.
}

// receivedNFT handles the "received_nft" queue data type.
// It retrieves metadata and sends a callback for a received NFT.
func receivedNFT(data models.DecodeData) error {
	nft := tonsub.NFT{}
	err := json.Unmarshal(data.Data, &nft)
	if err != nil {
		return err // Return error if deserialization fails.
	}

	// Retrieve additional JSON metadata from the URL.
	nftData, err := utils.FetchNFTMetaData(nft.Meta)
	if err != nil {
		return err // Return error if metadata retrieval fails.
	}

	// Execute callback REST request with the received NFT details including metadata.
	result, err := utils.CallbackREST(config.CallbackURL, models.ReceivedNFT{
		Type:       data.Type,
		Address:    nft.Address,
		Owner:      nft.OwnerAddress,
		Collection: nft.CollectionAddress,
		Message:    nft.Message,
		MetaData:   nftData,
		TxHash:     nft.TxHash,
	})
	if err != nil {
		return err // Return error if callback fails.
	}

	if result {
		return nil // Return nil if callback indicates success.
	}

	return errors.New("unsuccess") // Return a generic error for unsuccessful callback.
}
