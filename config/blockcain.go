package config

import "subs/utils/env"

var (
    // BlockchainWallet holds the blockchain wallet address.
    // This value is retrieved from the environment variable "BLOCKCHAIN_WALLET".
    // If the environment variable is not set, it defaults to an empty string.
    BlockchainWallet = env.GetEnvString("BLOCKCHAIN_WALLET", "")

    // BlockchainNetwork specifies the URL to the blockchain network's configuration.
    // It is retrieved from the environment variable "BLOCKCHAIN_NETWORK".
    // If not set, it defaults to "https://ton.org/global.config.json".
    BlockchainNetwork = env.GetEnvString("BLOCKCHAIN_NETWORK", "https://ton.org/global.config.json")

    // WalletWords holds an array of seed words used for wallet recovery or generation.
    // These words are retrieved from the environment variable "BLOCKCHAIN_WORDS",
    // split by commas. If the environment variable is not set, it defaults to an empty slice.
    BlockchainWords = env.GetEnvArrayString("BLOCKCHAIN_WORDS", ",", []string{})
)