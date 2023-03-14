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

var _ TraderSpi = &BaseTraderSpi{}

type BaseTraderSpi struct {
}

// 当客户端与交易后台建立起通信连接时（还未登录前），该方法被调用。
func (s *BaseTraderSpi) OnFrontConnected() {
}

// 当客户端与交易后台通信连接断开时，该方法被调用。当发生这个情况后，API会自动重新连接，客户端可不做处理。
///@param nReason 错误原因
///        0x1001 网络读失败
///        0x1002 网络写失败
///        0x2001 接收心跳超时
///        0x2002 发送心跳失败
///        0x2003 收到错误报文
func (s *BaseTraderSpi) OnFrontDisconnected(nReason int) {
}

// 心跳超时警告。当长时间未收到报文时，该方法被调用。
///@param nTimeLapse 距离上次接收报文的时间
func (s *BaseTraderSpi) OnHeartBeatWarning(nTimeLapse int) {
}

// 客户端认证响应
func (s *BaseTraderSpi) OnRspAuthenticate(pRspAuthenticateField *thost.CThostFtdcRspAuthenticateField, pRspInfo *thost.CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
}

// 登录请求响应
func (s *BaseTraderSpi) OnRspUserLogin(pRspUserLogin *thost.CThostFtdcRspUserLoginField, pRspInfo *thost.CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
}

// 登出请求响应
func (s *BaseTraderSpi) OnRspUserLogout(pUserLogout *thost.CThostFtdcUserLogoutField, pRspInfo *thost.CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
}

// 用户口令更新请求响应
func (s *BaseTraderSpi) OnRspUserPasswordUpdate(pUserPasswordUpdate *thost.CThostFtdcUserPasswordUpdateField, pRspInfo *thost.CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
}

// 资金账户口令更新请求响应
func (s *BaseTraderSpi) OnRspTradingAccountPasswordUpdate(pTradingAccountPasswordUpdate *thost.CThostFtdcTradingAccountPasswordUpdateField, pRspInfo *thost.CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
}

// 查询用户当前支持的认证模式的回复
func (s *BaseTraderSpi) OnRspUserAuthMethod(pRspUserAuthMethod *thost.CThostFtdcRspUserAuthMethodField, pRspInfo *thost.CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
}

// 获取图形验证码请求的回复
func (s *BaseTraderSpi) OnRspGenUserCaptcha(pRspGenUserCaptcha *thost.CThostFtdcRspGenUserCaptchaField, pRspInfo *thost.CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
}

// 获取短信验证码请求的回复
func (s *BaseTraderSpi) OnRspGenUserText(pRspGenUserText *thost.CThostFtdcRspGenUserTextField, pRspInfo *thost.CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
}

// 报单录入请求响应
func (s *BaseTraderSpi) OnRspOrderInsert(pInputOrder *thost.CThostFtdcInputOrderField, pRspInfo *thost.CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
}

// 预埋单录入请求响应
func (s *BaseTraderSpi) OnRspParkedOrderInsert(pParkedOrder *thost.CThostFtdcParkedOrderField, pRspInfo *thost.CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
}

// 预埋撤单录入请求响应
func (s *BaseTraderSpi) OnRspParkedOrderAction(pParkedOrderAction *thost.CThostFtdcParkedOrderActionField, pRspInfo *thost.CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
}

// 报单操作请求响应
func (s *BaseTraderSpi) OnRspOrderAction(pInputOrderAction *thost.CThostFtdcInputOrderActionField, pRspInfo *thost.CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
}

// 查询最大报单数量响应
func (s *BaseTraderSpi) OnRspQryMaxOrderVolume(pQryMaxOrderVolume *thost.CThostFtdcQryMaxOrderVolumeField, pRspInfo *thost.CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
}

// 投资者结算结果确认响应
func (s *BaseTraderSpi) OnRspSettlementInfoConfirm(pSettlementInfoConfirm *thost.CThostFtdcSettlementInfoConfirmField, pRspInfo *thost.CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
}

// 删除预埋单响应
func (s *BaseTraderSpi) OnRspRemoveParkedOrder(pRemoveParkedOrder *thost.CThostFtdcRemoveParkedOrderField, pRspInfo *thost.CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
}

// 删除预埋撤单响应
func (s *BaseTraderSpi) OnRspRemoveParkedOrderAction(pRemoveParkedOrderAction *thost.CThostFtdcRemoveParkedOrderActionField, pRspInfo *thost.CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
}

// 执行宣告录入请求响应
func (s *BaseTraderSpi) OnRspExecOrderInsert(pInputExecOrder *thost.CThostFtdcInputExecOrderField, pRspInfo *thost.CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
}

// 执行宣告操作请求响应
func (s *BaseTraderSpi) OnRspExecOrderAction(pInputExecOrderAction *thost.CThostFtdcInputExecOrderActionField, pRspInfo *thost.CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
}

// 询价录入请求响应
func (s *BaseTraderSpi) OnRspForQuoteInsert(pInputForQuote *thost.CThostFtdcInputForQuoteField, pRspInfo *thost.CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
}

// 报价录入请求响应
func (s *BaseTraderSpi) OnRspQuoteInsert(pInputQuote *thost.CThostFtdcInputQuoteField, pRspInfo *thost.CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
}

// 报价操作请求响应
func (s *BaseTraderSpi) OnRspQuoteAction(pInputQuoteAction *thost.CThostFtdcInputQuoteActionField, pRspInfo *thost.CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
}

// 批量报单操作请求响应
func (s *BaseTraderSpi) OnRspBatchOrderAction(pInputBatchOrderAction *thost.CThostFtdcInputBatchOrderActionField, pRspInfo *thost.CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
}

// 期权自对冲录入请求响应
func (s *BaseTraderSpi) OnRspOptionSelfCloseInsert(pInputOptionSelfClose *thost.CThostFtdcInputOptionSelfCloseField, pRspInfo *thost.CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
}

// 期权自对冲操作请求响应
func (s *BaseTraderSpi) OnRspOptionSelfCloseAction(pInputOptionSelfCloseAction *thost.CThostFtdcInputOptionSelfCloseActionField, pRspInfo *thost.CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
}

// 申请组合录入请求响应
func (s *BaseTraderSpi) OnRspCombActionInsert(pInputCombAction *thost.CThostFtdcInputCombActionField, pRspInfo *thost.CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
}

// 请求查询报单响应
func (s *BaseTraderSpi) OnRspQryOrder(pOrder *thost.CThostFtdcOrderField, pRspInfo *thost.CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
}

// 请求查询成交响应
func (s *BaseTraderSpi) OnRspQryTrade(pTrade *thost.CThostFtdcTradeField, pRspInfo *thost.CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
}

// 请求查询投资者持仓响应
func (s *BaseTraderSpi) OnRspQryInvestorPosition(pInvestorPosition *thost.CThostFtdcInvestorPositionField, pRspInfo *thost.CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
}

// 请求查询资金账户响应
func (s *BaseTraderSpi) OnRspQryTradingAccount(pTradingAccount *thost.CThostFtdcTradingAccountField, pRspInfo *thost.CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
}

// 请求查询投资者响应
func (s *BaseTraderSpi) OnRspQryInvestor(pInvestor *thost.CThostFtdcInvestorField, pRspInfo *thost.CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
}

// 请求查询交易编码响应
func (s *BaseTraderSpi) OnRspQryTradingCode(pTradingCode *thost.CThostFtdcTradingCodeField, pRspInfo *thost.CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
}

// 请求查询合约保证金率响应
func (s *BaseTraderSpi) OnRspQryInstrumentMarginRate(pInstrumentMarginRate *thost.CThostFtdcInstrumentMarginRateField, pRspInfo *thost.CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
}

// 请求查询合约手续费率响应
func (s *BaseTraderSpi) OnRspQryInstrumentCommissionRate(pInstrumentCommissionRate *thost.CThostFtdcInstrumentCommissionRateField, pRspInfo *thost.CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
}

// 请求查询交易所响应
func (s *BaseTraderSpi) OnRspQryExchange(pExchange *thost.CThostFtdcExchangeField, pRspInfo *thost.CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
}

// 请求查询产品响应
func (s *BaseTraderSpi) OnRspQryProduct(pProduct *thost.CThostFtdcProductField, pRspInfo *thost.CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
}

// 请求查询合约响应
func (s *BaseTraderSpi) OnRspQryInstrument(pInstrument *thost.CThostFtdcInstrumentField, pRspInfo *thost.CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
}

// 请求查询行情响应
func (s *BaseTraderSpi) OnRspQryDepthMarketData(pDepthMarketData *thost.CThostFtdcDepthMarketDataField, pRspInfo *thost.CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
}

// 请求查询投资者结算结果响应
func (s *BaseTraderSpi) OnRspQrySettlementInfo(pSettlementInfo *thost.CThostFtdcSettlementInfoField, pRspInfo *thost.CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
}

// 请求查询转帐银行响应
func (s *BaseTraderSpi) OnRspQryTransferBank(pTransferBank *thost.CThostFtdcTransferBankField, pRspInfo *thost.CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
}

// 请求查询投资者持仓明细响应
func (s *BaseTraderSpi) OnRspQryInvestorPositionDetail(pInvestorPositionDetail *thost.CThostFtdcInvestorPositionDetailField, pRspInfo *thost.CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
}

// 请求查询客户通知响应
func (s *BaseTraderSpi) OnRspQryNotice(pNotice *thost.CThostFtdcNoticeField, pRspInfo *thost.CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
}

// 请求查询结算信息确认响应
func (s *BaseTraderSpi) OnRspQrySettlementInfoConfirm(pSettlementInfoConfirm *thost.CThostFtdcSettlementInfoConfirmField, pRspInfo *thost.CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
}

// 请求查询投资者持仓明细响应
func (s *BaseTraderSpi) OnRspQryInvestorPositionCombineDetail(pInvestorPositionCombineDetail *thost.CThostFtdcInvestorPositionCombineDetailField, pRspInfo *thost.CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
}

// 查询保证金监管系统经纪公司资金账户密钥响应
func (s *BaseTraderSpi) OnRspQryCFMMCTradingAccountKey(pCFMMCTradingAccountKey *thost.CThostFtdcCFMMCTradingAccountKeyField, pRspInfo *thost.CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
}

// 请求查询仓单折抵信息响应
func (s *BaseTraderSpi) OnRspQryEWarrantOffset(pEWarrantOffset *thost.CThostFtdcEWarrantOffsetField, pRspInfo *thost.CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
}

// 请求查询投资者品种/跨品种保证金响应
func (s *BaseTraderSpi) OnRspQryInvestorProductGroupMargin(pInvestorProductGroupMargin *thost.CThostFtdcInvestorProductGroupMarginField, pRspInfo *thost.CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
}

// 请求查询交易所保证金率响应
func (s *BaseTraderSpi) OnRspQryExchangeMarginRate(pExchangeMarginRate *thost.CThostFtdcExchangeMarginRateField, pRspInfo *thost.CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
}

// 请求查询交易所调整保证金率响应
func (s *BaseTraderSpi) OnRspQryExchangeMarginRateAdjust(pExchangeMarginRateAdjust *thost.CThostFtdcExchangeMarginRateAdjustField, pRspInfo *thost.CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
}

// 请求查询汇率响应
func (s *BaseTraderSpi) OnRspQryExchangeRate(pExchangeRate *thost.CThostFtdcExchangeRateField, pRspInfo *thost.CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
}

// 请求查询二级代理操作员银期权限响应
func (s *BaseTraderSpi) OnRspQrySecAgentACIDMap(pSecAgentACIDMap *thost.CThostFtdcSecAgentACIDMapField, pRspInfo *thost.CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
}

// 请求查询产品报价汇率
func (s *BaseTraderSpi) OnRspQryProductExchRate(pProductExchRate *thost.CThostFtdcProductExchRateField, pRspInfo *thost.CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
}

// 请求查询产品组
func (s *BaseTraderSpi) OnRspQryProductGroup(pProductGroup *thost.CThostFtdcProductGroupField, pRspInfo *thost.CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
}

// 请求查询做市商合约手续费率响应
func (s *BaseTraderSpi) OnRspQryMMInstrumentCommissionRate(pMMInstrumentCommissionRate *thost.CThostFtdcMMInstrumentCommissionRateField, pRspInfo *thost.CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
}

// 请求查询做市商期权合约手续费响应
func (s *BaseTraderSpi) OnRspQryMMOptionInstrCommRate(pMMOptionInstrCommRate *thost.CThostFtdcMMOptionInstrCommRateField, pRspInfo *thost.CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
}

// 请求查询报单手续费响应
func (s *BaseTraderSpi) OnRspQryInstrumentOrderCommRate(pInstrumentOrderCommRate *thost.CThostFtdcInstrumentOrderCommRateField, pRspInfo *thost.CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
}

// 请求查询资金账户响应
func (s *BaseTraderSpi) OnRspQrySecAgentTradingAccount(pTradingAccount *thost.CThostFtdcTradingAccountField, pRspInfo *thost.CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
}

// 请求查询二级代理商资金校验模式响应
func (s *BaseTraderSpi) OnRspQrySecAgentCheckMode(pSecAgentCheckMode *thost.CThostFtdcSecAgentCheckModeField, pRspInfo *thost.CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
}

// 请求查询二级代理商信息响应
func (s *BaseTraderSpi) OnRspQrySecAgentTradeInfo(pSecAgentTradeInfo *thost.CThostFtdcSecAgentTradeInfoField, pRspInfo *thost.CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
}

// 请求查询期权交易成本响应
func (s *BaseTraderSpi) OnRspQryOptionInstrTradeCost(pOptionInstrTradeCost *thost.CThostFtdcOptionInstrTradeCostField, pRspInfo *thost.CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
}

// 请求查询期权合约手续费响应
func (s *BaseTraderSpi) OnRspQryOptionInstrCommRate(pOptionInstrCommRate *thost.CThostFtdcOptionInstrCommRateField, pRspInfo *thost.CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
}

// 请求查询执行宣告响应
func (s *BaseTraderSpi) OnRspQryExecOrder(pExecOrder *thost.CThostFtdcExecOrderField, pRspInfo *thost.CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
}

// 请求查询询价响应
func (s *BaseTraderSpi) OnRspQryForQuote(pForQuote *thost.CThostFtdcForQuoteField, pRspInfo *thost.CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
}

// 请求查询报价响应
func (s *BaseTraderSpi) OnRspQryQuote(pQuote *thost.CThostFtdcQuoteField, pRspInfo *thost.CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
}

// 请求查询期权自对冲响应
func (s *BaseTraderSpi) OnRspQryOptionSelfClose(pOptionSelfClose *thost.CThostFtdcOptionSelfCloseField, pRspInfo *thost.CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
}

// 请求查询投资单元响应
func (s *BaseTraderSpi) OnRspQryInvestUnit(pInvestUnit *thost.CThostFtdcInvestUnitField, pRspInfo *thost.CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
}

// 请求查询组合合约安全系数响应
func (s *BaseTraderSpi) OnRspQryCombInstrumentGuard(pCombInstrumentGuard *thost.CThostFtdcCombInstrumentGuardField, pRspInfo *thost.CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
}

// 请求查询申请组合响应
func (s *BaseTraderSpi) OnRspQryCombAction(pCombAction *thost.CThostFtdcCombActionField, pRspInfo *thost.CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
}

// 请求查询转帐流水响应
func (s *BaseTraderSpi) OnRspQryTransferSerial(pTransferSerial *thost.CThostFtdcTransferSerialField, pRspInfo *thost.CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
}

// 请求查询银期签约关系响应
func (s *BaseTraderSpi) OnRspQryAccountregister(pAccountregister *thost.CThostFtdcAccountregisterField, pRspInfo *thost.CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
}

// 错误应答
func (s *BaseTraderSpi) OnRspError(pRspInfo *thost.CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
}

// 报单通知
func (s *BaseTraderSpi) OnRtnOrder(pOrder *thost.CThostFtdcOrderField) {
}

// 成交通知
func (s *BaseTraderSpi) OnRtnTrade(pTrade *thost.CThostFtdcTradeField) {
}

// 报单录入错误回报
func (s *BaseTraderSpi) OnErrRtnOrderInsert(pInputOrder *thost.CThostFtdcInputOrderField, pRspInfo *thost.CThostFtdcRspInfoField) {
}

// 报单操作错误回报
func (s *BaseTraderSpi) OnErrRtnOrderAction(pOrderAction *thost.CThostFtdcOrderActionField, pRspInfo *thost.CThostFtdcRspInfoField) {
}

// 合约交易状态通知
func (s *BaseTraderSpi) OnRtnInstrumentStatus(pInstrumentStatus *thost.CThostFtdcInstrumentStatusField) {
}

// 交易所公告通知
func (s *BaseTraderSpi) OnRtnBulletin(pBulletin *thost.CThostFtdcBulletinField) {
}

// 交易通知
func (s *BaseTraderSpi) OnRtnTradingNotice(pTradingNoticeInfo *thost.CThostFtdcTradingNoticeInfoField) {
}

// 提示条件单校验错误
func (s *BaseTraderSpi) OnRtnErrorConditionalOrder(pErrorConditionalOrder *thost.CThostFtdcErrorConditionalOrderField) {
}

// 执行宣告通知
func (s *BaseTraderSpi) OnRtnExecOrder(pExecOrder *thost.CThostFtdcExecOrderField) {
}

// 执行宣告录入错误回报
func (s *BaseTraderSpi) OnErrRtnExecOrderInsert(pInputExecOrder *thost.CThostFtdcInputExecOrderField, pRspInfo *thost.CThostFtdcRspInfoField) {
}

// 执行宣告操作错误回报
func (s *BaseTraderSpi) OnErrRtnExecOrderAction(pExecOrderAction *thost.CThostFtdcExecOrderActionField, pRspInfo *thost.CThostFtdcRspInfoField) {
}

// 询价录入错误回报
func (s *BaseTraderSpi) OnErrRtnForQuoteInsert(pInputForQuote *thost.CThostFtdcInputForQuoteField, pRspInfo *thost.CThostFtdcRspInfoField) {
}

// 报价通知
func (s *BaseTraderSpi) OnRtnQuote(pQuote *thost.CThostFtdcQuoteField) {
}

// 报价录入错误回报
func (s *BaseTraderSpi) OnErrRtnQuoteInsert(pInputQuote *thost.CThostFtdcInputQuoteField, pRspInfo *thost.CThostFtdcRspInfoField) {
}

// 报价操作错误回报
func (s *BaseTraderSpi) OnErrRtnQuoteAction(pQuoteAction *thost.CThostFtdcQuoteActionField, pRspInfo *thost.CThostFtdcRspInfoField) {
}

// 询价通知
func (s *BaseTraderSpi) OnRtnForQuoteRsp(pForQuoteRsp *thost.CThostFtdcForQuoteRspField) {
}

// 保证金监控中心用户令牌
func (s *BaseTraderSpi) OnRtnCFMMCTradingAccountToken(pCFMMCTradingAccountToken *thost.CThostFtdcCFMMCTradingAccountTokenField) {
}

// 批量报单操作错误回报
func (s *BaseTraderSpi) OnErrRtnBatchOrderAction(pBatchOrderAction *thost.CThostFtdcBatchOrderActionField, pRspInfo *thost.CThostFtdcRspInfoField) {
}

// 期权自对冲通知
func (s *BaseTraderSpi) OnRtnOptionSelfClose(pOptionSelfClose *thost.CThostFtdcOptionSelfCloseField) {
}

// 期权自对冲录入错误回报
func (s *BaseTraderSpi) OnErrRtnOptionSelfCloseInsert(pInputOptionSelfClose *thost.CThostFtdcInputOptionSelfCloseField, pRspInfo *thost.CThostFtdcRspInfoField) {
}

// 期权自对冲操作错误回报
func (s *BaseTraderSpi) OnErrRtnOptionSelfCloseAction(pOptionSelfCloseAction *thost.CThostFtdcOptionSelfCloseActionField, pRspInfo *thost.CThostFtdcRspInfoField) {
}

// 申请组合通知
func (s *BaseTraderSpi) OnRtnCombAction(pCombAction *thost.CThostFtdcCombActionField) {
}

// 申请组合录入错误回报
func (s *BaseTraderSpi) OnErrRtnCombActionInsert(pInputCombAction *thost.CThostFtdcInputCombActionField, pRspInfo *thost.CThostFtdcRspInfoField) {
}

// 请求查询签约银行响应
func (s *BaseTraderSpi) OnRspQryContractBank(pContractBank *thost.CThostFtdcContractBankField, pRspInfo *thost.CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
}

// 请求查询预埋单响应
func (s *BaseTraderSpi) OnRspQryParkedOrder(pParkedOrder *thost.CThostFtdcParkedOrderField, pRspInfo *thost.CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
}

// 请求查询预埋撤单响应
func (s *BaseTraderSpi) OnRspQryParkedOrderAction(pParkedOrderAction *thost.CThostFtdcParkedOrderActionField, pRspInfo *thost.CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
}

// 请求查询交易通知响应
func (s *BaseTraderSpi) OnRspQryTradingNotice(pTradingNotice *thost.CThostFtdcTradingNoticeField, pRspInfo *thost.CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
}

// 请求查询经纪公司交易参数响应
func (s *BaseTraderSpi) OnRspQryBrokerTradingParams(pBrokerTradingParams *thost.CThostFtdcBrokerTradingParamsField, pRspInfo *thost.CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
}

// 请求查询经纪公司交易算法响应
func (s *BaseTraderSpi) OnRspQryBrokerTradingAlgos(pBrokerTradingAlgos *thost.CThostFtdcBrokerTradingAlgosField, pRspInfo *thost.CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
}

// 请求查询监控中心用户令牌
func (s *BaseTraderSpi) OnRspQueryCFMMCTradingAccountToken(pQueryCFMMCTradingAccountToken *thost.CThostFtdcQueryCFMMCTradingAccountTokenField, pRspInfo *thost.CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
}

// 银行发起银行资金转期货通知
func (s *BaseTraderSpi) OnRtnFromBankToFutureByBank(pRspTransfer *thost.CThostFtdcRspTransferField) {
}

// 银行发起期货资金转银行通知
func (s *BaseTraderSpi) OnRtnFromFutureToBankByBank(pRspTransfer *thost.CThostFtdcRspTransferField) {
}

// 银行发起冲正银行转期货通知
func (s *BaseTraderSpi) OnRtnRepealFromBankToFutureByBank(pRspRepeal *thost.CThostFtdcRspRepealField) {
}

// 银行发起冲正期货转银行通知
func (s *BaseTraderSpi) OnRtnRepealFromFutureToBankByBank(pRspRepeal *thost.CThostFtdcRspRepealField) {
}

// 期货发起银行资金转期货通知
func (s *BaseTraderSpi) OnRtnFromBankToFutureByFuture(pRspTransfer *thost.CThostFtdcRspTransferField) {
}

// 期货发起期货资金转银行通知
func (s *BaseTraderSpi) OnRtnFromFutureToBankByFuture(pRspTransfer *thost.CThostFtdcRspTransferField) {
}

// 系统运行时期货端手工发起冲正银行转期货请求，银行处理完毕后报盘发回的通知
func (s *BaseTraderSpi) OnRtnRepealFromBankToFutureByFutureManual(pRspRepeal *thost.CThostFtdcRspRepealField) {
}

// 系统运行时期货端手工发起冲正期货转银行请求，银行处理完毕后报盘发回的通知
func (s *BaseTraderSpi) OnRtnRepealFromFutureToBankByFutureManual(pRspRepeal *thost.CThostFtdcRspRepealField) {
}

// 期货发起查询银行余额通知
func (s *BaseTraderSpi) OnRtnQueryBankBalanceByFuture(pNotifyQueryAccount *thost.CThostFtdcNotifyQueryAccountField) {
}

// 期货发起银行资金转期货错误回报
func (s *BaseTraderSpi) OnErrRtnBankToFutureByFuture(pReqTransfer *thost.CThostFtdcReqTransferField, pRspInfo *thost.CThostFtdcRspInfoField) {
}

// 期货发起期货资金转银行错误回报
func (s *BaseTraderSpi) OnErrRtnFutureToBankByFuture(pReqTransfer *thost.CThostFtdcReqTransferField, pRspInfo *thost.CThostFtdcRspInfoField) {
}

// 系统运行时期货端手工发起冲正银行转期货错误回报
func (s *BaseTraderSpi) OnErrRtnRepealBankToFutureByFutureManual(pReqRepeal *thost.CThostFtdcReqRepealField, pRspInfo *thost.CThostFtdcRspInfoField) {
}

// 系统运行时期货端手工发起冲正期货转银行错误回报
func (s *BaseTraderSpi) OnErrRtnRepealFutureToBankByFutureManual(pReqRepeal *thost.CThostFtdcReqRepealField, pRspInfo *thost.CThostFtdcRspInfoField) {
}

// 期货发起查询银行余额错误回报
func (s *BaseTraderSpi) OnErrRtnQueryBankBalanceByFuture(pReqQueryAccount *thost.CThostFtdcReqQueryAccountField, pRspInfo *thost.CThostFtdcRspInfoField) {
}

// 期货发起冲正银行转期货请求，银行处理完毕后报盘发回的通知
func (s *BaseTraderSpi) OnRtnRepealFromBankToFutureByFuture(pRspRepeal *thost.CThostFtdcRspRepealField) {
}

// 期货发起冲正期货转银行请求，银行处理完毕后报盘发回的通知
func (s *BaseTraderSpi) OnRtnRepealFromFutureToBankByFuture(pRspRepeal *thost.CThostFtdcRspRepealField) {
}

// 期货发起银行资金转期货应答
func (s *BaseTraderSpi) OnRspFromBankToFutureByFuture(pReqTransfer *thost.CThostFtdcReqTransferField, pRspInfo *thost.CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
}

// 期货发起期货资金转银行应答
func (s *BaseTraderSpi) OnRspFromFutureToBankByFuture(pReqTransfer *thost.CThostFtdcReqTransferField, pRspInfo *thost.CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
}

// 期货发起查询银行余额应答
func (s *BaseTraderSpi) OnRspQueryBankAccountMoneyByFuture(pReqQueryAccount *thost.CThostFtdcReqQueryAccountField, pRspInfo *thost.CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
}

// 银行发起银期开户通知
func (s *BaseTraderSpi) OnRtnOpenAccountByBank(pOpenAccount *thost.CThostFtdcOpenAccountField) {
}

// 银行发起银期销户通知
func (s *BaseTraderSpi) OnRtnCancelAccountByBank(pCancelAccount *thost.CThostFtdcCancelAccountField) {
}

// 银行发起变更银行账号通知
func (s *BaseTraderSpi) OnRtnChangeAccountByBank(pChangeAccount *thost.CThostFtdcChangeAccountField) {
}

// 请求查询分类合约响应
func (s *BaseTraderSpi) OnRspQryClassifiedInstrument(pInstrument *thost.CThostFtdcInstrumentField, pRspInfo *thost.CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
}

// 请求组合优惠比例响应
func (s *BaseTraderSpi) OnRspQryCombPromotionParam(pCombPromotionParam *thost.CThostFtdcCombPromotionParamField, pRspInfo *thost.CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
}

// 投资者风险结算持仓查询响应
func (s *BaseTraderSpi) OnRspQryRiskSettleInvstPosition(pRiskSettleInvstPosition *thost.CThostFtdcRiskSettleInvstPositionField, pRspInfo *thost.CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
}

// 风险结算产品查询响应
func (s *BaseTraderSpi) OnRspQryRiskSettleProductStatus(pRiskSettleProductStatus *thost.CThostFtdcRiskSettleProductStatusField, pRspInfo *thost.CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
}
