package main

import (
	"net/http"
	"subs/utils"
	"subs/utils/msg"

	"github.com/gin-gonic/gin"
	"github.com/xssnick/tonutils-go/address"
	"github.com/xssnick/tonutils-go/ton/nft"
)

func HandlerNFTInfo(ctx *gin.Context) {

	defer func(ctx *gin.Context) {
		if r := recover(); r != nil {
			ctx.String(http.StatusBadRequest, "invalid address")
			ctx.Abort()
			return
		}
	}(ctx)

	addr := ctx.Params.ByName("addr")

	if len(addr) < 32 {
		ctx.String(http.StatusNotFound, "404 Not Found")
		ctx.Abort()
		return
	}

	item := nft.NewItemClient(s.Api, address.MustParseAddr(addr))

	// Retrieve NFT data, which contains various information related to the NFT.
	nftData, err := item.GetNFTData(s.Context)
	if err != nil {
		ctx.String(http.StatusBadRequest, err.Error())
		ctx.Abort()
		return
	}

	if nftData.CollectionAddress.Type() == address.NoneAddress {
		ctx.String(http.StatusBadRequest, "invalid address")
		ctx.Abort()
		return
	}

	// get info about our nft's collection
	collection := nft.NewCollectionClient(s.Api, nftData.CollectionAddress)

	// get full nft's content url using collection method that will merge base url with nft's data
	nftContent, err := collection.GetNFTContent(s.Context, nftData.Index, nftData.Content)
	if err != nil {
		ctx.String(http.StatusBadRequest, err.Error())
		ctx.Abort()
		return
	}

	off, ok := nftContent.(*nft.ContentOffchain)
	if !ok {
		ctx.String(http.StatusBadRequest, "invalid content in offchain")
		ctx.Abort()
		return
	}

	// Retrieve additional JSON metadata from the URL.
	metaData, err := utils.FetchNFTMetaData(off.URI)
	if err != nil {
		ctx.String(http.StatusBadRequest, err.Error())
		ctx.Abort()
		return
	}

	msg.Send(ctx, metaData)

}
