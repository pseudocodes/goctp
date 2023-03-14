package goctp

import (
	"github.com/pseudocodes/goctp/thost"
)

type TraderSpi interface {

	///当客户端与交易后台建立起通信连接时（还未登录前），该方法被调用。
	OnFrontConnected()

	///当客户端与交易后台通信连接断开时，该方法被调用。当发生这个情况后，API会自动重新连接，客户端可不做处理。
	///@param nReason 错误原因
	///        0x1001 网络读失败
	///        0x1002 网络写失败
	///        0x2001 接收心跳超时
	///        0x2002 发送心跳失败
	///        0x2003 收到错误报文
	OnFrontDisconnected(nReason int)

	///心跳超时警告。当长时间未收到报文时，该方法被调用。
	///@param nTimeLapse 距离上次接收报文的时间
	OnHeartBeatWarning(nTimeLapse int)

	///客户端认证响应
	OnRspAuthenticate(pRspAuthenticateField *thost.CThostFtdcRspAuthenticateField, pRspInfo *thost.CThostFtdcRspInfoField, nRequestID int, bIsLast bool)

	///登录请求响应
	OnRspUserLogin(pRspUserLogin *thost.CThostFtdcRspUserLoginField, pRspInfo *thost.CThostFtdcRspInfoField, nRequestID int, bIsLast bool)

	///登出请求响应
	OnRspUserLogout(pUserLogout *thost.CThostFtdcUserLogoutField, pRspInfo *thost.CThostFtdcRspInfoField, nRequestID int, bIsLast bool)

	///用户口令更新请求响应
	OnRspUserPasswordUpdate(pUserPasswordUpdate *thost.CThostFtdcUserPasswordUpdateField, pRspInfo *thost.CThostFtdcRspInfoField, nRequestID int, bIsLast bool)

	///资金账户口令更新请求响应
	OnRspTradingAccountPasswordUpdate(pTradingAccountPasswordUpdate *thost.CThostFtdcTradingAccountPasswordUpdateField, pRspInfo *thost.CThostFtdcRspInfoField, nRequestID int, bIsLast bool)

	///查询用户当前支持的认证模式的回复
	OnRspUserAuthMethod(pRspUserAuthMethod *thost.CThostFtdcRspUserAuthMethodField, pRspInfo *thost.CThostFtdcRspInfoField, nRequestID int, bIsLast bool)

	///获取图形验证码请求的回复
	OnRspGenUserCaptcha(pRspGenUserCaptcha *thost.CThostFtdcRspGenUserCaptchaField, pRspInfo *thost.CThostFtdcRspInfoField, nRequestID int, bIsLast bool)

	///获取短信验证码请求的回复
	OnRspGenUserText(pRspGenUserText *thost.CThostFtdcRspGenUserTextField, pRspInfo *thost.CThostFtdcRspInfoField, nRequestID int, bIsLast bool)

	///报单录入请求响应
	OnRspOrderInsert(pInputOrder *thost.CThostFtdcInputOrderField, pRspInfo *thost.CThostFtdcRspInfoField, nRequestID int, bIsLast bool)

	///预埋单录入请求响应
	OnRspParkedOrderInsert(pParkedOrder *thost.CThostFtdcParkedOrderField, pRspInfo *thost.CThostFtdcRspInfoField, nRequestID int, bIsLast bool)

	///预埋撤单录入请求响应
	OnRspParkedOrderAction(pParkedOrderAction *thost.CThostFtdcParkedOrderActionField, pRspInfo *thost.CThostFtdcRspInfoField, nRequestID int, bIsLast bool)

	///报单操作请求响应
	OnRspOrderAction(pInputOrderAction *thost.CThostFtdcInputOrderActionField, pRspInfo *thost.CThostFtdcRspInfoField, nRequestID int, bIsLast bool)

	///查询最大报单数量响应
	OnRspQryMaxOrderVolume(pQryMaxOrderVolume *thost.CThostFtdcQryMaxOrderVolumeField, pRspInfo *thost.CThostFtdcRspInfoField, nRequestID int, bIsLast bool)

	///投资者结算结果确认响应
	OnRspSettlementInfoConfirm(pSettlementInfoConfirm *thost.CThostFtdcSettlementInfoConfirmField, pRspInfo *thost.CThostFtdcRspInfoField, nRequestID int, bIsLast bool)

	///删除预埋单响应
	OnRspRemoveParkedOrder(pRemoveParkedOrder *thost.CThostFtdcRemoveParkedOrderField, pRspInfo *thost.CThostFtdcRspInfoField, nRequestID int, bIsLast bool)

	///删除预埋撤单响应
	OnRspRemoveParkedOrderAction(pRemoveParkedOrderAction *thost.CThostFtdcRemoveParkedOrderActionField, pRspInfo *thost.CThostFtdcRspInfoField, nRequestID int, bIsLast bool)

	///执行宣告录入请求响应
	OnRspExecOrderInsert(pInputExecOrder *thost.CThostFtdcInputExecOrderField, pRspInfo *thost.CThostFtdcRspInfoField, nRequestID int, bIsLast bool)

	///执行宣告操作请求响应
	OnRspExecOrderAction(pInputExecOrderAction *thost.CThostFtdcInputExecOrderActionField, pRspInfo *thost.CThostFtdcRspInfoField, nRequestID int, bIsLast bool)

	///询价录入请求响应
	OnRspForQuoteInsert(pInputForQuote *thost.CThostFtdcInputForQuoteField, pRspInfo *thost.CThostFtdcRspInfoField, nRequestID int, bIsLast bool)

	///报价录入请求响应
	OnRspQuoteInsert(pInputQuote *thost.CThostFtdcInputQuoteField, pRspInfo *thost.CThostFtdcRspInfoField, nRequestID int, bIsLast bool)

	///报价操作请求响应
	OnRspQuoteAction(pInputQuoteAction *thost.CThostFtdcInputQuoteActionField, pRspInfo *thost.CThostFtdcRspInfoField, nRequestID int, bIsLast bool)

	///批量报单操作请求响应
	OnRspBatchOrderAction(pInputBatchOrderAction *thost.CThostFtdcInputBatchOrderActionField, pRspInfo *thost.CThostFtdcRspInfoField, nRequestID int, bIsLast bool)

	///期权自对冲录入请求响应
	OnRspOptionSelfCloseInsert(pInputOptionSelfClose *thost.CThostFtdcInputOptionSelfCloseField, pRspInfo *thost.CThostFtdcRspInfoField, nRequestID int, bIsLast bool)

	///期权自对冲操作请求响应
	OnRspOptionSelfCloseAction(pInputOptionSelfCloseAction *thost.CThostFtdcInputOptionSelfCloseActionField, pRspInfo *thost.CThostFtdcRspInfoField, nRequestID int, bIsLast bool)

	///申请组合录入请求响应
	OnRspCombActionInsert(pInputCombAction *thost.CThostFtdcInputCombActionField, pRspInfo *thost.CThostFtdcRspInfoField, nRequestID int, bIsLast bool)

	///请求查询报单响应
	OnRspQryOrder(pOrder *thost.CThostFtdcOrderField, pRspInfo *thost.CThostFtdcRspInfoField, nRequestID int, bIsLast bool)

	///请求查询成交响应
	OnRspQryTrade(pTrade *thost.CThostFtdcTradeField, pRspInfo *thost.CThostFtdcRspInfoField, nRequestID int, bIsLast bool)

	///请求查询投资者持仓响应
	OnRspQryInvestorPosition(pInvestorPosition *thost.CThostFtdcInvestorPositionField, pRspInfo *thost.CThostFtdcRspInfoField, nRequestID int, bIsLast bool)

	///请求查询资金账户响应
	OnRspQryTradingAccount(pTradingAccount *thost.CThostFtdcTradingAccountField, pRspInfo *thost.CThostFtdcRspInfoField, nRequestID int, bIsLast bool)

	///请求查询投资者响应
	OnRspQryInvestor(pInvestor *thost.CThostFtdcInvestorField, pRspInfo *thost.CThostFtdcRspInfoField, nRequestID int, bIsLast bool)

	///请求查询交易编码响应
	OnRspQryTradingCode(pTradingCode *thost.CThostFtdcTradingCodeField, pRspInfo *thost.CThostFtdcRspInfoField, nRequestID int, bIsLast bool)

	///请求查询合约保证金率响应
	OnRspQryInstrumentMarginRate(pInstrumentMarginRate *thost.CThostFtdcInstrumentMarginRateField, pRspInfo *thost.CThostFtdcRspInfoField, nRequestID int, bIsLast bool)

	///请求查询合约手续费率响应
	OnRspQryInstrumentCommissionRate(pInstrumentCommissionRate *thost.CThostFtdcInstrumentCommissionRateField, pRspInfo *thost.CThostFtdcRspInfoField, nRequestID int, bIsLast bool)

	///请求查询交易所响应
	OnRspQryExchange(pExchange *thost.CThostFtdcExchangeField, pRspInfo *thost.CThostFtdcRspInfoField, nRequestID int, bIsLast bool)

	///请求查询产品响应
	OnRspQryProduct(pProduct *thost.CThostFtdcProductField, pRspInfo *thost.CThostFtdcRspInfoField, nRequestID int, bIsLast bool)

	///请求查询合约响应
	OnRspQryInstrument(pInstrument *thost.CThostFtdcInstrumentField, pRspInfo *thost.CThostFtdcRspInfoField, nRequestID int, bIsLast bool)

	///请求查询行情响应
	OnRspQryDepthMarketData(pDepthMarketData *thost.CThostFtdcDepthMarketDataField, pRspInfo *thost.CThostFtdcRspInfoField, nRequestID int, bIsLast bool)

	///请求查询投资者结算结果响应
	OnRspQrySettlementInfo(pSettlementInfo *thost.CThostFtdcSettlementInfoField, pRspInfo *thost.CThostFtdcRspInfoField, nRequestID int, bIsLast bool)

	///请求查询转帐银行响应
	OnRspQryTransferBank(pTransferBank *thost.CThostFtdcTransferBankField, pRspInfo *thost.CThostFtdcRspInfoField, nRequestID int, bIsLast bool)

	///请求查询投资者持仓明细响应
	OnRspQryInvestorPositionDetail(pInvestorPositionDetail *thost.CThostFtdcInvestorPositionDetailField, pRspInfo *thost.CThostFtdcRspInfoField, nRequestID int, bIsLast bool)

	///请求查询客户通知响应
	OnRspQryNotice(pNotice *thost.CThostFtdcNoticeField, pRspInfo *thost.CThostFtdcRspInfoField, nRequestID int, bIsLast bool)

	///请求查询结算信息确认响应
	OnRspQrySettlementInfoConfirm(pSettlementInfoConfirm *thost.CThostFtdcSettlementInfoConfirmField, pRspInfo *thost.CThostFtdcRspInfoField, nRequestID int, bIsLast bool)

	///请求查询投资者持仓明细响应
	OnRspQryInvestorPositionCombineDetail(pInvestorPositionCombineDetail *thost.CThostFtdcInvestorPositionCombineDetailField, pRspInfo *thost.CThostFtdcRspInfoField, nRequestID int, bIsLast bool)

	///查询保证金监管系统经纪公司资金账户密钥响应
	OnRspQryCFMMCTradingAccountKey(pCFMMCTradingAccountKey *thost.CThostFtdcCFMMCTradingAccountKeyField, pRspInfo *thost.CThostFtdcRspInfoField, nRequestID int, bIsLast bool)

	///请求查询仓单折抵信息响应
	OnRspQryEWarrantOffset(pEWarrantOffset *thost.CThostFtdcEWarrantOffsetField, pRspInfo *thost.CThostFtdcRspInfoField, nRequestID int, bIsLast bool)

	///请求查询投资者品种/跨品种保证金响应
	OnRspQryInvestorProductGroupMargin(pInvestorProductGroupMargin *thost.CThostFtdcInvestorProductGroupMarginField, pRspInfo *thost.CThostFtdcRspInfoField, nRequestID int, bIsLast bool)

	///请求查询交易所保证金率响应
	OnRspQryExchangeMarginRate(pExchangeMarginRate *thost.CThostFtdcExchangeMarginRateField, pRspInfo *thost.CThostFtdcRspInfoField, nRequestID int, bIsLast bool)

	///请求查询交易所调整保证金率响应
	OnRspQryExchangeMarginRateAdjust(pExchangeMarginRateAdjust *thost.CThostFtdcExchangeMarginRateAdjustField, pRspInfo *thost.CThostFtdcRspInfoField, nRequestID int, bIsLast bool)

	///请求查询汇率响应
	OnRspQryExchangeRate(pExchangeRate *thost.CThostFtdcExchangeRateField, pRspInfo *thost.CThostFtdcRspInfoField, nRequestID int, bIsLast bool)

	///请求查询二级代理操作员银期权限响应
	OnRspQrySecAgentACIDMap(pSecAgentACIDMap *thost.CThostFtdcSecAgentACIDMapField, pRspInfo *thost.CThostFtdcRspInfoField, nRequestID int, bIsLast bool)

	///请求查询产品报价汇率
	OnRspQryProductExchRate(pProductExchRate *thost.CThostFtdcProductExchRateField, pRspInfo *thost.CThostFtdcRspInfoField, nRequestID int, bIsLast bool)

	///请求查询产品组
	OnRspQryProductGroup(pProductGroup *thost.CThostFtdcProductGroupField, pRspInfo *thost.CThostFtdcRspInfoField, nRequestID int, bIsLast bool)

	///请求查询做市商合约手续费率响应
	OnRspQryMMInstrumentCommissionRate(pMMInstrumentCommissionRate *thost.CThostFtdcMMInstrumentCommissionRateField, pRspInfo *thost.CThostFtdcRspInfoField, nRequestID int, bIsLast bool)

	///请求查询做市商期权合约手续费响应
	OnRspQryMMOptionInstrCommRate(pMMOptionInstrCommRate *thost.CThostFtdcMMOptionInstrCommRateField, pRspInfo *thost.CThostFtdcRspInfoField, nRequestID int, bIsLast bool)

	///请求查询报单手续费响应
	OnRspQryInstrumentOrderCommRate(pInstrumentOrderCommRate *thost.CThostFtdcInstrumentOrderCommRateField, pRspInfo *thost.CThostFtdcRspInfoField, nRequestID int, bIsLast bool)

	///请求查询资金账户响应
	OnRspQrySecAgentTradingAccount(pTradingAccount *thost.CThostFtdcTradingAccountField, pRspInfo *thost.CThostFtdcRspInfoField, nRequestID int, bIsLast bool)

	///请求查询二级代理商资金校验模式响应
	OnRspQrySecAgentCheckMode(pSecAgentCheckMode *thost.CThostFtdcSecAgentCheckModeField, pRspInfo *thost.CThostFtdcRspInfoField, nRequestID int, bIsLast bool)

	///请求查询二级代理商信息响应
	OnRspQrySecAgentTradeInfo(pSecAgentTradeInfo *thost.CThostFtdcSecAgentTradeInfoField, pRspInfo *thost.CThostFtdcRspInfoField, nRequestID int, bIsLast bool)

	///请求查询期权交易成本响应
	OnRspQryOptionInstrTradeCost(pOptionInstrTradeCost *thost.CThostFtdcOptionInstrTradeCostField, pRspInfo *thost.CThostFtdcRspInfoField, nRequestID int, bIsLast bool)

	///请求查询期权合约手续费响应
	OnRspQryOptionInstrCommRate(pOptionInstrCommRate *thost.CThostFtdcOptionInstrCommRateField, pRspInfo *thost.CThostFtdcRspInfoField, nRequestID int, bIsLast bool)

	///请求查询执行宣告响应
	OnRspQryExecOrder(pExecOrder *thost.CThostFtdcExecOrderField, pRspInfo *thost.CThostFtdcRspInfoField, nRequestID int, bIsLast bool)

	///请求查询询价响应
	OnRspQryForQuote(pForQuote *thost.CThostFtdcForQuoteField, pRspInfo *thost.CThostFtdcRspInfoField, nRequestID int, bIsLast bool)

	///请求查询报价响应
	OnRspQryQuote(pQuote *thost.CThostFtdcQuoteField, pRspInfo *thost.CThostFtdcRspInfoField, nRequestID int, bIsLast bool)

	///请求查询期权自对冲响应
	OnRspQryOptionSelfClose(pOptionSelfClose *thost.CThostFtdcOptionSelfCloseField, pRspInfo *thost.CThostFtdcRspInfoField, nRequestID int, bIsLast bool)

	///请求查询投资单元响应
	OnRspQryInvestUnit(pInvestUnit *thost.CThostFtdcInvestUnitField, pRspInfo *thost.CThostFtdcRspInfoField, nRequestID int, bIsLast bool)

	///请求查询组合合约安全系数响应
	OnRspQryCombInstrumentGuard(pCombInstrumentGuard *thost.CThostFtdcCombInstrumentGuardField, pRspInfo *thost.CThostFtdcRspInfoField, nRequestID int, bIsLast bool)

	///请求查询申请组合响应
	OnRspQryCombAction(pCombAction *thost.CThostFtdcCombActionField, pRspInfo *thost.CThostFtdcRspInfoField, nRequestID int, bIsLast bool)

	///请求查询转帐流水响应
	OnRspQryTransferSerial(pTransferSerial *thost.CThostFtdcTransferSerialField, pRspInfo *thost.CThostFtdcRspInfoField, nRequestID int, bIsLast bool)

	///请求查询银期签约关系响应
	OnRspQryAccountregister(pAccountregister *thost.CThostFtdcAccountregisterField, pRspInfo *thost.CThostFtdcRspInfoField, nRequestID int, bIsLast bool)

	///错误应答
	OnRspError(pRspInfo *thost.CThostFtdcRspInfoField, nRequestID int, bIsLast bool)

	///报单通知
	OnRtnOrder(pOrder *thost.CThostFtdcOrderField)

	///成交通知
	OnRtnTrade(pTrade *thost.CThostFtdcTradeField)

	///报单录入错误回报
	OnErrRtnOrderInsert(pInputOrder *thost.CThostFtdcInputOrderField, pRspInfo *thost.CThostFtdcRspInfoField)

	///报单操作错误回报
	OnErrRtnOrderAction(pOrderAction *thost.CThostFtdcOrderActionField, pRspInfo *thost.CThostFtdcRspInfoField)

	///合约交易状态通知
	OnRtnInstrumentStatus(pInstrumentStatus *thost.CThostFtdcInstrumentStatusField)

	///交易所公告通知
	OnRtnBulletin(pBulletin *thost.CThostFtdcBulletinField)

	///交易通知
	OnRtnTradingNotice(pTradingNoticeInfo *thost.CThostFtdcTradingNoticeInfoField)

	///提示条件单校验错误
	OnRtnErrorConditionalOrder(pErrorConditionalOrder *thost.CThostFtdcErrorConditionalOrderField)

	///执行宣告通知
	OnRtnExecOrder(pExecOrder *thost.CThostFtdcExecOrderField)

	///执行宣告录入错误回报
	OnErrRtnExecOrderInsert(pInputExecOrder *thost.CThostFtdcInputExecOrderField, pRspInfo *thost.CThostFtdcRspInfoField)

	///执行宣告操作错误回报
	OnErrRtnExecOrderAction(pExecOrderAction *thost.CThostFtdcExecOrderActionField, pRspInfo *thost.CThostFtdcRspInfoField)

	///询价录入错误回报
	OnErrRtnForQuoteInsert(pInputForQuote *thost.CThostFtdcInputForQuoteField, pRspInfo *thost.CThostFtdcRspInfoField)

	///报价通知
	OnRtnQuote(pQuote *thost.CThostFtdcQuoteField)

	///报价录入错误回报
	OnErrRtnQuoteInsert(pInputQuote *thost.CThostFtdcInputQuoteField, pRspInfo *thost.CThostFtdcRspInfoField)

	///报价操作错误回报
	OnErrRtnQuoteAction(pQuoteAction *thost.CThostFtdcQuoteActionField, pRspInfo *thost.CThostFtdcRspInfoField)

	///询价通知
	OnRtnForQuoteRsp(pForQuoteRsp *thost.CThostFtdcForQuoteRspField)

	///保证金监控中心用户令牌
	OnRtnCFMMCTradingAccountToken(pCFMMCTradingAccountToken *thost.CThostFtdcCFMMCTradingAccountTokenField)

	///批量报单操作错误回报
	OnErrRtnBatchOrderAction(pBatchOrderAction *thost.CThostFtdcBatchOrderActionField, pRspInfo *thost.CThostFtdcRspInfoField)

	///期权自对冲通知
	OnRtnOptionSelfClose(pOptionSelfClose *thost.CThostFtdcOptionSelfCloseField)

	///期权自对冲录入错误回报
	OnErrRtnOptionSelfCloseInsert(pInputOptionSelfClose *thost.CThostFtdcInputOptionSelfCloseField, pRspInfo *thost.CThostFtdcRspInfoField)

	///期权自对冲操作错误回报
	OnErrRtnOptionSelfCloseAction(pOptionSelfCloseAction *thost.CThostFtdcOptionSelfCloseActionField, pRspInfo *thost.CThostFtdcRspInfoField)

	///申请组合通知
	OnRtnCombAction(pCombAction *thost.CThostFtdcCombActionField)

	///申请组合录入错误回报
	OnErrRtnCombActionInsert(pInputCombAction *thost.CThostFtdcInputCombActionField, pRspInfo *thost.CThostFtdcRspInfoField)

	///请求查询签约银行响应
	OnRspQryContractBank(pContractBank *thost.CThostFtdcContractBankField, pRspInfo *thost.CThostFtdcRspInfoField, nRequestID int, bIsLast bool)

	///请求查询预埋单响应
	OnRspQryParkedOrder(pParkedOrder *thost.CThostFtdcParkedOrderField, pRspInfo *thost.CThostFtdcRspInfoField, nRequestID int, bIsLast bool)

	///请求查询预埋撤单响应
	OnRspQryParkedOrderAction(pParkedOrderAction *thost.CThostFtdcParkedOrderActionField, pRspInfo *thost.CThostFtdcRspInfoField, nRequestID int, bIsLast bool)

	///请求查询交易通知响应
	OnRspQryTradingNotice(pTradingNotice *thost.CThostFtdcTradingNoticeField, pRspInfo *thost.CThostFtdcRspInfoField, nRequestID int, bIsLast bool)

	///请求查询经纪公司交易参数响应
	OnRspQryBrokerTradingParams(pBrokerTradingParams *thost.CThostFtdcBrokerTradingParamsField, pRspInfo *thost.CThostFtdcRspInfoField, nRequestID int, bIsLast bool)

	///请求查询经纪公司交易算法响应
	OnRspQryBrokerTradingAlgos(pBrokerTradingAlgos *thost.CThostFtdcBrokerTradingAlgosField, pRspInfo *thost.CThostFtdcRspInfoField, nRequestID int, bIsLast bool)

	///请求查询监控中心用户令牌
	OnRspQueryCFMMCTradingAccountToken(pQueryCFMMCTradingAccountToken *thost.CThostFtdcQueryCFMMCTradingAccountTokenField, pRspInfo *thost.CThostFtdcRspInfoField, nRequestID int, bIsLast bool)

	///银行发起银行资金转期货通知
	OnRtnFromBankToFutureByBank(pRspTransfer *thost.CThostFtdcRspTransferField)

	///银行发起期货资金转银行通知
	OnRtnFromFutureToBankByBank(pRspTransfer *thost.CThostFtdcRspTransferField)

	///银行发起冲正银行转期货通知
	OnRtnRepealFromBankToFutureByBank(pRspRepeal *thost.CThostFtdcRspRepealField)

	///银行发起冲正期货转银行通知
	OnRtnRepealFromFutureToBankByBank(pRspRepeal *thost.CThostFtdcRspRepealField)

	///期货发起银行资金转期货通知
	OnRtnFromBankToFutureByFuture(pRspTransfer *thost.CThostFtdcRspTransferField)

	///期货发起期货资金转银行通知
	OnRtnFromFutureToBankByFuture(pRspTransfer *thost.CThostFtdcRspTransferField)

	///系统运行时期货端手工发起冲正银行转期货请求，银行处理完毕后报盘发回的通知
	OnRtnRepealFromBankToFutureByFutureManual(pRspRepeal *thost.CThostFtdcRspRepealField)

	///系统运行时期货端手工发起冲正期货转银行请求，银行处理完毕后报盘发回的通知
	OnRtnRepealFromFutureToBankByFutureManual(pRspRepeal *thost.CThostFtdcRspRepealField)

	///期货发起查询银行余额通知
	OnRtnQueryBankBalanceByFuture(pNotifyQueryAccount *thost.CThostFtdcNotifyQueryAccountField)

	///期货发起银行资金转期货错误回报
	OnErrRtnBankToFutureByFuture(pReqTransfer *thost.CThostFtdcReqTransferField, pRspInfo *thost.CThostFtdcRspInfoField)

	///期货发起期货资金转银行错误回报
	OnErrRtnFutureToBankByFuture(pReqTransfer *thost.CThostFtdcReqTransferField, pRspInfo *thost.CThostFtdcRspInfoField)

	///系统运行时期货端手工发起冲正银行转期货错误回报
	OnErrRtnRepealBankToFutureByFutureManual(pReqRepeal *thost.CThostFtdcReqRepealField, pRspInfo *thost.CThostFtdcRspInfoField)

	///系统运行时期货端手工发起冲正期货转银行错误回报
	OnErrRtnRepealFutureToBankByFutureManual(pReqRepeal *thost.CThostFtdcReqRepealField, pRspInfo *thost.CThostFtdcRspInfoField)

	///期货发起查询银行余额错误回报
	OnErrRtnQueryBankBalanceByFuture(pReqQueryAccount *thost.CThostFtdcReqQueryAccountField, pRspInfo *thost.CThostFtdcRspInfoField)

	///期货发起冲正银行转期货请求，银行处理完毕后报盘发回的通知
	OnRtnRepealFromBankToFutureByFuture(pRspRepeal *thost.CThostFtdcRspRepealField)

	///期货发起冲正期货转银行请求，银行处理完毕后报盘发回的通知
	OnRtnRepealFromFutureToBankByFuture(pRspRepeal *thost.CThostFtdcRspRepealField)

	///期货发起银行资金转期货应答
	OnRspFromBankToFutureByFuture(pReqTransfer *thost.CThostFtdcReqTransferField, pRspInfo *thost.CThostFtdcRspInfoField, nRequestID int, bIsLast bool)

	///期货发起期货资金转银行应答
	OnRspFromFutureToBankByFuture(pReqTransfer *thost.CThostFtdcReqTransferField, pRspInfo *thost.CThostFtdcRspInfoField, nRequestID int, bIsLast bool)

	///期货发起查询银行余额应答
	OnRspQueryBankAccountMoneyByFuture(pReqQueryAccount *thost.CThostFtdcReqQueryAccountField, pRspInfo *thost.CThostFtdcRspInfoField, nRequestID int, bIsLast bool)

	///银行发起银期开户通知
	OnRtnOpenAccountByBank(pOpenAccount *thost.CThostFtdcOpenAccountField)

	///银行发起银期销户通知
	OnRtnCancelAccountByBank(pCancelAccount *thost.CThostFtdcCancelAccountField)

	///银行发起变更银行账号通知
	OnRtnChangeAccountByBank(pChangeAccount *thost.CThostFtdcChangeAccountField)

	///请求查询分类合约响应
	OnRspQryClassifiedInstrument(pInstrument *thost.CThostFtdcInstrumentField, pRspInfo *thost.CThostFtdcRspInfoField, nRequestID int, bIsLast bool)

	///请求组合优惠比例响应
	OnRspQryCombPromotionParam(pCombPromotionParam *thost.CThostFtdcCombPromotionParamField, pRspInfo *thost.CThostFtdcRspInfoField, nRequestID int, bIsLast bool)

	///投资者风险结算持仓查询响应
	OnRspQryRiskSettleInvstPosition(pRiskSettleInvstPosition *thost.CThostFtdcRiskSettleInvstPositionField, pRspInfo *thost.CThostFtdcRspInfoField, nRequestID int, bIsLast bool)

	///风险结算产品查询响应
	OnRspQryRiskSettleProductStatus(pRiskSettleProductStatus *thost.CThostFtdcRiskSettleProductStatusField, pRspInfo *thost.CThostFtdcRspInfoField, nRequestID int, bIsLast bool)
}

type TraderApi interface {

	///获取API的版本信息
	///@retrun 获取到的版本号
	GetApiVersion() string

	///删除接口对象本身
	///@remark 不再使用本接口对象时,调用该函数删除接口对象
	Release()

	///初始化
	///@remark 初始化运行环境,只有调用后,接口才开始工作
	Init()

	///等待接口线程结束运行
	///@return 线程退出代码
	Join() int

	///获取当前交易日
	///@retrun 获取到的交易日
	///@remark 只有登录成功后,才能得到正确的交易日
	GetTradingDay() string

	///注册前置机网络地址
	///@param pszFrontAddress：前置机网络地址。
	///@remark 网络地址的格式为：“protocol://ipaddress:port”，如：”tcp://127.0.0.1:17001”。
	///@remark “tcp”代表传输协议，“127.0.0.1”代表服务器地址。”17001”代表服务器端口号。
	RegisterFront(frontAddress string)

	///注册名字服务器网络地址
	///@param pszNsAddress：名字服务器网络地址。
	///@remark 网络地址的格式为：“protocol://ipaddress:port”，如：”tcp://127.0.0.1:12001”。
	///@remark “tcp”代表传输协议，“127.0.0.1”代表服务器地址。”12001”代表服务器端口号。
	///@remark RegisterNameServer优先于RegisterFront
	RegisterNameServer(nsAddress string)

	///注册名字服务器用户信息
	///@param pFensUserInfo：用户信息。
	RegisterFensUserInfo(pFensUserInfo *thost.CThostFtdcFensUserInfoField)

	///注册回调接口
	///@param pSpi 派生自回调接口类的实例
	RegisterSpi(spi TraderSpi)

	///订阅私有流。
	///@param nResumeType 私有流重传方式
	///        THOST_TERT_RESTART:从本交易日开始重传
	///        THOST_TERT_RESUME:从上次收到的续传
	///        THOST_TERT_QUICK:只传送登录后私有流的内容
	///@remark 该方法要在Init方法前调用。若不调用则不会收到私有流的数据。
	SubscribePrivateTopic(nResumeType thost.THOST_TE_RESUME_TYPE)

	///订阅公共流。
	///@param nResumeType 公共流重传方式
	///        THOST_TERT_RESTART:从本交易日开始重传
	///        THOST_TERT_RESUME:从上次收到的续传
	///        THOST_TERT_QUICK:只传送登录后公共流的内容
	///        THOST_TERT_NONE:取消订阅公共流
	///@remark 该方法要在Init方法前调用。若不调用则不会收到公共流的数据。
	SubscribePublicTopic(nResumeType thost.THOST_TE_RESUME_TYPE)

	///客户端认证请求
	ReqAuthenticate(pReqAuthenticateField *thost.CThostFtdcReqAuthenticateField, nRequestID int) int

	///注册用户终端信息，用于中继服务器多连接模式
	///需要在终端认证成功后，用户登录前调用该接口
	RegisterUserSystemInfo(pUserSystemInfo *thost.CThostFtdcUserSystemInfoField) int

	///上报用户终端信息，用于中继服务器操作员登录模式
	///操作员登录后，可以多次调用该接口上报客户信息
	SubmitUserSystemInfo(pUserSystemInfo *thost.CThostFtdcUserSystemInfoField) int

	///用户登录请求
	ReqUserLogin(pReqUserLoginField *thost.CThostFtdcReqUserLoginField, nRequestID int) int

	///登出请求
	ReqUserLogout(pUserLogout *thost.CThostFtdcUserLogoutField, nRequestID int) int

	///用户口令更新请求
	ReqUserPasswordUpdate(pUserPasswordUpdate *thost.CThostFtdcUserPasswordUpdateField, nRequestID int) int

	///资金账户口令更新请求
	ReqTradingAccountPasswordUpdate(pTradingAccountPasswordUpdate *thost.CThostFtdcTradingAccountPasswordUpdateField, nRequestID int) int

	///查询用户当前支持的认证模式
	ReqUserAuthMethod(pReqUserAuthMethod *thost.CThostFtdcReqUserAuthMethodField, nRequestID int) int

	///用户发出获取图形验证码请求
	ReqGenUserCaptcha(pReqGenUserCaptcha *thost.CThostFtdcReqGenUserCaptchaField, nRequestID int) int

	///用户发出获取短信验证码请求
	ReqGenUserText(pReqGenUserText *thost.CThostFtdcReqGenUserTextField, nRequestID int) int

	///用户发出带有图片验证码的登陆请求
	ReqUserLoginWithCaptcha(pReqUserLoginWithCaptcha *thost.CThostFtdcReqUserLoginWithCaptchaField, nRequestID int) int

	///用户发出带有短信验证码的登陆请求
	ReqUserLoginWithText(pReqUserLoginWithText *thost.CThostFtdcReqUserLoginWithTextField, nRequestID int) int

	///用户发出带有动态口令的登陆请求
	ReqUserLoginWithOTP(pReqUserLoginWithOTP *thost.CThostFtdcReqUserLoginWithOTPField, nRequestID int) int

	///报单录入请求
	ReqOrderInsert(pInputOrder *thost.CThostFtdcInputOrderField, nRequestID int) int

	///预埋单录入请求
	ReqParkedOrderInsert(pParkedOrder *thost.CThostFtdcParkedOrderField, nRequestID int) int

	///预埋撤单录入请求
	ReqParkedOrderAction(pParkedOrderAction *thost.CThostFtdcParkedOrderActionField, nRequestID int) int

	///报单操作请求
	ReqOrderAction(pInputOrderAction *thost.CThostFtdcInputOrderActionField, nRequestID int) int

	///查询最大报单数量请求
	ReqQryMaxOrderVolume(pQryMaxOrderVolume *thost.CThostFtdcQryMaxOrderVolumeField, nRequestID int) int

	///投资者结算结果确认
	ReqSettlementInfoConfirm(pSettlementInfoConfirm *thost.CThostFtdcSettlementInfoConfirmField, nRequestID int) int

	///请求删除预埋单
	ReqRemoveParkedOrder(pRemoveParkedOrder *thost.CThostFtdcRemoveParkedOrderField, nRequestID int) int

	///请求删除预埋撤单
	ReqRemoveParkedOrderAction(pRemoveParkedOrderAction *thost.CThostFtdcRemoveParkedOrderActionField, nRequestID int) int

	///执行宣告录入请求
	ReqExecOrderInsert(pInputExecOrder *thost.CThostFtdcInputExecOrderField, nRequestID int) int

	///执行宣告操作请求
	ReqExecOrderAction(pInputExecOrderAction *thost.CThostFtdcInputExecOrderActionField, nRequestID int) int

	///询价录入请求
	ReqForQuoteInsert(pInputForQuote *thost.CThostFtdcInputForQuoteField, nRequestID int) int

	///报价录入请求
	ReqQuoteInsert(pInputQuote *thost.CThostFtdcInputQuoteField, nRequestID int) int

	///报价操作请求
	ReqQuoteAction(pInputQuoteAction *thost.CThostFtdcInputQuoteActionField, nRequestID int) int

	///批量报单操作请求
	ReqBatchOrderAction(pInputBatchOrderAction *thost.CThostFtdcInputBatchOrderActionField, nRequestID int) int

	///期权自对冲录入请求
	ReqOptionSelfCloseInsert(pInputOptionSelfClose *thost.CThostFtdcInputOptionSelfCloseField, nRequestID int) int

	///期权自对冲操作请求
	ReqOptionSelfCloseAction(pInputOptionSelfCloseAction *thost.CThostFtdcInputOptionSelfCloseActionField, nRequestID int) int

	///申请组合录入请求
	ReqCombActionInsert(pInputCombAction *thost.CThostFtdcInputCombActionField, nRequestID int) int

	///请求查询报单
	ReqQryOrder(pQryOrder *thost.CThostFtdcQryOrderField, nRequestID int) int

	///请求查询成交
	ReqQryTrade(pQryTrade *thost.CThostFtdcQryTradeField, nRequestID int) int

	///请求查询投资者持仓
	ReqQryInvestorPosition(pQryInvestorPosition *thost.CThostFtdcQryInvestorPositionField, nRequestID int) int

	///请求查询资金账户
	ReqQryTradingAccount(pQryTradingAccount *thost.CThostFtdcQryTradingAccountField, nRequestID int) int

	///请求查询投资者
	ReqQryInvestor(pQryInvestor *thost.CThostFtdcQryInvestorField, nRequestID int) int

	///请求查询交易编码
	ReqQryTradingCode(pQryTradingCode *thost.CThostFtdcQryTradingCodeField, nRequestID int) int

	///请求查询合约保证金率
	ReqQryInstrumentMarginRate(pQryInstrumentMarginRate *thost.CThostFtdcQryInstrumentMarginRateField, nRequestID int) int

	///请求查询合约手续费率
	ReqQryInstrumentCommissionRate(pQryInstrumentCommissionRate *thost.CThostFtdcQryInstrumentCommissionRateField, nRequestID int) int

	///请求查询交易所
	ReqQryExchange(pQryExchange *thost.CThostFtdcQryExchangeField, nRequestID int) int

	///请求查询产品
	ReqQryProduct(pQryProduct *thost.CThostFtdcQryProductField, nRequestID int) int

	///请求查询合约
	ReqQryInstrument(pQryInstrument *thost.CThostFtdcQryInstrumentField, nRequestID int) int

	///请求查询行情
	ReqQryDepthMarketData(pQryDepthMarketData *thost.CThostFtdcQryDepthMarketDataField, nRequestID int) int

	///请求查询投资者结算结果
	ReqQrySettlementInfo(pQrySettlementInfo *thost.CThostFtdcQrySettlementInfoField, nRequestID int) int

	///请求查询转帐银行
	ReqQryTransferBank(pQryTransferBank *thost.CThostFtdcQryTransferBankField, nRequestID int) int

	///请求查询投资者持仓明细
	ReqQryInvestorPositionDetail(pQryInvestorPositionDetail *thost.CThostFtdcQryInvestorPositionDetailField, nRequestID int) int

	///请求查询客户通知
	ReqQryNotice(pQryNotice *thost.CThostFtdcQryNoticeField, nRequestID int) int

	///请求查询结算信息确认
	ReqQrySettlementInfoConfirm(pQrySettlementInfoConfirm *thost.CThostFtdcQrySettlementInfoConfirmField, nRequestID int) int

	///请求查询投资者持仓明细
	ReqQryInvestorPositionCombineDetail(pQryInvestorPositionCombineDetail *thost.CThostFtdcQryInvestorPositionCombineDetailField, nRequestID int) int

	///请求查询保证金监管系统经纪公司资金账户密钥
	ReqQryCFMMCTradingAccountKey(pQryCFMMCTradingAccountKey *thost.CThostFtdcQryCFMMCTradingAccountKeyField, nRequestID int) int

	///请求查询仓单折抵信息
	ReqQryEWarrantOffset(pQryEWarrantOffset *thost.CThostFtdcQryEWarrantOffsetField, nRequestID int) int

	///请求查询投资者品种/跨品种保证金
	ReqQryInvestorProductGroupMargin(pQryInvestorProductGroupMargin *thost.CThostFtdcQryInvestorProductGroupMarginField, nRequestID int) int

	///请求查询交易所保证金率
	ReqQryExchangeMarginRate(pQryExchangeMarginRate *thost.CThostFtdcQryExchangeMarginRateField, nRequestID int) int

	///请求查询交易所调整保证金率
	ReqQryExchangeMarginRateAdjust(pQryExchangeMarginRateAdjust *thost.CThostFtdcQryExchangeMarginRateAdjustField, nRequestID int) int

	///请求查询汇率
	ReqQryExchangeRate(pQryExchangeRate *thost.CThostFtdcQryExchangeRateField, nRequestID int) int

	///请求查询二级代理操作员银期权限
	ReqQrySecAgentACIDMap(pQrySecAgentACIDMap *thost.CThostFtdcQrySecAgentACIDMapField, nRequestID int) int

	///请求查询产品报价汇率
	ReqQryProductExchRate(pQryProductExchRate *thost.CThostFtdcQryProductExchRateField, nRequestID int) int

	///请求查询产品组
	ReqQryProductGroup(pQryProductGroup *thost.CThostFtdcQryProductGroupField, nRequestID int) int

	///请求查询做市商合约手续费率
	ReqQryMMInstrumentCommissionRate(pQryMMInstrumentCommissionRate *thost.CThostFtdcQryMMInstrumentCommissionRateField, nRequestID int) int

	///请求查询做市商期权合约手续费
	ReqQryMMOptionInstrCommRate(pQryMMOptionInstrCommRate *thost.CThostFtdcQryMMOptionInstrCommRateField, nRequestID int) int

	///请求查询报单手续费
	ReqQryInstrumentOrderCommRate(pQryInstrumentOrderCommRate *thost.CThostFtdcQryInstrumentOrderCommRateField, nRequestID int) int

	///请求查询资金账户
	ReqQrySecAgentTradingAccount(pQryTradingAccount *thost.CThostFtdcQryTradingAccountField, nRequestID int) int

	///请求查询二级代理商资金校验模式
	ReqQrySecAgentCheckMode(pQrySecAgentCheckMode *thost.CThostFtdcQrySecAgentCheckModeField, nRequestID int) int

	///请求查询二级代理商信息
	ReqQrySecAgentTradeInfo(pQrySecAgentTradeInfo *thost.CThostFtdcQrySecAgentTradeInfoField, nRequestID int) int

	///请求查询期权交易成本
	ReqQryOptionInstrTradeCost(pQryOptionInstrTradeCost *thost.CThostFtdcQryOptionInstrTradeCostField, nRequestID int) int

	///请求查询期权合约手续费
	ReqQryOptionInstrCommRate(pQryOptionInstrCommRate *thost.CThostFtdcQryOptionInstrCommRateField, nRequestID int) int

	///请求查询执行宣告
	ReqQryExecOrder(pQryExecOrder *thost.CThostFtdcQryExecOrderField, nRequestID int) int

	///请求查询询价
	ReqQryForQuote(pQryForQuote *thost.CThostFtdcQryForQuoteField, nRequestID int) int

	///请求查询报价
	ReqQryQuote(pQryQuote *thost.CThostFtdcQryQuoteField, nRequestID int) int

	///请求查询期权自对冲
	ReqQryOptionSelfClose(pQryOptionSelfClose *thost.CThostFtdcQryOptionSelfCloseField, nRequestID int) int

	///请求查询投资单元
	ReqQryInvestUnit(pQryInvestUnit *thost.CThostFtdcQryInvestUnitField, nRequestID int) int

	///请求查询组合合约安全系数
	ReqQryCombInstrumentGuard(pQryCombInstrumentGuard *thost.CThostFtdcQryCombInstrumentGuardField, nRequestID int) int

	///请求查询申请组合
	ReqQryCombAction(pQryCombAction *thost.CThostFtdcQryCombActionField, nRequestID int) int

	///请求查询转帐流水
	ReqQryTransferSerial(pQryTransferSerial *thost.CThostFtdcQryTransferSerialField, nRequestID int) int

	///请求查询银期签约关系
	ReqQryAccountregister(pQryAccountregister *thost.CThostFtdcQryAccountregisterField, nRequestID int) int

	///请求查询签约银行
	ReqQryContractBank(pQryContractBank *thost.CThostFtdcQryContractBankField, nRequestID int) int

	///请求查询预埋单
	ReqQryParkedOrder(pQryParkedOrder *thost.CThostFtdcQryParkedOrderField, nRequestID int) int

	///请求查询预埋撤单
	ReqQryParkedOrderAction(pQryParkedOrderAction *thost.CThostFtdcQryParkedOrderActionField, nRequestID int) int

	///请求查询交易通知
	ReqQryTradingNotice(pQryTradingNotice *thost.CThostFtdcQryTradingNoticeField, nRequestID int) int

	///请求查询经纪公司交易参数
	ReqQryBrokerTradingParams(pQryBrokerTradingParams *thost.CThostFtdcQryBrokerTradingParamsField, nRequestID int) int

	///请求查询经纪公司交易算法
	ReqQryBrokerTradingAlgos(pQryBrokerTradingAlgos *thost.CThostFtdcQryBrokerTradingAlgosField, nRequestID int) int

	///请求查询监控中心用户令牌
	ReqQueryCFMMCTradingAccountToken(pQueryCFMMCTradingAccountToken *thost.CThostFtdcQueryCFMMCTradingAccountTokenField, nRequestID int) int

	///期货发起银行资金转期货请求
	ReqFromBankToFutureByFuture(pReqTransfer *thost.CThostFtdcReqTransferField, nRequestID int) int

	///期货发起期货资金转银行请求
	ReqFromFutureToBankByFuture(pReqTransfer *thost.CThostFtdcReqTransferField, nRequestID int) int

	///期货发起查询银行余额请求
	ReqQueryBankAccountMoneyByFuture(pReqQueryAccount *thost.CThostFtdcReqQueryAccountField, nRequestID int) int

	///请求查询分类合约
	ReqQryClassifiedInstrument(pQryClassifiedInstrument *thost.CThostFtdcQryClassifiedInstrumentField, nRequestID int) int

	///请求组合优惠比例
	ReqQryCombPromotionParam(pQryCombPromotionParam *thost.CThostFtdcQryCombPromotionParamField, nRequestID int) int

	///投资者风险结算持仓查询
	ReqQryRiskSettleInvstPosition(pQryRiskSettleInvstPosition *thost.CThostFtdcQryRiskSettleInvstPositionField, nRequestID int) int

	///风险结算产品查询
	ReqQryRiskSettleProductStatus(pQryRiskSettleProductStatus *thost.CThostFtdcQryRiskSettleProductStatusField, nRequestID int) int
}
