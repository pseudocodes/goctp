// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package goctp

import (
	"github.com/pseudocodes/goctp/thost"
)

var _ MdSpi = &BaseMdSpi{}

type BaseMdSpi struct {
}

// 当客户端与交易后台建立起通信连接时（还未登录前），该方法被调用。
func (s *BaseMdSpi) OnFrontConnected() {
}

// 当客户端与交易后台通信连接断开时，该方法被调用。当发生这个情况后，API会自动重新连接，客户端可不做处理。
///@param nReason 错误原因
///        0x1001 网络读失败
///        0x1002 网络写失败
///        0x2001 接收心跳超时
///        0x2002 发送心跳失败
///        0x2003 收到错误报文
func (s *BaseMdSpi) OnFrontDisconnected(nReason int) {
}

// 心跳超时警告。当长时间未收到报文时，该方法被调用。
///@param nTimeLapse 距离上次接收报文的时间
func (s *BaseMdSpi) OnHeartBeatWarning(nTimeLapse int) {
}

// 登录请求响应
func (s *BaseMdSpi) OnRspUserLogin(pRspUserLogin *thost.CThostFtdcRspUserLoginField, pRspInfo *thost.CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
}

// 登出请求响应
func (s *BaseMdSpi) OnRspUserLogout(pUserLogout *thost.CThostFtdcUserLogoutField, pRspInfo *thost.CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
}

// 请求查询组播合约响应
func (s *BaseMdSpi) OnRspQryMulticastInstrument(pMulticastInstrument *thost.CThostFtdcMulticastInstrumentField, pRspInfo *thost.CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
}

// 错误应答
func (s *BaseMdSpi) OnRspError(pRspInfo *thost.CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
}

// 订阅行情应答
func (s *BaseMdSpi) OnRspSubMarketData(pSpecificInstrument *thost.CThostFtdcSpecificInstrumentField, pRspInfo *thost.CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
}

// 取消订阅行情应答
func (s *BaseMdSpi) OnRspUnSubMarketData(pSpecificInstrument *thost.CThostFtdcSpecificInstrumentField, pRspInfo *thost.CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
}

// 订阅询价应答
func (s *BaseMdSpi) OnRspSubForQuoteRsp(pSpecificInstrument *thost.CThostFtdcSpecificInstrumentField, pRspInfo *thost.CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
}

// 取消订阅询价应答
func (s *BaseMdSpi) OnRspUnSubForQuoteRsp(pSpecificInstrument *thost.CThostFtdcSpecificInstrumentField, pRspInfo *thost.CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
}

// 深度行情通知
func (s *BaseMdSpi) OnRtnDepthMarketData(pDepthMarketData *thost.CThostFtdcDepthMarketDataField) {
}

// 询价通知
func (s *BaseMdSpi) OnRtnForQuoteRsp(pForQuoteRsp *thost.CThostFtdcForQuoteRspField) {
}
