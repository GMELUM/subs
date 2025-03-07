package main

import (
	"encoding/json"
	"subs/config"
	"subs/shared/models"
	"time"

	"github.com/elum-utils/queue"
	"github.com/elum-utils/tonsub"
	"github.com/elum-utils/wallet"
)

func HandlerQueueOuter(item queue.Item, delay func(t time.Duration)) {

	var data = models.DecodeData{}
	err := json.Unmarshal(item.Data, &data)
	if err != nil {
		return
	}

	// Determine the type of data and handle accordingly.
	switch data.Type {
	case "send_nft":
		err = SendNFT(data)
		if err != nil {
			delay(time.Second * 2)
			return // Exit if handling "success_send_nft" fails.
		}
	}

	// Delete the item from the queue after successful handling.
	o.Delete(item.ID)

}

func SendNFT(data models.DecodeData) error {

	nft := tonsub.NFT{}
	err := json.Unmarshal(data.Data, &nft)
	if err != nil {
		return err
	}

	txHash, err := w.TransferNFT(config.BlockchainWallet, wallet.TransactionNFT{
		AddressNFT:    nft.Address,
		AddressTarget: nft.OwnerAddress,
		Message:       nft.Message,
	})
	if err != nil {
		return err
	}

	bytes, err := json.Marshal(models.EncodeData{
		Type: "success_send_nft",
		Data: tonsub.NFT{
			Address:      nft.Address,
			OwnerAddress: nft.OwnerAddress,
			Message:      nft.Message,
			TxHash:       txHash,
		},
	})
	if err != nil {
		return err
	}

	err = i.Add(bytes)
	if err != nil {
		return err
	}

	return nil

}
