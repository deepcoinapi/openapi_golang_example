package marketCtrl

import (
	"apiRequest/consts"
	"apiRequest/signature"
	"apiRequest/structs"
	"fmt"
)

type MarketCtrl struct {
	env structs.Env
}

func NewMarketCtrl(env *structs.Env) *MarketCtrl {
	return &MarketCtrl{
		env: *env,
	}
}

func (m *MarketCtrl) GetMarketBooks() {
	requestURL := fmt.Sprintf(m.env.Url+consts.MARKET_BOOKS+"?instId=%s", "BTC-USDT")
	requestPath := fmt.Sprintf(consts.MARKET_BOOKS+"?instId=%s", "BTC-USDT")
	signature.DoHttp(requestURL, consts.HTTP_METHOD_GET, requestPath, "", &m.env)
}

func (m *MarketCtrl) GetMarketCandles() {
	requestURL := fmt.Sprintf(m.env.Url+consts.MARKET_CANDLES+"?instId=%s", "BTC-USDT")
	requestPath := fmt.Sprintf(consts.MARKET_CANDLES+"?instId=%s", "BTC-USDT")
	signature.DoHttp(requestURL, consts.HTTP_METHOD_GET, requestPath, "", &m.env)
}

func (m *MarketCtrl) GetMarketTickers() {
	requestURL := fmt.Sprintf(m.env.Url+consts.MARKET_TICKERS+"?instType=%s&uly=%s", consts.SPOT, "SAGE-USDT")
	requestPath := fmt.Sprintf(consts.MARKET_TICKERS+"?instType=%s&uly=%s", consts.SPOT, "SAGE-USDT")
	signature.DoHttp(requestURL, consts.HTTP_METHOD_GET, requestPath, "", &m.env)
}

func (m *MarketCtrl) GetMarketInstruments() {
	requestURL := fmt.Sprintf(m.env.Url+consts.MARKET_INSTRUMENTS+"?instType=%s&uly=%s", consts.SWAP, "BTC-USDT")
	requestPath := fmt.Sprintf(consts.MARKET_INSTRUMENTS+"?instType=%s&uly=%s", consts.SWAP, "BTC-USDT")
	signature.DoHttp(requestURL, consts.HTTP_METHOD_GET, requestPath, "", &m.env)
}

// GetMarketIndexCandles 指数K线
func (m *MarketCtrl) GetMarketIndexCandles() {
	requestURL := fmt.Sprintf(m.env.Url+consts.MARKET_INDEX_CANDLES+"?instId=%s&bar=%s", "BTC-USDT", "1m")
	requestPath := fmt.Sprintf(consts.MARKET_INDEX_CANDLES+"?instId=%s&bar=%s", "BTC-USDT", "1m")
	signature.DoHttp(requestURL, consts.HTTP_METHOD_GET, requestPath, "", &m.env)
}

// GetMarketMarkCandles 标价K线
func (m *MarketCtrl) GetMarketMarkCandles() {
	requestURL := fmt.Sprintf(m.env.Url+consts.MARKET_MARK_CANDLES+"?instId=%s&bar=%s", "BTC-USDT-SWAP", "1m")
	requestPath := fmt.Sprintf(consts.MARKET_MARK_CANDLES+"?instId=%s&bar=%s", "BTC-USDT-SWAP", "1m")
	signature.DoHttp(requestURL, consts.HTTP_METHOD_GET, requestPath, "", &m.env)
}

// GetMarketPositionGrade 持仓等级
func (m *MarketCtrl) GetMarketPositionGrade() {
	requestURL := fmt.Sprintf(m.env.Url+consts.MARKET_POSITION_GRADE+"?instId=%s", "BTC-USDT-SWAP")
	requestPath := fmt.Sprintf(consts.MARKET_POSITION_GRADE+"?instId=%s", "BTC-USDT-SWAP")
	signature.DoHttp(requestURL, consts.HTTP_METHOD_GET, requestPath, "", &m.env)
}

// GetMarketTrades 成交记录
func (m *MarketCtrl) GetMarketTrades() {
	requestURL := fmt.Sprintf(m.env.Url+consts.MARKET_TRADES+"?instId=%s&limit=%d", "BTC-USDT-SWAP", 100)
	requestPath := fmt.Sprintf(consts.MARKET_TRADES+"?instId=%s&limit=%d", "BTC-USDT-SWAP", 100)
	signature.DoHttp(requestURL, consts.HTTP_METHOD_GET, requestPath, "", &m.env)
}

// GetMarketSysTime 系统时间
func (m *MarketCtrl) GetMarketSysTime() {
	requestURL := m.env.Url + consts.MARKET_SYS_TIME
	requestPath := consts.MARKET_SYS_TIME
	signature.DoHttp(requestURL, consts.HTTP_METHOD_GET, requestPath, "", &m.env)
}

// GetMarketSysPing 联通性检测
func (m *MarketCtrl) GetMarketSysPing() {
	requestURL := m.env.Url + consts.MARKET_SYS_PING
	requestPath := consts.MARKET_SYS_PING
	signature.DoHttp(requestURL, consts.HTTP_METHOD_GET, requestPath, "", &m.env)
}

// GetMarketBookSpread 深度点差
func (m *MarketCtrl) GetMarketBookSpread() {
	requestURL := fmt.Sprintf(m.env.Url+consts.MARKET_BOOK_SPREAD+"?instId=%s&value=%s&vType=%s", "BTC-USDT-SWAP", "100000", "0")
	requestPath := fmt.Sprintf(consts.MARKET_BOOK_SPREAD+"?instId=%s&value=%s&vType=%s", "BTC-USDT-SWAP", "100000", "0")
	signature.DoHttp(requestURL, consts.HTTP_METHOD_GET, requestPath, "", &m.env)
}

// GetMarketHandicapKline1m 盘口K线
func (m *MarketCtrl) GetMarketHandicapKline1m() {
	requestURL := fmt.Sprintf(m.env.Url+consts.MARKET_HANDICAP_KLINE1M+"?instId=%s&startTime=%d&endTime=%d&limit=%d", "BTC-USDT-SWAP", 1700000000, 1700003600, 60)
	requestPath := fmt.Sprintf(consts.MARKET_HANDICAP_KLINE1M+"?instId=%s&startTime=%d&endTime=%d&limit=%d", "BTC-USDT-SWAP", 1700000000, 1700003600, 60)
	signature.DoHttp(requestURL, consts.HTTP_METHOD_GET, requestPath, "", &m.env)
}

// GetMarketHandicapOrderbook 盘口订单簿
func (m *MarketCtrl) GetMarketHandicapOrderbook() {
	requestURL := fmt.Sprintf(m.env.Url+consts.MARKET_HANDICAP_ORDERBOOK+"?instId=%s&startTime=%d&endTime=%d&limit=%d", "BTC-USDT-SWAP", 1700000000, 1700003600, 60)
	requestPath := fmt.Sprintf(consts.MARKET_HANDICAP_ORDERBOOK+"?instId=%s&startTime=%d&endTime=%d&limit=%d", "BTC-USDT-SWAP", 1700000000, 1700003600, 60)
	signature.DoHttp(requestURL, consts.HTTP_METHOD_GET, requestPath, "", &m.env)
}

// GetMarketHandicapTrade 盘口历史成交记录
func (m *MarketCtrl) GetMarketHandicapTrade() {
	requestURL := fmt.Sprintf(m.env.Url+consts.MARKET_HANDICAP_TRADE+"?instId=%s&startTime=%d&endTime=%d&limit=%d", "BTC-USDT-SWAP", 1700000000, 1700003600, 60)
	requestPath := fmt.Sprintf(consts.MARKET_HANDICAP_TRADE+"?instId=%s&startTime=%d&endTime=%d&limit=%d", "BTC-USDT-SWAP", 1700000000, 1700003600, 60)
	signature.DoHttp(requestURL, consts.HTTP_METHOD_GET, requestPath, "", &m.env)
}

// GetMarketFundingRate 资金费率
func (m *MarketCtrl) GetMarketFundingRate() {
	requestURL := fmt.Sprintf(m.env.Url+consts.MARKET_FUNDING_RATE+"?instId=%s", "BTC-USDT-SWAP")
	requestPath := fmt.Sprintf(consts.MARKET_FUNDING_RATE+"?instId=%s", "BTC-USDT-SWAP")
	signature.DoHttp(requestURL, consts.HTTP_METHOD_GET, requestPath, "", &m.env)
}

// GetMarketCurrentFundingRate 当前资金费率
func (m *MarketCtrl) GetMarketCurrentFundingRate() {
	requestURL := fmt.Sprintf(m.env.Url+consts.MARKET_CURRENT_FUNDING_RATE+"?instId=%s", "BTCUSDT")
	requestPath := fmt.Sprintf(consts.MARKET_CURRENT_FUNDING_RATE+"?instId=%s", "BTCUSDT")
	signature.DoHttp(requestURL, consts.HTTP_METHOD_GET, requestPath, "", &m.env)
}

// GetMarketFundingRateHistory 资金费率历史
func (m *MarketCtrl) GetMarketFundingRateHistory() {
	requestURL := fmt.Sprintf(m.env.Url+consts.MARKET_FUNDING_RATE_HISTORY+"?instId=%s&page=%d&size=%d", "BTC-USDT-SWAP", 1, 20)
	requestPath := fmt.Sprintf(consts.MARKET_FUNDING_RATE_HISTORY+"?instId=%s&page=%d&size=%d", "BTC-USDT-SWAP", 1, 20)
	signature.DoHttp(requestURL, consts.HTTP_METHOD_GET, requestPath, "", &m.env)
}
