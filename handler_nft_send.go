package main

import (
	"encoding/json"
	"net/http"
	"subs/shared/models"
	"subs/utils/msg"

	"github.com/elum-utils/tonsub"
	"github.com/gin-gonic/gin"
)

// NFTSendBody defines the expected JSON structure for NFT send requests.
type NFTSendBody struct {
	models.SendNFT
}

// NFTSendResponse defines the structure of the response for NFT send requests.
type NFTSendResponse struct {
	Result bool `json:"result"` // Indicates if the NFT send request was successful.
}

// HandlerNFTSend handles HTTP POST requests to send an NFT.
func HandlerNFTSend(ctx *gin.Context) {

	defer func(ctx *gin.Context) {
		if r := recover(); r != nil {
			ctx.String(http.StatusBadRequest, "invalid data")
			ctx.Abort()
			return
		}
	}(ctx)

	var body NFTSendBody
	// Bind and validate incoming JSON request against NFTSendBody structure.
	if err := ctx.ShouldBindJSON(&body); err != nil {
		msg.InvalidFields(ctx) // Respond with an error if validation fails.
		return
	}

	// Prepare the data for sending by marshaling it to JSON.
	bytes, err := json.Marshal(models.EncodeData{
		Type: "send_nft",
		Data: tonsub.NFT{
			Address:      body.Address,
			OwnerAddress: body.Owner,
			Message:      body.Message,
		},
	})
	if err != nil {
		msg.InvalidFields(ctx) // Respond with an error if marshalling fails.
		return
	}

	// Attempt to add the marshaled data to an external queue.
	err = o.Add(bytes)
	msg.Send(ctx, NFTSendResponse{
		Result: err == nil, // Respond with success if no error occurred when adding to the queue.
	})
}
