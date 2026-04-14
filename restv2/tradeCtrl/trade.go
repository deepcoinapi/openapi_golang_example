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
	requestURL := fmt.Sprintf(t.env.Url+consts.TRADE_HISTORY_ORDER+"?instType=%s", consts.SWAP)
	requestPath := fmt.Sprintf(consts.TRADE_HISTORY_ORDER+"?instType=%s", consts.SWAP)
	signature.DoHttp(requestURL, consts.HTTP_METHOD_GET, requestPath, "", &t.env)
}

func (t *TradeCtrl) PendingOrder() {
	requestURL := fmt.Sprintf(t.env.Url+consts.TRADE_PENDING_ORDER+"?page=%v", 1)
	requestPath := fmt.Sprintf(consts.TRADE_PENDING_ORDER+"?page=%v", 1)
	signature.DoHttp(requestURL, consts.HTTP_METHOD_GET, requestPath, "", &t.env)
}

func (t *TradeCtrl) QueryOrderByOrderSysID() {
	requestURL := fmt.Sprintf(t.env.Url+consts.TRADE_ORDER_BY_ID+"?instId=%s&ordId=%s", "BTC-USDT-SWAP", "1000597765910558")
	requestPath := fmt.Sprintf(consts.TRADE_ORDER_BY_ID+"?instId=%s&ordId=%s", "BTC-USDT-SWAP", "1000597765910558")
	signature.DoHttp(requestURL, consts.HTTP_METHOD_GET, requestPath, "", &t.env)
}

func (t *TradeCtrl) FinishQueryOrderByOrderSysID() {
	requestURL := fmt.Sprintf(t.env.Url+consts.TRADE_FINISH_ORDER_BY_ID+"?instId=%s&ordId=%s", "BTC-USDT-SWAP", "1000587866272245")
	requestPath := fmt.Sprintf(consts.TRADE_FINISH_ORDER_BY_ID+"?instId=%s&ordId=%s", "BTC-USDT-SWAP", "1000587866272245")
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
	replaceOrder.OrderSysID = "1000597765908207"
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
		OrdIds []string `json:"ordIds"`
	}

	cancelOrders := &batchCancelOrderRequest{
		OrdIds: []string{"1000587865918838", "1000587865914949"},
	}

	requestBody, err := json.Marshal(cancelOrders)
	if err != nil {
		fmt.Println("JSON编码错误:", err)
		return
	}

	requestURL := t.env.Url + consts.TRADE_BATCH_CANCEL_ORDER
	signature.DoHttp(requestURL, consts.HTTP_METHOD_POST, consts.TRADE_BATCH_CANCEL_ORDER, string(requestBody), &t.env)
}

func (t *TradeCtrl) ReplaceOrderSlTp() {
	type replaceOrderRequest struct {
		OrderSysID  string  `json:"ordId"`
		TpTriggerPx float64 `json:"tpTriggerPx"`
		SlTriggerPx float64 `json:"slTriggerPx"`
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
		InstId string `json:"instId,omitempty"`
		OrdId  string `json:"ordId,omitempty"`
	}

	req := &cancelPositionSLTPRequest{
		InstId: "BTC-USDT-SWAP",
		OrdId:  "1000587866272245",
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
		OrdId       string  `json:"ordId,omitempty"`
		PosSide     string  `json:"posSide,omitempty"`
		MrgPosition string  `json:"mrgPosition,omitempty"`
		TdMode      string  `json:"tdMode,omitempty"`
		TPTriggerPx float64 `json:"tpTriggerPx,omitempty"`
		SLTriggerPx float64 `json:"slTriggerPx,omitempty"`
	}

	req := &modifyPositionSLTPRequest{
		InstId:      "BTC-USDT-SWAP",
		OrdId:       "1000587866272245",
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
		PositionIds []string `json:"posIds,omitempty"`
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

// BatchOrderQuery 批量查询订单
func (t *TradeCtrl) BatchOrderQuery() {
	type orderQueryItem struct {
		InstId  string `json:"instId,omitempty"`
		OrdId   string `json:"ordId,omitempty"`
		ClOrdId string `json:"clOrdId,omitempty"`
	}

	type batchOrderQueryRequest struct {
		Orders []orderQueryItem `json:"orders"`
	}

	req := &batchOrderQueryRequest{
		Orders: []orderQueryItem{
			{
				InstId: "BTC-USDT-SWAP",
				OrdId:  "1000597586292096",
			},
			{
				InstId: "ETH-USDT-SWAP",
				OrdId:  "1000597586292104",
			},
		},
	}

	requestBody, err := json.Marshal(req)
	if err != nil {
		fmt.Println(err)
		return
	}

	requestURL := t.env.Url + consts.TRADE_BATCH_ORDER_QUERY
	signature.DoHttp(requestURL, consts.HTTP_METHOD_POST, consts.TRADE_BATCH_ORDER_QUERY, string(requestBody), &t.env)
}

// TriggerOrdersPending 查询未触发条件单
func (t *TradeCtrl) TriggerOrdersPending() {
	requestURL := fmt.Sprintf(t.env.Url+consts.TRADE_TRIGGER_ORDERS_PENDING+"?instId=%s&orderType=%s&limit=%d", "BTC-USDT-SWAP", "limit", 100)
	requestPath := fmt.Sprintf(consts.TRADE_TRIGGER_ORDERS_PENDING+"?instId=%s&orderType=%s&limit=%d", "BTC-USDT-SWAP", "limit", 100)
	signature.DoHttp(requestURL, consts.HTTP_METHOD_GET, requestPath, "", &t.env)
}

// TriggerOrdersHistory 查询已触发条件单
func (t *TradeCtrl) TriggerOrdersHistory() {
	requestURL := fmt.Sprintf(t.env.Url+consts.TRADE_TRIGGER_ORDERS_HISTORY+"?instId=%s&OrderType=%s&limit=%d", "BTC-USDT-SWAP", "limit", 100)
	requestPath := fmt.Sprintf(consts.TRADE_TRIGGER_ORDERS_HISTORY+"?instId=%s&OrderType=%s&limit=%d", "BTC-USDT-SWAP", "limit", 100)
	signature.DoHttp(requestURL, consts.HTTP_METHOD_GET, requestPath, "", &t.env)
}

// TraceOrder 追踪出场委托单
func (t *TradeCtrl) TraceOrder() {
	type traceOrderRequest struct {
		InstId       string `json:"instId,omitempty"`
		RetracePoint string `json:"retracePoint,omitempty"`
		TriggerPrice string `json:"triggerPrice,omitempty"`
		PosSide      string `json:"posSide,omitempty"`
	}

	req := &traceOrderRequest{
		InstId:       "BTC-USDT-SWAP",
		RetracePoint: "100",
		TriggerPrice: "95000",
		PosSide:      consts.POSITION_SIDE_LONG,
	}

	requestBody, err := json.Marshal(req)
	if err != nil {
		fmt.Println(err)
		return
	}

	requestURL := t.env.Url + consts.TRADE_TRACE_ORDER
	signature.DoHttp(requestURL, consts.HTTP_METHOD_POST, consts.TRADE_TRACE_ORDER, string(requestBody), &t.env)
}

// TraceOrderList 查看追踪出场委托单
func (t *TradeCtrl) TraceOrderList() {
	requestURL := fmt.Sprintf(t.env.Url+consts.TRADE_TRACE_ORDER_LIST+"?instId=%s", "BTC-USDT-SWAP")
	requestPath := fmt.Sprintf(consts.TRADE_TRACE_ORDER_LIST+"?instId=%s", "BTC-USDT-SWAP")
	signature.DoHttp(requestURL, consts.HTTP_METHOD_GET, requestPath, "", &t.env)
}
