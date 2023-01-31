package client

// DONTCOVER

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net"
	"os"
	"path/filepath"

	"github.com/FiboChain/fbc/app/crypto/hd"
	ethermint "github.com/FiboChain/fbc/app/types"
	"github.com/FiboChain/fbc/libs/cosmos-sdk/client/flags"
	clientkeys "github.com/FiboChain/fbc/libs/cosmos-sdk/client/keys"
	"github.com/FiboChain/fbc/libs/cosmos-sdk/codec"
	"github.com/FiboChain/fbc/libs/cosmos-sdk/crypto/keys"
	"github.com/FiboChain/fbc/libs/cosmos-sdk/server"
	srvconfig "github.com/FiboChain/fbc/libs/cosmos-sdk/server/config"
	sdk "github.com/FiboChain/fbc/libs/cosmos-sdk/types"
	"github.com/FiboChain/fbc/libs/cosmos-sdk/types/module"
	authexported "github.com/FiboChain/fbc/libs/cosmos-sdk/x/auth/exported"
	authtypes "github.com/FiboChain/fbc/libs/cosmos-sdk/x/auth/types"
	"github.com/FiboChain/fbc/libs/cosmos-sdk/x/crisis"
	genutiltypes "github.com/FiboChain/fbc/libs/cosmos-sdk/x/genutil/types"
	"github.com/FiboChain/fbc/libs/cosmos-sdk/x/mint"
	tmconfig "github.com/FiboChain/fbc/libs/tendermint/config"
	tmcrypto "github.com/FiboChain/fbc/libs/tendermint/crypto"
	tmos "github.com/FiboChain/fbc/libs/tendermint/libs/os"
	tmrand "github.com/FiboChain/fbc/libs/tendermint/libs/rand"
	tmtypes "github.com/FiboChain/fbc/libs/tendermint/types"
	tmtime "github.com/FiboChain/fbc/libs/tendermint/types/time"
	"github.com/FiboChain/fbc/x/common"
	"github.com/FiboChain/fbc/x/genutil"
	"github.com/FiboChain/fbc/x/gov"
	stakingtypes "github.com/FiboChain/fbc/x/staking/types"
	ethcrypto "github.com/ethereum/go-ethereum/crypto"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	flagNodeDirPrefix     = "node-dir-prefix"
	flagNumValidators     = "v"
	flagOutputDir         = "output-dir"
	flagNodeDaemonHome    = "node-daemon-home"
	flagNodeCLIHome       = "node-cli-home"
	flagStartingIPAddress = "starting-ip-address"
	flagCoinDenom         = "coin-denom"
	flagKeyAlgo           = "algo"
	flagIPAddrs           = "ip-addrs"
	flagBaseport          = "base-port"
	flagLocal             = "local"
	flagEqualVotingPower  = "equal-voting-power"
	flagNumRPCs           = "r"
)

const nodeDirPerm = 0755

// mnemonicList contains some hard-coded mnemonic (generated by entropy bytes that bit size is 128).
var mnemonicList = []string{
	"puzzle glide follow cruel say burst deliver wild tragic galaxy lumber offer",
	"palace cube bitter light woman side pave cereal donor bronze twice work",
	"antique onion adult slot sad dizzy sure among cement demise submit scare",
	"lazy cause kite fence gravity regret visa fuel tone clerk motor rent",
}

// TestnetCmd initializes all files for tendermint testnet and application
func TestnetCmd(ctx *server.Context, cdc *codec.Codec,
	mbm module.BasicManager, genAccIterator authtypes.GenesisAccountIterator,
) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "testnet",
		Short: "Initialize files for an FBchain testnet",
		Long: `testnet will create "v" number of directories and populate each with
necessary files (private validator, genesis, config, etc.).

Note, strict routability for addresses is turned off in the config file.`,

		Example: "fbchaind testnet --v 4 --keyring-backend test --output-dir ./output --starting-ip-address 192.168.10.2",
		RunE: func(cmd *cobra.Command, _ []string) error {
			config := ctx.Config

			outputDir, _ := cmd.Flags().GetString(flagOutputDir)
			keyringBackend, _ := cmd.Flags().GetString(flags.FlagKeyringBackend)
			chainID, _ := cmd.Flags().GetString(flags.FlagChainID)
			minGasPrices, _ := cmd.Flags().GetString(server.FlagMinGasPrices)
			nodeDirPrefix, _ := cmd.Flags().GetString(flagNodeDirPrefix)
			nodeDaemonHome, _ := cmd.Flags().GetString(flagNodeDaemonHome)
			nodeCLIHome, _ := cmd.Flags().GetString(flagNodeCLIHome)
			startingIPAddress, _ := cmd.Flags().GetString(flagStartingIPAddress)
			ipAddresses, _ := cmd.Flags().GetStringSlice(flagIPAddrs)
			numValidators, _ := cmd.Flags().GetInt(flagNumValidators)
			numRPCs, _ := cmd.Flags().GetInt(flagNumRPCs)
			coinDenom, _ := cmd.Flags().GetString(flagCoinDenom)
			algo, _ := cmd.Flags().GetString(flagKeyAlgo)
			isLocal := viper.GetBool(flagLocal)
			isEqualVotingPower, _ := cmd.Flags().GetBool(flagEqualVotingPower)
			return InitTestnet(
				cmd, config, cdc, mbm, genAccIterator, outputDir, chainID, coinDenom, minGasPrices,
				nodeDirPrefix, nodeDaemonHome, nodeCLIHome, startingIPAddress, ipAddresses, keyringBackend,
				algo, numValidators, isLocal, numRPCs, isEqualVotingPower)
		},
	}

	cmd.Flags().Int(flagNumValidators, 4, "Number of validators to initialize the testnet with")
	cmd.Flags().StringP(flagOutputDir, "o", "./build", "Directory to store initialization data for the testnet")
	cmd.Flags().String(flagNodeDirPrefix, "node", "Prefix the directory name for each node with (node results in node0, node1, ...)")
	cmd.Flags().String(flagNodeDaemonHome, "fbchaind", "Home directory of the node's daemon configuration")
	cmd.Flags().String(flagNodeCLIHome, "fbchaincli", "Home directory of the node's cli configuration")
	cmd.Flags().String(flagStartingIPAddress, "192.168.0.1", "Starting IP address (192.168.0.1 results in persistent peers list ID0@192.168.0.1:46656, ID1@192.168.0.2:46656, ...)")
	cmd.Flags().StringSlice(flagIPAddrs, []string{}, "List of IP addresses to use (i.e. `192.168.0.1,172.168.0.1` results in persistent peers list ID0@192.168.0.1:46656, ID1@172.168.0.1)")
	cmd.Flags().String(flags.FlagChainID, "", "genesis file chain-id, if left blank will be randomly created")
	cmd.Flags().String(flagCoinDenom, ethermint.NativeToken, "Coin denomination used for staking, governance, mint, crisis and evm parameters")
	cmd.Flags().String(server.FlagMinGasPrices, fmt.Sprintf("0.000006%s", ethermint.NativeToken), "Minimum gas prices to accept for transactions; All fees in a tx must meet this minimum (e.g. 0.01aphoton,0.001stake)")
	cmd.Flags().String(flags.FlagKeyringBackend, flags.DefaultKeyringBackend, "Select keyring's backend (os|file|test)")
	cmd.Flags().String(flagKeyAlgo, string(hd.EthSecp256k1), "Key signing algorithm to generate keys for")
	cmd.Flags().Int(flagBaseport, 26656, "testnet base port")
	cmd.Flags().BoolP(flagLocal, "l", false, "run all nodes on local host")
	cmd.Flags().Int(flagNumRPCs, 0, "Number of RPC nodes to initialize the testnet with")
	cmd.Flags().BoolP(flagEqualVotingPower, "", false, "Create validators with equal voting power")

	return cmd
}

// InitTestnet initializes the testnet configuration
// 1.initializes the "validator" node configuration;
// 2.based the "validator" node, initializes the "rpc" node.
func InitTestnet(
	cmd *cobra.Command,
	config *tmconfig.Config,
	cdc *codec.Codec,
	mbm module.BasicManager,
	genAccIterator authtypes.GenesisAccountIterator,
	outputDir,
	chainID,
	coinDenom,
	minGasPrices,
	nodeDirPrefix,
	nodeDaemonHome,
	nodeCLIHome,
	startingIPAddress string,
	ipAddresses []string,
	keyringBackend,
	algo string,
	numValidators int,
	isLocal bool,
	numRPCs int,
	isEqualVotingPower bool,
) error {

	if chainID == "" {
		chainID = fmt.Sprintf("fbc-%d", tmrand.Int63n(9999999999999)+1)
	}

	if !ethermint.IsValidChainID(chainID) {
		return fmt.Errorf("invalid chain-id: %s", chainID)
	}

	if err := sdk.ValidateDenom(coinDenom); err != nil {
		return err
	}

	numNodes := numValidators + numRPCs
	if len(ipAddresses) != 0 {
		numNodes = len(ipAddresses)
	}

	nodeIDs := make([]string, numNodes)
	valPubKeys := make([]tmcrypto.PubKey, numNodes)

	simappConfig := srvconfig.DefaultConfig()
	simappConfig.MinGasPrices = minGasPrices

	var (
		genAccounts []authexported.GenesisAccount
		genFiles    []string
	)

	inBuf := bufio.NewReader(cmd.InOrStdin())
	// generate private keys, node IDs, and initial transactions
	for i := 0; i < numNodes; i++ {
		nodeDirName := fmt.Sprintf("%s%d", nodeDirPrefix, i)
		nodeDir := filepath.Join(outputDir, nodeDirName, nodeDaemonHome)
		clientDir := filepath.Join(outputDir, nodeDirName, nodeCLIHome)
		gentxsDir := filepath.Join(outputDir, "gentxs")

		// generate private keys, node IDs for all nodes
		config.SetRoot(nodeDir)
		config.RPC.ListenAddress = "tcp://0.0.0.0:26657"

		if err := os.MkdirAll(filepath.Join(nodeDir, "config"), nodeDirPerm); err != nil {
			_ = os.RemoveAll(outputDir)
			return err
		}

		config.Moniker = nodeDirName

		var err error
		nodeIDs[i], valPubKeys[i], err = genutil.InitializeNodeValidatorFilesByIndex(config, i)
		if err != nil {
			_ = os.RemoveAll(outputDir)
			return err
		}

		genFiles = append(genFiles, config.GenesisFile())
		if i >= numValidators {
			// rpc nodes do not need to add initial transactions, key_seeds etc.
			continue
		}

		// validator nodes add initial transactions
		if err := os.MkdirAll(clientDir, nodeDirPerm); err != nil {
			_ = os.RemoveAll(outputDir)
			return err
		}

		var ip string
		port := viper.GetInt(flagBaseport)

		if isLocal {
			ip, err = getIP(0, startingIPAddress)
			port += i * 100
		} else {
			if len(ipAddresses) == 0 {
				ip, err = getIP(i, startingIPAddress)
				if err != nil {
					_ = os.RemoveAll(outputDir)
					return err
				}
			} else {
				ip = ipAddresses[i]
			}
		}

		memo := fmt.Sprintf("%s@%s:%d", nodeIDs[i], ip, port)

		kb, err := keys.NewKeyring(
			sdk.KeyringServiceName(),
			keyringBackend,
			clientDir,
			inBuf,
			hd.EthSecp256k1Options()...,
		)
		if err != nil {
			return err
		}

		cmd.Printf(
			"Password for account '%s' :\n", nodeDirName,
		)

		keyPass := clientkeys.DefaultKeyPass
		mnemonic := ""
		if i < len(mnemonicList) {
			mnemonic = mnemonicList[i]
		}
		addr, secret, err := GenerateSaveCoinKey(kb, nodeDirName, keyPass, true, keys.SigningAlgo(algo), mnemonic)
		if err != nil {
			_ = os.RemoveAll(outputDir)
			return err
		}

		fmt.Printf("nodeDir: %s\nnodeDirName: %s\naddr: %s\nmnenonics: %s\n--------------------------------------\n",
			clientDir, nodeDirName, addr, secret)
		info := map[string]string{"secret": secret}

		cliPrint, err := json.Marshal(info)
		if err != nil {
			return err
		}

		// save private key seed words
		if err := writeFile(fmt.Sprintf("%v.json", "key_seed"), clientDir, cliPrint); err != nil {
			return err
		}

		coins := sdk.NewCoins(
			sdk.NewCoin(coinDenom, sdk.NewDec(9000000)),
		)

		genAccounts = append(genAccounts, ethermint.EthAccount{
			BaseAccount: authtypes.NewBaseAccount(addr, coins, nil, 0, 0),
			CodeHash:    ethcrypto.Keccak256(nil),
		})

		//make and save create validator tx
		sequence := uint64(0)
		msgCreateVal := stakingtypes.NewMsgCreateValidator(
			sdk.ValAddress(addr), valPubKeys[i],
			stakingtypes.NewDescription(nodeDirName, "", "", ""),
			sdk.NewDecCoinFromDec(common.NativeToken, stakingtypes.DefaultMinSelfDelegation),
		)
		if err := makeTxAndWriteFile(msgCreateVal, inBuf, chainID, kb, nodeDirName, sequence, memo, cdc, gentxsDir, outputDir); err != nil {
			return err
		}

		if !isEqualVotingPower {
			//make and save deposit tx
			sequence++
			msgDeposit := stakingtypes.NewMsgDeposit(addr, sdk.NewDecCoinFromDec(common.NativeToken, sdk.NewDec(10000*int64(i+1))))
			if err := makeTxAndWriteFile(msgDeposit, inBuf, chainID, kb, nodeDirName, sequence, "", cdc, gentxsDir, outputDir); err != nil {
				return err
			}

			//make and save add shares tx
			sequence++
			msgAddShares := stakingtypes.NewMsgAddShares(addr, []sdk.ValAddress{sdk.ValAddress(addr)})
			if err := makeTxAndWriteFile(msgAddShares, inBuf, chainID, kb, nodeDirName, sequence, "", cdc, gentxsDir, outputDir); err != nil {
				return err
			}
		}

		srvconfig.WriteConfigFile(filepath.Join(nodeDir, "config/app.toml"), simappConfig)
	}

	if err := initGenFiles(cdc, mbm, chainID, coinDenom, genAccounts, genFiles, numNodes); err != nil {
		return err
	}

	err := collectGenFiles(
		cdc, config, chainID, nodeIDs, valPubKeys, numNodes,
		outputDir, nodeDirPrefix, nodeDaemonHome, genAccIterator,
	)
	if err != nil {
		return err
	}

	cmd.Printf("Successfully initialized %d validator nodes directories, %d rpc nodes directories\n", numValidators, numRPCs)
	return nil
}

func makeTxAndWriteFile(msg sdk.Msg, inBuf *bufio.Reader, chainID string, kb keys.Keybase,
	nodeDirName string, sequence uint64, memo string, cdc *codec.Codec, gentxsDir string,
	outputDir string) error {
	tx := authtypes.NewStdTx([]sdk.Msg{msg}, authtypes.StdFee{}, []authtypes.StdSignature{}, memo)
	txBldr := authtypes.NewTxBuilderFromCLI(inBuf).WithChainID(chainID).WithMemo(memo).WithKeybase(kb).WithSequence(sequence)

	signedTx, err := txBldr.SignStdTx(nodeDirName, clientkeys.DefaultKeyPass, tx, false)
	if err != nil {
		_ = os.RemoveAll(outputDir)
		return err
	}

	txBytes, err := cdc.MarshalJSON(signedTx)
	if err != nil {
		_ = os.RemoveAll(outputDir)
		return err
	}

	// gather gentxs folder
	if err := writeFile(fmt.Sprintf("%v-%d.json", nodeDirName, sequence), gentxsDir, txBytes); err != nil {
		_ = os.RemoveAll(outputDir)
		return err
	}

	return nil
}

func initGenFiles(
	cdc *codec.Codec, mbm module.BasicManager,
	chainID, coinDenom string,
	genAccounts []authexported.GenesisAccount,
	genFiles []string, numValidators int,
) error {

	appGenState := mbm.DefaultGenesis()

	// set the accounts in the genesis state
	var authGenState authtypes.GenesisState
	cdc.MustUnmarshalJSON(appGenState[authtypes.ModuleName], &authGenState)

	authGenState.Accounts = genAccounts
	appGenState[authtypes.ModuleName] = cdc.MustMarshalJSON(authGenState)

	var govGenState gov.GenesisState
	cdc.MustUnmarshalJSON(appGenState[gov.ModuleName], &govGenState)

	govGenState.DepositParams.MinDeposit[0].Denom = coinDenom
	appGenState[gov.ModuleName] = cdc.MustMarshalJSON(govGenState)

	var mintGenState mint.GenesisState
	cdc.MustUnmarshalJSON(appGenState[mint.ModuleName], &mintGenState)

	mintGenState.Params.MintDenom = coinDenom
	appGenState[mint.ModuleName] = cdc.MustMarshalJSON(mintGenState)

	var crisisGenState crisis.GenesisState
	cdc.MustUnmarshalJSON(appGenState[crisis.ModuleName], &crisisGenState)

	crisisGenState.ConstantFee.Denom = coinDenom
	appGenState[crisis.ModuleName] = cdc.MustMarshalJSON(crisisGenState)

	appGenStateJSON, err := codec.MarshalJSONIndent(cdc, appGenState)
	if err != nil {
		return err
	}

	genDoc := tmtypes.GenesisDoc{
		ChainID:    chainID,
		AppState:   appGenStateJSON,
		Validators: nil,
	}

	// generate empty genesis files for each validator and save
	for i := 0; i < numValidators; i++ {
		if err := genDoc.SaveAs(genFiles[i]); err != nil {
			return err
		}
	}
	return nil
}

// GenerateSaveCoinKey returns the address of a public key, along with the secret
// phrase to recover the private key.
func GenerateSaveCoinKey(keybase keys.Keybase, keyName, keyPass string, overwrite bool, algo keys.SigningAlgo, mnemonic string) (sdk.AccAddress, string, error) {
	// ensure no overwrite
	if !overwrite {
		_, err := keybase.Get(keyName)
		if err == nil {
			return sdk.AccAddress([]byte{}), "", fmt.Errorf(
				"key already exists, overwrite is disabled")
		}
	}

	// generate a private key, with recovery phrase
	// If mnemonic is not "", secret is this mnemonic, or secret is random mnemonic.
	info, secret, err := keybase.CreateMnemonic(keyName, keys.English, keyPass, algo, mnemonic)
	if err != nil {
		return sdk.AccAddress([]byte{}), "", err
	}

	return sdk.AccAddress(info.GetPubKey().Address()), secret, nil
}

func collectGenFiles(
	cdc *codec.Codec, config *tmconfig.Config, chainID string,
	nodeIDs []string, valPubKeys []tmcrypto.PubKey,
	numValidators int, outputDir, nodeDirPrefix, nodeDaemonHome string,
	genAccIterator authtypes.GenesisAccountIterator,
) error {

	var appState json.RawMessage
	genTime := tmtime.Now()

	for i := 0; i < numValidators; i++ {
		nodeDirName := fmt.Sprintf("%s%d", nodeDirPrefix, i)
		nodeDir := filepath.Join(outputDir, nodeDirName, nodeDaemonHome)
		gentxsDir := filepath.Join(outputDir, "gentxs")
		config.Moniker = nodeDirName

		config.SetRoot(nodeDir)
		// set node's port
		port := viper.GetInt(flagBaseport)
		p2pPort := port + i*100
		rpcPort := port + i*100 + 1
		config.P2P.ListenAddress = fmt.Sprintf("tcp://0.0.0.0:%d", p2pPort)
		config.RPC.ListenAddress = fmt.Sprintf("tcp://0.0.0.0:%d", rpcPort)

		nodeID, valPubKey := nodeIDs[i], valPubKeys[i]
		initCfg := genutiltypes.NewInitConfig(chainID, gentxsDir, nodeID, nodeID, valPubKey)

		genDoc, err := tmtypes.GenesisDocFromFile(config.GenesisFile())
		if err != nil {
			return err
		}

		nodeAppState, err := genutil.GenAppStateFromConfig(cdc, config, initCfg, *genDoc, genAccIterator)
		if err != nil {
			return err
		}

		if appState == nil {
			// set the canonical application state (they should not differ)
			appState = nodeAppState
		}

		genFile := config.GenesisFile()

		// overwrite each validator's genesis file to have a canonical genesis time
		if err := genutil.ExportGenesisFileWithTime(genFile, chainID, nil, appState, genTime); err != nil {
			return err
		}
	}

	return nil
}

func getIP(i int, startingIPAddr string) (ip string, err error) {
	if len(startingIPAddr) == 0 {
		ip, err = server.ExternalIP()
		if err != nil {
			return "", err
		}
		return ip, nil
	}
	return calculateIP(startingIPAddr, i)
}

func calculateIP(ip string, i int) (string, error) {
	ipv4 := net.ParseIP(ip).To4()
	if ipv4 == nil {
		return "", fmt.Errorf("%v: non ipv4 address", ip)
	}

	for j := 0; j < i; j++ {
		ipv4[3]++
	}

	return ipv4.String(), nil
}

func writeFile(name string, dir string, contents []byte) error {
	writePath := filepath.Join(dir)
	file := filepath.Join(writePath, name)

	err := tmos.EnsureDir(writePath, 0755)
	if err != nil {
		return err
	}

	err = tmos.WriteFile(file, contents, 0644)
	if err != nil {
		return err
	}

	return nil
}
