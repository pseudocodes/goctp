package goctp

import (
	"github.com/pseudocodes/goctp/thost"
)

type TraderApiLite struct {
	*traderApi
}

func CreateTraderApiLite(options ...TraderOption) *TraderApiLite {
	api := &TraderApiLite{}
	traderApi := CreateTraderApi(options...)
	api.traderApi = traderApi
	return api
}

// 客户端认证请求
func (t *TraderApiLite) ReqAuthenticate(pReqAuthenticateField *ReqAuthenticateField, nRequestID int) int {
	var f0 = toCThostFtdcReqAuthenticateField(pReqAuthenticateField)

	return t.traderApi.ReqAuthenticate((f0), int(nRequestID))
}

// 用户登录请求
func (t *TraderApiLite) ReqUserLogin(pReqUserLoginField *ReqUserLoginField, nRequestID int) int {
	var f0 = toCThostFtdcReqUserLoginField(pReqUserLoginField)

	return t.traderApi.ReqUserLogin((f0), int(nRequestID))
}

// 用户口令更新请求
func (t *TraderApiLite) ReqUserPasswordUpdate(pUserPasswordUpdate *UserPasswordUpdateField, nRequestID int) int {
	var f0 = toCThostFtdcUserPasswordUpdateField(pUserPasswordUpdate)

	return t.traderApi.ReqUserPasswordUpdate((f0), int(nRequestID))
}

// 资金账户口令更新请求
func (t *TraderApiLite) ReqTradingAccountPasswordUpdate(pTradingAccountPasswordUpdate *TradingAccountPasswordUpdateField, nRequestID int) int {
	var f0 = toCThostFtdcTradingAccountPasswordUpdateField(pTradingAccountPasswordUpdate)

	return t.traderApi.ReqTradingAccountPasswordUpdate((f0), int(nRequestID))
}

// 报单录入请求
func (t *TraderApiLite) ReqOrderInsert(pInputOrder *InputOrderField, nRequestID int) int {
	var f0 = toCThostFtdcInputOrderField(pInputOrder)

	return t.traderApi.ReqOrderInsert((f0), int(nRequestID))
}

// 报单操作请求
func (t *TraderApiLite) ReqOrderAction(pInputOrderAction *InputOrderActionField, nRequestID int) int {
	var f0 = toCThostFtdcInputOrderActionField(pInputOrderAction)

	return t.traderApi.ReqOrderAction((f0), int(nRequestID))
}

// 投资者结算结果确认
func (t *TraderApiLite) ReqSettlementInfoConfirm(pSettlementInfoConfirm *SettlementInfoConfirmField, nRequestID int) int {
	var f0 = toCThostFtdcSettlementInfoConfirmField(pSettlementInfoConfirm)

	return t.traderApi.ReqSettlementInfoConfirm((f0), int(nRequestID))
}

// 请求查询报单
func (t *TraderApiLite) ReqQryOrder(pQryOrder *QryOrderField, nRequestID int) int {
	var f0 = toCThostFtdcQryOrderField(pQryOrder)

	return t.traderApi.ReqQryOrder((f0), int(nRequestID))
}

// 请求查询成交
func (t *TraderApiLite) ReqQryTrade(pQryTrade *QryTradeField, nRequestID int) int {
	var f0 = toCThostFtdcQryTradeField(pQryTrade)

	return t.traderApi.ReqQryTrade((f0), int(nRequestID))
}

// 请求查询投资者持仓
func (t *TraderApiLite) ReqQryInvestorPosition(pQryInvestorPosition *QryInvestorPositionField, nRequestID int) int {
	var f0 = toCThostFtdcQryInvestorPositionField(pQryInvestorPosition)

	return t.traderApi.ReqQryInvestorPosition((f0), int(nRequestID))
}

// 请求查询资金账户
func (t *TraderApiLite) ReqQryTradingAccount(pQryTradingAccount *QryTradingAccountField, nRequestID int) int {
	var f0 = toCThostFtdcQryTradingAccountField(pQryTradingAccount)

	return t.traderApi.ReqQryTradingAccount((f0), int(nRequestID))
}

// 请求查询合约保证金率
func (t *TraderApiLite) ReqQryInstrumentMarginRate(pQryInstrumentMarginRate *QryInstrumentMarginRateField, nRequestID int) int {
	var f0 = toCThostFtdcQryInstrumentMarginRateField(pQryInstrumentMarginRate)

	return t.traderApi.ReqQryInstrumentMarginRate((f0), int(nRequestID))
}

// 请求查询合约手续费率
func (t *TraderApiLite) ReqQryInstrumentCommissionRate(pQryInstrumentCommissionRate *QryInstrumentCommissionRateField, nRequestID int) int {
	var f0 = toCThostFtdcQryInstrumentCommissionRateField(pQryInstrumentCommissionRate)

	return t.traderApi.ReqQryInstrumentCommissionRate((f0), int(nRequestID))
}

// 请求查询合约
func (t *TraderApiLite) ReqQryInstrument(pQryInstrument *QryInstrumentField, nRequestID int) int {
	var f0 = toCThostFtdcQryInstrumentField(pQryInstrument)

	return t.traderApi.ReqQryInstrument((f0), int(nRequestID))
}

// 请求查询投资者结算结果
func (t *TraderApiLite) ReqQrySettlementInfo(pQrySettlementInfo *QrySettlementInfoField, nRequestID int) int {
	var f0 = toCThostFtdcQrySettlementInfoField(pQrySettlementInfo)

	return t.traderApi.ReqQrySettlementInfo((f0), int(nRequestID))
}

// 请求查询结算信息确认
func (t *TraderApiLite) ReqQrySettlementInfoConfirm(pQrySettlementInfoConfirm *QrySettlementInfoConfirmField, nRequestID int) int {
	var f0 = toCThostFtdcQrySettlementInfoConfirmField(pQrySettlementInfoConfirm)

	return t.traderApi.ReqQrySettlementInfoConfirm((f0), int(nRequestID))
}

// 请求查询转帐流水
func (t *TraderApiLite) ReqQryTransferSerial(pQryTransferSerial *QryTransferSerialField, nRequestID int) int {
	var f0 = toCThostFtdcQryTransferSerialField(pQryTransferSerial)

	return t.traderApi.ReqQryTransferSerial((f0), int(nRequestID))
}

// 请求查询银期签约关系
func (t *TraderApiLite) ReqQryAccountregister(pQryAccountregister *QryAccountregisterField, nRequestID int) int {
	var f0 = toCThostFtdcQryAccountregisterField(pQryAccountregister)

	return t.traderApi.ReqQryAccountregister((f0), int(nRequestID))
}

// 请求查询签约银行
func (t *TraderApiLite) ReqQryContractBank(pQryContractBank *QryContractBankField, nRequestID int) int {
	var f0 = toCThostFtdcQryContractBankField(pQryContractBank)

	return t.traderApi.ReqQryContractBank((f0), int(nRequestID))
}

// 请求查询经纪公司交易参数
func (t *TraderApiLite) ReqQryBrokerTradingParams(pQryBrokerTradingParams *QryBrokerTradingParamsField, nRequestID int) int {
	var f0 = toCThostFtdcQryBrokerTradingParamsField(pQryBrokerTradingParams)

	return t.traderApi.ReqQryBrokerTradingParams((f0), int(nRequestID))
}

// 期货发起银行资金转期货请求
func (t *TraderApiLite) ReqFromBankToFutureByFuture(pReqTransfer *ReqTransferField, nRequestID int) int {
	var f0 = toCThostFtdcReqTransferField(pReqTransfer)

	return t.traderApi.ReqFromBankToFutureByFuture((f0), int(nRequestID))
}

// 期货发起期货资金转银行请求
func (t *TraderApiLite) ReqFromFutureToBankByFuture(pReqTransfer *ReqTransferField, nRequestID int) int {
	var f0 = toCThostFtdcReqTransferField(pReqTransfer)

	return t.traderApi.ReqFromFutureToBankByFuture((f0), int(nRequestID))
}

type TraderSpiLite struct {
	BaseTraderSpi

	//当客户端与交易后台建立起通信连接时（还未登录前），该方法被调用。
	OnFrontConnectedCallback func()

	//当客户端与交易后台通信连接断开时，该方法被调用。当发生这个情况后，API会自动重新连接，客户端可不做处理。
	///@param nReason 错误原因
	///        0x1001 网络读失败
	///        0x1002 网络写失败
	///        0x2001 接收心跳超时
	///        0x2002 发送心跳失败
	///        0x2003 收到错误报文
	OnFrontDisconnectedCallback func(nReason int)

	//心跳超时警告。当长时间未收到报文时，该方法被调用。
	///@param nTimeLapse 距离上次接收报文的时间
	OnHeartBeatWarningCallback func(nTimeLapse int)

	//客户端认证响应
	OnRspAuthenticateCallback func(pRspAuthenticateField *RspAuthenticateField, pRspInfo *RspInfoField, nRequestID int, bIsLast bool)

	//登录请求响应
	OnRspUserLoginCallback func(pRspUserLogin *RspUserLoginField, pRspInfo *RspInfoField, nRequestID int, bIsLast bool)

	//用户口令更新请求响应
	OnRspUserPasswordUpdateCallback func(pUserPasswordUpdate *UserPasswordUpdateField, pRspInfo *RspInfoField, nRequestID int, bIsLast bool)

	//资金账户口令更新请求响应
	OnRspTradingAccountPasswordUpdateCallback func(pTradingAccountPasswordUpdate *TradingAccountPasswordUpdateField, pRspInfo *RspInfoField, nRequestID int, bIsLast bool)

	//报单录入请求响应
	OnRspOrderInsertCallback func(pInputOrder *InputOrderField, pRspInfo *RspInfoField, nRequestID int, bIsLast bool)

	//报单操作请求响应
	OnRspOrderActionCallback func(pInputOrderAction *InputOrderActionField, pRspInfo *RspInfoField, nRequestID int, bIsLast bool)

	//投资者结算结果确认响应
	OnRspSettlementInfoConfirmCallback func(pSettlementInfoConfirm *SettlementInfoConfirmField, pRspInfo *RspInfoField, nRequestID int, bIsLast bool)

	//请求查询报单响应
	OnRspQryOrderCallback func(pOrder *OrderField, pRspInfo *RspInfoField, nRequestID int, bIsLast bool)

	//请求查询投资者持仓响应
	OnRspQryInvestorPositionCallback func(pInvestorPosition *InvestorPositionField, pRspInfo *RspInfoField, nRequestID int, bIsLast bool)

	//请求查询资金账户响应
	OnRspQryTradingAccountCallback func(pTradingAccount *TradingAccountField, pRspInfo *RspInfoField, nRequestID int, bIsLast bool)

	//请求查询合约保证金率响应
	OnRspQryInstrumentMarginRateCallback func(pInstrumentMarginRate *InstrumentMarginRateField, pRspInfo *RspInfoField, nRequestID int, bIsLast bool)

	//请求查询合约手续费率响应
	OnRspQryInstrumentCommissionRateCallback func(pInstrumentCommissionRate *InstrumentCommissionRateField, pRspInfo *RspInfoField, nRequestID int, bIsLast bool)

	//请求查询合约响应
	OnRspQryInstrumentCallback func(pInstrument *InstrumentField, pRspInfo *RspInfoField, nRequestID int, bIsLast bool)

	//请求查询行情响应
	OnRspQryDepthMarketDataCallback func(pDepthMarketData *DepthMarketDataField, pRspInfo *RspInfoField, nRequestID int, bIsLast bool)

	//请求查询投资者结算结果响应
	OnRspQrySettlementInfoCallback func(pSettlementInfo *SettlementInfoField, pRspInfo *RspInfoField, nRequestID int, bIsLast bool)

	//请求查询转帐流水响应
	OnRspQryTransferSerialCallback func(pTransferSerial *TransferSerialField, pRspInfo *RspInfoField, nRequestID int, bIsLast bool)

	//请求查询银期签约关系响应
	OnRspQryAccountregisterCallback func(pAccountregister *AccountregisterField, pRspInfo *RspInfoField, nRequestID int, bIsLast bool)

	//错误应答
	OnRspErrorCallback func(pRspInfo *RspInfoField, nRequestID int, bIsLast bool)

	//报单通知
	OnRtnOrderCallback func(pOrder *OrderField)

	//成交通知
	OnRtnTradeCallback func(pTrade *TradeField)

	//报单录入错误回报
	OnErrRtnOrderInsertCallback func(pInputOrder *InputOrderField, pRspInfo *RspInfoField)

	//报单操作错误回报
	OnErrRtnOrderActionCallback func(pOrderAction *OrderActionField, pRspInfo *RspInfoField)

	//合约交易状态通知
	OnRtnInstrumentStatusCallback func(pInstrumentStatus *InstrumentStatusField)

	//交易通知
	OnRtnTradingNoticeCallback func(pTradingNoticeInfo *TradingNoticeInfoField)

	//请求查询签约银行响应
	OnRspQryContractBankCallback func(pContractBank *ContractBankField, pRspInfo *RspInfoField, nRequestID int, bIsLast bool)

	//请求查询经纪公司交易参数响应
	OnRspQryBrokerTradingParamsCallback func(pBrokerTradingParams *BrokerTradingParamsField, pRspInfo *RspInfoField, nRequestID int, bIsLast bool)

	//期货发起银行资金转期货通知
	OnRtnFromBankToFutureByFutureCallback func(pRspTransfer *RspTransferField)

	//期货发起期货资金转银行通知
	OnRtnFromFutureToBankByFutureCallback func(pRspTransfer *RspTransferField)
}

// 当客户端与交易后台建立起通信连接时（还未登录前），该方法被调用。
func (s *TraderSpiLite) OnFrontConnected() {

	if s.OnFrontConnectedCallback != nil {
		s.OnFrontConnectedCallback()
	}
}

// 当客户端与交易后台通信连接断开时，该方法被调用。当发生这个情况后，API会自动重新连接，客户端可不做处理。
// /@param nReason 错误原因
// /        0x1001 网络读失败
// /        0x1002 网络写失败
// /        0x2001 接收心跳超时
// /        0x2002 发送心跳失败
// /        0x2003 收到错误报文
func (s *TraderSpiLite) OnFrontDisconnected(nReason int) {

	if s.OnFrontDisconnectedCallback != nil {
		s.OnFrontDisconnectedCallback(int(nReason))
	}
}

// 心跳超时警告。当长时间未收到报文时，该方法被调用。
// /@param nTimeLapse 距离上次接收报文的时间
func (s *TraderSpiLite) OnHeartBeatWarning(nTimeLapse int) {

	if s.OnHeartBeatWarningCallback != nil {
		s.OnHeartBeatWarningCallback(int(nTimeLapse))
	}
}

// 客户端认证响应
func (s *TraderSpiLite) OnRspAuthenticate(pRspAuthenticateField *thost.CThostFtdcRspAuthenticateField, pRspInfo *thost.CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {

	var f0 *RspAuthenticateField
	if pRspAuthenticateField != nil {
		f0 = fromCThostFtdcRspAuthenticateField(pRspAuthenticateField)
	}

	var f1 *RspInfoField
	if pRspInfo != nil {
		f1 = fromCThostFtdcRspInfoField(pRspInfo)
	}

	if s.OnRspAuthenticateCallback != nil {
		s.OnRspAuthenticateCallback((f0), (f1), int(nRequestID), bool(bIsLast))
	}
}

// 登录请求响应
func (s *TraderSpiLite) OnRspUserLogin(pRspUserLogin *thost.CThostFtdcRspUserLoginField, pRspInfo *thost.CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {

	var f0 *RspUserLoginField
	if pRspUserLogin != nil {
		f0 = fromCThostFtdcRspUserLoginField(pRspUserLogin)
	}

	var f1 *RspInfoField
	if pRspInfo != nil {
		f1 = fromCThostFtdcRspInfoField(pRspInfo)
	}

	if s.OnRspUserLoginCallback != nil {
		s.OnRspUserLoginCallback((f0), (f1), int(nRequestID), bool(bIsLast))
	}
}

// 用户口令更新请求响应
func (s *TraderSpiLite) OnRspUserPasswordUpdate(pUserPasswordUpdate *thost.CThostFtdcUserPasswordUpdateField, pRspInfo *thost.CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {

	var f0 *UserPasswordUpdateField
	if pUserPasswordUpdate != nil {
		f0 = fromCThostFtdcUserPasswordUpdateField(pUserPasswordUpdate)
	}

	var f1 *RspInfoField
	if pRspInfo != nil {
		f1 = fromCThostFtdcRspInfoField(pRspInfo)
	}

	if s.OnRspUserPasswordUpdateCallback != nil {
		s.OnRspUserPasswordUpdateCallback((f0), (f1), int(nRequestID), bool(bIsLast))
	}
}

// 资金账户口令更新请求响应
func (s *TraderSpiLite) OnRspTradingAccountPasswordUpdate(pTradingAccountPasswordUpdate *thost.CThostFtdcTradingAccountPasswordUpdateField, pRspInfo *thost.CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {

	var f0 *TradingAccountPasswordUpdateField
	if pTradingAccountPasswordUpdate != nil {
		f0 = fromCThostFtdcTradingAccountPasswordUpdateField(pTradingAccountPasswordUpdate)
	}

	var f1 *RspInfoField
	if pRspInfo != nil {
		f1 = fromCThostFtdcRspInfoField(pRspInfo)
	}

	if s.OnRspTradingAccountPasswordUpdateCallback != nil {
		s.OnRspTradingAccountPasswordUpdateCallback((f0), (f1), int(nRequestID), bool(bIsLast))
	}
}

// 报单录入请求响应
func (s *TraderSpiLite) OnRspOrderInsert(pInputOrder *thost.CThostFtdcInputOrderField, pRspInfo *thost.CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {

	var f0 *InputOrderField
	if pInputOrder != nil {
		f0 = fromCThostFtdcInputOrderField(pInputOrder)
	}

	var f1 *RspInfoField
	if pRspInfo != nil {
		f1 = fromCThostFtdcRspInfoField(pRspInfo)
	}

	if s.OnRspOrderInsertCallback != nil {
		s.OnRspOrderInsertCallback((f0), (f1), int(nRequestID), bool(bIsLast))
	}
}

// 报单操作请求响应
func (s *TraderSpiLite) OnRspOrderAction(pInputOrderAction *thost.CThostFtdcInputOrderActionField, pRspInfo *thost.CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {

	var f0 *InputOrderActionField
	if pInputOrderAction != nil {
		f0 = fromCThostFtdcInputOrderActionField(pInputOrderAction)
	}

	var f1 *RspInfoField
	if pRspInfo != nil {
		f1 = fromCThostFtdcRspInfoField(pRspInfo)
	}

	if s.OnRspOrderActionCallback != nil {
		s.OnRspOrderActionCallback((f0), (f1), int(nRequestID), bool(bIsLast))
	}
}

// 投资者结算结果确认响应
func (s *TraderSpiLite) OnRspSettlementInfoConfirm(pSettlementInfoConfirm *thost.CThostFtdcSettlementInfoConfirmField, pRspInfo *thost.CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {

	var f0 *SettlementInfoConfirmField
	if pSettlementInfoConfirm != nil {
		f0 = fromCThostFtdcSettlementInfoConfirmField(pSettlementInfoConfirm)
	}

	var f1 *RspInfoField
	if pRspInfo != nil {
		f1 = fromCThostFtdcRspInfoField(pRspInfo)
	}

	if s.OnRspSettlementInfoConfirmCallback != nil {
		s.OnRspSettlementInfoConfirmCallback((f0), (f1), int(nRequestID), bool(bIsLast))
	}
}

// 请求查询报单响应
func (s *TraderSpiLite) OnRspQryOrder(pOrder *thost.CThostFtdcOrderField, pRspInfo *thost.CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {

	var f0 *OrderField
	if pOrder != nil {
		f0 = fromCThostFtdcOrderField(pOrder)
	}

	var f1 *RspInfoField
	if pRspInfo != nil {
		f1 = fromCThostFtdcRspInfoField(pRspInfo)
	}

	if s.OnRspQryOrderCallback != nil {
		s.OnRspQryOrderCallback((f0), (f1), int(nRequestID), bool(bIsLast))
	}
}

// 请求查询投资者持仓响应
func (s *TraderSpiLite) OnRspQryInvestorPosition(pInvestorPosition *thost.CThostFtdcInvestorPositionField, pRspInfo *thost.CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {

	var f0 *InvestorPositionField
	if pInvestorPosition != nil {
		f0 = fromCThostFtdcInvestorPositionField(pInvestorPosition)
	}

	var f1 *RspInfoField
	if pRspInfo != nil {
		f1 = fromCThostFtdcRspInfoField(pRspInfo)
	}

	if s.OnRspQryInvestorPositionCallback != nil {
		s.OnRspQryInvestorPositionCallback((f0), (f1), int(nRequestID), bool(bIsLast))
	}
}

// 请求查询资金账户响应
func (s *TraderSpiLite) OnRspQryTradingAccount(pTradingAccount *thost.CThostFtdcTradingAccountField, pRspInfo *thost.CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {

	var f0 *TradingAccountField
	if pTradingAccount != nil {
		f0 = fromCThostFtdcTradingAccountField(pTradingAccount)
	}

	var f1 *RspInfoField
	if pRspInfo != nil {
		f1 = fromCThostFtdcRspInfoField(pRspInfo)
	}

	if s.OnRspQryTradingAccountCallback != nil {
		s.OnRspQryTradingAccountCallback((f0), (f1), int(nRequestID), bool(bIsLast))
	}
}

// 请求查询合约保证金率响应
func (s *TraderSpiLite) OnRspQryInstrumentMarginRate(pInstrumentMarginRate *thost.CThostFtdcInstrumentMarginRateField, pRspInfo *thost.CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {

	var f0 *InstrumentMarginRateField
	if pInstrumentMarginRate != nil {
		f0 = fromCThostFtdcInstrumentMarginRateField(pInstrumentMarginRate)
	}

	var f1 *RspInfoField
	if pRspInfo != nil {
		f1 = fromCThostFtdcRspInfoField(pRspInfo)
	}

	if s.OnRspQryInstrumentMarginRateCallback != nil {
		s.OnRspQryInstrumentMarginRateCallback((f0), (f1), int(nRequestID), bool(bIsLast))
	}
}

// 请求查询合约手续费率响应
func (s *TraderSpiLite) OnRspQryInstrumentCommissionRate(pInstrumentCommissionRate *thost.CThostFtdcInstrumentCommissionRateField, pRspInfo *thost.CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {

	var f0 *InstrumentCommissionRateField
	if pInstrumentCommissionRate != nil {
		f0 = fromCThostFtdcInstrumentCommissionRateField(pInstrumentCommissionRate)
	}

	var f1 *RspInfoField
	if pRspInfo != nil {
		f1 = fromCThostFtdcRspInfoField(pRspInfo)
	}

	if s.OnRspQryInstrumentCommissionRateCallback != nil {
		s.OnRspQryInstrumentCommissionRateCallback((f0), (f1), int(nRequestID), bool(bIsLast))
	}
}

// 请求查询合约响应
func (s *TraderSpiLite) OnRspQryInstrument(pInstrument *thost.CThostFtdcInstrumentField, pRspInfo *thost.CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {

	var f0 *InstrumentField
	if pInstrument != nil {
		f0 = fromCThostFtdcInstrumentField(pInstrument)
	}

	var f1 *RspInfoField
	if pRspInfo != nil {
		f1 = fromCThostFtdcRspInfoField(pRspInfo)
	}

	if s.OnRspQryInstrumentCallback != nil {
		s.OnRspQryInstrumentCallback((f0), (f1), int(nRequestID), bool(bIsLast))
	}
}

// 请求查询行情响应
func (s *TraderSpiLite) OnRspQryDepthMarketData(pDepthMarketData *thost.CThostFtdcDepthMarketDataField, pRspInfo *thost.CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {

	var f0 *DepthMarketDataField
	if pDepthMarketData != nil {
		f0 = fromCThostFtdcDepthMarketDataField(pDepthMarketData)
	}

	var f1 *RspInfoField
	if pRspInfo != nil {
		f1 = fromCThostFtdcRspInfoField(pRspInfo)
	}

	if s.OnRspQryDepthMarketDataCallback != nil {
		s.OnRspQryDepthMarketDataCallback((f0), (f1), int(nRequestID), bool(bIsLast))
	}
}

// 请求查询投资者结算结果响应
func (s *TraderSpiLite) OnRspQrySettlementInfo(pSettlementInfo *thost.CThostFtdcSettlementInfoField, pRspInfo *thost.CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {

	var f0 *SettlementInfoField
	if pSettlementInfo != nil {
		f0 = fromCThostFtdcSettlementInfoField(pSettlementInfo)
	}

	var f1 *RspInfoField
	if pRspInfo != nil {
		f1 = fromCThostFtdcRspInfoField(pRspInfo)
	}

	if s.OnRspQrySettlementInfoCallback != nil {
		s.OnRspQrySettlementInfoCallback((f0), (f1), int(nRequestID), bool(bIsLast))
	}
}

// 请求查询转帐流水响应
func (s *TraderSpiLite) OnRspQryTransferSerial(pTransferSerial *thost.CThostFtdcTransferSerialField, pRspInfo *thost.CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {

	var f0 *TransferSerialField
	if pTransferSerial != nil {
		f0 = fromCThostFtdcTransferSerialField(pTransferSerial)
	}

	var f1 *RspInfoField
	if pRspInfo != nil {
		f1 = fromCThostFtdcRspInfoField(pRspInfo)
	}

	if s.OnRspQryTransferSerialCallback != nil {
		s.OnRspQryTransferSerialCallback((f0), (f1), int(nRequestID), bool(bIsLast))
	}
}

// 请求查询银期签约关系响应
func (s *TraderSpiLite) OnRspQryAccountregister(pAccountregister *thost.CThostFtdcAccountregisterField, pRspInfo *thost.CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {

	var f0 *AccountregisterField
	if pAccountregister != nil {
		f0 = fromCThostFtdcAccountregisterField(pAccountregister)
	}

	var f1 *RspInfoField
	if pRspInfo != nil {
		f1 = fromCThostFtdcRspInfoField(pRspInfo)
	}

	if s.OnRspQryAccountregisterCallback != nil {
		s.OnRspQryAccountregisterCallback((f0), (f1), int(nRequestID), bool(bIsLast))
	}
}

// 错误应答
func (s *TraderSpiLite) OnRspError(pRspInfo *thost.CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {

	var f0 *RspInfoField
	if pRspInfo != nil {
		f0 = fromCThostFtdcRspInfoField(pRspInfo)
	}

	if s.OnRspErrorCallback != nil {
		s.OnRspErrorCallback((f0), int(nRequestID), bool(bIsLast))
	}
}

// 报单通知
func (s *TraderSpiLite) OnRtnOrder(pOrder *thost.CThostFtdcOrderField) {

	var f0 *OrderField
	if pOrder != nil {
		f0 = fromCThostFtdcOrderField(pOrder)
	}

	if s.OnRtnOrderCallback != nil {
		s.OnRtnOrderCallback((f0))
	}
}

// 成交通知
func (s *TraderSpiLite) OnRtnTrade(pTrade *thost.CThostFtdcTradeField) {

	var f0 *TradeField
	if pTrade != nil {
		f0 = fromCThostFtdcTradeField(pTrade)
	}

	if s.OnRtnTradeCallback != nil {
		s.OnRtnTradeCallback((f0))
	}
}

// 报单录入错误回报
func (s *TraderSpiLite) OnErrRtnOrderInsert(pInputOrder *thost.CThostFtdcInputOrderField, pRspInfo *thost.CThostFtdcRspInfoField) {

	var f0 *InputOrderField
	if pInputOrder != nil {
		f0 = fromCThostFtdcInputOrderField(pInputOrder)
	}

	var f1 *RspInfoField
	if pRspInfo != nil {
		f1 = fromCThostFtdcRspInfoField(pRspInfo)
	}

	if s.OnErrRtnOrderInsertCallback != nil {
		s.OnErrRtnOrderInsertCallback((f0), (f1))
	}
}

// 报单操作错误回报
func (s *TraderSpiLite) OnErrRtnOrderAction(pOrderAction *thost.CThostFtdcOrderActionField, pRspInfo *thost.CThostFtdcRspInfoField) {

	var f0 *OrderActionField
	if pOrderAction != nil {
		f0 = fromCThostFtdcOrderActionField(pOrderAction)
	}

	var f1 *RspInfoField
	if pRspInfo != nil {
		f1 = fromCThostFtdcRspInfoField(pRspInfo)
	}

	if s.OnErrRtnOrderActionCallback != nil {
		s.OnErrRtnOrderActionCallback((f0), (f1))
	}
}

// 合约交易状态通知
func (s *TraderSpiLite) OnRtnInstrumentStatus(pInstrumentStatus *thost.CThostFtdcInstrumentStatusField) {

	var f0 *InstrumentStatusField
	if pInstrumentStatus != nil {
		f0 = fromCThostFtdcInstrumentStatusField(pInstrumentStatus)
	}

	if s.OnRtnInstrumentStatusCallback != nil {
		s.OnRtnInstrumentStatusCallback((f0))
	}
}

// 交易通知
func (s *TraderSpiLite) OnRtnTradingNotice(pTradingNoticeInfo *thost.CThostFtdcTradingNoticeInfoField) {

	var f0 *TradingNoticeInfoField
	if pTradingNoticeInfo != nil {
		f0 = fromCThostFtdcTradingNoticeInfoField(pTradingNoticeInfo)
	}

	if s.OnRtnTradingNoticeCallback != nil {
		s.OnRtnTradingNoticeCallback((f0))
	}
}

// 请求查询签约银行响应
func (s *TraderSpiLite) OnRspQryContractBank(pContractBank *thost.CThostFtdcContractBankField, pRspInfo *thost.CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {

	var f0 *ContractBankField
	if pContractBank != nil {
		f0 = fromCThostFtdcContractBankField(pContractBank)
	}

	var f1 *RspInfoField
	if pRspInfo != nil {
		f1 = fromCThostFtdcRspInfoField(pRspInfo)
	}

	if s.OnRspQryContractBankCallback != nil {
		s.OnRspQryContractBankCallback((f0), (f1), int(nRequestID), bool(bIsLast))
	}
}

// 请求查询经纪公司交易参数响应
func (s *TraderSpiLite) OnRspQryBrokerTradingParams(pBrokerTradingParams *thost.CThostFtdcBrokerTradingParamsField, pRspInfo *thost.CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {

	var f0 *BrokerTradingParamsField
	if pBrokerTradingParams != nil {
		f0 = fromCThostFtdcBrokerTradingParamsField(pBrokerTradingParams)
	}

	var f1 *RspInfoField
	if pRspInfo != nil {
		f1 = fromCThostFtdcRspInfoField(pRspInfo)
	}

	if s.OnRspQryBrokerTradingParamsCallback != nil {
		s.OnRspQryBrokerTradingParamsCallback((f0), (f1), int(nRequestID), bool(bIsLast))
	}
}

// 期货发起银行资金转期货通知
func (s *TraderSpiLite) OnRtnFromBankToFutureByFuture(pRspTransfer *thost.CThostFtdcRspTransferField) {

	var f0 *RspTransferField
	if pRspTransfer != nil {
		f0 = fromCThostFtdcRspTransferField(pRspTransfer)
	}

	if s.OnRtnFromBankToFutureByFutureCallback != nil {
		s.OnRtnFromBankToFutureByFutureCallback((f0))
	}
}

// 期货发起期货资金转银行通知
func (s *TraderSpiLite) OnRtnFromFutureToBankByFuture(pRspTransfer *thost.CThostFtdcRspTransferField) {

	var f0 *RspTransferField
	if pRspTransfer != nil {
		f0 = fromCThostFtdcRspTransferField(pRspTransfer)
	}

	if s.OnRtnFromFutureToBankByFutureCallback != nil {
		s.OnRtnFromFutureToBankByFutureCallback((f0))
	}
}

func (s *TraderSpiLite) SetOnFrontConnected(f func()) {
	s.OnFrontConnectedCallback = f
}

func (s *TraderSpiLite) SetOnFrontDisconnected(f func(int)) {
	s.OnFrontDisconnectedCallback = f
}

func (s *TraderSpiLite) SetOnHeartBeatWarning(f func(int)) {
	s.OnHeartBeatWarningCallback = f
}

func (s *TraderSpiLite) SetOnRspAuthenticate(f func(*RspAuthenticateField, *RspInfoField, int, bool)) {
	s.OnRspAuthenticateCallback = f
}

func (s *TraderSpiLite) SetOnRspUserLogin(f func(*RspUserLoginField, *RspInfoField, int, bool)) {
	s.OnRspUserLoginCallback = f
}

func (s *TraderSpiLite) SetOnRspUserPasswordUpdate(f func(*UserPasswordUpdateField, *RspInfoField, int, bool)) {
	s.OnRspUserPasswordUpdateCallback = f
}

func (s *TraderSpiLite) SetOnRspTradingAccountPasswordUpdate(f func(*TradingAccountPasswordUpdateField, *RspInfoField, int, bool)) {
	s.OnRspTradingAccountPasswordUpdateCallback = f
}

func (s *TraderSpiLite) SetOnRspOrderInsert(f func(*InputOrderField, *RspInfoField, int, bool)) {
	s.OnRspOrderInsertCallback = f
}

func (s *TraderSpiLite) SetOnRspOrderAction(f func(*InputOrderActionField, *RspInfoField, int, bool)) {
	s.OnRspOrderActionCallback = f
}

func (s *TraderSpiLite) SetOnRspSettlementInfoConfirm(f func(*SettlementInfoConfirmField, *RspInfoField, int, bool)) {
	s.OnRspSettlementInfoConfirmCallback = f
}

func (s *TraderSpiLite) SetOnRspQryOrder(f func(*OrderField, *RspInfoField, int, bool)) {
	s.OnRspQryOrderCallback = f
}

func (s *TraderSpiLite) SetOnRspQryInvestorPosition(f func(*InvestorPositionField, *RspInfoField, int, bool)) {
	s.OnRspQryInvestorPositionCallback = f
}

func (s *TraderSpiLite) SetOnRspQryTradingAccount(f func(*TradingAccountField, *RspInfoField, int, bool)) {
	s.OnRspQryTradingAccountCallback = f
}

func (s *TraderSpiLite) SetOnRspQryInstrumentMarginRate(f func(*InstrumentMarginRateField, *RspInfoField, int, bool)) {
	s.OnRspQryInstrumentMarginRateCallback = f
}

func (s *TraderSpiLite) SetOnRspQryInstrumentCommissionRate(f func(*InstrumentCommissionRateField, *RspInfoField, int, bool)) {
	s.OnRspQryInstrumentCommissionRateCallback = f
}

func (s *TraderSpiLite) SetOnRspQryInstrument(f func(*InstrumentField, *RspInfoField, int, bool)) {
	s.OnRspQryInstrumentCallback = f
}

func (s *TraderSpiLite) SetOnRspQryDepthMarketData(f func(*DepthMarketDataField, *RspInfoField, int, bool)) {
	s.OnRspQryDepthMarketDataCallback = f
}

func (s *TraderSpiLite) SetOnRspQrySettlementInfo(f func(*SettlementInfoField, *RspInfoField, int, bool)) {
	s.OnRspQrySettlementInfoCallback = f
}

func (s *TraderSpiLite) SetOnRspQryTransferSerial(f func(*TransferSerialField, *RspInfoField, int, bool)) {
	s.OnRspQryTransferSerialCallback = f
}

func (s *TraderSpiLite) SetOnRspQryAccountregister(f func(*AccountregisterField, *RspInfoField, int, bool)) {
	s.OnRspQryAccountregisterCallback = f
}

func (s *TraderSpiLite) SetOnRspError(f func(*RspInfoField, int, bool)) {
	s.OnRspErrorCallback = f
}

func (s *TraderSpiLite) SetOnRtnOrder(f func(*OrderField)) {
	s.OnRtnOrderCallback = f
}

func (s *TraderSpiLite) SetOnRtnTrade(f func(*TradeField)) {
	s.OnRtnTradeCallback = f
}

func (s *TraderSpiLite) SetOnErrRtnOrderInsert(f func(*InputOrderField, *RspInfoField)) {
	s.OnErrRtnOrderInsertCallback = f
}

func (s *TraderSpiLite) SetOnErrRtnOrderAction(f func(*OrderActionField, *RspInfoField)) {
	s.OnErrRtnOrderActionCallback = f
}

func (s *TraderSpiLite) SetOnRtnInstrumentStatus(f func(*InstrumentStatusField)) {
	s.OnRtnInstrumentStatusCallback = f
}

func (s *TraderSpiLite) SetOnRtnTradingNotice(f func(*TradingNoticeInfoField)) {
	s.OnRtnTradingNoticeCallback = f
}

func (s *TraderSpiLite) SetOnRspQryContractBank(f func(*ContractBankField, *RspInfoField, int, bool)) {
	s.OnRspQryContractBankCallback = f
}

func (s *TraderSpiLite) SetOnRspQryBrokerTradingParams(f func(*BrokerTradingParamsField, *RspInfoField, int, bool)) {
	s.OnRspQryBrokerTradingParamsCallback = f
}

func (s *TraderSpiLite) SetOnRtnFromBankToFutureByFuture(f func(*RspTransferField)) {
	s.OnRtnFromBankToFutureByFutureCallback = f
}

func (s *TraderSpiLite) SetOnRtnFromFutureToBankByFuture(f func(*RspTransferField)) {
	s.OnRtnFromFutureToBankByFutureCallback = f
}
