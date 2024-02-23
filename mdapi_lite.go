package goctp

import "github.com/pseudocodes/goctp/thost"

type MdApiLite struct {
	*mdApi
}

func CreateMdApiLite(options ...MdOption) *MdApiLite {
	api := &MdApiLite{}
	mdApi := CreateMdApi(options...)
	api.mdApi = mdApi
	return api
}

// 用户登录请求
func (t *MdApiLite) ReqUserLogin(pReqUserLoginField *ReqUserLoginField, nRequestID int) int {
	var f0 = toCThostFtdcReqUserLoginField(pReqUserLoginField)

	return t.mdApi.ReqUserLogin((f0), int(nRequestID))
}

type MdSpiLite struct {
	BaseMdSpi

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

	//登录请求响应
	OnRspUserLoginCallback func(pRspUserLogin *RspUserLoginField, pRspInfo *RspInfoField, nRequestID int, bIsLast bool)

	//错误应答
	OnRspErrorCallback func(pRspInfo *RspInfoField, nRequestID int, bIsLast bool)

	//深度行情通知
	OnRtnDepthMarketDataCallback func(pDepthMarketData *DepthMarketDataField)

	// 订阅行情应答
	OnRspSubMarketDataCallback func(specificInstrument string, pRspInfo *RspInfoField, nRequestID int, bIsLast bool)
}

// 当客户端与交易后台建立起通信连接时（还未登录前），该方法被调用。
func (s *MdSpiLite) OnFrontConnected() {

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
func (s *MdSpiLite) OnFrontDisconnected(nReason int) {

	if s.OnFrontDisconnectedCallback != nil {
		s.OnFrontDisconnectedCallback(int(nReason))
	}
}

// 心跳超时警告。当长时间未收到报文时，该方法被调用。
// /@param nTimeLapse 距离上次接收报文的时间
func (s *MdSpiLite) OnHeartBeatWarning(nTimeLapse int) {

	if s.OnHeartBeatWarningCallback != nil {
		s.OnHeartBeatWarningCallback(int(nTimeLapse))
	}
}

// 登录请求响应
func (s *MdSpiLite) OnRspUserLogin(pRspUserLogin *thost.CThostFtdcRspUserLoginField, pRspInfo *thost.CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
	f0 := fromCThostFtdcRspUserLoginField(pRspUserLogin)

	var f1 *RspInfoField
	if pRspInfo != nil {
		f1 = fromCThostFtdcRspInfoField(pRspInfo)
	}

	if s.OnRspUserLoginCallback != nil {
		s.OnRspUserLoginCallback((f0), (f1), int(nRequestID), bool(bIsLast))
	}
}

// 错误应答
func (s *MdSpiLite) OnRspError(pRspInfo *thost.CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {

	var f0 *RspInfoField
	if pRspInfo != nil {
		f0 = fromCThostFtdcRspInfoField(pRspInfo)
	}

	if s.OnRspErrorCallback != nil {
		s.OnRspErrorCallback((f0), int(nRequestID), bool(bIsLast))
	}
}

// 深度行情通知
func (s *MdSpiLite) OnRtnDepthMarketData(pDepthMarketData *thost.CThostFtdcDepthMarketDataField) {
	f0 := fromCThostFtdcDepthMarketDataField(pDepthMarketData)

	if s.OnRtnDepthMarketDataCallback != nil {
		s.OnRtnDepthMarketDataCallback((f0))
	}
}

// 深度行情通知
func (s *MdSpiLite) OnRspSubMarketData(pSpecificInstrument *thost.CThostFtdcSpecificInstrumentField, pRspInfo *thost.CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
	var f0 = bytes2String(pSpecificInstrument.InstrumentID[:])
	var f1 *RspInfoField
	if pRspInfo != nil {
		f1 = fromCThostFtdcRspInfoField(pRspInfo)
	}

	if s.OnRspSubMarketDataCallback != nil {
		s.OnRspSubMarketDataCallback(f0, f1, nRequestID, bIsLast)
	}
}

func (s *MdSpiLite) SetOnFrontConnected(f func()) {
	s.OnFrontConnectedCallback = f
}

func (s *MdSpiLite) SetOnFrontDisconnected(f func(int)) {
	s.OnFrontDisconnectedCallback = f
}

func (s *MdSpiLite) SetOnHeartBeatWarning(f func(int)) {
	s.OnHeartBeatWarningCallback = f
}

func (s *MdSpiLite) SetOnRspUserLogin(f func(*RspUserLoginField, *RspInfoField, int, bool)) {
	s.OnRspUserLoginCallback = f
}

func (s *MdSpiLite) SetOnRspError(f func(*RspInfoField, int, bool)) {
	s.OnRspErrorCallback = f
}

func (s *MdSpiLite) SetOnRtnDepthMarketData(f func(*DepthMarketDataField)) {
	s.OnRtnDepthMarketDataCallback = f
}

func (s *MdSpiLite) SetOnRspSubMarketData(f func(string, *RspInfoField, int, bool)) {
	s.OnRspSubMarketDataCallback = f
}
