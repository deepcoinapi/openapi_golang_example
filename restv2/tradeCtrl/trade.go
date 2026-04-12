package tradeCtrl

import (
	"apiRequest/consts"
	"apiRequest/signature"
	"apiRequest/structs"
	"encoding/json"
	"fmt"
)

type TradeCtrl struct {
	env structs.Env
}

func NewTradeCtrl(env *structs.Env) *TradeCtrl {
	return &TradeCtrl{
		env: *env,
	}
}

func (t *TradeCtrl) Order() {
	type orderRequest struct {
		InstId         string `json:"instId,omitempty"`
		TdMode         string `json:"tdMode,omitempty"`
		Ccy            string `json:"ccy,omitempty"`
		Side           string `json:"side,omitempty"`
		PosSide        string `json:"posSide,omitempty"`
		MrgPosition    string `json:"mrgPosition,omitempty"`
		ClosePosId     string `json:"closePosId,omitempty"`
		OrdType        string `json:"ordType,omitempty"`
		Sz             string `json:"sz,omitempty"`
		Px             string `json:"px,omitempty"`
		ReduceOnly     string `json:"reduceOnly,omitempty"`
		TgtCcy         string `json:"tgtCcy,omitempty"`
		TPTriggerPrice string `json:"tPTriggerPrice,omitempty"`
		SLTriggerPrice string `json:"sLTriggerPrice,omitempty"`
	}

	order := &orderRequest{}
	order.InstId = "BTC-USDT-SWAP"
	order.Ccy = "USDT"
	order.Side = consts.SIDE_BUY
	order.OrdType = consts.ORDER_TYPE_MARKET
	order.Sz = "90000"
	order.Px = "1"
	order.PosSide = consts.POSITION_SIDE_LONG
	order.TdMode = consts.CROSS
	order.MrgPosition = consts.SPLIT
	// order.ClosePosId = "1000583388205658"
	order.TPTriggerPrice = "100000"
	// order.SLTriggerPrice = "80000"

	requestBody, err := json.Marshal(order)
	if err != nil {
		fmt.Println(err)
		return
	}

	requestURL := t.env.Url + consts.TRADE_ORDER
	signature.DoHttp(requestURL, consts.HTTP_METHOD_POST, consts.TRADE_ORDER, string(requestBody), &t.env)
}

func (t *TradeCtrl) CancelOrder() {
	type cancelOrderRequest struct {
		InstId  string `json:"instId,omitempty"`
		OrdId   string `json:"ordId,omitempty"`
		ClOrdId string `json:"clOrdId,omitempty"`
	}

	cancelOrder := &cancelOrderRequest{}
	cancelOrder.InstId = "BTC-USDT-SWAP"
	cancelOrder.OrdId = "1000587866272245"

	requestBody, err := json.Marshal(cancelOrder)
	if err != nil {
		fmt.Println(err)
		return
	}

	requestURL := t.env.Url + consts.TRADE_CANCEL_ORDER
	signature.DoHttp(requestURL, consts.HTTP_METHOD_POST, consts.TRADE_CANCEL_ORDER, string(requestBody), &t.env)
}

func (t *TradeCtrl) TradeFills() {
	requestURL := fmt.Sprintf(t.env.Url+consts.TRADE_FILLS+"?instType=%s&instId=%s", consts.SWAP, "BTC-USDT-SWAP")
	requestPath := fmt.Sprintf(consts.TRADE_FILLS+"?instType=%s&instId=%s", consts.SWAP, "BTC-USDT-SWAP")
	signature.DoHttp(requestURL, consts.HTTP_METHOD_GET, requestPath, "", &t.env)
}

func (t *TradeCtrl) HistoryOrder() {
	requestURL := fmt.Sprintf(t.env.Url+consts.TRADE_HISTORY_ORDER+"?instType=%s&ordType=market", consts.SWAP)
	requestPath := fmt.Sprintf(consts.TRADE_HISTORY_ORDER+"?instType=%s&ordType=market", consts.SWAP)
	signature.DoHttp(requestURL, consts.HTTP_METHOD_GET, requestPath, "", &t.env)
}

func (t *TradeCtrl) SpotHistoryOrder() {
	requestURL := fmt.Sprintf(t.env.Url+consts.TRADE_HISTORY_ORDER+"?instType=%s&ordId=%s", consts.SPOT, "1000750232272249")
	requestPath := fmt.Sprintf(consts.TRADE_HISTORY_ORDER+"?instType=%s&ordId=%s", consts.SPOT, "1000750232272249")
	signature.DoHttp(requestURL, consts.HTTP_METHOD_GET, requestPath, "", &t.env)
}

func (t *TradeCtrl) PendingOrder() {
	requestURL := fmt.Sprintf(t.env.Url+consts.TRADE_PENDING_ORDER+"?instType=%s&ordId=%s", consts.SWAP, "1000187178097137")
	requestPath := fmt.Sprintf(consts.TRADE_PENDING_ORDER+"?instType=%s&ordId=%s", consts.SWAP, "1000187178097137")
	signature.DoHttp(requestURL, consts.HTTP_METHOD_GET, requestPath, "", &t.env)
}

func (t *TradeCtrl) SpotPendingOrder() {
	requestURL := fmt.Sprintf(t.env.Url+consts.TRADE_PENDING_ORDER+"?instType=%s", consts.SPOT)
	requestPath := fmt.Sprintf(consts.TRADE_PENDING_ORDER+"?instType=%s", consts.SPOT)
	signature.DoHttp(requestURL, consts.HTTP_METHOD_GET, requestPath, "", &t.env)
}

func (t *TradeCtrl) GetPosition() {
	requestURL := fmt.Sprintf(t.env.Url+consts.TRADE_POSITION+"?instType=%s&posId=%s", consts.SWAP, "1000439562104988")
	requestPath := fmt.Sprintf(consts.TRADE_POSITION+"?instType=%s&posId=%s", consts.SWAP, "1000439562104988")
	signature.DoHttp(requestURL, consts.HTTP_METHOD_GET, requestPath, "", &t.env)
}

func (t *TradeCtrl) SwapQueryOrderByOrderSysID() {
	requestURL := fmt.Sprintf(t.env.Url+consts.TRADE_ORDER_BY_ID+"?instId=%s&ordId=%s", "BTC-USDT-SWAP", "1000587866808715")
	requestPath := fmt.Sprintf(consts.TRADE_ORDER_BY_ID+"?instId=%s&ordId=%s", "BTC-USDT-SWAP", "1000587866808715")
	signature.DoHttp(requestURL, consts.HTTP_METHOD_GET, requestPath, "", &t.env)
}

func (t *TradeCtrl) SpotQueryOrderByOrderSysID() {
	requestURL := fmt.Sprintf(t.env.Url+consts.TRADE_ORDER_BY_ID+"?instId=%s&ordId=%s", "BTC-USDT", "1000188018518136")
	requestPath := fmt.Sprintf(consts.TRADE_ORDER_BY_ID+"?instId=%s&ordId=%s", "BTC-USDT", "1000188018518136")
	signature.DoHttp(requestURL, consts.HTTP_METHOD_GET, requestPath, "", &t.env)
}

func (t *TradeCtrl) SwapFinishQueryOrderByOrderSysID() {
	requestURL := fmt.Sprintf(t.env.Url+consts.TRADE_FINISH_ORDER_BY_ID+"?instId=%s&ordId=%s", "BTC-USDT-SWAP", "1000587866272245")
	requestPath := fmt.Sprintf(consts.TRADE_FINISH_ORDER_BY_ID+"?instId=%s&ordId=%s", "BTC-USDT-SWAP", "1000587866272245")
	signature.DoHttp(requestURL, consts.HTTP_METHOD_GET, requestPath, "", &t.env)
}

func (t *TradeCtrl) SpotFinishQueryOrderByOrderSysID() {
	requestURL := fmt.Sprintf(t.env.Url+consts.TRADE_FINISH_ORDER_BY_ID+"?instId=%s&ordId=%s", "BTC-USDT", "1000188021378725")
	requestPath := fmt.Sprintf(consts.TRADE_FINISH_ORDER_BY_ID+"?instId=%s&ordId=%s", "BTC-USDT", "1000188021378725")
	signature.DoHttp(requestURL, consts.HTTP_METHOD_GET, requestPath, "", &t.env)
}

func (t *TradeCtrl) GetFundingRate() {
	requestURL := fmt.Sprintf(t.env.Url+consts.TRADE_FUNDING_RATE+"?instType=%s", "SwapU")
	requestPath := fmt.Sprintf(consts.TRADE_FUNDING_RATE+"?instType=%s", "SwapU")
	signature.DoHttp(requestURL, consts.HTTP_METHOD_GET, requestPath, "", &t.env)
}

func (t *TradeCtrl) ReplaceOrder() {
	type replaceOrderRequest struct {
		OrderSysID   string  `json:"orderSysID"`
		Price        float64 `json:"price"`
		Volume       float64 `json:"volume"`
		ProductGroup string  `json:"productGroup"`
		TpTriggerPx  float64 `json:"tpTriggerPx"`
		SlTriggerPx  float64 `json:"slTriggerPx"`
	}

	replaceOrder := &replaceOrderRequest{}
	replaceOrder.OrderSysID = "1000587867035933"
	replaceOrder.TpTriggerPx = 110003
	replaceOrder.SlTriggerPx = 70002

	requestBody, err := json.Marshal(replaceOrder)
	if err != nil {
		fmt.Println(err)
		return
	}

	requestURL := t.env.Url + consts.TRADE_REPLACE_ORDER
	signature.DoHttp(requestURL, consts.HTTP_METHOD_POST, consts.TRADE_REPLACE_ORDER, string(requestBody), &t.env)
}

func (t *TradeCtrl) BatchCancelOrder() {
	type batchCancelOrderRequest struct {
		OrderSysIDs []string `json:"orderSysIDs"`
	}

	cancelOrders := &batchCancelOrderRequest{
		OrderSysIDs: []string{"1000587865918838", "1000587865914949"},
	}

	requestBody, err := json.Marshal(cancelOrders)
	if err != nil {
		fmt.Println("JSON编码错误:", err)
		return
	}

	requestURL := t.env.Url + consts.TRADE_BATCH_CANCEL_ORDER
	signature.DoHttp(requestURL, consts.HTTP_METHOD_POST, consts.TRADE_BATCH_CANCEL_ORDER, string(requestBody), &t.env)
}

func (t *TradeCtrl) SwapQueryPendingOrders() {
	requestURL := fmt.Sprintf(t.env.Url+consts.TRADE_PENDING_ORDER_V2+"?instId=%s&index=%d&limit=%d", "BTC-USDT-SWAP", 1, 100)
	requestPath := fmt.Sprintf(consts.TRADE_PENDING_ORDER_V2+"?instId=%s&index=%d&limit=%d", "BTC-USDT-SWAP", 1, 100)
	signature.DoHttp(requestURL, consts.HTTP_METHOD_GET, requestPath, "", &t.env)
}

func (t *TradeCtrl) SwapCalcelAllOrders() {
	type calcelAllOrdersRequest struct {
		InstrumentID  string `json:"instrumentID"`
		ProductGroup  string `json:"productGroup"`
		IsCrossMargin int32  `json:"IsCrossMargin"`
		IsMergeMode   int32  `json:"IsMergeMode"`
	}

	req := &calcelAllOrdersRequest{
		InstrumentID:  "BTCUSDT",
		ProductGroup:  "SwapU",
		IsCrossMargin: 1,
		IsMergeMode:   0,
	}

	requestBody, err := json.Marshal(req)
	if err != nil {
		fmt.Println("JSON编码错误:", err)
		return
	}

	requestURL := t.env.Url + consts.TRADE_SWAP_CANCEL_ALL
	signature.DoHttp(requestURL, consts.HTTP_METHOD_POST, consts.TRADE_SWAP_CANCEL_ALL, string(requestBody), &t.env)
}

func (t *TradeCtrl) SwapReplaceOrderSlTp() {
	type replaceOrderRequest struct {
		OrderSysID   string  `json:"orderSysID"`
		ProductGroup string  `json:"productGroup"`
		TpTriggerPx  float64 `json:"tpTriggerPx"`
		SlTriggerPx  float64 `json:"slTriggerPx"`
	}

	replaceOrder := &replaceOrderRequest{}
	replaceOrder.OrderSysID = "1000588112470603"
	replaceOrder.SlTriggerPx = 80005

	requestBody, err := json.Marshal(replaceOrder)
	if err != nil {
		fmt.Println(err)
		return
	}

	requestURL := t.env.Url + consts.TRADE_REPLACE_ORDER_SLTP
	signature.DoHttp(requestURL, consts.HTTP_METHOD_POST, consts.TRADE_REPLACE_ORDER_SLTP, string(requestBody), &t.env)
}

func (t *TradeCtrl) SwapReplacePositionSlTp() {
	type replacePosSlTpRequest struct {
		OrderSysID   string  `json:"orderSysID"`
		ProductGroup string  `json:"productGroup"`
		TpTriggerPx  float64 `json:"tpTriggerPx"`
		SlTriggerPx  float64 `json:"slTriggerPx"`
		Volume       float64 `json:"volume"`
	}

	replaceOrder := &replacePosSlTpRequest{}
	replaceOrder.OrderSysID = "1000588139491613"

	replaceOrder.TpTriggerPx = 100006
	replaceOrder.SlTriggerPx = 80006
	replaceOrder.Volume = 0
	requestBody, err := json.Marshal(replaceOrder)
	if err != nil {
		fmt.Println(err)
		return
	}

	requestURL := t.env.Url + consts.TRADE_REPLACE_POS_SLTP
	signature.DoHttp(requestURL, consts.HTTP_METHOD_POST, consts.TRADE_REPLACE_POS_SLTP, string(requestBody), &t.env)
}

// =================== V2 交易接口示例 ===================

// BatchOrders 批量下单
func (t *TradeCtrl) BatchOrders() {
	type batchOrderItem struct {
		InstId      string `json:"instId,omitempty"`
		TdMode      string `json:"tdMode,omitempty"`
		Side        string `json:"side,omitempty"`
		OrdType     string `json:"ordType,omitempty"`
		Sz          string `json:"sz,omitempty"`
		Px          string `json:"px,omitempty"`
		PosSide     string `json:"posSide,omitempty"`
		MrgPosition string `json:"mrgPosition,omitempty"`
	}

	type batchOrderRequest struct {
		Orders []batchOrderItem `json:"orders"`
	}

	req := &batchOrderRequest{
		Orders: []batchOrderItem{
			{
				InstId:      "BTC-USDT-SWAP",
				TdMode:      consts.CROSS,
				Side:        consts.SIDE_BUY,
				OrdType:     consts.ORDER_TYPE_LIMIT,
				Sz:          "1",
				Px:          "65000",
				PosSide:     consts.POSITION_SIDE_LONG,
				MrgPosition: consts.MERGE,
			},
			{
				InstId:      "ETH-USDT-SWAP",
				TdMode:      consts.CROSS,
				Side:        consts.SIDE_SELL,
				OrdType:     consts.ORDER_TYPE_MARKET,
				Sz:          "2",
				PosSide:     consts.POSITION_SIDE_SHORT,
				MrgPosition: consts.MERGE,
			},
		},
	}

	requestBody, err := json.Marshal(req)
	if err != nil {
		fmt.Println(err)
		return
	}

	requestURL := t.env.Url + consts.TRADE_BATCH_ORDERS
	signature.DoHttp(requestURL, consts.HTTP_METHOD_POST, consts.TRADE_BATCH_ORDERS, string(requestBody), &t.env)
}

// TriggerOrder 条件委托单
func (t *TradeCtrl) TriggerOrder() {
	type triggerOrderRequest struct {
		InstId       string  `json:"instId,omitempty"`
		Sz           string  `json:"sz,omitempty"`
		Side         string  `json:"side,omitempty"`
		PosSide      string  `json:"posSide,omitempty"`
		Px           string  `json:"px,omitempty"`
		OrderType    string  `json:"orderType,omitempty"`
		TriggerPrice string  `json:"triggerPrice,omitempty"`
		MrgPosition  string  `json:"mrgPosition,omitempty"`
		TdMode       string  `json:"tdMode,omitempty"`
		TPTriggerPx  float64 `json:"tpTriggerPx,omitempty"`
		SLTriggerPx  float64 `json:"slTriggerPx,omitempty"`
	}

	req := &triggerOrderRequest{
		InstId:       "BTC-USDT-SWAP",
		Sz:           "1",
		Side:         consts.SIDE_BUY,
		PosSide:      consts.POSITION_SIDE_LONG,
		OrderType:    consts.ORDER_TYPE_MARKET,
		TriggerPrice: "150000",
		MrgPosition:  consts.MERGE,
		TdMode:       consts.CROSS,
		TPTriggerPx:  160000,
		SLTriggerPx:  140000,
	}

	requestBody, err := json.Marshal(req)
	if err != nil {
		fmt.Println(err)
		return
	}

	requestURL := t.env.Url + consts.TRADE_TRIGGER_ORDER
	signature.DoHttp(requestURL, consts.HTTP_METHOD_POST, consts.TRADE_TRIGGER_ORDER, string(requestBody), &t.env)
}

// CancelTriggerOrder 撤销条件单
func (t *TradeCtrl) CancelTriggerOrder() {
	type cancelTriggerOrderRequest struct {
		InstId string `json:"instId,omitempty"`
		OrdId  string `json:"ordId,omitempty"`
	}

	req := &cancelTriggerOrderRequest{
		InstId: "BTC-USDT-SWAP",
		OrdId:  "1000587866272245",
	}

	requestBody, err := json.Marshal(req)
	if err != nil {
		fmt.Println(err)
		return
	}

	requestURL := t.env.Url + consts.TRADE_CANCEL_TRIGGER_ORDER
	signature.DoHttp(requestURL, consts.HTTP_METHOD_POST, consts.TRADE_CANCEL_TRIGGER_ORDER, string(requestBody), &t.env)
}

// BatchCancelOrderV2 V2批量撤单
func (t *TradeCtrl) BatchCancelOrderV2() {
	type batchCancelItem struct {
		InstId string `json:"instId,omitempty"`
		OrdId  string `json:"ordId,omitempty"`
	}

	type batchCancelRequest struct {
		Orders []batchCancelItem `json:"orders"`
	}

	req := &batchCancelRequest{
		Orders: []batchCancelItem{
			{InstId: "BTC-USDT-SWAP", OrdId: "1000587866272245"},
			{InstId: "ETH-USDT-SWAP", OrdId: "1000587866272246"},
		},
	}

	requestBody, err := json.Marshal(req)
	if err != nil {
		fmt.Println(err)
		return
	}

	requestURL := t.env.Url + consts.TRADE_BATCH_CANCEL_ORDER
	signature.DoHttp(requestURL, consts.HTTP_METHOD_POST, consts.TRADE_BATCH_CANCEL_ORDER, string(requestBody), &t.env)
}

// ReplaceOrderV2 改单
func (t *TradeCtrl) ReplaceOrderV2() {
	type replaceOrderRequest struct {
		OrdId string  `json:"ordId,omitempty"`
		Px    float64 `json:"px,omitempty"`
		Sz    float64 `json:"sz,omitempty"`
	}

	req := &replaceOrderRequest{
		OrdId: "1000587866272245",
		Px:    66000,
		Sz:    1.5,
	}

	requestBody, err := json.Marshal(req)
	if err != nil {
		fmt.Println(err)
		return
	}

	requestURL := t.env.Url + consts.TRADE_REPLACE_ORDER
	signature.DoHttp(requestURL, consts.HTTP_METHOD_POST, consts.TRADE_REPLACE_ORDER, string(requestBody), &t.env)
}

// ReplaceOrderSLTPV2 修改订单止盈止损
func (t *TradeCtrl) ReplaceOrderSLTPV2() {
	type replaceSLTPRequest struct {
		OrdId       string  `json:"ordId,omitempty"`
		TPTriggerPx float64 `json:"tpTriggerPx,omitempty"`
		SLTriggerPx float64 `json:"slTriggerPx,omitempty"`
	}

	req := &replaceSLTPRequest{
		OrdId:       "1000587866272245",
		TPTriggerPx: 170000,
		SLTriggerPx: 130000,
	}

	requestBody, err := json.Marshal(req)
	if err != nil {
		fmt.Println(err)
		return
	}

	requestURL := t.env.Url + consts.TRADE_REPLACE_ORDER_SLTP
	signature.DoHttp(requestURL, consts.HTTP_METHOD_POST, consts.TRADE_REPLACE_ORDER_SLTP, string(requestBody), &t.env)
}

// SetPositionSLTP 设置持仓止盈止损
func (t *TradeCtrl) SetPositionSLTP() {
	type setPositionSLTPRequest struct {
		InstId      string  `json:"instId,omitempty"`
		PosSide     string  `json:"posSide,omitempty"`
		MrgPosition string  `json:"mrgPosition,omitempty"`
		TdMode      string  `json:"tdMode,omitempty"`
		TPTriggerPx float64 `json:"tpTriggerPx,omitempty"`
		SLTriggerPx float64 `json:"slTriggerPx,omitempty"`
	}

	req := &setPositionSLTPRequest{
		InstId:      "BTC-USDT-SWAP",
		PosSide:     consts.POSITION_SIDE_LONG,
		MrgPosition: consts.MERGE,
		TdMode:      consts.CROSS,
		TPTriggerPx: 170000,
		SLTriggerPx: 130000,
	}

	requestBody, err := json.Marshal(req)
	if err != nil {
		fmt.Println(err)
		return
	}

	requestURL := t.env.Url + consts.TRADE_SET_POSITION_SLTP
	signature.DoHttp(requestURL, consts.HTTP_METHOD_POST, consts.TRADE_SET_POSITION_SLTP, string(requestBody), &t.env)
}

// CancelPositionSLTP 取消持仓止盈止损
func (t *TradeCtrl) CancelPositionSLTP() {
	type cancelPositionSLTPRequest struct {
		InstId      string `json:"instId,omitempty"`
		PosSide     string `json:"posSide,omitempty"`
		MrgPosition string `json:"mrgPosition,omitempty"`
		TdMode      string `json:"tdMode,omitempty"`
	}

	req := &cancelPositionSLTPRequest{
		InstId:      "BTC-USDT-SWAP",
		PosSide:     consts.POSITION_SIDE_LONG,
		MrgPosition: consts.MERGE,
		TdMode:      consts.CROSS,
	}

	requestBody, err := json.Marshal(req)
	if err != nil {
		fmt.Println(err)
		return
	}

	requestURL := t.env.Url + consts.TRADE_CANCEL_POSITION_SLTP
	signature.DoHttp(requestURL, consts.HTTP_METHOD_POST, consts.TRADE_CANCEL_POSITION_SLTP, string(requestBody), &t.env)
}

// ModifyPositionSLTP 修改持仓止盈止损
func (t *TradeCtrl) ModifyPositionSLTP() {
	type modifyPositionSLTPRequest struct {
		InstId      string  `json:"instId,omitempty"`
		PosSide     string  `json:"posSide,omitempty"`
		MrgPosition string  `json:"mrgPosition,omitempty"`
		TdMode      string  `json:"tdMode,omitempty"`
		TPTriggerPx float64 `json:"tpTriggerPx,omitempty"`
		SLTriggerPx float64 `json:"slTriggerPx,omitempty"`
	}

	req := &modifyPositionSLTPRequest{
		InstId:      "BTC-USDT-SWAP",
		PosSide:     consts.POSITION_SIDE_LONG,
		MrgPosition: consts.MERGE,
		TdMode:      consts.CROSS,
		TPTriggerPx: 180000,
		SLTriggerPx: 120000,
	}

	requestBody, err := json.Marshal(req)
	if err != nil {
		fmt.Println(err)
		return
	}

	requestURL := t.env.Url + consts.TRADE_MODIFY_POSITION_SLTP
	signature.DoHttp(requestURL, consts.HTTP_METHOD_POST, consts.TRADE_MODIFY_POSITION_SLTP, string(requestBody), &t.env)
}

// CancelOrderAll 一键撤单
func (t *TradeCtrl) CancelOrderAll() {
	type cancelOrderAllRequest struct {
		InstId        string `json:"instId,omitempty"`
		IsCrossMargin int    `json:"IsCrossMargin,omitempty"`
		IsMergeMode   int    `json:"IsMergeMode,omitempty"`
	}

	req := &cancelOrderAllRequest{
		InstId:        "BTC-USDT-SWAP",
		IsCrossMargin: 1,
		IsMergeMode:   1,
	}

	requestBody, err := json.Marshal(req)
	if err != nil {
		fmt.Println(err)
		return
	}

	requestURL := t.env.Url + consts.TRADE_CANCEL_ORDER_ALL
	signature.DoHttp(requestURL, consts.HTTP_METHOD_POST, consts.TRADE_CANCEL_ORDER_ALL, string(requestBody), &t.env)
}

// CancelTriggerOrderAll 一键撤销条件单
func (t *TradeCtrl) CancelTriggerOrderAll() {
	type cancelTriggerAllRequest struct {
		InstId        string `json:"instId,omitempty"`
		IsCrossMargin int    `json:"IsCrossMargin,omitempty"`
		IsMergeMode   int    `json:"IsMergeMode,omitempty"`
	}

	req := &cancelTriggerAllRequest{
		InstId:        "BTC-USDT-SWAP",
		IsCrossMargin: 1,
		IsMergeMode:   1,
	}

	requestBody, err := json.Marshal(req)
	if err != nil {
		fmt.Println(err)
		return
	}

	requestURL := t.env.Url + consts.TRADE_CANCEL_TRIGGER_ALL
	signature.DoHttp(requestURL, consts.HTTP_METHOD_POST, consts.TRADE_CANCEL_TRIGGER_ALL, string(requestBody), &t.env)
}

// BatchClosePosition 批量平仓
func (t *TradeCtrl) BatchClosePosition() {
	type batchClosePositionRequest struct {
		InstId string `json:"instId,omitempty"`
	}

	req := &batchClosePositionRequest{
		InstId: "BTC-USDT-SWAP",
	}

	requestBody, err := json.Marshal(req)
	if err != nil {
		fmt.Println(err)
		return
	}

	requestURL := t.env.Url + consts.TRADE_BATCH_CLOSE_POSITION
	signature.DoHttp(requestURL, consts.HTTP_METHOD_POST, consts.TRADE_BATCH_CLOSE_POSITION, string(requestBody), &t.env)
}

// ClosePositionByIds 按仓位ID平仓
func (t *TradeCtrl) ClosePositionByIds() {
	type closePositionByIdsRequest struct {
		InstId      string   `json:"instId,omitempty"`
		PositionIds []string `json:"positionIds,omitempty"`
	}

	req := &closePositionByIdsRequest{
		InstId:      "BTC-USDT-SWAP",
		PositionIds: []string{"1000439562104988", "1000439562104989"},
	}

	requestBody, err := json.Marshal(req)
	if err != nil {
		fmt.Println(err)
		return
	}

	requestURL := t.env.Url + consts.TRADE_CLOSE_POSITION_BY_IDS
	signature.DoHttp(requestURL, consts.HTTP_METHOD_POST, consts.TRADE_CLOSE_POSITION_BY_IDS, string(requestBody), &t.env)
}

// OrderList 订单列表
func (t *TradeCtrl) OrderList() {
	requestURL := fmt.Sprintf(t.env.Url+consts.TRADE_ORDER_LIST+"?instId=%s&ordType=%s", "BTC-USDT-SWAP", "limit")
	requestPath := fmt.Sprintf(consts.TRADE_ORDER_LIST+"?instId=%s&ordType=%s", "BTC-USDT-SWAP", "limit")
	signature.DoHttp(requestURL, consts.HTTP_METHOD_GET, requestPath, "", &t.env)
}
