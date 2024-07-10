package service

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	tronClient "github.com/fbsobreira/gotron-sdk/pkg/client"
	"google.golang.org/grpc"
	"log"
	"math/big"
	"strings"
)

// https://data-seed-prebsc-1-s1.binance.org:8545/
const (
	//infuraURL     = "https://mainnet.infura.io/v3/fa99da8fb08c4b94b7e9a29f6d7f7c09" // Replace with your Infura project URL
	//infuraURL     = "https://rpchttp.kadsea.org"
	erc20ABI = `[{"constant":true,"inputs":[{"name":"_owner","type":"address"}],"name":"balanceOf","outputs":[{"name":"balance","type":"uint256"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":true,"inputs":[],"name":"decimals","outputs":[{"name":"","type":"uint8"}],"payable":false,"stateMutability":"view","type":"function"}]`
	//tokenAddress  = "0x8Bb00438D421ecf26d802277A8f699D94771E092" // Replace with your ERC20 token contract address
	//walletAddress = "0x2b83877aCE845279f59919aeb912946C8b5Abe26" // Replace with the wallet address you want to check

)

type RpcUtil struct {
}

func NewRpcUtil() *RpcUtil {
	return &RpcUtil{}
}

func (r *RpcUtil) GetEvmBalance(rpc, tokenAddress, account string) (string, error) {
	// Connect to the Ethereum network
	client, err := ethclient.Dial(rpc)
	if err != nil {
		//log.Fatalf("Failed to connect to the Ethereum client: %v", err)
		return "", err
	}

	// Special address for ETH
	ethAddress := "0xEeeeeEeeeEeEeeEeEeEeeEEEeeeeEeeeeeeeEEeE"

	var balance *big.Float

	if tokenAddress == ethAddress {
		// Get ETH balance
		balance, err = getEthBalance(client, account)
		if err != nil {
			//log.Fatalf("Failed to get ETH balance: %v", err)
			return "", err
		}
	} else {
		// Get ERC-20 token balance
		balance, err = getTokenBalance(client, tokenAddress, account)
		if err != nil {
			//log.Fatalf("Failed to get token balance: %v", err)
			return "", err
		}
	}
	balanceStr := balance.Text('f', 18)

	return balanceStr, nil

}

func getEthBalance(client *ethclient.Client, walletAddress string) (*big.Float, error) {
	address := common.HexToAddress(walletAddress)
	balance, err := client.BalanceAt(context.Background(), address, nil)
	if err != nil {
		return nil, err
	}
	return new(big.Float).Quo(new(big.Float).SetInt(balance), big.NewFloat(1e18)), nil
}

func getTokenBalance(client *ethclient.Client, tokenAddress, walletAddress string) (*big.Float, error) {
	// Parse the contract ABI
	parsedABI, err := abi.JSON(strings.NewReader(erc20ABI))
	if err != nil {
		return nil, fmt.Errorf("failed to parse contract ABI: %w", err)
	}
	fmt.Println("Parsed contract ABI")

	// Get the token decimals
	contractAddress := common.HexToAddress(tokenAddress)
	data, err := parsedABI.Pack("decimals")
	if err != nil {
		return nil, fmt.Errorf("failed to pack method call for decimals: %w", err)
	}
	callMsg := ethereum.CallMsg{
		To:   &contractAddress,
		Data: data,
	}
	decimalsResult, err := client.CallContract(context.Background(), callMsg, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to call contract for decimals: %w", err)
	}

	var decimals uint8
	err = parsedABI.UnpackIntoInterface(&decimals, "decimals", decimalsResult)
	if err != nil {
		return nil, fmt.Errorf("failed to unpack decimals result: %w", err)
	}

	// Get the token balance
	ownerAddress := common.HexToAddress(walletAddress)
	data, err = parsedABI.Pack("balanceOf", ownerAddress)
	if err != nil {
		return nil, fmt.Errorf("failed to pack method call for balanceOf: %w", err)
	}
	callMsg = ethereum.CallMsg{
		To:   &contractAddress,
		Data: data,
	}
	result, err := client.CallContract(context.Background(), callMsg, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to call contract for balanceOf: %w", err)
	}

	var tokenBalance = new(big.Int)
	err = parsedABI.UnpackIntoInterface(&tokenBalance, "balanceOf", result)
	if err != nil {
		return nil, fmt.Errorf("failed to unpack balanceOf result: %w", err)
	}

	// Convert the token balance to ETH units
	ethBalance := new(big.Float).Quo(new(big.Float).SetInt(tokenBalance), big.NewFloat(float64(1e18)))

	return ethBalance, nil
}

func (r *RpcUtil) GetTronBalance(rpc, tokenAddress, account string) (string, error) {
	client := tronClient.NewGrpcClient(rpc)

	err := client.Start(grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect to Tron client: %v", err)
	}
	defer client.Stop()

	if tokenAddress == "0xEeeeeEeeeEeEeeEeEeEeeEEEeeeeEeeeeeeeEEeE" {
		accountInfo, err := client.GetAccount(account)
		if err != nil {
			log.Fatalf("Failed to get account info: %v", err)
		}

		// TRX 精度处理
		trxBalance := new(big.Float).Quo(big.NewFloat(float64(accountInfo.Balance)), big.NewFloat(1e6))
		fmt.Printf("TRX Balance: %s\n", trxBalance.Text('f', 6))

		return trxBalance.Text('f', 6), nil
	}

	// 获取 TRX 余额

	trx20Balance, err := client.TRC20ContractBalance(account, tokenAddress)
	if err != nil {
		log.Fatalf("Failed to get trx20Balance info: %v", err)
	}
	// TRX20 精度处理
	trx20BalanceFloat := new(big.Float).Quo(new(big.Float).SetInt(trx20Balance), big.NewFloat(1e18))
	fmt.Printf("TRX20 Balance: %s\n", trx20BalanceFloat.Text('f', 18))
	return trx20BalanceFloat.Text('f', 18), nil
}
