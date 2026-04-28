package assetCtrl

import (
	"apiRequest/consts"
	"apiRequest/signature"
	"apiRequest/structs"
	"encoding/json"
	"fmt"
)

type AssetCtrl struct {
	env structs.Env
}

func NewAssetCtrl(env *structs.Env) *AssetCtrl {
	return &AssetCtrl{
		env: *env,
	}
}

func (a *AssetCtrl) GetDepositList() {
	requestURL := fmt.Sprintf(a.env.Url+consts.ASSET_DEPOSIT_LIST+"?ccy=USDT&&size=%s", "50")
	requestPath := fmt.Sprintf(consts.ASSET_DEPOSIT_LIST+"?ccy=USDT&&size=%s", "50")
	signature.DoHttp(requestURL, consts.HTTP_METHOD_GET, requestPath, "", &a.env)
}

func (a *AssetCtrl) GetWithdrawList() {
	requestURL := fmt.Sprintf(a.env.Url+consts.ASSET_WITHDARW_LIST+"?size=%s", "1")
	requestPath := fmt.Sprintf(consts.ASSET_WITHDARW_LIST+"?size=%s", "1")
	signature.DoHttp(requestURL, consts.HTTP_METHOD_GET, requestPath, "", &a.env)
}

func (a *AssetCtrl) GetInternalTransferSupport() {
	requestURL := fmt.Sprintf(a.env.Url + consts.INTERNAL_TRANSFER_SUPPORT)
	requestPath := fmt.Sprintf(consts.INTERNAL_TRANSFER_SUPPORT)
	signature.DoHttp(requestURL, consts.HTTP_METHOD_GET, requestPath, "", &a.env)
}

func (a *AssetCtrl) GetInternalTransferHistory() {
	requestURL := fmt.Sprintf(a.env.Url + consts.INTERNAL_TRANSFER_HISTORY)
	requestPath := fmt.Sprintf(consts.INTERNAL_TRANSFER_HISTORY)
	signature.DoHttp(requestURL, consts.HTTP_METHOD_GET, requestPath, "", &a.env)
}

func (a *AssetCtrl) PostInternalTransfer() {
	type internalTransferRequest struct {
		Amount          string `json:"amount"`
		Coin            string `json:"coin"`
		ReceiverAccount string `json:"receiverAccount,omitempty"`
		AccountType     string `json:"accountType,omitempty"`
		ReceiverUID     string `json:"receiverUID"`
	}

	req := &internalTransferRequest{}
	req.Amount = "10"
	req.Coin = "USDT"
	req.ReceiverUID = "36007196"

	requestBody, err := json.Marshal(req)
	if err != nil {
		fmt.Println(err)
		return
	}

	requestURL := a.env.Url + consts.INTERNAL_TRANSFER
	signature.DoHttp(requestURL, consts.HTTP_METHOD_POST, consts.INTERNAL_TRANSFER, string(requestBody), &a.env)
}

// SubAccountTransfer 子母账号划转
func (a *AssetCtrl) SubAccountTransfer() {
	type subAccountTransferRequest struct {
		FromUid string `json:"fromUid"`
		ToUid   string `json:"toUid"`
		FromId  string `json:"fromId"`
		ToId    string `json:"toId"`
		Amount  string `json:"amount"`
		Coin    string `json:"coin"`
	}

	req := &subAccountTransferRequest{
		FromUid: "36033735", // 划出uid
		ToUid:   "36034023", // 划入uid
		FromId:  "7",        // 划出账户id (1:现货账户 2:钱包账户 5:反向合约账户 7:正向合约账户)
		ToId:    "7",        // 划入账户id (1:现货账户 2:钱包账户 5:反向合约账户 7:正向合约账户)
		Amount:  "10",       // 金额
		Coin:    "UDST",     // 币种
	}

	requestBody, err := json.Marshal(req)
	if err != nil {
		fmt.Println(err)
		return
	}

	requestURL := a.env.Url + consts.SUB_ACCOUNT_TRANSFER
	signature.DoHttp(requestURL, consts.HTTP_METHOD_POST, consts.SUB_ACCOUNT_TRANSFER, string(requestBody), &a.env)
}

// SubAccountTransferRecord 子母账号划转记录
func (a *AssetCtrl) SubAccountTransferRecord() {
	requestURL := fmt.Sprintf(a.env.Url+consts.SUB_ACCOUNT_TRANSFER_RECORD+"?coin=%s&fromId=%s&toId=%s&relationType=%s&page=%d&size=%d", "USDT", "7", "7", "1", 1, 20)
	requestPath := fmt.Sprintf(consts.SUB_ACCOUNT_TRANSFER_RECORD+"?coin=%s&fromId=%s&toId=%s&relationType=%s&page=%d&size=%d", "USDT", "7", "7", "1", 1, 20)
	signature.DoHttp(requestURL, consts.HTTP_METHOD_GET, requestPath, "", &a.env)
}

// SubAccountList 查询子账号列表
func (a *AssetCtrl) SubAccountList() {
	requestURL := fmt.Sprintf(a.env.Url + consts.SUB_ACCOUNT_LIST)
	requestPath := fmt.Sprintf(consts.SUB_ACCOUNT_LIST)
	signature.DoHttp(requestURL, consts.HTTP_METHOD_GET, requestPath, "", &a.env)
}

// RechargeChainList 充值链列表
func (a *AssetCtrl) RechargeChainList() {
	requestURL := fmt.Sprintf(a.env.Url+consts.RECHARGE_CHAIN_LIST+"?ccy=%s&lang=%s", "USDT", "zh")
	requestPath := fmt.Sprintf(consts.RECHARGE_CHAIN_LIST+"?ccy=%s&lang=%s", "USDT", "zh")
	signature.DoHttp(requestURL, consts.HTTP_METHOD_GET, requestPath, "", &a.env)
}

// AssetTransfer 资金划转
func (a *AssetCtrl) AssetTransfer() {
	type assetTransferRequest struct {
		CurrencyId string `json:"ccy"`
		Amount     string `json:"amount"`
		FromId     int    `json:"from_id"`
		ToId       int    `json:"to_id"`
		Uid        int    `json:"uid"`
	}

	req := &assetTransferRequest{
		CurrencyId: "USDT",   // 币种ID
		Amount:     "10",     // 划转金额
		FromId:     7,        // 转出账户ID (7:正向合约账户)
		ToId:       1,        // 转入账户ID (1:现货账户)
		Uid:        36007196, // 用户ID
	}

	requestBody, err := json.Marshal(req)
	if err != nil {
		fmt.Println(err)
		return
	}

	requestURL := a.env.Url + consts.ASSET_TRANSFER
	signature.DoHttp(requestURL, consts.HTTP_METHOD_POST, consts.ASSET_TRANSFER, string(requestBody), &a.env)
}
