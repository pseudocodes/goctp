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

// 响应信息
type RspInfoField struct {
	///错误代码
	ErrorID int32
	///错误信息
	ErrorMsg string
}

// 合约状态
type InstrumentStatusField struct {
	///交易所代码
	ExchangeID string
	///保留的无效字段
	reserve1 string
	///结算组代码
	SettlementGroupID string
	///保留的无效字段
	reserve2 string
	///合约交易状态
	InstrumentStatus byte
	///交易阶段编号
	TradingSegmentSN int
	///进入本状态时间
	EnterTime string
	///进入本状态原因
	EnterReason byte
	///合约在交易所的代码
	ExchangeInstID string
	///合约代码
	InstrumentID string
}

// 查询合约保证金率
type QryInstrumentMarginRateField struct {
	///经纪公司代码
	BrokerID string
	///投资者代码
	InvestorID string
	///保留的无效字段
	reserve1 string
	///投机套保标志
	HedgeFlag byte
	///交易所代码
	ExchangeID string
	///投资单元代码
	InvestUnitID string
	///合约代码
	InstrumentID string
}

// 查询签约银行响应
type ContractBankField struct {
	///经纪公司代码
	BrokerID string
	///银行代码
	BankID string
	///银行分中心代码
	BankBrchID string
	///银行名称
	BankName string
}

// 查询投资者持仓
type QryInvestorPositionField struct {
	///经纪公司代码
	BrokerID string
	///投资者代码
	InvestorID string
	///保留的无效字段
	reserve1 string
	///交易所代码
	ExchangeID string
	///投资单元代码
	InvestUnitID string
	///合约代码
	InstrumentID string
}

// 查询结算信息确认域
type QrySettlementInfoConfirmField struct {
	///经纪公司代码
	BrokerID string
	///投资者代码
	InvestorID string
	///投资者帐号
	AccountID string
	///币种代码
	CurrencyID string
}

// 银期转账交易流水表
type TransferSerialField struct {
	///平台流水号
	PlateSerial int
	///交易发起方日期
	TradeDate string
	///交易日期
	TradingDay string
	///交易时间
	TradeTime string
	///交易代码
	TradeCode string
	///会话编号
	SessionID int
	///银行编码
	BankID string
	///银行分支机构编码
	BankBranchID string
	///银行帐号类型
	BankAccType byte
	///银行帐号
	BankAccount string
	///银行流水号
	BankSerial string
	///期货公司编码
	BrokerID string
	///期商分支机构代码
	BrokerBranchID string
	///期货公司帐号类型
	FutureAccType byte
	///投资者帐号
	AccountID string
	///投资者代码
	InvestorID string
	///期货公司流水号
	FutureSerial int
	///证件类型
	IdCardType byte
	///证件号码
	IdentifiedCardNo string
	///币种代码
	CurrencyID string
	///交易金额
	TradeAmount float64
	///应收客户费用
	CustFee float64
	///应收期货公司费用
	BrokerFee float64
	///有效标志
	AvailabilityFlag byte
	///操作员
	OperatorCode string
	///新银行帐号
	BankNewAccount string
	///错误代码
	ErrorID int
	///错误信息
	ErrorMsg string
}

// 投资者结算结果确认信息
type SettlementInfoConfirmField struct {
	///经纪公司代码
	BrokerID string
	///投资者代码
	InvestorID string
	///确认日期
	ConfirmDate string
	///确认时间
	ConfirmTime string
	///结算编号
	SettlementID int
	///投资者帐号
	AccountID string
	///币种代码
	CurrencyID string
}

// 资金账户口令变更域
type TradingAccountPasswordUpdateField struct {
	///经纪公司代码
	BrokerID string
	///投资者帐号
	AccountID string
	///原来的口令
	OldPassword string
	///新的口令
	NewPassword string
	///币种代码
	CurrencyID string
}

// 经纪公司交易参数
type BrokerTradingParamsField struct {
	///经纪公司代码
	BrokerID string
	///投资者代码
	InvestorID string
	///保证金价格类型
	MarginPriceType byte
	///盈亏算法
	Algorithm byte
	///可用是否包含平仓盈利
	AvailIncludeCloseProfit byte
	///币种代码
	CurrencyID string
	///期权权利金价格类型
	OptionRoyaltyPriceType byte
	///投资者帐号
	AccountID string
}

// 查询合约
type QryInstrumentField struct {
	///保留的无效字段
	reserve1 string
	///交易所代码
	ExchangeID string
	///保留的无效字段
	reserve2 string
	///保留的无效字段
	reserve3 string
	///合约代码
	InstrumentID string
	///合约在交易所的代码
	ExchangeInstID string
	///产品代码
	ProductID string
}

// 客户端认证请求
type ReqAuthenticateField struct {
	///经纪公司代码
	BrokerID string
	///用户代码
	UserID string
	///用户端产品信息
	UserProductInfo string
	///认证码
	AuthCode string
	///App代码
	AppID string
}

// 资金账户
type TradingAccountField struct {
	///经纪公司代码
	BrokerID string
	///投资者帐号
	AccountID string
	///上次质押金额
	PreMortgage float64
	///上次信用额度
	PreCredit float64
	///上次存款额
	PreDeposit float64
	///上次结算准备金
	PreBalance float64
	///上次占用的保证金
	PreMargin float64
	///利息基数
	InterestBase float64
	///利息收入
	Interest float64
	///入金金额
	Deposit float64
	///出金金额
	Withdraw float64
	///冻结的保证金
	FrozenMargin float64
	///冻结的资金
	FrozenCash float64
	///冻结的手续费
	FrozenCommission float64
	///当前保证金总额
	CurrMargin float64
	///资金差额
	CashIn float64
	///手续费
	Commission float64
	///平仓盈亏
	CloseProfit float64
	///持仓盈亏
	PositionProfit float64
	///期货结算准备金
	Balance float64
	///可用资金
	Available float64
	///可取资金
	WithdrawQuota float64
	///基本准备金
	Reserve float64
	///交易日
	TradingDay string
	///结算编号
	SettlementID int
	///信用额度
	Credit float64
	///质押金额
	Mortgage float64
	///交易所保证金
	ExchangeMargin float64
	///投资者交割保证金
	DeliveryMargin float64
	///交易所交割保证金
	ExchangeDeliveryMargin float64
	///保底期货结算准备金
	ReserveBalance float64
	///币种代码
	CurrencyID string
	///上次货币质入金额
	PreFundMortgageIn float64
	///上次货币质出金额
	PreFundMortgageOut float64
	///货币质入金额
	FundMortgageIn float64
	///货币质出金额
	FundMortgageOut float64
	///货币质押余额
	FundMortgageAvailable float64
	///可质押货币金额
	MortgageableFund float64
	///特殊产品占用保证金
	SpecProductMargin float64
	///特殊产品冻结保证金
	SpecProductFrozenMargin float64
	///特殊产品手续费
	SpecProductCommission float64
	///特殊产品冻结手续费
	SpecProductFrozenCommission float64
	///特殊产品持仓盈亏
	SpecProductPositionProfit float64
	///特殊产品平仓盈亏
	SpecProductCloseProfit float64
	///根据持仓盈亏算法计算的特殊产品持仓盈亏
	SpecProductPositionProfitByAlg float64
	///特殊产品交易所保证金
	SpecProductExchangeMargin float64
	///业务类型
	BizType byte
	///延时换汇冻结金额
	FrozenSwap float64
	///剩余换汇额度
	RemainSwap float64
}

// 输入报单操作
type InputOrderActionField struct {
	///经纪公司代码
	BrokerID string
	///投资者代码
	InvestorID string
	///报单操作引用
	OrderActionRef int
	///报单引用
	OrderRef string
	///请求编号
	RequestID int
	///前置编号
	FrontID int
	///会话编号
	SessionID int
	///交易所代码
	ExchangeID string
	///报单编号
	OrderSysID string
	///操作标志
	ActionFlag byte
	///价格
	LimitPrice float64
	///数量变化
	VolumeChange int
	///用户代码
	UserID string
	///保留的无效字段
	reserve1 string
	///投资单元代码
	InvestUnitID string
	///保留的无效字段
	reserve2 string
	///Mac地址
	MacAddress string
	///合约代码
	InstrumentID string
	///IP地址
	IPAddress string
}

// 请求查询转帐流水
type QryTransferSerialField struct {
	///经纪公司代码
	BrokerID string
	///投资者帐号
	AccountID string
	///银行编码
	BankID string
	///币种代码
	CurrencyID string
}

// 查询报单
type QryOrderField struct {
	///经纪公司代码
	BrokerID string
	///投资者代码
	InvestorID string
	///保留的无效字段
	reserve1 string
	///交易所代码
	ExchangeID string
	///报单编号
	OrderSysID string
	///开始时间
	InsertTimeStart string
	///结束时间
	InsertTimeEnd string
	///投资单元代码
	InvestUnitID string
	///合约代码
	InstrumentID string
}

// 投资者结算结果
type SettlementInfoField struct {
	///交易日
	TradingDay string
	///结算编号
	SettlementID int
	///经纪公司代码
	BrokerID string
	///投资者代码
	InvestorID string
	///序号
	SequenceNo int
	///消息正文
	Content string
	///投资者帐号
	AccountID string
	///币种代码
	CurrencyID string
}

// 深度行情
type DepthMarketDataField struct {
	///交易日
	TradingDay string
	///保留的无效字段
	reserve1 string
	///交易所代码
	ExchangeID string
	///保留的无效字段
	reserve2 string
	///最新价
	LastPrice float64
	///上次结算价
	PreSettlementPrice float64
	///昨收盘
	PreClosePrice float64
	///昨持仓量
	PreOpenInterest float64
	///今开盘
	OpenPrice float64
	///最高价
	HighestPrice float64
	///最低价
	LowestPrice float64
	///数量
	Volume int
	///成交金额
	Turnover float64
	///持仓量
	OpenInterest float64
	///今收盘
	ClosePrice float64
	///本次结算价
	SettlementPrice float64
	///涨停板价
	UpperLimitPrice float64
	///跌停板价
	LowerLimitPrice float64
	///昨虚实度
	PreDelta float64
	///今虚实度
	CurrDelta float64
	///最后修改时间
	UpdateTime string
	///最后修改毫秒
	UpdateMillisec int
	///申买价一
	BidPrice1 float64
	///申买量一
	BidVolume1 int
	///申卖价一
	AskPrice1 float64
	///申卖量一
	AskVolume1 int
	///申买价二
	BidPrice2 float64
	///申买量二
	BidVolume2 int
	///申卖价二
	AskPrice2 float64
	///申卖量二
	AskVolume2 int
	///申买价三
	BidPrice3 float64
	///申买量三
	BidVolume3 int
	///申卖价三
	AskPrice3 float64
	///申卖量三
	AskVolume3 int
	///申买价四
	BidPrice4 float64
	///申买量四
	BidVolume4 int
	///申卖价四
	AskPrice4 float64
	///申卖量四
	AskVolume4 int
	///申买价五
	BidPrice5 float64
	///申买量五
	BidVolume5 int
	///申卖价五
	AskPrice5 float64
	///申卖量五
	AskVolume5 int
	///当日均价
	AveragePrice float64
	///业务日期
	ActionDay string
	///合约代码
	InstrumentID string
	///合约在交易所的代码
	ExchangeInstID string
	///上带价
	BandingUpperPrice float64
	///下带价
	BandingLowerPrice float64
}

// 合约手续费率
type InstrumentCommissionRateField struct {
	///保留的无效字段
	reserve1 string
	///投资者范围
	InvestorRange byte
	///经纪公司代码
	BrokerID string
	///投资者代码
	InvestorID string
	///开仓手续费率
	OpenRatioByMoney float64
	///开仓手续费
	OpenRatioByVolume float64
	///平仓手续费率
	CloseRatioByMoney float64
	///平仓手续费
	CloseRatioByVolume float64
	///平今手续费率
	CloseTodayRatioByMoney float64
	///平今手续费
	CloseTodayRatioByVolume float64
	///交易所代码
	ExchangeID string
	///业务类型
	BizType byte
	///投资单元代码
	InvestUnitID string
	///合约代码
	InstrumentID string
}

// 输入报单
type InputOrderField struct {
	///经纪公司代码
	BrokerID string
	///投资者代码
	InvestorID string
	///保留的无效字段
	reserve1 string
	///报单引用
	OrderRef string
	///用户代码
	UserID string
	///报单价格条件
	OrderPriceType byte
	///买卖方向
	Direction byte
	///组合开平标志
	CombOffsetFlag string
	///组合投机套保标志
	CombHedgeFlag string
	///价格
	LimitPrice float64
	///数量
	VolumeTotalOriginal int
	///有效期类型
	TimeCondition byte
	///GTD日期
	GTDDate string
	///成交量类型
	VolumeCondition byte
	///最小成交量
	MinVolume int
	///触发条件
	ContingentCondition byte
	///止损价
	StopPrice float64
	///强平原因
	ForceCloseReason byte
	///自动挂起标志
	IsAutoSuspend int
	///业务单元
	BusinessUnit string
	///请求编号
	RequestID int
	///用户强评标志
	UserForceClose int
	///互换单标志
	IsSwapOrder int
	///交易所代码
	ExchangeID string
	///投资单元代码
	InvestUnitID string
	///资金账号
	AccountID string
	///币种代码
	CurrencyID string
	///交易编码
	ClientID string
	///保留的无效字段
	reserve2 string
	///Mac地址
	MacAddress string
	///合约代码
	InstrumentID string
	///IP地址
	IPAddress string
}

// 成交
type TradeField struct {
	///经纪公司代码
	BrokerID string
	///投资者代码
	InvestorID string
	///保留的无效字段
	reserve1 string
	///报单引用
	OrderRef string
	///用户代码
	UserID string
	///交易所代码
	ExchangeID string
	///成交编号
	TradeID string
	///买卖方向
	Direction byte
	///报单编号
	OrderSysID string
	///会员代码
	ParticipantID string
	///客户代码
	ClientID string
	///交易角色
	TradingRole byte
	///保留的无效字段
	reserve2 string
	///开平标志
	OffsetFlag byte
	///投机套保标志
	HedgeFlag byte
	///价格
	Price float64
	///数量
	Volume int
	///成交时期
	TradeDate string
	///成交时间
	TradeTime string
	///成交类型
	TradeType byte
	///成交价来源
	PriceSource byte
	///交易所交易员代码
	TraderID string
	///本地报单编号
	OrderLocalID string
	///结算会员编号
	ClearingPartID string
	///业务单元
	BusinessUnit string
	///序号
	SequenceNo int
	///交易日
	TradingDay string
	///结算编号
	SettlementID int
	///经纪公司报单编号
	BrokerOrderSeq int
	///成交来源
	TradeSource byte
	///投资单元代码
	InvestUnitID string
	///合约代码
	InstrumentID string
	///合约在交易所的代码
	ExchangeInstID string
}

// 银行发起银行资金转期货响应
type RspTransferField struct {
	///业务功能码
	TradeCode string
	///银行代码
	BankID string
	///银行分支机构代码
	BankBranchID string
	///期商代码
	BrokerID string
	///期商分支机构代码
	BrokerBranchID string
	///交易日期
	TradeDate string
	///交易时间
	TradeTime string
	///银行流水号
	BankSerial string
	///交易系统日期
	TradingDay string
	///银期平台消息流水号
	PlateSerial int
	///最后分片标志
	LastFragment byte
	///会话号
	SessionID int
	///客户姓名
	CustomerName string
	///证件类型
	IdCardType byte
	///证件号码
	IdentifiedCardNo string
	///客户类型
	CustType byte
	///银行帐号
	BankAccount string
	///银行密码
	BankPassWord string
	///投资者帐号
	AccountID string
	///期货密码
	Password string
	///安装编号
	InstallID int
	///期货公司流水号
	FutureSerial int
	///用户标识
	UserID string
	///验证客户证件号码标志
	VerifyCertNoFlag byte
	///币种代码
	CurrencyID string
	///转帐金额
	TradeAmount float64
	///期货可取金额
	FutureFetchAmount float64
	///费用支付标志
	FeePayFlag byte
	///应收客户费用
	CustFee float64
	///应收期货公司费用
	BrokerFee float64
	///发送方给接收方的消息
	Message string
	///摘要
	Digest string
	///银行帐号类型
	BankAccType byte
	///渠道标志
	DeviceID string
	///期货单位帐号类型
	BankSecuAccType byte
	///期货公司银行编码
	BrokerIDByBank string
	///期货单位帐号
	BankSecuAcc string
	///银行密码标志
	BankPwdFlag byte
	///期货资金密码核对标志
	SecuPwdFlag byte
	///交易柜员
	OperNo string
	///请求编号
	RequestID int
	///交易ID
	TID int
	///转账交易状态
	TransferStatus byte
	///错误代码
	ErrorID int
	///错误信息
	ErrorMsg string
	///长客户姓名
	LongCustomerName string
}

// 用户事件通知信息
type TradingNoticeInfoField struct {
	///经纪公司代码
	BrokerID string
	///投资者代码
	InvestorID string
	///发送时间
	SendTime string
	///消息正文
	FieldContent string
	///序列系列号
	SequenceSeries int16
	///序列号
	SequenceNo int
	///投资单元代码
	InvestUnitID string
}

// 投资者持仓
type InvestorPositionField struct {
	///保留的无效字段
	reserve1 string
	///经纪公司代码
	BrokerID string
	///投资者代码
	InvestorID string
	///持仓多空方向
	PosiDirection byte
	///投机套保标志
	HedgeFlag byte
	///持仓日期
	PositionDate byte
	///上日持仓
	YdPosition int
	///今日持仓
	Position int
	///多头冻结
	LongFrozen int
	///空头冻结
	ShortFrozen int
	///开仓冻结金额
	LongFrozenAmount float64
	///开仓冻结金额
	ShortFrozenAmount float64
	///开仓量
	OpenVolume int
	///平仓量
	CloseVolume int
	///开仓金额
	OpenAmount float64
	///平仓金额
	CloseAmount float64
	///持仓成本
	PositionCost float64
	///上次占用的保证金
	PreMargin float64
	///占用的保证金
	UseMargin float64
	///冻结的保证金
	FrozenMargin float64
	///冻结的资金
	FrozenCash float64
	///冻结的手续费
	FrozenCommission float64
	///资金差额
	CashIn float64
	///手续费
	Commission float64
	///平仓盈亏
	CloseProfit float64
	///持仓盈亏
	PositionProfit float64
	///上次结算价
	PreSettlementPrice float64
	///本次结算价
	SettlementPrice float64
	///交易日
	TradingDay string
	///结算编号
	SettlementID int
	///开仓成本
	OpenCost float64
	///交易所保证金
	ExchangeMargin float64
	///组合成交形成的持仓
	CombPosition int
	///组合多头冻结
	CombLongFrozen int
	///组合空头冻结
	CombShortFrozen int
	///逐日盯市平仓盈亏
	CloseProfitByDate float64
	///逐笔对冲平仓盈亏
	CloseProfitByTrade float64
	///今日持仓
	TodayPosition int
	///保证金率
	MarginRateByMoney float64
	///保证金率(按手数)
	MarginRateByVolume float64
	///执行冻结
	StrikeFrozen int
	///执行冻结金额
	StrikeFrozenAmount float64
	///放弃执行冻结
	AbandonFrozen int
	///交易所代码
	ExchangeID string
	///执行冻结的昨仓
	YdStrikeFrozen int
	///投资单元代码
	InvestUnitID string
	///大商所持仓成本差值，只有大商所使用
	PositionCostOffset float64
	///tas持仓手数
	TasPosition int
	///tas持仓成本
	TasPositionCost float64
	///合约代码
	InstrumentID string
}

// 报单
type OrderField struct {
	///经纪公司代码
	BrokerID string
	///投资者代码
	InvestorID string
	///保留的无效字段
	reserve1 string
	///报单引用
	OrderRef string
	///用户代码
	UserID string
	///报单价格条件
	OrderPriceType byte
	///买卖方向
	Direction byte
	///组合开平标志
	CombOffsetFlag string
	///组合投机套保标志
	CombHedgeFlag string
	///价格
	LimitPrice float64
	///数量
	VolumeTotalOriginal int
	///有效期类型
	TimeCondition byte
	///GTD日期
	GTDDate string
	///成交量类型
	VolumeCondition byte
	///最小成交量
	MinVolume int
	///触发条件
	ContingentCondition byte
	///止损价
	StopPrice float64
	///强平原因
	ForceCloseReason byte
	///自动挂起标志
	IsAutoSuspend int
	///业务单元
	BusinessUnit string
	///请求编号
	RequestID int
	///本地报单编号
	OrderLocalID string
	///交易所代码
	ExchangeID string
	///会员代码
	ParticipantID string
	///客户代码
	ClientID string
	///保留的无效字段
	reserve2 string
	///交易所交易员代码
	TraderID string
	///安装编号
	InstallID int
	///报单提交状态
	OrderSubmitStatus byte
	///报单提示序号
	NotifySequence int
	///交易日
	TradingDay string
	///结算编号
	SettlementID int
	///报单编号
	OrderSysID string
	///报单来源
	OrderSource byte
	///报单状态
	OrderStatus byte
	///报单类型
	OrderType byte
	///今成交数量
	VolumeTraded int
	///剩余数量
	VolumeTotal int
	///报单日期
	InsertDate string
	///委托时间
	InsertTime string
	///激活时间
	ActiveTime string
	///挂起时间
	SuspendTime string
	///最后修改时间
	UpdateTime string
	///撤销时间
	CancelTime string
	///最后修改交易所交易员代码
	ActiveTraderID string
	///结算会员编号
	ClearingPartID string
	///序号
	SequenceNo int
	///前置编号
	FrontID int
	///会话编号
	SessionID int
	///用户端产品信息
	UserProductInfo string
	///状态信息
	StatusMsg string
	///用户强评标志
	UserForceClose int
	///操作用户代码
	ActiveUserID string
	///经纪公司报单编号
	BrokerOrderSeq int
	///相关报单
	RelativeOrderSysID string
	///郑商所成交数量
	ZCETotalTradedVolume int
	///互换单标志
	IsSwapOrder int
	///营业部编号
	BranchID string
	///投资单元代码
	InvestUnitID string
	///资金账号
	AccountID string
	///币种代码
	CurrencyID string
	///保留的无效字段
	reserve3 string
	///Mac地址
	MacAddress string
	///合约代码
	InstrumentID string
	///合约在交易所的代码
	ExchangeInstID string
	///IP地址
	IPAddress string
}

// 客户开销户信息表
type AccountregisterField struct {
	///交易日期
	TradeDay string
	///银行编码
	BankID string
	///银行分支机构编码
	BankBranchID string
	///银行帐号
	BankAccount string
	///期货公司编码
	BrokerID string
	///期货公司分支机构编码
	BrokerBranchID string
	///投资者帐号
	AccountID string
	///证件类型
	IdCardType byte
	///证件号码
	IdentifiedCardNo string
	///客户姓名
	CustomerName string
	///币种代码
	CurrencyID string
	///开销户类别
	OpenOrDestroy byte
	///签约日期
	RegDate string
	///解约日期
	OutDate string
	///交易ID
	TID int
	///客户类型
	CustType byte
	///银行帐号类型
	BankAccType byte
	///长客户姓名
	LongCustomerName string
}

// 合约保证金率
type InstrumentMarginRateField struct {
	///保留的无效字段
	reserve1 string
	///投资者范围
	InvestorRange byte
	///经纪公司代码
	BrokerID string
	///投资者代码
	InvestorID string
	///投机套保标志
	HedgeFlag byte
	///多头保证金率
	LongMarginRatioByMoney float64
	///多头保证金费
	LongMarginRatioByVolume float64
	///空头保证金率
	ShortMarginRatioByMoney float64
	///空头保证金费
	ShortMarginRatioByVolume float64
	///是否相对交易所收取
	IsRelative int
	///交易所代码
	ExchangeID string
	///投资单元代码
	InvestUnitID string
	///合约代码
	InstrumentID string
}

// 查询资金账户
type QryTradingAccountField struct {
	///经纪公司代码
	BrokerID string
	///投资者代码
	InvestorID string
	///币种代码
	CurrencyID string
	///业务类型
	BizType byte
	///投资者帐号
	AccountID string
}

// 请求查询银期签约关系
type QryAccountregisterField struct {
	///经纪公司代码
	BrokerID string
	///投资者帐号
	AccountID string
	///银行编码
	BankID string
	///银行分支机构编码
	BankBranchID string
	///币种代码
	CurrencyID string
}

// 查询投资者结算结果
type QrySettlementInfoField struct {
	///经纪公司代码
	BrokerID string
	///投资者代码
	InvestorID string
	///交易日
	TradingDay string
	///投资者帐号
	AccountID string
	///币种代码
	CurrencyID string
}

// 用户登录请求
type ReqUserLoginField struct {
	///交易日
	TradingDay string
	///经纪公司代码
	BrokerID string
	///用户代码
	UserID string
	///密码
	Password string
	///用户端产品信息
	UserProductInfo string
	///接口端产品信息
	InterfaceProductInfo string
	///协议信息
	ProtocolInfo string
	///Mac地址
	MacAddress string
	///动态密码
	OneTimePassword string
	///保留的无效字段
	reserve1 string
	///登录备注
	LoginRemark string
	///终端IP端口
	ClientIPPort int
	///终端IP地址
	ClientIPAddress string
}

// 查询经纪公司交易参数
type QryBrokerTradingParamsField struct {
	///经纪公司代码
	BrokerID string
	///投资者代码
	InvestorID string
	///币种代码
	CurrencyID string
	///投资者帐号
	AccountID string
}

// 查询签约银行请求
type QryContractBankField struct {
	///经纪公司代码
	BrokerID string
	///银行代码
	BankID string
	///银行分中心代码
	BankBrchID string
}

// 用户口令变更
type UserPasswordUpdateField struct {
	///经纪公司代码
	BrokerID string
	///用户代码
	UserID string
	///原来的口令
	OldPassword string
	///新的口令
	NewPassword string
}

// 用户登录应答
type RspUserLoginField struct {
	///交易日
	TradingDay string
	///登录成功时间
	LoginTime string
	///经纪公司代码
	BrokerID string
	///用户代码
	UserID string
	///交易系统名称
	SystemName string
	///前置编号
	FrontID int
	///会话编号
	SessionID int
	///最大报单引用
	MaxOrderRef string
	///上期所时间
	SHFETime string
	///大商所时间
	DCETime string
	///郑商所时间
	CZCETime string
	///中金所时间
	FFEXTime string
	///能源中心时间
	INETime string
}

// 合约
type InstrumentField struct {
	///保留的无效字段
	reserve1 string
	///交易所代码
	ExchangeID string
	///合约名称
	InstrumentName string
	///保留的无效字段
	reserve2 string
	///保留的无效字段
	reserve3 string
	///产品类型
	ProductClass byte
	///交割年份
	DeliveryYear int
	///交割月
	DeliveryMonth int
	///市价单最大下单量
	MaxMarketOrderVolume int
	///市价单最小下单量
	MinMarketOrderVolume int
	///限价单最大下单量
	MaxLimitOrderVolume int
	///限价单最小下单量
	MinLimitOrderVolume int
	///合约数量乘数
	VolumeMultiple int
	///最小变动价位
	PriceTick float64
	///创建日
	CreateDate string
	///上市日
	OpenDate string
	///到期日
	ExpireDate string
	///开始交割日
	StartDelivDate string
	///结束交割日
	EndDelivDate string
	///合约生命周期状态
	InstLifePhase byte
	///当前是否交易
	IsTrading int
	///持仓类型
	PositionType byte
	///持仓日期类型
	PositionDateType byte
	///多头保证金率
	LongMarginRatio float64
	///空头保证金率
	ShortMarginRatio float64
	///是否使用大额单边保证金算法
	MaxMarginSideAlgorithm byte
	///保留的无效字段
	reserve4 string
	///执行价
	StrikePrice float64
	///期权类型
	OptionsType byte
	///合约基础商品乘数
	UnderlyingMultiple float64
	///组合类型
	CombinationType byte
	///合约代码
	InstrumentID string
	///合约在交易所的代码
	ExchangeInstID string
	///产品代码
	ProductID string
	///基础商品代码
	UnderlyingInstrID string
}

// 报单操作
type OrderActionField struct {
	///经纪公司代码
	BrokerID string
	///投资者代码
	InvestorID string
	///报单操作引用
	OrderActionRef int
	///报单引用
	OrderRef string
	///请求编号
	RequestID int
	///前置编号
	FrontID int
	///会话编号
	SessionID int
	///交易所代码
	ExchangeID string
	///报单编号
	OrderSysID string
	///操作标志
	ActionFlag byte
	///价格
	LimitPrice float64
	///数量变化
	VolumeChange int
	///操作日期
	ActionDate string
	///操作时间
	ActionTime string
	///交易所交易员代码
	TraderID string
	///安装编号
	InstallID int
	///本地报单编号
	OrderLocalID string
	///操作本地编号
	ActionLocalID string
	///会员代码
	ParticipantID string
	///客户代码
	ClientID string
	///业务单元
	BusinessUnit string
	///报单操作状态
	OrderActionStatus byte
	///用户代码
	UserID string
	///状态信息
	StatusMsg string
	///保留的无效字段
	reserve1 string
	///营业部编号
	BranchID string
	///投资单元代码
	InvestUnitID string
	///保留的无效字段
	reserve2 string
	///Mac地址
	MacAddress string
	///合约代码
	InstrumentID string
	///IP地址
	IPAddress string
}

// 转账请求
type ReqTransferField struct {
	///业务功能码
	TradeCode string
	///银行代码
	BankID string
	///银行分支机构代码
	BankBranchID string
	///期商代码
	BrokerID string
	///期商分支机构代码
	BrokerBranchID string
	///交易日期
	TradeDate string
	///交易时间
	TradeTime string
	///银行流水号
	BankSerial string
	///交易系统日期
	TradingDay string
	///银期平台消息流水号
	PlateSerial int
	///最后分片标志
	LastFragment byte
	///会话号
	SessionID int
	///客户姓名
	CustomerName string
	///证件类型
	IdCardType byte
	///证件号码
	IdentifiedCardNo string
	///客户类型
	CustType byte
	///银行帐号
	BankAccount string
	///银行密码
	BankPassWord string
	///投资者帐号
	AccountID string
	///期货密码
	Password string
	///安装编号
	InstallID int
	///期货公司流水号
	FutureSerial int
	///用户标识
	UserID string
	///验证客户证件号码标志
	VerifyCertNoFlag byte
	///币种代码
	CurrencyID string
	///转帐金额
	TradeAmount float64
	///期货可取金额
	FutureFetchAmount float64
	///费用支付标志
	FeePayFlag byte
	///应收客户费用
	CustFee float64
	///应收期货公司费用
	BrokerFee float64
	///发送方给接收方的消息
	Message string
	///摘要
	Digest string
	///银行帐号类型
	BankAccType byte
	///渠道标志
	DeviceID string
	///期货单位帐号类型
	BankSecuAccType byte
	///期货公司银行编码
	BrokerIDByBank string
	///期货单位帐号
	BankSecuAcc string
	///银行密码标志
	BankPwdFlag byte
	///期货资金密码核对标志
	SecuPwdFlag byte
	///交易柜员
	OperNo string
	///请求编号
	RequestID int
	///交易ID
	TID int
	///转账交易状态
	TransferStatus byte
	///长客户姓名
	LongCustomerName string
}

// 客户端认证响应
type RspAuthenticateField struct {
	///经纪公司代码
	BrokerID string
	///用户代码
	UserID string
	///用户端产品信息
	UserProductInfo string
	///App代码
	AppID string
	///App类型
	AppType byte
}

// 查询手续费率
type QryInstrumentCommissionRateField struct {
	///经纪公司代码
	BrokerID string
	///投资者代码
	InvestorID string
	///保留的无效字段
	reserve1 string
	///交易所代码
	ExchangeID string
	///投资单元代码
	InvestUnitID string
	///合约代码
	InstrumentID string
}

// 查询成交
type QryTradeField struct {
	///经纪公司代码
	BrokerID string
	///投资者代码
	InvestorID string
	///保留的无效字段
	reserve1 string
	///交易所代码
	ExchangeID string
	///成交编号
	TradeID string
	///开始时间
	TradeTimeStart string
	///结束时间
	TradeTimeEnd string
	///投资单元代码
	InvestUnitID string
	///合约代码
	InstrumentID string
}
