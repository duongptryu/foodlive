package eventsmartcontract

import (
	"context"
	"fmt"
	"foodlive/component"
	"foodlive/config"
	flabi "foodlive/contract/foodlive"
	"foodlive/modules/historypayment/historypaymentstore"
	"foodlive/modules/order/orderbiz"
	"foodlive/modules/order/ordermodel"
	"foodlive/modules/order/orderstore"
	"foodlive/modules/ordertracking/ordertrackingstore"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	log "github.com/sirupsen/logrus"
	"math"
	"math/big"
	"strings"
)

type EventWatcher struct {
	client          *ethclient.Client
	ContractAbi     *abi.ABI
	ContractAddress common.Address
}

func NewEventWatcher(config *config.AppConfig) *EventWatcher {
	client, err := ethclient.Dial("wss://rinkeby.infura.io/ws/v3/00c89dc659554a84854d037a57011004")
	if err != nil {
		log.Fatal(err)
	}
	contractAddress := common.HexToAddress("0x16431A94153CDB004513782c1cdC38172EFE5B86")

	contractAbi, err := abi.JSON(strings.NewReader(flabi.FoodliveMetaData.ABI))
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Connect to smart contract log successfully")
	return &EventWatcher{
		client:          client,
		ContractAbi:     &contractAbi,
		ContractAddress: contractAddress,
	}
}

func (e *EventWatcher) Watch(appCtx component.AppContext) {
	query := ethereum.FilterQuery{
		Addresses: []common.Address{e.ContractAddress},
	}
	logs := make(chan types.Log)
	sub, err := e.client.SubscribeFilterLogs(context.Background(), query, logs)
	if err != nil {
		log.Fatal(err)
	}

	for {
		select {
		case err := <-sub.Err():
			log.Fatal(err)
		case vLog := <-logs:
			e.confirmPaymentCrypto(appCtx, &vLog)
		}
	}
}

type PaymentOrderRaw struct {
	Arg0 *big.Int
	Arg1 *big.Int
}

func (e *EventWatcher) confirmPaymentCrypto(appCtx component.AppContext, vlog *types.Log) {
	var result PaymentOrderRaw
	err := e.ContractAbi.UnpackIntoInterface(&result, "PaymentOrder", vlog.Data)
	if err != nil {
		log.Error(err)
		return
	}

	orderPaymentEvent := ordermodel.PaymentOrderEvent{
		OrderId:     int(result.Arg0.Int64()),
		Value:       fmt.Sprintf("%.18f", float64(result.Arg1.Int64())/math.Pow10(18)),
		Hash:        vlog.TxHash.String(),
		BlockNumber: vlog.BlockNumber,
	}

	orderStore := orderstore.NewSqlStore(appCtx.GetDatabase())
	orderTrackingStore := ordertrackingstore.NewSqlStore(appCtx.GetDatabase())
	historyPaymentStore := historypaymentstore.NewSqlStore(appCtx.GetDatabase())

	biz := orderbiz.NewConfirmPaymentCryptoBiz(orderStore, orderTrackingStore, historyPaymentStore)

	if err := biz.ConfirmCryptoPayment(context.Background(), &orderPaymentEvent); err != nil {
		log.Error(err)
		return
	}
}
