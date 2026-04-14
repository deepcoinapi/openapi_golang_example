package main

import (
	"apiRequest/accountCtrl"
	assetCtrl "apiRequest/asset"
	"apiRequest/consts"
	"apiRequest/copytrading"
	"apiRequest/marketCtrl"
	"apiRequest/structs"
	"apiRequest/tradeCtrl"
	"apiRequest/ws"
	"fmt"
	"log"
	"os"

	"github.com/spf13/viper"
)

func getConfigs() *structs.Env {
	configName := consts.MASTER

	fmt.Println(configName)

	viper.SetConfigName(configName)
	viper.AddConfigPath("./config")

	viper.SetConfigType("yaml")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Failed to read config file: %v", err)
	}

	url := viper.GetString("api.url")
	key := viper.GetString("api.key")
	secretKey := viper.GetString("api.secret_key")
	passphrase := viper.GetString("api.passphrase")

	return &structs.Env{
		Url:        url,
		Key:        key,
		SecretKey:  secretKey,
		Passphrase: passphrase,
	}
}

type CommandHandler func()

type CommandRegistry struct {
	handlers map[string]CommandHandler
}

func NewCommandRegistry(env *structs.Env) *CommandRegistry {
	account := accountCtrl.NewAccountCtrl(env)
	market := marketCtrl.NewMarketCtrl(env)
	trade := tradeCtrl.NewTradeCtrl(env)
	ws := ws.NewPrivateWsCtrl(env)
	copytrading := copytrading.NewCopyTradingCtrl(env)
	asset := assetCtrl.NewAssetCtrl(env)

	return &CommandRegistry{
		handlers: map[string]CommandHandler{
			// account
			"getAccountBalance": account.GetAccountBalance, //✅
			"getUid":            account.GetUid,            //✅
			"getAccountBills":   account.GetAccountBills,   //✅
			"setLeverage":       account.SetLeverage,       //✅
			"getPositions":      account.GetPositions,      //✅

			// market
			"getMarketBooks":              market.GetMarketBooks,              //✅
			"getMarketCandles":            market.GetMarketCandles,            //✅
			"getMarketInstruments":        market.GetMarketInstruments,        //✅
			"getMarketTickers":            market.GetMarketTickers,            //✅
			"getMarketIndexCandles":       market.GetMarketIndexCandles,       //✅
			"getMarketTrades":             market.GetMarketTrades,             //✅
			"getMarketMarkCandles":        market.GetMarketMarkCandles,        //✅
			"getMarketPositionGrade":      market.GetMarketPositionGrade,      //✅
			"getMarketBookSpread":         market.GetMarketBookSpread,         //✅
			"getMarketSysTime":            market.GetMarketSysTime,            //✅
			"getMarketHandicapKline1m":    market.GetMarketHandicapKline1m,    //✅
			"getMarketHandicapOrderbook":  market.GetMarketHandicapOrderbook,  //✅
			"getMarketHandicapTrade":      market.GetMarketHandicapTrade,      //✅
			"getMarketFundingRate":        market.GetMarketFundingRate,        //✅
			"getMarketCurrentFundingRate": market.GetMarketCurrentFundingRate, //✅
			"getMarketFundingRateHistory": market.GetMarketFundingRateHistory, //✅
			"getMarketSysPing":            market.GetMarketSysPing,

			// trade
			"order":                        trade.Order,                        //✅
			"batchOrders":                  trade.BatchOrders,                  //✅
			"replace-order":                trade.ReplaceOrder,                 //✅
			"cancelOrder":                  trade.CancelOrder,                  //✅
			"batch-cancel-order":           trade.BatchCancelOrder,             //✅
			"cancelTriggerOrder":           trade.CancelTriggerOrder,           //✅
			"cancelOrderAll":               trade.CancelOrderAll,               //✅
			"cancelTriggerOrderAll":        trade.CancelTriggerOrderAll,        //✅
			"tradeFills":                   trade.TradeFills,                   //✅
			"queryOrderByOrderSysID":       trade.QueryOrderByOrderSysID,       //✅
			"batchOrderQuery":              trade.BatchOrderQuery,              //✅
			"finishQueryOrderByOrderSysID": trade.FinishQueryOrderByOrderSysID, //✅
			"historyOrder":                 trade.HistoryOrder,                 //✅
			"pendingOrder":                 trade.PendingOrder,                 //✅
			"triggerOrder":                 trade.TriggerOrder,                 //✅
			"batchClosePosition":           trade.BatchClosePosition,           //✅
			"replaceOrderSlTp":             trade.ReplaceOrderSlTp,             //✅
			"closePositionByIds":           trade.ClosePositionByIds,           //✅
			"setPositionSLTP":              trade.SetPositionSLTP,              //✅
			"cancelPositionSLTP":           trade.CancelPositionSLTP,           //✅
			"modifyPositionSLTP":           trade.ModifyPositionSLTP,           //✅
			"triggerOrdersPending":         trade.TriggerOrdersPending,         //✅
			"triggerOrdersHistory":         trade.TriggerOrdersHistory,         //✅
			"traceOrder":                   trade.TraceOrder,                   //✅
			"traceOrderList":               trade.TraceOrderList,               //✅

			// WebSocket
			"getListenKey":    ws.GetListenKey, //✅
			"extendListenKey": ws.Extend,       //✅

			// copytrading
			"leader-settings":    copytrading.LeaderSettings,     //✅
			"support-contracts":  copytrading.SupportContracts,   //✅
			"set-contracts":      copytrading.SetContracts,       //✅
			"leader-position":    copytrading.LeaderPosition,     //✅
			"estimate-profit":    copytrading.EstimateProfit,     //✅
			"history-profit":     copytrading.HistoryProfit,      //✅
			"follower-rank":      copytrading.FollowerRank,       //✅
			"positionType":       copytrading.PositionType,       //✅
			"updatePositionType": copytrading.UpdatePositionType, //✅

			// internal-transfer
			"getInternalTransferSupport": asset.GetInternalTransferSupport, //✅
			"postInternalTransfer":       asset.PostInternalTransfer,       //✅
			"getInternalTransferHistory": asset.GetInternalTransferHistory, //✅

			// subAccount
			"subAccountTransfer": asset.SubAccountTransfer, //✅

			// asset
			"depositList":  asset.GetDepositList,
			"withdrawList": asset.GetWithdrawList,
		},
	}
}

func (r *CommandRegistry) Execute(command string) error {
	handler, exists := r.handlers[command]
	if !exists {
		return fmt.Errorf("unknown command: %s", command)
	}
	handler()
	return nil
}

func main() {
	args := os.Args
	if len(args) < 2 {
		fmt.Println("command is required")
		return
	}

	env := getConfigs()
	registry := NewCommandRegistry(env)

	if err := registry.Execute(args[1]); err != nil {
		fmt.Println(err)
	}
}
