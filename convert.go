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
	"strings"

	"github.com/pseudocodes/goctp/thost"
	"golang.org/x/text/encoding/simplifiedchinese"
)

//用户登录请求
func toCThostFtdcReqUserLoginField(r *ReqUserLoginField) *thost.CThostFtdcReqUserLoginField {

	f := &thost.CThostFtdcReqUserLoginField{}

	copy(f.TradingDay[:len(f.TradingDay)-1], []byte(r.TradingDay))
	copy(f.BrokerID[:len(f.BrokerID)-1], []byte(r.BrokerID))
	copy(f.UserID[:len(f.UserID)-1], []byte(r.UserID))
	copy(f.Password[:len(f.Password)-1], []byte(r.Password))
	copy(f.UserProductInfo[:len(f.UserProductInfo)-1], []byte(r.UserProductInfo))
	copy(f.InterfaceProductInfo[:len(f.InterfaceProductInfo)-1], []byte(r.InterfaceProductInfo))
	copy(f.ProtocolInfo[:len(f.ProtocolInfo)-1], []byte(r.ProtocolInfo))
	copy(f.MacAddress[:len(f.MacAddress)-1], []byte(r.MacAddress))
	copy(f.OneTimePassword[:len(f.OneTimePassword)-1], []byte(r.OneTimePassword))

	copy(f.LoginRemark[:len(f.LoginRemark)-1], []byte(r.LoginRemark))
	f.ClientIPPort = thost.TThostFtdcIPPortType(r.ClientIPPort)
	copy(f.ClientIPAddress[:len(f.ClientIPAddress)-1], []byte(r.ClientIPAddress))

	return f
}

//客户端认证请求
func toCThostFtdcReqAuthenticateField(r *ReqAuthenticateField) *thost.CThostFtdcReqAuthenticateField {

	f := &thost.CThostFtdcReqAuthenticateField{}

	copy(f.BrokerID[:len(f.BrokerID)-1], []byte(r.BrokerID))
	copy(f.UserID[:len(f.UserID)-1], []byte(r.UserID))
	copy(f.UserProductInfo[:len(f.UserProductInfo)-1], []byte(r.UserProductInfo))
	copy(f.AuthCode[:len(f.AuthCode)-1], []byte(r.AuthCode))
	copy(f.AppID[:len(f.AppID)-1], []byte(r.AppID))

	return f
}

//查询投资者结算结果
func toCThostFtdcQrySettlementInfoField(r *QrySettlementInfoField) *thost.CThostFtdcQrySettlementInfoField {

	f := &thost.CThostFtdcQrySettlementInfoField{}

	copy(f.BrokerID[:len(f.BrokerID)-1], []byte(r.BrokerID))
	copy(f.InvestorID[:len(f.InvestorID)-1], []byte(r.InvestorID))
	copy(f.TradingDay[:len(f.TradingDay)-1], []byte(r.TradingDay))
	copy(f.AccountID[:len(f.AccountID)-1], []byte(r.AccountID))
	copy(f.CurrencyID[:len(f.CurrencyID)-1], []byte(r.CurrencyID))

	return f
}

//投资者结算结果
func toCThostFtdcSettlementInfoField(r *SettlementInfoField) *thost.CThostFtdcSettlementInfoField {

	f := &thost.CThostFtdcSettlementInfoField{}

	copy(f.TradingDay[:len(f.TradingDay)-1], []byte(r.TradingDay))
	f.SettlementID = thost.TThostFtdcSettlementIDType(r.SettlementID)
	copy(f.BrokerID[:len(f.BrokerID)-1], []byte(r.BrokerID))
	copy(f.InvestorID[:len(f.InvestorID)-1], []byte(r.InvestorID))
	f.SequenceNo = thost.TThostFtdcSequenceNoType(r.SequenceNo)
	copy(f.Content[:len(f.Content)-1], []byte(r.Content))
	copy(f.AccountID[:len(f.AccountID)-1], []byte(r.AccountID))
	copy(f.CurrencyID[:len(f.CurrencyID)-1], []byte(r.CurrencyID))

	return f
}

//查询结算信息确认域
func toCThostFtdcQrySettlementInfoConfirmField(r *QrySettlementInfoConfirmField) *thost.CThostFtdcQrySettlementInfoConfirmField {

	f := &thost.CThostFtdcQrySettlementInfoConfirmField{}

	copy(f.BrokerID[:len(f.BrokerID)-1], []byte(r.BrokerID))
	copy(f.InvestorID[:len(f.InvestorID)-1], []byte(r.InvestorID))
	copy(f.AccountID[:len(f.AccountID)-1], []byte(r.AccountID))
	copy(f.CurrencyID[:len(f.CurrencyID)-1], []byte(r.CurrencyID))

	return f
}

//投资者结算结果确认信息
func toCThostFtdcSettlementInfoConfirmField(r *SettlementInfoConfirmField) *thost.CThostFtdcSettlementInfoConfirmField {

	f := &thost.CThostFtdcSettlementInfoConfirmField{}

	copy(f.BrokerID[:len(f.BrokerID)-1], []byte(r.BrokerID))
	copy(f.InvestorID[:len(f.InvestorID)-1], []byte(r.InvestorID))
	copy(f.ConfirmDate[:len(f.ConfirmDate)-1], []byte(r.ConfirmDate))
	copy(f.ConfirmTime[:len(f.ConfirmTime)-1], []byte(r.ConfirmTime))
	f.SettlementID = thost.TThostFtdcSettlementIDType(r.SettlementID)
	copy(f.AccountID[:len(f.AccountID)-1], []byte(r.AccountID))
	copy(f.CurrencyID[:len(f.CurrencyID)-1], []byte(r.CurrencyID))

	return f
}

//查询合约
func toCThostFtdcQryInstrumentField(r *QryInstrumentField) *thost.CThostFtdcQryInstrumentField {

	f := &thost.CThostFtdcQryInstrumentField{}

	copy(f.ExchangeID[:len(f.ExchangeID)-1], []byte(r.ExchangeID))

	copy(f.InstrumentID[:len(f.InstrumentID)-1], []byte(r.InstrumentID))
	copy(f.ExchangeInstID[:len(f.ExchangeInstID)-1], []byte(r.ExchangeInstID))
	copy(f.ProductID[:len(f.ProductID)-1], []byte(r.ProductID))

	return f
}

//查询资金账户
func toCThostFtdcQryTradingAccountField(r *QryTradingAccountField) *thost.CThostFtdcQryTradingAccountField {

	f := &thost.CThostFtdcQryTradingAccountField{}

	copy(f.BrokerID[:len(f.BrokerID)-1], []byte(r.BrokerID))
	copy(f.InvestorID[:len(f.InvestorID)-1], []byte(r.InvestorID))
	copy(f.CurrencyID[:len(f.CurrencyID)-1], []byte(r.CurrencyID))
	f.BizType = thost.TThostFtdcBizTypeType(r.BizType)
	copy(f.AccountID[:len(f.AccountID)-1], []byte(r.AccountID))

	return f
}

//查询投资者持仓
func toCThostFtdcQryInvestorPositionField(r *QryInvestorPositionField) *thost.CThostFtdcQryInvestorPositionField {

	f := &thost.CThostFtdcQryInvestorPositionField{}

	copy(f.BrokerID[:len(f.BrokerID)-1], []byte(r.BrokerID))
	copy(f.InvestorID[:len(f.InvestorID)-1], []byte(r.InvestorID))

	copy(f.ExchangeID[:len(f.ExchangeID)-1], []byte(r.ExchangeID))
	copy(f.InvestUnitID[:len(f.InvestUnitID)-1], []byte(r.InvestUnitID))
	copy(f.InstrumentID[:len(f.InstrumentID)-1], []byte(r.InstrumentID))

	return f
}

//输入报单操作
func toCThostFtdcInputOrderActionField(r *InputOrderActionField) *thost.CThostFtdcInputOrderActionField {

	f := &thost.CThostFtdcInputOrderActionField{}

	copy(f.BrokerID[:len(f.BrokerID)-1], []byte(r.BrokerID))
	copy(f.InvestorID[:len(f.InvestorID)-1], []byte(r.InvestorID))
	f.OrderActionRef = thost.TThostFtdcOrderActionRefType(r.OrderActionRef)
	copy(f.OrderRef[:len(f.OrderRef)-1], []byte(r.OrderRef))
	f.RequestID = thost.TThostFtdcRequestIDType(r.RequestID)
	f.FrontID = thost.TThostFtdcFrontIDType(r.FrontID)
	f.SessionID = thost.TThostFtdcSessionIDType(r.SessionID)
	copy(f.ExchangeID[:len(f.ExchangeID)-1], []byte(r.ExchangeID))
	copy(f.OrderSysID[:len(f.OrderSysID)-1], []byte(r.OrderSysID))
	f.ActionFlag = thost.TThostFtdcActionFlagType(r.ActionFlag)
	f.LimitPrice = thost.TThostFtdcPriceType(r.LimitPrice)
	f.VolumeChange = thost.TThostFtdcVolumeType(r.VolumeChange)
	copy(f.UserID[:len(f.UserID)-1], []byte(r.UserID))

	copy(f.InvestUnitID[:len(f.InvestUnitID)-1], []byte(r.InvestUnitID))

	copy(f.MacAddress[:len(f.MacAddress)-1], []byte(r.MacAddress))
	copy(f.InstrumentID[:len(f.InstrumentID)-1], []byte(r.InstrumentID))
	copy(f.IPAddress[:len(f.IPAddress)-1], []byte(r.IPAddress))

	return f
}

//查询报单
func toCThostFtdcQryOrderField(r *QryOrderField) *thost.CThostFtdcQryOrderField {

	f := &thost.CThostFtdcQryOrderField{}

	copy(f.BrokerID[:len(f.BrokerID)-1], []byte(r.BrokerID))
	copy(f.InvestorID[:len(f.InvestorID)-1], []byte(r.InvestorID))

	copy(f.ExchangeID[:len(f.ExchangeID)-1], []byte(r.ExchangeID))
	copy(f.OrderSysID[:len(f.OrderSysID)-1], []byte(r.OrderSysID))
	copy(f.InsertTimeStart[:len(f.InsertTimeStart)-1], []byte(r.InsertTimeStart))
	copy(f.InsertTimeEnd[:len(f.InsertTimeEnd)-1], []byte(r.InsertTimeEnd))
	copy(f.InvestUnitID[:len(f.InvestUnitID)-1], []byte(r.InvestUnitID))
	copy(f.InstrumentID[:len(f.InstrumentID)-1], []byte(r.InstrumentID))

	return f
}

//查询成交
func toCThostFtdcQryTradeField(r *QryTradeField) *thost.CThostFtdcQryTradeField {

	f := &thost.CThostFtdcQryTradeField{}

	copy(f.BrokerID[:len(f.BrokerID)-1], []byte(r.BrokerID))
	copy(f.InvestorID[:len(f.InvestorID)-1], []byte(r.InvestorID))

	copy(f.ExchangeID[:len(f.ExchangeID)-1], []byte(r.ExchangeID))
	copy(f.TradeID[:len(f.TradeID)-1], []byte(r.TradeID))
	copy(f.TradeTimeStart[:len(f.TradeTimeStart)-1], []byte(r.TradeTimeStart))
	copy(f.TradeTimeEnd[:len(f.TradeTimeEnd)-1], []byte(r.TradeTimeEnd))
	copy(f.InvestUnitID[:len(f.InvestUnitID)-1], []byte(r.InvestUnitID))
	copy(f.InstrumentID[:len(f.InstrumentID)-1], []byte(r.InstrumentID))

	return f
}

//输入报单
func toCThostFtdcInputOrderField(r *InputOrderField) *thost.CThostFtdcInputOrderField {

	f := &thost.CThostFtdcInputOrderField{}

	copy(f.BrokerID[:len(f.BrokerID)-1], []byte(r.BrokerID))
	copy(f.InvestorID[:len(f.InvestorID)-1], []byte(r.InvestorID))

	copy(f.OrderRef[:len(f.OrderRef)-1], []byte(r.OrderRef))
	copy(f.UserID[:len(f.UserID)-1], []byte(r.UserID))
	f.OrderPriceType = thost.TThostFtdcOrderPriceTypeType(r.OrderPriceType)
	f.Direction = thost.TThostFtdcDirectionType(r.Direction)
	copy(f.CombOffsetFlag[:len(f.CombOffsetFlag)-1], []byte(r.CombOffsetFlag))
	copy(f.CombHedgeFlag[:len(f.CombHedgeFlag)-1], []byte(r.CombHedgeFlag))
	f.LimitPrice = thost.TThostFtdcPriceType(r.LimitPrice)
	f.VolumeTotalOriginal = thost.TThostFtdcVolumeType(r.VolumeTotalOriginal)
	f.TimeCondition = thost.TThostFtdcTimeConditionType(r.TimeCondition)
	copy(f.GTDDate[:len(f.GTDDate)-1], []byte(r.GTDDate))
	f.VolumeCondition = thost.TThostFtdcVolumeConditionType(r.VolumeCondition)
	f.MinVolume = thost.TThostFtdcVolumeType(r.MinVolume)
	f.ContingentCondition = thost.TThostFtdcContingentConditionType(r.ContingentCondition)
	f.StopPrice = thost.TThostFtdcPriceType(r.StopPrice)
	f.ForceCloseReason = thost.TThostFtdcForceCloseReasonType(r.ForceCloseReason)
	f.IsAutoSuspend = thost.TThostFtdcBoolType(r.IsAutoSuspend)
	copy(f.BusinessUnit[:len(f.BusinessUnit)-1], []byte(r.BusinessUnit))
	f.RequestID = thost.TThostFtdcRequestIDType(r.RequestID)
	f.UserForceClose = thost.TThostFtdcBoolType(r.UserForceClose)
	f.IsSwapOrder = thost.TThostFtdcBoolType(r.IsSwapOrder)
	copy(f.ExchangeID[:len(f.ExchangeID)-1], []byte(r.ExchangeID))
	copy(f.InvestUnitID[:len(f.InvestUnitID)-1], []byte(r.InvestUnitID))
	copy(f.AccountID[:len(f.AccountID)-1], []byte(r.AccountID))
	copy(f.CurrencyID[:len(f.CurrencyID)-1], []byte(r.CurrencyID))
	copy(f.ClientID[:len(f.ClientID)-1], []byte(r.ClientID))

	copy(f.MacAddress[:len(f.MacAddress)-1], []byte(r.MacAddress))
	copy(f.InstrumentID[:len(f.InstrumentID)-1], []byte(r.InstrumentID))
	copy(f.IPAddress[:len(f.IPAddress)-1], []byte(r.IPAddress))

	return f
}

//查询合约保证金率
func toCThostFtdcQryInstrumentMarginRateField(r *QryInstrumentMarginRateField) *thost.CThostFtdcQryInstrumentMarginRateField {

	f := &thost.CThostFtdcQryInstrumentMarginRateField{}

	copy(f.BrokerID[:len(f.BrokerID)-1], []byte(r.BrokerID))
	copy(f.InvestorID[:len(f.InvestorID)-1], []byte(r.InvestorID))

	f.HedgeFlag = thost.TThostFtdcHedgeFlagType(r.HedgeFlag)
	copy(f.ExchangeID[:len(f.ExchangeID)-1], []byte(r.ExchangeID))
	copy(f.InvestUnitID[:len(f.InvestUnitID)-1], []byte(r.InvestUnitID))
	copy(f.InstrumentID[:len(f.InstrumentID)-1], []byte(r.InstrumentID))

	return f
}

//查询手续费率
func toCThostFtdcQryInstrumentCommissionRateField(r *QryInstrumentCommissionRateField) *thost.CThostFtdcQryInstrumentCommissionRateField {

	f := &thost.CThostFtdcQryInstrumentCommissionRateField{}

	copy(f.BrokerID[:len(f.BrokerID)-1], []byte(r.BrokerID))
	copy(f.InvestorID[:len(f.InvestorID)-1], []byte(r.InvestorID))

	copy(f.ExchangeID[:len(f.ExchangeID)-1], []byte(r.ExchangeID))
	copy(f.InvestUnitID[:len(f.InvestUnitID)-1], []byte(r.InvestUnitID))
	copy(f.InstrumentID[:len(f.InstrumentID)-1], []byte(r.InstrumentID))

	return f
}

//响应信息
func fromCThostFtdcRspInfoField(r *thost.CThostFtdcRspInfoField) *RspInfoField {

	f := &RspInfoField{

		ErrorID:  int32(r.ErrorID),
		ErrorMsg: bytes2String(r.ErrorMsg[:]),
	}

	return f
}

//深度行情
func fromCThostFtdcDepthMarketDataField(r *thost.CThostFtdcDepthMarketDataField) *DepthMarketDataField {

	f := &DepthMarketDataField{

		TradingDay: bytes2String(r.TradingDay[:]),

		ExchangeID: bytes2String(r.ExchangeID[:]),

		LastPrice:          float64(r.LastPrice),
		PreSettlementPrice: float64(r.PreSettlementPrice),
		PreClosePrice:      float64(r.PreClosePrice),
		PreOpenInterest:    float64(r.PreOpenInterest),
		OpenPrice:          float64(r.OpenPrice),
		HighestPrice:       float64(r.HighestPrice),
		LowestPrice:        float64(r.LowestPrice),
		Volume:             int(r.Volume),
		Turnover:           float64(r.Turnover),
		OpenInterest:       float64(r.OpenInterest),
		ClosePrice:         float64(r.ClosePrice),
		SettlementPrice:    float64(r.SettlementPrice),
		UpperLimitPrice:    float64(r.UpperLimitPrice),
		LowerLimitPrice:    float64(r.LowerLimitPrice),
		PreDelta:           float64(r.PreDelta),
		CurrDelta:          float64(r.CurrDelta),
		UpdateTime:         bytes2String(r.UpdateTime[:]),
		UpdateMillisec:     int(r.UpdateMillisec),
		BidPrice1:          float64(r.BidPrice1),
		BidVolume1:         int(r.BidVolume1),
		AskPrice1:          float64(r.AskPrice1),
		AskVolume1:         int(r.AskVolume1),
		BidPrice2:          float64(r.BidPrice2),
		BidVolume2:         int(r.BidVolume2),
		AskPrice2:          float64(r.AskPrice2),
		AskVolume2:         int(r.AskVolume2),
		BidPrice3:          float64(r.BidPrice3),
		BidVolume3:         int(r.BidVolume3),
		AskPrice3:          float64(r.AskPrice3),
		AskVolume3:         int(r.AskVolume3),
		BidPrice4:          float64(r.BidPrice4),
		BidVolume4:         int(r.BidVolume4),
		AskPrice4:          float64(r.AskPrice4),
		AskVolume4:         int(r.AskVolume4),
		BidPrice5:          float64(r.BidPrice5),
		BidVolume5:         int(r.BidVolume5),
		AskPrice5:          float64(r.AskPrice5),
		AskVolume5:         int(r.AskVolume5),
		AveragePrice:       float64(r.AveragePrice),
		ActionDay:          bytes2String(r.ActionDay[:]),
		InstrumentID:       bytes2String(r.InstrumentID[:]),
		ExchangeInstID:     bytes2String(r.ExchangeInstID[:]),
		BandingUpperPrice:  float64(r.BandingUpperPrice),
		BandingLowerPrice:  float64(r.BandingLowerPrice),
	}

	return f
}

//用户登录应答
func fromCThostFtdcRspUserLoginField(r *thost.CThostFtdcRspUserLoginField) *RspUserLoginField {

	f := &RspUserLoginField{

		TradingDay:  bytes2String(r.TradingDay[:]),
		LoginTime:   bytes2String(r.LoginTime[:]),
		BrokerID:    bytes2String(r.BrokerID[:]),
		UserID:      bytes2String(r.UserID[:]),
		SystemName:  bytes2String(r.SystemName[:]),
		FrontID:     int(r.FrontID),
		SessionID:   int(r.SessionID),
		MaxOrderRef: bytes2String(r.MaxOrderRef[:]),
		SHFETime:    bytes2String(r.SHFETime[:]),
		DCETime:     bytes2String(r.DCETime[:]),
		CZCETime:    bytes2String(r.CZCETime[:]),
		FFEXTime:    bytes2String(r.FFEXTime[:]),
		INETime:     bytes2String(r.INETime[:]),
	}

	return f
}

//投资者结算结果确认信息
func fromCThostFtdcSettlementInfoConfirmField(r *thost.CThostFtdcSettlementInfoConfirmField) *SettlementInfoConfirmField {

	f := &SettlementInfoConfirmField{

		BrokerID:     bytes2String(r.BrokerID[:]),
		InvestorID:   bytes2String(r.InvestorID[:]),
		ConfirmDate:  bytes2String(r.ConfirmDate[:]),
		ConfirmTime:  bytes2String(r.ConfirmTime[:]),
		SettlementID: int(r.SettlementID),
		AccountID:    bytes2String(r.AccountID[:]),
		CurrencyID:   bytes2String(r.CurrencyID[:]),
	}

	return f
}

//客户端认证响应
func fromCThostFtdcRspAuthenticateField(r *thost.CThostFtdcRspAuthenticateField) *RspAuthenticateField {

	f := &RspAuthenticateField{

		BrokerID:        bytes2String(r.BrokerID[:]),
		UserID:          bytes2String(r.UserID[:]),
		UserProductInfo: bytes2String(r.UserProductInfo[:]),
		AppID:           bytes2String(r.AppID[:]),
		AppType:         byte(r.AppType),
	}

	return f
}

//合约
func fromCThostFtdcInstrumentField(r *thost.CThostFtdcInstrumentField) *InstrumentField {

	f := &InstrumentField{

		ExchangeID:     bytes2String(r.ExchangeID[:]),
		InstrumentName: bytes2String(r.InstrumentName[:]),

		ProductClass:           byte(r.ProductClass),
		DeliveryYear:           int(r.DeliveryYear),
		DeliveryMonth:          int(r.DeliveryMonth),
		MaxMarketOrderVolume:   int(r.MaxMarketOrderVolume),
		MinMarketOrderVolume:   int(r.MinMarketOrderVolume),
		MaxLimitOrderVolume:    int(r.MaxLimitOrderVolume),
		MinLimitOrderVolume:    int(r.MinLimitOrderVolume),
		VolumeMultiple:         int(r.VolumeMultiple),
		PriceTick:              float64(r.PriceTick),
		CreateDate:             bytes2String(r.CreateDate[:]),
		OpenDate:               bytes2String(r.OpenDate[:]),
		ExpireDate:             bytes2String(r.ExpireDate[:]),
		StartDelivDate:         bytes2String(r.StartDelivDate[:]),
		EndDelivDate:           bytes2String(r.EndDelivDate[:]),
		InstLifePhase:          byte(r.InstLifePhase),
		IsTrading:              int(r.IsTrading),
		PositionType:           byte(r.PositionType),
		PositionDateType:       byte(r.PositionDateType),
		LongMarginRatio:        float64(r.LongMarginRatio),
		ShortMarginRatio:       float64(r.ShortMarginRatio),
		MaxMarginSideAlgorithm: byte(r.MaxMarginSideAlgorithm),

		StrikePrice:        float64(r.StrikePrice),
		OptionsType:        byte(r.OptionsType),
		UnderlyingMultiple: float64(r.UnderlyingMultiple),
		CombinationType:    byte(r.CombinationType),
		InstrumentID:       bytes2String(r.InstrumentID[:]),
		ExchangeInstID:     bytes2String(r.ExchangeInstID[:]),
		ProductID:          bytes2String(r.ProductID[:]),
		UnderlyingInstrID:  bytes2String(r.UnderlyingInstrID[:]),
	}

	return f
}

//投资者结算结果
func fromCThostFtdcSettlementInfoField(r *thost.CThostFtdcSettlementInfoField) *SettlementInfoField {

	f := &SettlementInfoField{

		TradingDay:   bytes2String(r.TradingDay[:]),
		SettlementID: int(r.SettlementID),
		BrokerID:     bytes2String(r.BrokerID[:]),
		InvestorID:   bytes2String(r.InvestorID[:]),
		SequenceNo:   int(r.SequenceNo),
		Content:      bytes2String(r.Content[:]),
		AccountID:    bytes2String(r.AccountID[:]),
		CurrencyID:   bytes2String(r.CurrencyID[:]),
	}

	return f
}

//资金账户
func fromCThostFtdcTradingAccountField(r *thost.CThostFtdcTradingAccountField) *TradingAccountField {

	f := &TradingAccountField{

		BrokerID:                       bytes2String(r.BrokerID[:]),
		AccountID:                      bytes2String(r.AccountID[:]),
		PreMortgage:                    float64(r.PreMortgage),
		PreCredit:                      float64(r.PreCredit),
		PreDeposit:                     float64(r.PreDeposit),
		PreBalance:                     float64(r.PreBalance),
		PreMargin:                      float64(r.PreMargin),
		InterestBase:                   float64(r.InterestBase),
		Interest:                       float64(r.Interest),
		Deposit:                        float64(r.Deposit),
		Withdraw:                       float64(r.Withdraw),
		FrozenMargin:                   float64(r.FrozenMargin),
		FrozenCash:                     float64(r.FrozenCash),
		FrozenCommission:               float64(r.FrozenCommission),
		CurrMargin:                     float64(r.CurrMargin),
		CashIn:                         float64(r.CashIn),
		Commission:                     float64(r.Commission),
		CloseProfit:                    float64(r.CloseProfit),
		PositionProfit:                 float64(r.PositionProfit),
		Balance:                        float64(r.Balance),
		Available:                      float64(r.Available),
		WithdrawQuota:                  float64(r.WithdrawQuota),
		Reserve:                        float64(r.Reserve),
		TradingDay:                     bytes2String(r.TradingDay[:]),
		SettlementID:                   int(r.SettlementID),
		Credit:                         float64(r.Credit),
		Mortgage:                       float64(r.Mortgage),
		ExchangeMargin:                 float64(r.ExchangeMargin),
		DeliveryMargin:                 float64(r.DeliveryMargin),
		ExchangeDeliveryMargin:         float64(r.ExchangeDeliveryMargin),
		ReserveBalance:                 float64(r.ReserveBalance),
		CurrencyID:                     bytes2String(r.CurrencyID[:]),
		PreFundMortgageIn:              float64(r.PreFundMortgageIn),
		PreFundMortgageOut:             float64(r.PreFundMortgageOut),
		FundMortgageIn:                 float64(r.FundMortgageIn),
		FundMortgageOut:                float64(r.FundMortgageOut),
		FundMortgageAvailable:          float64(r.FundMortgageAvailable),
		MortgageableFund:               float64(r.MortgageableFund),
		SpecProductMargin:              float64(r.SpecProductMargin),
		SpecProductFrozenMargin:        float64(r.SpecProductFrozenMargin),
		SpecProductCommission:          float64(r.SpecProductCommission),
		SpecProductFrozenCommission:    float64(r.SpecProductFrozenCommission),
		SpecProductPositionProfit:      float64(r.SpecProductPositionProfit),
		SpecProductCloseProfit:         float64(r.SpecProductCloseProfit),
		SpecProductPositionProfitByAlg: float64(r.SpecProductPositionProfitByAlg),
		SpecProductExchangeMargin:      float64(r.SpecProductExchangeMargin),
		BizType:                        byte(r.BizType),
		FrozenSwap:                     float64(r.FrozenSwap),
		RemainSwap:                     float64(r.RemainSwap),
	}

	return f
}

//投资者持仓
func fromCThostFtdcInvestorPositionField(r *thost.CThostFtdcInvestorPositionField) *InvestorPositionField {

	f := &InvestorPositionField{

		BrokerID:           bytes2String(r.BrokerID[:]),
		InvestorID:         bytes2String(r.InvestorID[:]),
		PosiDirection:      byte(r.PosiDirection),
		HedgeFlag:          byte(r.HedgeFlag),
		PositionDate:       byte(r.PositionDate),
		YdPosition:         int(r.YdPosition),
		Position:           int(r.Position),
		LongFrozen:         int(r.LongFrozen),
		ShortFrozen:        int(r.ShortFrozen),
		LongFrozenAmount:   float64(r.LongFrozenAmount),
		ShortFrozenAmount:  float64(r.ShortFrozenAmount),
		OpenVolume:         int(r.OpenVolume),
		CloseVolume:        int(r.CloseVolume),
		OpenAmount:         float64(r.OpenAmount),
		CloseAmount:        float64(r.CloseAmount),
		PositionCost:       float64(r.PositionCost),
		PreMargin:          float64(r.PreMargin),
		UseMargin:          float64(r.UseMargin),
		FrozenMargin:       float64(r.FrozenMargin),
		FrozenCash:         float64(r.FrozenCash),
		FrozenCommission:   float64(r.FrozenCommission),
		CashIn:             float64(r.CashIn),
		Commission:         float64(r.Commission),
		CloseProfit:        float64(r.CloseProfit),
		PositionProfit:     float64(r.PositionProfit),
		PreSettlementPrice: float64(r.PreSettlementPrice),
		SettlementPrice:    float64(r.SettlementPrice),
		TradingDay:         bytes2String(r.TradingDay[:]),
		SettlementID:       int(r.SettlementID),
		OpenCost:           float64(r.OpenCost),
		ExchangeMargin:     float64(r.ExchangeMargin),
		CombPosition:       int(r.CombPosition),
		CombLongFrozen:     int(r.CombLongFrozen),
		CombShortFrozen:    int(r.CombShortFrozen),
		CloseProfitByDate:  float64(r.CloseProfitByDate),
		CloseProfitByTrade: float64(r.CloseProfitByTrade),
		TodayPosition:      int(r.TodayPosition),
		MarginRateByMoney:  float64(r.MarginRateByMoney),
		MarginRateByVolume: float64(r.MarginRateByVolume),
		StrikeFrozen:       int(r.StrikeFrozen),
		StrikeFrozenAmount: float64(r.StrikeFrozenAmount),
		AbandonFrozen:      int(r.AbandonFrozen),
		ExchangeID:         bytes2String(r.ExchangeID[:]),
		YdStrikeFrozen:     int(r.YdStrikeFrozen),
		InvestUnitID:       bytes2String(r.InvestUnitID[:]),
		PositionCostOffset: float64(r.PositionCostOffset),
		TasPosition:        int(r.TasPosition),
		TasPositionCost:    float64(r.TasPositionCost),
		InstrumentID:       bytes2String(r.InstrumentID[:]),
	}

	return f
}

//报单
func fromCThostFtdcOrderField(r *thost.CThostFtdcOrderField) *OrderField {

	f := &OrderField{

		BrokerID:   bytes2String(r.BrokerID[:]),
		InvestorID: bytes2String(r.InvestorID[:]),

		OrderRef:            bytes2String(r.OrderRef[:]),
		UserID:              bytes2String(r.UserID[:]),
		OrderPriceType:      byte(r.OrderPriceType),
		Direction:           byte(r.Direction),
		CombOffsetFlag:      bytes2String(r.CombOffsetFlag[:]),
		CombHedgeFlag:       bytes2String(r.CombHedgeFlag[:]),
		LimitPrice:          float64(r.LimitPrice),
		VolumeTotalOriginal: int(r.VolumeTotalOriginal),
		TimeCondition:       byte(r.TimeCondition),
		GTDDate:             bytes2String(r.GTDDate[:]),
		VolumeCondition:     byte(r.VolumeCondition),
		MinVolume:           int(r.MinVolume),
		ContingentCondition: byte(r.ContingentCondition),
		StopPrice:           float64(r.StopPrice),
		ForceCloseReason:    byte(r.ForceCloseReason),
		IsAutoSuspend:       int(r.IsAutoSuspend),
		BusinessUnit:        bytes2String(r.BusinessUnit[:]),
		RequestID:           int(r.RequestID),
		OrderLocalID:        bytes2String(r.OrderLocalID[:]),
		ExchangeID:          bytes2String(r.ExchangeID[:]),
		ParticipantID:       bytes2String(r.ParticipantID[:]),
		ClientID:            bytes2String(r.ClientID[:]),

		TraderID:             bytes2String(r.TraderID[:]),
		InstallID:            int(r.InstallID),
		OrderSubmitStatus:    byte(r.OrderSubmitStatus),
		NotifySequence:       int(r.NotifySequence),
		TradingDay:           bytes2String(r.TradingDay[:]),
		SettlementID:         int(r.SettlementID),
		OrderSysID:           bytes2String(r.OrderSysID[:]),
		OrderSource:          byte(r.OrderSource),
		OrderStatus:          byte(r.OrderStatus),
		OrderType:            byte(r.OrderType),
		VolumeTraded:         int(r.VolumeTraded),
		VolumeTotal:          int(r.VolumeTotal),
		InsertDate:           bytes2String(r.InsertDate[:]),
		InsertTime:           bytes2String(r.InsertTime[:]),
		ActiveTime:           bytes2String(r.ActiveTime[:]),
		SuspendTime:          bytes2String(r.SuspendTime[:]),
		UpdateTime:           bytes2String(r.UpdateTime[:]),
		CancelTime:           bytes2String(r.CancelTime[:]),
		ActiveTraderID:       bytes2String(r.ActiveTraderID[:]),
		ClearingPartID:       bytes2String(r.ClearingPartID[:]),
		SequenceNo:           int(r.SequenceNo),
		FrontID:              int(r.FrontID),
		SessionID:            int(r.SessionID),
		UserProductInfo:      bytes2String(r.UserProductInfo[:]),
		StatusMsg:            bytes2String(r.StatusMsg[:]),
		UserForceClose:       int(r.UserForceClose),
		ActiveUserID:         bytes2String(r.ActiveUserID[:]),
		BrokerOrderSeq:       int(r.BrokerOrderSeq),
		RelativeOrderSysID:   bytes2String(r.RelativeOrderSysID[:]),
		ZCETotalTradedVolume: int(r.ZCETotalTradedVolume),
		IsSwapOrder:          int(r.IsSwapOrder),
		BranchID:             bytes2String(r.BranchID[:]),
		InvestUnitID:         bytes2String(r.InvestUnitID[:]),
		AccountID:            bytes2String(r.AccountID[:]),
		CurrencyID:           bytes2String(r.CurrencyID[:]),

		MacAddress:     bytes2String(r.MacAddress[:]),
		InstrumentID:   bytes2String(r.InstrumentID[:]),
		ExchangeInstID: bytes2String(r.ExchangeInstID[:]),
		IPAddress:      bytes2String(r.IPAddress[:]),
	}

	return f
}

//成交
func fromCThostFtdcTradeField(r *thost.CThostFtdcTradeField) *TradeField {

	f := &TradeField{

		BrokerID:   bytes2String(r.BrokerID[:]),
		InvestorID: bytes2String(r.InvestorID[:]),

		OrderRef:      bytes2String(r.OrderRef[:]),
		UserID:        bytes2String(r.UserID[:]),
		ExchangeID:    bytes2String(r.ExchangeID[:]),
		TradeID:       bytes2String(r.TradeID[:]),
		Direction:     byte(r.Direction),
		OrderSysID:    bytes2String(r.OrderSysID[:]),
		ParticipantID: bytes2String(r.ParticipantID[:]),
		ClientID:      bytes2String(r.ClientID[:]),
		TradingRole:   byte(r.TradingRole),

		OffsetFlag:     byte(r.OffsetFlag),
		HedgeFlag:      byte(r.HedgeFlag),
		Price:          float64(r.Price),
		Volume:         int(r.Volume),
		TradeDate:      bytes2String(r.TradeDate[:]),
		TradeTime:      bytes2String(r.TradeTime[:]),
		TradeType:      byte(r.TradeType),
		PriceSource:    byte(r.PriceSource),
		TraderID:       bytes2String(r.TraderID[:]),
		OrderLocalID:   bytes2String(r.OrderLocalID[:]),
		ClearingPartID: bytes2String(r.ClearingPartID[:]),
		BusinessUnit:   bytes2String(r.BusinessUnit[:]),
		SequenceNo:     int(r.SequenceNo),
		TradingDay:     bytes2String(r.TradingDay[:]),
		SettlementID:   int(r.SettlementID),
		BrokerOrderSeq: int(r.BrokerOrderSeq),
		TradeSource:    byte(r.TradeSource),
		InvestUnitID:   bytes2String(r.InvestUnitID[:]),
		InstrumentID:   bytes2String(r.InstrumentID[:]),
		ExchangeInstID: bytes2String(r.ExchangeInstID[:]),
	}

	return f
}

//输入报单操作
func fromCThostFtdcInputOrderActionField(r *thost.CThostFtdcInputOrderActionField) *InputOrderActionField {

	f := &InputOrderActionField{

		BrokerID:       bytes2String(r.BrokerID[:]),
		InvestorID:     bytes2String(r.InvestorID[:]),
		OrderActionRef: int(r.OrderActionRef),
		OrderRef:       bytes2String(r.OrderRef[:]),
		RequestID:      int(r.RequestID),
		FrontID:        int(r.FrontID),
		SessionID:      int(r.SessionID),
		ExchangeID:     bytes2String(r.ExchangeID[:]),
		OrderSysID:     bytes2String(r.OrderSysID[:]),
		ActionFlag:     byte(r.ActionFlag),
		LimitPrice:     float64(r.LimitPrice),
		VolumeChange:   int(r.VolumeChange),
		UserID:         bytes2String(r.UserID[:]),

		InvestUnitID: bytes2String(r.InvestUnitID[:]),

		MacAddress:   bytes2String(r.MacAddress[:]),
		InstrumentID: bytes2String(r.InstrumentID[:]),
		IPAddress:    bytes2String(r.IPAddress[:]),
	}

	return f
}

//输入报单
func fromCThostFtdcInputOrderField(r *thost.CThostFtdcInputOrderField) *InputOrderField {

	f := &InputOrderField{

		BrokerID:   bytes2String(r.BrokerID[:]),
		InvestorID: bytes2String(r.InvestorID[:]),

		OrderRef:            bytes2String(r.OrderRef[:]),
		UserID:              bytes2String(r.UserID[:]),
		OrderPriceType:      byte(r.OrderPriceType),
		Direction:           byte(r.Direction),
		CombOffsetFlag:      bytes2String(r.CombOffsetFlag[:]),
		CombHedgeFlag:       bytes2String(r.CombHedgeFlag[:]),
		LimitPrice:          float64(r.LimitPrice),
		VolumeTotalOriginal: int(r.VolumeTotalOriginal),
		TimeCondition:       byte(r.TimeCondition),
		GTDDate:             bytes2String(r.GTDDate[:]),
		VolumeCondition:     byte(r.VolumeCondition),
		MinVolume:           int(r.MinVolume),
		ContingentCondition: byte(r.ContingentCondition),
		StopPrice:           float64(r.StopPrice),
		ForceCloseReason:    byte(r.ForceCloseReason),
		IsAutoSuspend:       int(r.IsAutoSuspend),
		BusinessUnit:        bytes2String(r.BusinessUnit[:]),
		RequestID:           int(r.RequestID),
		UserForceClose:      int(r.UserForceClose),
		IsSwapOrder:         int(r.IsSwapOrder),
		ExchangeID:          bytes2String(r.ExchangeID[:]),
		InvestUnitID:        bytes2String(r.InvestUnitID[:]),
		AccountID:           bytes2String(r.AccountID[:]),
		CurrencyID:          bytes2String(r.CurrencyID[:]),
		ClientID:            bytes2String(r.ClientID[:]),

		MacAddress:   bytes2String(r.MacAddress[:]),
		InstrumentID: bytes2String(r.InstrumentID[:]),
		IPAddress:    bytes2String(r.IPAddress[:]),
	}

	return f
}

//报单操作
func fromCThostFtdcOrderActionField(r *thost.CThostFtdcOrderActionField) *OrderActionField {

	f := &OrderActionField{

		BrokerID:          bytes2String(r.BrokerID[:]),
		InvestorID:        bytes2String(r.InvestorID[:]),
		OrderActionRef:    int(r.OrderActionRef),
		OrderRef:          bytes2String(r.OrderRef[:]),
		RequestID:         int(r.RequestID),
		FrontID:           int(r.FrontID),
		SessionID:         int(r.SessionID),
		ExchangeID:        bytes2String(r.ExchangeID[:]),
		OrderSysID:        bytes2String(r.OrderSysID[:]),
		ActionFlag:        byte(r.ActionFlag),
		LimitPrice:        float64(r.LimitPrice),
		VolumeChange:      int(r.VolumeChange),
		ActionDate:        bytes2String(r.ActionDate[:]),
		ActionTime:        bytes2String(r.ActionTime[:]),
		TraderID:          bytes2String(r.TraderID[:]),
		InstallID:         int(r.InstallID),
		OrderLocalID:      bytes2String(r.OrderLocalID[:]),
		ActionLocalID:     bytes2String(r.ActionLocalID[:]),
		ParticipantID:     bytes2String(r.ParticipantID[:]),
		ClientID:          bytes2String(r.ClientID[:]),
		BusinessUnit:      bytes2String(r.BusinessUnit[:]),
		OrderActionStatus: byte(r.OrderActionStatus),
		UserID:            bytes2String(r.UserID[:]),
		StatusMsg:         bytes2String(r.StatusMsg[:]),

		BranchID:     bytes2String(r.BranchID[:]),
		InvestUnitID: bytes2String(r.InvestUnitID[:]),

		MacAddress:   bytes2String(r.MacAddress[:]),
		InstrumentID: bytes2String(r.InstrumentID[:]),
		IPAddress:    bytes2String(r.IPAddress[:]),
	}

	return f
}

//合约保证金率
func fromCThostFtdcInstrumentMarginRateField(r *thost.CThostFtdcInstrumentMarginRateField) *InstrumentMarginRateField {

	f := &InstrumentMarginRateField{

		InvestorRange:            byte(r.InvestorRange),
		BrokerID:                 bytes2String(r.BrokerID[:]),
		InvestorID:               bytes2String(r.InvestorID[:]),
		HedgeFlag:                byte(r.HedgeFlag),
		LongMarginRatioByMoney:   float64(r.LongMarginRatioByMoney),
		LongMarginRatioByVolume:  float64(r.LongMarginRatioByVolume),
		ShortMarginRatioByMoney:  float64(r.ShortMarginRatioByMoney),
		ShortMarginRatioByVolume: float64(r.ShortMarginRatioByVolume),
		IsRelative:               int(r.IsRelative),
		ExchangeID:               bytes2String(r.ExchangeID[:]),
		InvestUnitID:             bytes2String(r.InvestUnitID[:]),
		InstrumentID:             bytes2String(r.InstrumentID[:]),
	}

	return f
}

//合约手续费率
func fromCThostFtdcInstrumentCommissionRateField(r *thost.CThostFtdcInstrumentCommissionRateField) *InstrumentCommissionRateField {

	f := &InstrumentCommissionRateField{

		InvestorRange:           byte(r.InvestorRange),
		BrokerID:                bytes2String(r.BrokerID[:]),
		InvestorID:              bytes2String(r.InvestorID[:]),
		OpenRatioByMoney:        float64(r.OpenRatioByMoney),
		OpenRatioByVolume:       float64(r.OpenRatioByVolume),
		CloseRatioByMoney:       float64(r.CloseRatioByMoney),
		CloseRatioByVolume:      float64(r.CloseRatioByVolume),
		CloseTodayRatioByMoney:  float64(r.CloseTodayRatioByMoney),
		CloseTodayRatioByVolume: float64(r.CloseTodayRatioByVolume),
		ExchangeID:              bytes2String(r.ExchangeID[:]),
		BizType:                 byte(r.BizType),
		InvestUnitID:            bytes2String(r.InvestUnitID[:]),
		InstrumentID:            bytes2String(r.InstrumentID[:]),
	}

	return f
}

//合约状态
func fromCThostFtdcInstrumentStatusField(r *thost.CThostFtdcInstrumentStatusField) *InstrumentStatusField {

	f := &InstrumentStatusField{

		ExchangeID: bytes2String(r.ExchangeID[:]),

		SettlementGroupID: bytes2String(r.SettlementGroupID[:]),

		InstrumentStatus: byte(r.InstrumentStatus),
		TradingSegmentSN: int(r.TradingSegmentSN),
		EnterTime:        bytes2String(r.EnterTime[:]),
		EnterReason:      byte(r.EnterReason),
		ExchangeInstID:   bytes2String(r.ExchangeInstID[:]),
		InstrumentID:     bytes2String(r.InstrumentID[:]),
	}

	return f
}

//-------------------------------------------------------------------

func bytes2String(t []byte) string {
	msg, _ := simplifiedchinese.GB18030.NewDecoder().Bytes(t)
	return strings.Trim(string(msg), "\u0000")
}
