package main

import (
	"fmt"
	"log"
	"subs/config"
	"subs/shared/middleware"
	"time"

	"github.com/elum-utils/queue"
	"github.com/elum-utils/tonsub"
	"github.com/elum-utils/wallet"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

var (
	err error

	s           *tonsub.Sub
	o           *queue.Queue
	i           *queue.Queue
	w           *wallet.Wallet
	collections = map[string]string{
		"EQBG-g6ahkAUGWpefWbx-D_9sQ8oWbvy6puuq78U2c4NUDFS": "Plush Pepes",
		"EQD9ikZq6xPgKjzmdBG0G0S80RvUJjbwgHrPZXDKc_wsE84w": "Durov Caps",
		"EQBI07PXew94YQz7GwN72nPNGF6htSTOJkuU4Kx_bjTZv32U": "Swiss Watches",
		"EQCrGA9slCoksgD-NyRDjtHySKN0Ts8k6hdueJkUkZZdD4_K": "Signet Rings",
		"EQDL7HMbca0FufrjHFcRoiLkEiOXkXoO_vH2gVUN8JNp4khK": "Ion Gems",
		"EQATuUGdvrjLvTWE5ppVFOVCqU2dlCLUnKTsu0n1JYm9la10": "Scared Cats",
		"EQACcQpR2fmdeENWdE2YGQWHVxSTyA8Zq4_k7rk_IaxCRXNe": "Vintage Cigars",
		"EQA4i58iuS9DUYRtUZ97sZo5mnkbiYUBpWXQOe3dEUCcP1W8": "Precious Peaches",
		"EQDJsN9OJBhKGZoWZWtkEpzkCfIu16Z9UzTWbYjeLpuHdT5f": "Perfume Bottles",
		"EQDIReleOkTxCD4g_XEm8xj0LYNg6-zMsTGAAwCA-vEbkGBu": "Astral Shards",
		"EQC6zjid8vJNEWqcXk10XjsdDLRKbcPZzbHusuEW6FokOWIm": "Lol Pops",
		"EQAtFU9GrGfix4UG9DOivN58QxvgBJUaAZ_pdZBZCmbhKo4P": "Magic Potions",
		"EQDumy3bnZYzV4bSWMSSZkmXqx50XuH5d9RlX_yEi2FNlivk": "Eternal Roses",
		"EQD6mH9bwbn6S3M_tCRWOvqAIW8M34kRwbI01niGLRPeDPsl": "Spy Agarics",
		"EQCZxxFMS-y1hcGADL6EPB7usNstQqD9u-yBaYpXVVMr75NF": "Mini Oscars",
		"EQDTro-ogJbS7o-OBD6bt2NysPt7SnGm5zfuRXGB1nE_rbGa": "Kissed Frogs",
		"EQCt2C3yCRNX267B3l6h1QsU6agm4ZgTAb7NpVGiFKlBXOAA": "Genie Lamps",
		"EQCNsmpHqRSY_Dxnyh6P0MMO7zcABf8sVvG0wr245pBzO3B3": "Voodoo Dolls",
		"EQBzZLNIr4lie0pTfrbRsANJOtFYwY5gmngRfs84Ras5-aVN": "Eternal Candles",
		"EQBw2tO5UaJ4c_YXt3I8y5KD0k37staZxedV2O5HmryiK0dN": "Sharp Tongues",
		"EQCyAMkb6bNyNlKPH0tJbubk1VVjASqyq9sZwkJ8AbxMkxxU": "Trapped Hearts",
		"EQAtgbhSHOqTxjuRLLOAab6T66FPcQWTNd_DT3VgCG1-tHJw": "Crystal Balls",
		"EQAwzubeoJwnqmmBuTPpnUSurRzWPB8ERzcfzx55Z2YjE0jx": "Jelly Bunnies",
		"EQBaOL8mH5YywkXjkps65X1OLPNH7pns4YcfLmaVpFaoNKZn": "Skull Flowers",
		"EQDQ6DjRabTYSAxf2xrZsnsXtqcIm1bj9dF5x_h8lNjWPmH4": "Evil Eyes",
		"EQBAXR68f1UgRhToFR_bXY1zPJy5O6sm2St0CRTo92BTxGiH": "Star Notepads",
		"EQC2lsUy1SKxJEJBwj5ZCfVnLPvAqDqy5c26Xg8xS_pDTXGk": "Lunar Snakes",
		"EQBD8aBKC4NsnYMqtkCfPQk2EVnieynJQp1UgZVyx1VmR5Ml": "Witch Hats",
		"EQCefrjhCD2_7HRIr2lmwt9ZaqeG_tdseBvADC66833kBS3y": "Homemade Cakes",
		"EQBMcfMAZlMUr1W3X8kdEw3fJMUAaWH4-XcmE5R5RfFIY0E2": "Desk Calendars",
		"EQB4x3sT1DVdODzay3H-4VJIdOooS5-kTgyKcYMZWogPOsiq": "Berry Boxes",
		"EQCa1I09fE9UoTV6awM6QC9-fkv51hoii24w1tJoFfigG_ax": "Party Sparklers",
		"EQBG2o0lp-6Oy86NGEJm717BeTDAw_F5ELkgaX2l9UsfavWE": "Flying Brooms",
		"EQBCe75G0AhjqC64B7H_BHP0wgfONX_x98rszmsEwndDVAjG": "Ginger Cookies",
		"EQA8DCWyCWyywgOKYORerRoSVevWrUQ_FjKQgNihxY1227x7": "Spiced Wines",
		"EQC65Yy6N04vHoeCJ0yo4qll5eu-ZaWbS5nxsdilxymhmSus": "Hypno Lollipops",
		"EQB8zLzEOFQK3qTyMYgPD8BuzmNwblnouqaB80PW-s2E7nct": "Hanging Stars",
		"EQBdlKhLzezYFMCWSWTCnhpKC1uyczaBUOj3EtPjcatUsTrC": "Mad Pumpkins",
		"EQBT9PbZBR6FGcZBSnwgo-DLpc0r7_X_8dlhG5UA6v9l9uJM": "Cookie Hearts",
		"EQAaTIR7oJyowDiumYLVN0oe61kGE3I6EPEn7WgHPGuWAeCy": "Santa Hats",
		"EQCE80Aln8YfldnQLwWMvOfloLGgmPY0eGDJz9ufG3gRui3D": "Loot Bags",
		"EQB6AtBPOuTtQml8oSA7X8ZqJ5QmcOYYqoz92sQYXGUQrxyB": "Hex Pots",
		"EQCBK_JBASAA5XVz1D17Pn--kQaMWm0b9wReVtsEdRO4Tgy9": "Jester Hats",
		"EQDqHwSzU4I_U44vSM9EDP4HGGKWy9yWjbzkpCa3K8iMBEVD": "Love Candles",
		"EQCehrkZtKDtVe0qyvBAsrHx3hW-hroQyDrS_MZOOVYth2DG": "Jingle Bells",
		"EQA3-i1IUFjWyDhaIoCGdYUB4nt2IYaT3T-95CHPrSvV3AfX": "Bunny Muffins",
		"EQAIM-5QzZGXYTSZR1RGeT2g9rNpYmNPQ09_HtvaInHaTyPX": "Snow Mittens",
		"EQD7yDu2WCgd9Uzx1dF_DQkWK7IZJJ4Mp9M9g1rGUUiQE43m": "Love Potions",
		"EQCwEFfUbbR-22fn3VgxUpBil7bwBQqEHm7wgQYbWY9c08YJ": "B-Day Candles",
		"EQCWh1lPltyTwCWxCXm4umL5tPZoXR8kTIcT-pd0JqoadLHo": "Diamond Rings",
		"EQC1gud6QO8NdJjVrqr7qFBMO0oQsktkvzhmIRoMKo8vxiyL": "Toy Bears",
		"EQADvJxMxCHA7fRlYjoceBORf7RwKs0rzjVaKepQACMnZzG7": "Top Hats",
		"EQBcNxMCTyEHkcQ5cK3fO_3Ebjf6JcA5JJ_OJV4npDN-604P": "Sleigh Bells",
		"EQCDBbQYbv3n91TwywBRD9YrJNuNVmbD3Sprpq6hWIDHVu4p": "Sakura Flowers",
		"EQDr4xn5_GoCzDxhGJMek7fv3nm6W7bhRvlDSBjcNZul52tZ": "Record Players",
	}
)

func init() {

	// Load environment variables from a .env file if it exists, skip if not found
	_ = godotenv.Load()

	o, err = queue.New(queue.Config{
		LocalFile: fmt.Sprintf("%v%v", config.Volume, "_events"),
	})
	if err != nil {
		panic(err.Error())
	}

	i, err = queue.New(queue.Config{
		LocalFile: fmt.Sprintf("%v%v", config.Volume, "_success"),
	})
	if err != nil {
		panic(err.Error())
	}

	w, err = wallet.New(config.BlockchainWords, config.BlockchainNetwork)
	if err != nil {
		panic(err.Error())
	}

	s, err = tonsub.New(config.BlockchainWallet, config.BlockchainNetwork)
	if err != nil {
		panic(err.Error())
	}

	o.Listener(HandlerQueueOuter)
	i.Listener(HandlerQueueInner)
	s.OnNFT(HandlerTransactionNFT)

}

func main() {

	// Use defer and recover to handle panics gracefully.
	// Defer ensures the function is called at the end of main, recovering from any panic that might occur during runtime.
	defer func() {
		if r := recover(); r != nil {
			log.Println("Recovered from panic:", r) // Log the panic information.
		}
	}()

	gin.SetMode(gin.ReleaseMode)

	// Create a new Gin engine instance with default middleware: logger and recovery.
	engine := gin.New()

	// Configure CORS (Cross-Origin Resource Sharing) to manage requests from different domains.
	engine.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},                                       // Allow requests from any origin. In production, it's better to specify allowed origins.
		AllowMethods:     []string{"GET", "POST"},                             // Allow only GET and POST requests to come through.
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"}, // Specify which headers are allowed in requests.
		ExposeHeaders:    []string{"Content-Length"},                          // Headers that can be exposed to the client.
		AllowCredentials: false,                                               // Disable credentials support for security.
		MaxAge:           12 * time.Hour,                                      // Set preflight request cache duration.
	}))

	engine.POST("nft.send", middleware.Secret, HandlerNFTSend)

	// Attempt to run the server on the specified host and port.
	// fmt.Sprintf is used to create a formatted string for the address.
	if err := engine.Run(fmt.Sprintf("%v:%v", config.Host, config.Port)); err != nil {
		log.Fatalf("Failed to run server: %v", err) // Log any error that occurs while starting the server and exits the application.
	}

}
