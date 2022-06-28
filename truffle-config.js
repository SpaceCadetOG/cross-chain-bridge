require('babel-register')
require('babel-polyfill')
require("dotenv").config();
const HDWalletProvider = require("truffle-hdwallet-provider-privkey");

let _rinkeby="https://eth-rinkeby.alchemyapi.io/v2/bdCqPP7gytsAuNNCxzm5QToEbadyrdP6"
let _mumbai="https://polygon-mumbai.g.alchemy.com/v2/GwzyKOknlRxTsmd_Cza6bnwXwZ8BU-um"
let fuji="https://api.avax-test.network/ext/bc/C/rpc"


module.exports = {
	networks: {
		development: {
			host: "127.0.0.1",
			port: 8545,
			network_id: "*" // Match any network id
		},
		rinkeby: {
			provider: function () {
				return new HDWalletProvider(
					[process.env.PRIVATE_KEY], // Private Key
					_rinkeby // URL to Ethereum Node
				)
			},
			network_id: 4
    },
    
    mumbai: {
			provider: function () {
				return new HDWalletProvider(
					[process.env.PRIVATE_KEY], // Private Key
					_mumbai // URL to Ethereum Node
				)
			},
			network_id: 80001
    },
    
    fuji: {
			provider: function () {
				return new HDWalletProvider(
					[process.env.PRIVATE_KEY], // Private Key
					_mumbai // URL to Ethereum Node
				)
			},
			network_id: 43113
		},
    
		bsc_testnet: {
			provider: function () {
				return new HDWalletProvider(
					[process.env.PRIVATE_KEY], // Private Key
					`https://data-seed-prebsc-1-s1.binance.org:8545` // URL to Binance Node
				)
			},
			network_id: 97
		}
	},

	contracts_directory: './src/contracts',
	contracts_build_directory: './src/abis',

	compilers: {
		solc: {
			version: '0.8.9',
			optimizer: {
				enabled: true,
				runs: 200
			}
		}
	},
}
