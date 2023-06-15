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
	"bytes"
	"strings"

	"github.com/pseudocodes/goctp/thost"
	"golang.org/x/text/encoding/simplifiedchinese"
)

// 报单操作
func toCThostFtdcOrderActionField(r *OrderActionField) *thost.CThostFtdcOrderActionField {

	f := &thost.CThostFtdcOrderActionField{}

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
	copy(f.ActionDate[:len(f.ActionDate)-1], []byte(r.ActionDate))
	copy(f.ActionTime[:len(f.ActionTime)-1], []byte(r.ActionTime))
	copy(f.TraderID[:len(f.TraderID)-1], []byte(r.TraderID))
	f.InstallID = thost.TThostFtdcInstallIDType(r.InstallID)
	copy(f.OrderLocalID[:len(f.OrderLocalID)-1], []byte(r.OrderLocalID))
	copy(f.ActionLocalID[:len(f.ActionLocalID)-1], []byte(r.ActionLocalID))
	copy(f.ParticipantID[:len(f.ParticipantID)-1], []byte(r.ParticipantID))
	copy(f.ClientID[:len(f.ClientID)-1], []byte(r.ClientID))
	copy(f.BusinessUnit[:len(f.BusinessUnit)-1], []byte(r.BusinessUnit))
	f.OrderActionStatus = thost.TThostFtdcOrderActionStatusType(r.OrderActionStatus)
	copy(f.UserID[:len(f.UserID)-1], []byte(r.UserID))
	copy(f.StatusMsg[:len(f.StatusMsg)-1], []byte(r.StatusMsg))

	copy(f.BranchID[:len(f.BranchID)-1], []byte(r.BranchID))
	copy(f.InvestUnitID[:len(f.InvestUnitID)-1], []byte(r.InvestUnitID))

	copy(f.MacAddress[:len(f.MacAddress)-1], []byte(r.MacAddress))
	copy(f.InstrumentID[:len(f.InstrumentID)-1], []byte(r.InstrumentID))
	copy(f.IPAddress[:len(f.IPAddress)-1], []byte(r.IPAddress))

	return f
}

// 查询报单
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

// 客户开销户信息表
func toCThostFtdcAccountregisterField(r *AccountregisterField) *thost.CThostFtdcAccountregisterField {

	f := &thost.CThostFtdcAccountregisterField{}

	copy(f.TradeDay[:len(f.TradeDay)-1], []byte(r.TradeDay))
	copy(f.BankID[:len(f.BankID)-1], []byte(r.BankID))
	copy(f.BankBranchID[:len(f.BankBranchID)-1], []byte(r.BankBranchID))
	copy(f.BankAccount[:len(f.BankAccount)-1], []byte(r.BankAccount))
	copy(f.BrokerID[:len(f.BrokerID)-1], []byte(r.BrokerID))
	copy(f.BrokerBranchID[:len(f.BrokerBranchID)-1], []byte(r.BrokerBranchID))
	copy(f.AccountID[:len(f.AccountID)-1], []byte(r.AccountID))
	f.IdCardType = thost.TThostFtdcIdCardTypeType(r.IdCardType)
	copy(f.IdentifiedCardNo[:len(f.IdentifiedCardNo)-1], []byte(r.IdentifiedCardNo))
	copy(f.CustomerName[:len(f.CustomerName)-1], []byte(r.CustomerName))
	copy(f.CurrencyID[:len(f.CurrencyID)-1], []byte(r.CurrencyID))
	f.OpenOrDestroy = thost.TThostFtdcOpenOrDestroyType(r.OpenOrDestroy)
	copy(f.RegDate[:len(f.RegDate)-1], []byte(r.RegDate))
	copy(f.OutDate[:len(f.OutDate)-1], []byte(r.OutDate))
	f.TID = thost.TThostFtdcTIDType(r.TID)
	f.CustType = thost.TThostFtdcCustTypeType(r.CustType)
	f.BankAccType = thost.TThostFtdcBankAccTypeType(r.BankAccType)
	copy(f.LongCustomerName[:len(f.LongCustomerName)-1], []byte(r.LongCustomerName))

	return f
}

// 投资者结算结果确认信息
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

// 查询投资者持仓
func toCThostFtdcQryInvestorPositionField(r *QryInvestorPositionField) *thost.CThostFtdcQryInvestorPositionField {

	f := &thost.CThostFtdcQryInvestorPositionField{}

	copy(f.BrokerID[:len(f.BrokerID)-1], []byte(r.BrokerID))
	copy(f.InvestorID[:len(f.InvestorID)-1], []byte(r.InvestorID))

	copy(f.ExchangeID[:len(f.ExchangeID)-1], []byte(r.ExchangeID))
	copy(f.InvestUnitID[:len(f.InvestUnitID)-1], []byte(r.InvestUnitID))
	copy(f.InstrumentID[:len(f.InstrumentID)-1], []byte(r.InstrumentID))

	return f
}

// 输入报单操作
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

// 银行发起银行资金转期货响应
func toCThostFtdcRspTransferField(r *RspTransferField) *thost.CThostFtdcRspTransferField {

	f := &thost.CThostFtdcRspTransferField{}

	copy(f.TradeCode[:len(f.TradeCode)-1], []byte(r.TradeCode))
	copy(f.BankID[:len(f.BankID)-1], []byte(r.BankID))
	copy(f.BankBranchID[:len(f.BankBranchID)-1], []byte(r.BankBranchID))
	copy(f.BrokerID[:len(f.BrokerID)-1], []byte(r.BrokerID))
	copy(f.BrokerBranchID[:len(f.BrokerBranchID)-1], []byte(r.BrokerBranchID))
	copy(f.TradeDate[:len(f.TradeDate)-1], []byte(r.TradeDate))
	copy(f.TradeTime[:len(f.TradeTime)-1], []byte(r.TradeTime))
	copy(f.BankSerial[:len(f.BankSerial)-1], []byte(r.BankSerial))
	copy(f.TradingDay[:len(f.TradingDay)-1], []byte(r.TradingDay))
	f.PlateSerial = thost.TThostFtdcSerialType(r.PlateSerial)
	f.LastFragment = thost.TThostFtdcLastFragmentType(r.LastFragment)
	f.SessionID = thost.TThostFtdcSessionIDType(r.SessionID)
	copy(f.CustomerName[:len(f.CustomerName)-1], []byte(r.CustomerName))
	f.IdCardType = thost.TThostFtdcIdCardTypeType(r.IdCardType)
	copy(f.IdentifiedCardNo[:len(f.IdentifiedCardNo)-1], []byte(r.IdentifiedCardNo))
	f.CustType = thost.TThostFtdcCustTypeType(r.CustType)
	copy(f.BankAccount[:len(f.BankAccount)-1], []byte(r.BankAccount))
	copy(f.BankPassWord[:len(f.BankPassWord)-1], []byte(r.BankPassWord))
	copy(f.AccountID[:len(f.AccountID)-1], []byte(r.AccountID))
	copy(f.Password[:len(f.Password)-1], []byte(r.Password))
	f.InstallID = thost.TThostFtdcInstallIDType(r.InstallID)
	f.FutureSerial = thost.TThostFtdcFutureSerialType(r.FutureSerial)
	copy(f.UserID[:len(f.UserID)-1], []byte(r.UserID))
	f.VerifyCertNoFlag = thost.TThostFtdcYesNoIndicatorType(r.VerifyCertNoFlag)
	copy(f.CurrencyID[:len(f.CurrencyID)-1], []byte(r.CurrencyID))
	f.TradeAmount = thost.TThostFtdcTradeAmountType(r.TradeAmount)
	f.FutureFetchAmount = thost.TThostFtdcTradeAmountType(r.FutureFetchAmount)
	f.FeePayFlag = thost.TThostFtdcFeePayFlagType(r.FeePayFlag)
	f.CustFee = thost.TThostFtdcCustFeeType(r.CustFee)
	f.BrokerFee = thost.TThostFtdcFutureFeeType(r.BrokerFee)
	copy(f.Message[:len(f.Message)-1], []byte(r.Message))
	copy(f.Digest[:len(f.Digest)-1], []byte(r.Digest))
	f.BankAccType = thost.TThostFtdcBankAccTypeType(r.BankAccType)
	copy(f.DeviceID[:len(f.DeviceID)-1], []byte(r.DeviceID))
	f.BankSecuAccType = thost.TThostFtdcBankAccTypeType(r.BankSecuAccType)
	copy(f.BrokerIDByBank[:len(f.BrokerIDByBank)-1], []byte(r.BrokerIDByBank))
	copy(f.BankSecuAcc[:len(f.BankSecuAcc)-1], []byte(r.BankSecuAcc))
	f.BankPwdFlag = thost.TThostFtdcPwdFlagType(r.BankPwdFlag)
	f.SecuPwdFlag = thost.TThostFtdcPwdFlagType(r.SecuPwdFlag)
	copy(f.OperNo[:len(f.OperNo)-1], []byte(r.OperNo))
	f.RequestID = thost.TThostFtdcRequestIDType(r.RequestID)
	f.TID = thost.TThostFtdcTIDType(r.TID)
	f.TransferStatus = thost.TThostFtdcTransferStatusType(r.TransferStatus)
	f.ErrorID = thost.TThostFtdcErrorIDType(r.ErrorID)
	copy(f.ErrorMsg[:len(f.ErrorMsg)-1], []byte(r.ErrorMsg))
	copy(f.LongCustomerName[:len(f.LongCustomerName)-1], []byte(r.LongCustomerName))

	return f
}

// 请求查询银期签约关系
func toCThostFtdcQryAccountregisterField(r *QryAccountregisterField) *thost.CThostFtdcQryAccountregisterField {

	f := &thost.CThostFtdcQryAccountregisterField{}

	copy(f.BrokerID[:len(f.BrokerID)-1], []byte(r.BrokerID))
	copy(f.AccountID[:len(f.AccountID)-1], []byte(r.AccountID))
	copy(f.BankID[:len(f.BankID)-1], []byte(r.BankID))
	copy(f.BankBranchID[:len(f.BankBranchID)-1], []byte(r.BankBranchID))
	copy(f.CurrencyID[:len(f.CurrencyID)-1], []byte(r.CurrencyID))

	return f
}

// 合约
func toCThostFtdcInstrumentField(r *InstrumentField) *thost.CThostFtdcInstrumentField {

	f := &thost.CThostFtdcInstrumentField{}

	copy(f.ExchangeID[:len(f.ExchangeID)-1], []byte(r.ExchangeID))
	copy(f.InstrumentName[:len(f.InstrumentName)-1], []byte(r.InstrumentName))

	f.ProductClass = thost.TThostFtdcProductClassType(r.ProductClass)
	f.DeliveryYear = thost.TThostFtdcYearType(r.DeliveryYear)
	f.DeliveryMonth = thost.TThostFtdcMonthType(r.DeliveryMonth)
	f.MaxMarketOrderVolume = thost.TThostFtdcVolumeType(r.MaxMarketOrderVolume)
	f.MinMarketOrderVolume = thost.TThostFtdcVolumeType(r.MinMarketOrderVolume)
	f.MaxLimitOrderVolume = thost.TThostFtdcVolumeType(r.MaxLimitOrderVolume)
	f.MinLimitOrderVolume = thost.TThostFtdcVolumeType(r.MinLimitOrderVolume)
	f.VolumeMultiple = thost.TThostFtdcVolumeMultipleType(r.VolumeMultiple)
	f.PriceTick = thost.TThostFtdcPriceType(r.PriceTick)
	copy(f.CreateDate[:len(f.CreateDate)-1], []byte(r.CreateDate))
	copy(f.OpenDate[:len(f.OpenDate)-1], []byte(r.OpenDate))
	copy(f.ExpireDate[:len(f.ExpireDate)-1], []byte(r.ExpireDate))
	copy(f.StartDelivDate[:len(f.StartDelivDate)-1], []byte(r.StartDelivDate))
	copy(f.EndDelivDate[:len(f.EndDelivDate)-1], []byte(r.EndDelivDate))
	f.InstLifePhase = thost.TThostFtdcInstLifePhaseType(r.InstLifePhase)
	f.IsTrading = thost.TThostFtdcBoolType(r.IsTrading)
	f.PositionType = thost.TThostFtdcPositionTypeType(r.PositionType)
	f.PositionDateType = thost.TThostFtdcPositionDateTypeType(r.PositionDateType)
	f.LongMarginRatio = thost.TThostFtdcRatioType(r.LongMarginRatio)
	f.ShortMarginRatio = thost.TThostFtdcRatioType(r.ShortMarginRatio)
	f.MaxMarginSideAlgorithm = thost.TThostFtdcMaxMarginSideAlgorithmType(r.MaxMarginSideAlgorithm)

	f.StrikePrice = thost.TThostFtdcPriceType(r.StrikePrice)
	f.OptionsType = thost.TThostFtdcOptionsTypeType(r.OptionsType)
	f.UnderlyingMultiple = thost.TThostFtdcUnderlyingMultipleType(r.UnderlyingMultiple)
	f.CombinationType = thost.TThostFtdcCombinationTypeType(r.CombinationType)
	copy(f.InstrumentID[:len(f.InstrumentID)-1], []byte(r.InstrumentID))
	copy(f.ExchangeInstID[:len(f.ExchangeInstID)-1], []byte(r.ExchangeInstID))
	copy(f.ProductID[:len(f.ProductID)-1], []byte(r.ProductID))
	copy(f.UnderlyingInstrID[:len(f.UnderlyingInstrID)-1], []byte(r.UnderlyingInstrID))

	return f
}

// 资金账户口令变更域
func toCThostFtdcTradingAccountPasswordUpdateField(r *TradingAccountPasswordUpdateField) *thost.CThostFtdcTradingAccountPasswordUpdateField {

	f := &thost.CThostFtdcTradingAccountPasswordUpdateField{}

	copy(f.BrokerID[:len(f.BrokerID)-1], []byte(r.BrokerID))
	copy(f.AccountID[:len(f.AccountID)-1], []byte(r.AccountID))
	copy(f.OldPassword[:len(f.OldPassword)-1], []byte(r.OldPassword))
	copy(f.NewPassword[:len(f.NewPassword)-1], []byte(r.NewPassword))
	copy(f.CurrencyID[:len(f.CurrencyID)-1], []byte(r.CurrencyID))

	return f
}

// 响应信息
func toCThostFtdcRspInfoField(r *RspInfoField) *thost.CThostFtdcRspInfoField {

	f := &thost.CThostFtdcRspInfoField{}

	f.ErrorID = thost.TThostFtdcErrorIDType(r.ErrorID)
	copy(f.ErrorMsg[:len(f.ErrorMsg)-1], []byte(r.ErrorMsg))

	return f
}

// 查询签约银行请求
func toCThostFtdcQryContractBankField(r *QryContractBankField) *thost.CThostFtdcQryContractBankField {

	f := &thost.CThostFtdcQryContractBankField{}

	copy(f.BrokerID[:len(f.BrokerID)-1], []byte(r.BrokerID))
	copy(f.BankID[:len(f.BankID)-1], []byte(r.BankID))
	copy(f.BankBrchID[:len(f.BankBrchID)-1], []byte(r.BankBrchID))

	return f
}

// 查询经纪公司交易参数
func toCThostFtdcQryBrokerTradingParamsField(r *QryBrokerTradingParamsField) *thost.CThostFtdcQryBrokerTradingParamsField {

	f := &thost.CThostFtdcQryBrokerTradingParamsField{}

	copy(f.BrokerID[:len(f.BrokerID)-1], []byte(r.BrokerID))
	copy(f.InvestorID[:len(f.InvestorID)-1], []byte(r.InvestorID))
	copy(f.CurrencyID[:len(f.CurrencyID)-1], []byte(r.CurrencyID))
	copy(f.AccountID[:len(f.AccountID)-1], []byte(r.AccountID))

	return f
}

// 查询手续费率
func toCThostFtdcQryInstrumentCommissionRateField(r *QryInstrumentCommissionRateField) *thost.CThostFtdcQryInstrumentCommissionRateField {

	f := &thost.CThostFtdcQryInstrumentCommissionRateField{}

	copy(f.BrokerID[:len(f.BrokerID)-1], []byte(r.BrokerID))
	copy(f.InvestorID[:len(f.InvestorID)-1], []byte(r.InvestorID))

	copy(f.ExchangeID[:len(f.ExchangeID)-1], []byte(r.ExchangeID))
	copy(f.InvestUnitID[:len(f.InvestUnitID)-1], []byte(r.InvestUnitID))
	copy(f.InstrumentID[:len(f.InstrumentID)-1], []byte(r.InstrumentID))

	return f
}

// 客户端认证响应
func toCThostFtdcRspAuthenticateField(r *RspAuthenticateField) *thost.CThostFtdcRspAuthenticateField {

	f := &thost.CThostFtdcRspAuthenticateField{}

	copy(f.BrokerID[:len(f.BrokerID)-1], []byte(r.BrokerID))
	copy(f.UserID[:len(f.UserID)-1], []byte(r.UserID))
	copy(f.UserProductInfo[:len(f.UserProductInfo)-1], []byte(r.UserProductInfo))
	copy(f.AppID[:len(f.AppID)-1], []byte(r.AppID))
	f.AppType = thost.TThostFtdcAppTypeType(r.AppType)

	return f
}

// 银期转账交易流水表
func toCThostFtdcTransferSerialField(r *TransferSerialField) *thost.CThostFtdcTransferSerialField {

	f := &thost.CThostFtdcTransferSerialField{}

	f.PlateSerial = thost.TThostFtdcPlateSerialType(r.PlateSerial)
	copy(f.TradeDate[:len(f.TradeDate)-1], []byte(r.TradeDate))
	copy(f.TradingDay[:len(f.TradingDay)-1], []byte(r.TradingDay))
	copy(f.TradeTime[:len(f.TradeTime)-1], []byte(r.TradeTime))
	copy(f.TradeCode[:len(f.TradeCode)-1], []byte(r.TradeCode))
	f.SessionID = thost.TThostFtdcSessionIDType(r.SessionID)
	copy(f.BankID[:len(f.BankID)-1], []byte(r.BankID))
	copy(f.BankBranchID[:len(f.BankBranchID)-1], []byte(r.BankBranchID))
	f.BankAccType = thost.TThostFtdcBankAccTypeType(r.BankAccType)
	copy(f.BankAccount[:len(f.BankAccount)-1], []byte(r.BankAccount))
	copy(f.BankSerial[:len(f.BankSerial)-1], []byte(r.BankSerial))
	copy(f.BrokerID[:len(f.BrokerID)-1], []byte(r.BrokerID))
	copy(f.BrokerBranchID[:len(f.BrokerBranchID)-1], []byte(r.BrokerBranchID))
	f.FutureAccType = thost.TThostFtdcFutureAccTypeType(r.FutureAccType)
	copy(f.AccountID[:len(f.AccountID)-1], []byte(r.AccountID))
	copy(f.InvestorID[:len(f.InvestorID)-1], []byte(r.InvestorID))
	f.FutureSerial = thost.TThostFtdcFutureSerialType(r.FutureSerial)
	f.IdCardType = thost.TThostFtdcIdCardTypeType(r.IdCardType)
	copy(f.IdentifiedCardNo[:len(f.IdentifiedCardNo)-1], []byte(r.IdentifiedCardNo))
	copy(f.CurrencyID[:len(f.CurrencyID)-1], []byte(r.CurrencyID))
	f.TradeAmount = thost.TThostFtdcTradeAmountType(r.TradeAmount)
	f.CustFee = thost.TThostFtdcCustFeeType(r.CustFee)
	f.BrokerFee = thost.TThostFtdcFutureFeeType(r.BrokerFee)
	f.AvailabilityFlag = thost.TThostFtdcAvailabilityFlagType(r.AvailabilityFlag)
	copy(f.OperatorCode[:len(f.OperatorCode)-1], []byte(r.OperatorCode))
	copy(f.BankNewAccount[:len(f.BankNewAccount)-1], []byte(r.BankNewAccount))
	f.ErrorID = thost.TThostFtdcErrorIDType(r.ErrorID)
	copy(f.ErrorMsg[:len(f.ErrorMsg)-1], []byte(r.ErrorMsg))

	return f
}

// 查询投资者结算结果
func toCThostFtdcQrySettlementInfoField(r *QrySettlementInfoField) *thost.CThostFtdcQrySettlementInfoField {

	f := &thost.CThostFtdcQrySettlementInfoField{}

	copy(f.BrokerID[:len(f.BrokerID)-1], []byte(r.BrokerID))
	copy(f.InvestorID[:len(f.InvestorID)-1], []byte(r.InvestorID))
	copy(f.TradingDay[:len(f.TradingDay)-1], []byte(r.TradingDay))
	copy(f.AccountID[:len(f.AccountID)-1], []byte(r.AccountID))
	copy(f.CurrencyID[:len(f.CurrencyID)-1], []byte(r.CurrencyID))

	return f
}

// 查询成交
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

// 用户登录应答
func toCThostFtdcRspUserLoginField(r *RspUserLoginField) *thost.CThostFtdcRspUserLoginField {

	f := &thost.CThostFtdcRspUserLoginField{}

	copy(f.TradingDay[:len(f.TradingDay)-1], []byte(r.TradingDay))
	copy(f.LoginTime[:len(f.LoginTime)-1], []byte(r.LoginTime))
	copy(f.BrokerID[:len(f.BrokerID)-1], []byte(r.BrokerID))
	copy(f.UserID[:len(f.UserID)-1], []byte(r.UserID))
	copy(f.SystemName[:len(f.SystemName)-1], []byte(r.SystemName))
	f.FrontID = thost.TThostFtdcFrontIDType(r.FrontID)
	f.SessionID = thost.TThostFtdcSessionIDType(r.SessionID)
	copy(f.MaxOrderRef[:len(f.MaxOrderRef)-1], []byte(r.MaxOrderRef))
	copy(f.SHFETime[:len(f.SHFETime)-1], []byte(r.SHFETime))
	copy(f.DCETime[:len(f.DCETime)-1], []byte(r.DCETime))
	copy(f.CZCETime[:len(f.CZCETime)-1], []byte(r.CZCETime))
	copy(f.FFEXTime[:len(f.FFEXTime)-1], []byte(r.FFEXTime))
	copy(f.INETime[:len(f.INETime)-1], []byte(r.INETime))

	return f
}

// 经纪公司交易参数
func toCThostFtdcBrokerTradingParamsField(r *BrokerTradingParamsField) *thost.CThostFtdcBrokerTradingParamsField {

	f := &thost.CThostFtdcBrokerTradingParamsField{}

	copy(f.BrokerID[:len(f.BrokerID)-1], []byte(r.BrokerID))
	copy(f.InvestorID[:len(f.InvestorID)-1], []byte(r.InvestorID))
	f.MarginPriceType = thost.TThostFtdcMarginPriceTypeType(r.MarginPriceType)
	f.Algorithm = thost.TThostFtdcAlgorithmType(r.Algorithm)
	f.AvailIncludeCloseProfit = thost.TThostFtdcIncludeCloseProfitType(r.AvailIncludeCloseProfit)
	copy(f.CurrencyID[:len(f.CurrencyID)-1], []byte(r.CurrencyID))
	f.OptionRoyaltyPriceType = thost.TThostFtdcOptionRoyaltyPriceTypeType(r.OptionRoyaltyPriceType)
	copy(f.AccountID[:len(f.AccountID)-1], []byte(r.AccountID))

	return f
}

// 请求查询转帐流水
func toCThostFtdcQryTransferSerialField(r *QryTransferSerialField) *thost.CThostFtdcQryTransferSerialField {

	f := &thost.CThostFtdcQryTransferSerialField{}

	copy(f.BrokerID[:len(f.BrokerID)-1], []byte(r.BrokerID))
	copy(f.AccountID[:len(f.AccountID)-1], []byte(r.AccountID))
	copy(f.BankID[:len(f.BankID)-1], []byte(r.BankID))
	copy(f.CurrencyID[:len(f.CurrencyID)-1], []byte(r.CurrencyID))

	return f
}

// 查询签约银行响应
func toCThostFtdcContractBankField(r *ContractBankField) *thost.CThostFtdcContractBankField {

	f := &thost.CThostFtdcContractBankField{}

	copy(f.BrokerID[:len(f.BrokerID)-1], []byte(r.BrokerID))
	copy(f.BankID[:len(f.BankID)-1], []byte(r.BankID))
	copy(f.BankBrchID[:len(f.BankBrchID)-1], []byte(r.BankBrchID))
	copy(f.BankName[:len(f.BankName)-1], []byte(r.BankName))

	return f
}

// 查询结算信息确认域
func toCThostFtdcQrySettlementInfoConfirmField(r *QrySettlementInfoConfirmField) *thost.CThostFtdcQrySettlementInfoConfirmField {

	f := &thost.CThostFtdcQrySettlementInfoConfirmField{}

	copy(f.BrokerID[:len(f.BrokerID)-1], []byte(r.BrokerID))
	copy(f.InvestorID[:len(f.InvestorID)-1], []byte(r.InvestorID))
	copy(f.AccountID[:len(f.AccountID)-1], []byte(r.AccountID))
	copy(f.CurrencyID[:len(f.CurrencyID)-1], []byte(r.CurrencyID))

	return f
}

// 查询合约
func toCThostFtdcQryInstrumentField(r *QryInstrumentField) *thost.CThostFtdcQryInstrumentField {

	f := &thost.CThostFtdcQryInstrumentField{}

	copy(f.ExchangeID[:len(f.ExchangeID)-1], []byte(r.ExchangeID))

	copy(f.InstrumentID[:len(f.InstrumentID)-1], []byte(r.InstrumentID))
	copy(f.ExchangeInstID[:len(f.ExchangeInstID)-1], []byte(r.ExchangeInstID))
	copy(f.ProductID[:len(f.ProductID)-1], []byte(r.ProductID))

	return f
}

// 资金账户
func toCThostFtdcTradingAccountField(r *TradingAccountField) *thost.CThostFtdcTradingAccountField {

	f := &thost.CThostFtdcTradingAccountField{}

	copy(f.BrokerID[:len(f.BrokerID)-1], []byte(r.BrokerID))
	copy(f.AccountID[:len(f.AccountID)-1], []byte(r.AccountID))
	f.PreMortgage = thost.TThostFtdcMoneyType(r.PreMortgage)
	f.PreCredit = thost.TThostFtdcMoneyType(r.PreCredit)
	f.PreDeposit = thost.TThostFtdcMoneyType(r.PreDeposit)
	f.PreBalance = thost.TThostFtdcMoneyType(r.PreBalance)
	f.PreMargin = thost.TThostFtdcMoneyType(r.PreMargin)
	f.InterestBase = thost.TThostFtdcMoneyType(r.InterestBase)
	f.Interest = thost.TThostFtdcMoneyType(r.Interest)
	f.Deposit = thost.TThostFtdcMoneyType(r.Deposit)
	f.Withdraw = thost.TThostFtdcMoneyType(r.Withdraw)
	f.FrozenMargin = thost.TThostFtdcMoneyType(r.FrozenMargin)
	f.FrozenCash = thost.TThostFtdcMoneyType(r.FrozenCash)
	f.FrozenCommission = thost.TThostFtdcMoneyType(r.FrozenCommission)
	f.CurrMargin = thost.TThostFtdcMoneyType(r.CurrMargin)
	f.CashIn = thost.TThostFtdcMoneyType(r.CashIn)
	f.Commission = thost.TThostFtdcMoneyType(r.Commission)
	f.CloseProfit = thost.TThostFtdcMoneyType(r.CloseProfit)
	f.PositionProfit = thost.TThostFtdcMoneyType(r.PositionProfit)
	f.Balance = thost.TThostFtdcMoneyType(r.Balance)
	f.Available = thost.TThostFtdcMoneyType(r.Available)
	f.WithdrawQuota = thost.TThostFtdcMoneyType(r.WithdrawQuota)
	f.Reserve = thost.TThostFtdcMoneyType(r.Reserve)
	copy(f.TradingDay[:len(f.TradingDay)-1], []byte(r.TradingDay))
	f.SettlementID = thost.TThostFtdcSettlementIDType(r.SettlementID)
	f.Credit = thost.TThostFtdcMoneyType(r.Credit)
	f.Mortgage = thost.TThostFtdcMoneyType(r.Mortgage)
	f.ExchangeMargin = thost.TThostFtdcMoneyType(r.ExchangeMargin)
	f.DeliveryMargin = thost.TThostFtdcMoneyType(r.DeliveryMargin)
	f.ExchangeDeliveryMargin = thost.TThostFtdcMoneyType(r.ExchangeDeliveryMargin)
	f.ReserveBalance = thost.TThostFtdcMoneyType(r.ReserveBalance)
	copy(f.CurrencyID[:len(f.CurrencyID)-1], []byte(r.CurrencyID))
	f.PreFundMortgageIn = thost.TThostFtdcMoneyType(r.PreFundMortgageIn)
	f.PreFundMortgageOut = thost.TThostFtdcMoneyType(r.PreFundMortgageOut)
	f.FundMortgageIn = thost.TThostFtdcMoneyType(r.FundMortgageIn)
	f.FundMortgageOut = thost.TThostFtdcMoneyType(r.FundMortgageOut)
	f.FundMortgageAvailable = thost.TThostFtdcMoneyType(r.FundMortgageAvailable)
	f.MortgageableFund = thost.TThostFtdcMoneyType(r.MortgageableFund)
	f.SpecProductMargin = thost.TThostFtdcMoneyType(r.SpecProductMargin)
	f.SpecProductFrozenMargin = thost.TThostFtdcMoneyType(r.SpecProductFrozenMargin)
	f.SpecProductCommission = thost.TThostFtdcMoneyType(r.SpecProductCommission)
	f.SpecProductFrozenCommission = thost.TThostFtdcMoneyType(r.SpecProductFrozenCommission)
	f.SpecProductPositionProfit = thost.TThostFtdcMoneyType(r.SpecProductPositionProfit)
	f.SpecProductCloseProfit = thost.TThostFtdcMoneyType(r.SpecProductCloseProfit)
	f.SpecProductPositionProfitByAlg = thost.TThostFtdcMoneyType(r.SpecProductPositionProfitByAlg)
	f.SpecProductExchangeMargin = thost.TThostFtdcMoneyType(r.SpecProductExchangeMargin)
	f.BizType = thost.TThostFtdcBizTypeType(r.BizType)
	f.FrozenSwap = thost.TThostFtdcMoneyType(r.FrozenSwap)
	f.RemainSwap = thost.TThostFtdcMoneyType(r.RemainSwap)

	return f
}

// 深度行情
func toCThostFtdcDepthMarketDataField(r *DepthMarketDataField) *thost.CThostFtdcDepthMarketDataField {

	f := &thost.CThostFtdcDepthMarketDataField{}

	copy(f.TradingDay[:len(f.TradingDay)-1], []byte(r.TradingDay))

	copy(f.ExchangeID[:len(f.ExchangeID)-1], []byte(r.ExchangeID))

	f.LastPrice = thost.TThostFtdcPriceType(r.LastPrice)
	f.PreSettlementPrice = thost.TThostFtdcPriceType(r.PreSettlementPrice)
	f.PreClosePrice = thost.TThostFtdcPriceType(r.PreClosePrice)
	f.PreOpenInterest = thost.TThostFtdcLargeVolumeType(r.PreOpenInterest)
	f.OpenPrice = thost.TThostFtdcPriceType(r.OpenPrice)
	f.HighestPrice = thost.TThostFtdcPriceType(r.HighestPrice)
	f.LowestPrice = thost.TThostFtdcPriceType(r.LowestPrice)
	f.Volume = thost.TThostFtdcVolumeType(r.Volume)
	f.Turnover = thost.TThostFtdcMoneyType(r.Turnover)
	f.OpenInterest = thost.TThostFtdcLargeVolumeType(r.OpenInterest)
	f.ClosePrice = thost.TThostFtdcPriceType(r.ClosePrice)
	f.SettlementPrice = thost.TThostFtdcPriceType(r.SettlementPrice)
	f.UpperLimitPrice = thost.TThostFtdcPriceType(r.UpperLimitPrice)
	f.LowerLimitPrice = thost.TThostFtdcPriceType(r.LowerLimitPrice)
	f.PreDelta = thost.TThostFtdcRatioType(r.PreDelta)
	f.CurrDelta = thost.TThostFtdcRatioType(r.CurrDelta)
	copy(f.UpdateTime[:len(f.UpdateTime)-1], []byte(r.UpdateTime))
	f.UpdateMillisec = thost.TThostFtdcMillisecType(r.UpdateMillisec)
	f.BidPrice1 = thost.TThostFtdcPriceType(r.BidPrice1)
	f.BidVolume1 = thost.TThostFtdcVolumeType(r.BidVolume1)
	f.AskPrice1 = thost.TThostFtdcPriceType(r.AskPrice1)
	f.AskVolume1 = thost.TThostFtdcVolumeType(r.AskVolume1)
	f.BidPrice2 = thost.TThostFtdcPriceType(r.BidPrice2)
	f.BidVolume2 = thost.TThostFtdcVolumeType(r.BidVolume2)
	f.AskPrice2 = thost.TThostFtdcPriceType(r.AskPrice2)
	f.AskVolume2 = thost.TThostFtdcVolumeType(r.AskVolume2)
	f.BidPrice3 = thost.TThostFtdcPriceType(r.BidPrice3)
	f.BidVolume3 = thost.TThostFtdcVolumeType(r.BidVolume3)
	f.AskPrice3 = thost.TThostFtdcPriceType(r.AskPrice3)
	f.AskVolume3 = thost.TThostFtdcVolumeType(r.AskVolume3)
	f.BidPrice4 = thost.TThostFtdcPriceType(r.BidPrice4)
	f.BidVolume4 = thost.TThostFtdcVolumeType(r.BidVolume4)
	f.AskPrice4 = thost.TThostFtdcPriceType(r.AskPrice4)
	f.AskVolume4 = thost.TThostFtdcVolumeType(r.AskVolume4)
	f.BidPrice5 = thost.TThostFtdcPriceType(r.BidPrice5)
	f.BidVolume5 = thost.TThostFtdcVolumeType(r.BidVolume5)
	f.AskPrice5 = thost.TThostFtdcPriceType(r.AskPrice5)
	f.AskVolume5 = thost.TThostFtdcVolumeType(r.AskVolume5)
	f.AveragePrice = thost.TThostFtdcPriceType(r.AveragePrice)
	copy(f.ActionDay[:len(f.ActionDay)-1], []byte(r.ActionDay))
	copy(f.InstrumentID[:len(f.InstrumentID)-1], []byte(r.InstrumentID))
	copy(f.ExchangeInstID[:len(f.ExchangeInstID)-1], []byte(r.ExchangeInstID))
	f.BandingUpperPrice = thost.TThostFtdcPriceType(r.BandingUpperPrice)
	f.BandingLowerPrice = thost.TThostFtdcPriceType(r.BandingLowerPrice)

	return f
}

// 用户登录请求
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

// 报单
func toCThostFtdcOrderField(r *OrderField) *thost.CThostFtdcOrderField {

	f := &thost.CThostFtdcOrderField{}

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
	copy(f.OrderLocalID[:len(f.OrderLocalID)-1], []byte(r.OrderLocalID))
	copy(f.ExchangeID[:len(f.ExchangeID)-1], []byte(r.ExchangeID))
	copy(f.ParticipantID[:len(f.ParticipantID)-1], []byte(r.ParticipantID))
	copy(f.ClientID[:len(f.ClientID)-1], []byte(r.ClientID))

	copy(f.TraderID[:len(f.TraderID)-1], []byte(r.TraderID))
	f.InstallID = thost.TThostFtdcInstallIDType(r.InstallID)
	f.OrderSubmitStatus = thost.TThostFtdcOrderSubmitStatusType(r.OrderSubmitStatus)
	f.NotifySequence = thost.TThostFtdcSequenceNoType(r.NotifySequence)
	copy(f.TradingDay[:len(f.TradingDay)-1], []byte(r.TradingDay))
	f.SettlementID = thost.TThostFtdcSettlementIDType(r.SettlementID)
	copy(f.OrderSysID[:len(f.OrderSysID)-1], []byte(r.OrderSysID))
	f.OrderSource = thost.TThostFtdcOrderSourceType(r.OrderSource)
	f.OrderStatus = thost.TThostFtdcOrderStatusType(r.OrderStatus)
	f.OrderType = thost.TThostFtdcOrderTypeType(r.OrderType)
	f.VolumeTraded = thost.TThostFtdcVolumeType(r.VolumeTraded)
	f.VolumeTotal = thost.TThostFtdcVolumeType(r.VolumeTotal)
	copy(f.InsertDate[:len(f.InsertDate)-1], []byte(r.InsertDate))
	copy(f.InsertTime[:len(f.InsertTime)-1], []byte(r.InsertTime))
	copy(f.ActiveTime[:len(f.ActiveTime)-1], []byte(r.ActiveTime))
	copy(f.SuspendTime[:len(f.SuspendTime)-1], []byte(r.SuspendTime))
	copy(f.UpdateTime[:len(f.UpdateTime)-1], []byte(r.UpdateTime))
	copy(f.CancelTime[:len(f.CancelTime)-1], []byte(r.CancelTime))
	copy(f.ActiveTraderID[:len(f.ActiveTraderID)-1], []byte(r.ActiveTraderID))
	copy(f.ClearingPartID[:len(f.ClearingPartID)-1], []byte(r.ClearingPartID))
	f.SequenceNo = thost.TThostFtdcSequenceNoType(r.SequenceNo)
	f.FrontID = thost.TThostFtdcFrontIDType(r.FrontID)
	f.SessionID = thost.TThostFtdcSessionIDType(r.SessionID)
	copy(f.UserProductInfo[:len(f.UserProductInfo)-1], []byte(r.UserProductInfo))
	copy(f.StatusMsg[:len(f.StatusMsg)-1], []byte(r.StatusMsg))
	f.UserForceClose = thost.TThostFtdcBoolType(r.UserForceClose)
	copy(f.ActiveUserID[:len(f.ActiveUserID)-1], []byte(r.ActiveUserID))
	f.BrokerOrderSeq = thost.TThostFtdcSequenceNoType(r.BrokerOrderSeq)
	copy(f.RelativeOrderSysID[:len(f.RelativeOrderSysID)-1], []byte(r.RelativeOrderSysID))
	f.ZCETotalTradedVolume = thost.TThostFtdcVolumeType(r.ZCETotalTradedVolume)
	f.IsSwapOrder = thost.TThostFtdcBoolType(r.IsSwapOrder)
	copy(f.BranchID[:len(f.BranchID)-1], []byte(r.BranchID))
	copy(f.InvestUnitID[:len(f.InvestUnitID)-1], []byte(r.InvestUnitID))
	copy(f.AccountID[:len(f.AccountID)-1], []byte(r.AccountID))
	copy(f.CurrencyID[:len(f.CurrencyID)-1], []byte(r.CurrencyID))

	copy(f.MacAddress[:len(f.MacAddress)-1], []byte(r.MacAddress))
	copy(f.InstrumentID[:len(f.InstrumentID)-1], []byte(r.InstrumentID))
	copy(f.ExchangeInstID[:len(f.ExchangeInstID)-1], []byte(r.ExchangeInstID))
	copy(f.IPAddress[:len(f.IPAddress)-1], []byte(r.IPAddress))

	return f
}

// 输入报单
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

// 投资者持仓
func toCThostFtdcInvestorPositionField(r *InvestorPositionField) *thost.CThostFtdcInvestorPositionField {

	f := &thost.CThostFtdcInvestorPositionField{}

	copy(f.BrokerID[:len(f.BrokerID)-1], []byte(r.BrokerID))
	copy(f.InvestorID[:len(f.InvestorID)-1], []byte(r.InvestorID))
	f.PosiDirection = thost.TThostFtdcPosiDirectionType(r.PosiDirection)
	f.HedgeFlag = thost.TThostFtdcHedgeFlagType(r.HedgeFlag)
	f.PositionDate = thost.TThostFtdcPositionDateType(r.PositionDate)
	f.YdPosition = thost.TThostFtdcVolumeType(r.YdPosition)
	f.Position = thost.TThostFtdcVolumeType(r.Position)
	f.LongFrozen = thost.TThostFtdcVolumeType(r.LongFrozen)
	f.ShortFrozen = thost.TThostFtdcVolumeType(r.ShortFrozen)
	f.LongFrozenAmount = thost.TThostFtdcMoneyType(r.LongFrozenAmount)
	f.ShortFrozenAmount = thost.TThostFtdcMoneyType(r.ShortFrozenAmount)
	f.OpenVolume = thost.TThostFtdcVolumeType(r.OpenVolume)
	f.CloseVolume = thost.TThostFtdcVolumeType(r.CloseVolume)
	f.OpenAmount = thost.TThostFtdcMoneyType(r.OpenAmount)
	f.CloseAmount = thost.TThostFtdcMoneyType(r.CloseAmount)
	f.PositionCost = thost.TThostFtdcMoneyType(r.PositionCost)
	f.PreMargin = thost.TThostFtdcMoneyType(r.PreMargin)
	f.UseMargin = thost.TThostFtdcMoneyType(r.UseMargin)
	f.FrozenMargin = thost.TThostFtdcMoneyType(r.FrozenMargin)
	f.FrozenCash = thost.TThostFtdcMoneyType(r.FrozenCash)
	f.FrozenCommission = thost.TThostFtdcMoneyType(r.FrozenCommission)
	f.CashIn = thost.TThostFtdcMoneyType(r.CashIn)
	f.Commission = thost.TThostFtdcMoneyType(r.Commission)
	f.CloseProfit = thost.TThostFtdcMoneyType(r.CloseProfit)
	f.PositionProfit = thost.TThostFtdcMoneyType(r.PositionProfit)
	f.PreSettlementPrice = thost.TThostFtdcPriceType(r.PreSettlementPrice)
	f.SettlementPrice = thost.TThostFtdcPriceType(r.SettlementPrice)
	copy(f.TradingDay[:len(f.TradingDay)-1], []byte(r.TradingDay))
	f.SettlementID = thost.TThostFtdcSettlementIDType(r.SettlementID)
	f.OpenCost = thost.TThostFtdcMoneyType(r.OpenCost)
	f.ExchangeMargin = thost.TThostFtdcMoneyType(r.ExchangeMargin)
	f.CombPosition = thost.TThostFtdcVolumeType(r.CombPosition)
	f.CombLongFrozen = thost.TThostFtdcVolumeType(r.CombLongFrozen)
	f.CombShortFrozen = thost.TThostFtdcVolumeType(r.CombShortFrozen)
	f.CloseProfitByDate = thost.TThostFtdcMoneyType(r.CloseProfitByDate)
	f.CloseProfitByTrade = thost.TThostFtdcMoneyType(r.CloseProfitByTrade)
	f.TodayPosition = thost.TThostFtdcVolumeType(r.TodayPosition)
	f.MarginRateByMoney = thost.TThostFtdcRatioType(r.MarginRateByMoney)
	f.MarginRateByVolume = thost.TThostFtdcRatioType(r.MarginRateByVolume)
	f.StrikeFrozen = thost.TThostFtdcVolumeType(r.StrikeFrozen)
	f.StrikeFrozenAmount = thost.TThostFtdcMoneyType(r.StrikeFrozenAmount)
	f.AbandonFrozen = thost.TThostFtdcVolumeType(r.AbandonFrozen)
	copy(f.ExchangeID[:len(f.ExchangeID)-1], []byte(r.ExchangeID))
	f.YdStrikeFrozen = thost.TThostFtdcVolumeType(r.YdStrikeFrozen)
	copy(f.InvestUnitID[:len(f.InvestUnitID)-1], []byte(r.InvestUnitID))
	f.PositionCostOffset = thost.TThostFtdcMoneyType(r.PositionCostOffset)
	f.TasPosition = thost.TThostFtdcVolumeType(r.TasPosition)
	f.TasPositionCost = thost.TThostFtdcMoneyType(r.TasPositionCost)
	copy(f.InstrumentID[:len(f.InstrumentID)-1], []byte(r.InstrumentID))

	return f
}

// 合约手续费率
func toCThostFtdcInstrumentCommissionRateField(r *InstrumentCommissionRateField) *thost.CThostFtdcInstrumentCommissionRateField {

	f := &thost.CThostFtdcInstrumentCommissionRateField{}

	f.InvestorRange = thost.TThostFtdcInvestorRangeType(r.InvestorRange)
	copy(f.BrokerID[:len(f.BrokerID)-1], []byte(r.BrokerID))
	copy(f.InvestorID[:len(f.InvestorID)-1], []byte(r.InvestorID))
	f.OpenRatioByMoney = thost.TThostFtdcRatioType(r.OpenRatioByMoney)
	f.OpenRatioByVolume = thost.TThostFtdcRatioType(r.OpenRatioByVolume)
	f.CloseRatioByMoney = thost.TThostFtdcRatioType(r.CloseRatioByMoney)
	f.CloseRatioByVolume = thost.TThostFtdcRatioType(r.CloseRatioByVolume)
	f.CloseTodayRatioByMoney = thost.TThostFtdcRatioType(r.CloseTodayRatioByMoney)
	f.CloseTodayRatioByVolume = thost.TThostFtdcRatioType(r.CloseTodayRatioByVolume)
	copy(f.ExchangeID[:len(f.ExchangeID)-1], []byte(r.ExchangeID))
	f.BizType = thost.TThostFtdcBizTypeType(r.BizType)
	copy(f.InvestUnitID[:len(f.InvestUnitID)-1], []byte(r.InvestUnitID))
	copy(f.InstrumentID[:len(f.InstrumentID)-1], []byte(r.InstrumentID))

	return f
}

// 转账请求
func toCThostFtdcReqTransferField(r *ReqTransferField) *thost.CThostFtdcReqTransferField {

	f := &thost.CThostFtdcReqTransferField{}

	copy(f.TradeCode[:len(f.TradeCode)-1], []byte(r.TradeCode))
	copy(f.BankID[:len(f.BankID)-1], []byte(r.BankID))
	copy(f.BankBranchID[:len(f.BankBranchID)-1], []byte(r.BankBranchID))
	copy(f.BrokerID[:len(f.BrokerID)-1], []byte(r.BrokerID))
	copy(f.BrokerBranchID[:len(f.BrokerBranchID)-1], []byte(r.BrokerBranchID))
	copy(f.TradeDate[:len(f.TradeDate)-1], []byte(r.TradeDate))
	copy(f.TradeTime[:len(f.TradeTime)-1], []byte(r.TradeTime))
	copy(f.BankSerial[:len(f.BankSerial)-1], []byte(r.BankSerial))
	copy(f.TradingDay[:len(f.TradingDay)-1], []byte(r.TradingDay))
	f.PlateSerial = thost.TThostFtdcSerialType(r.PlateSerial)
	f.LastFragment = thost.TThostFtdcLastFragmentType(r.LastFragment)
	f.SessionID = thost.TThostFtdcSessionIDType(r.SessionID)
	copy(f.CustomerName[:len(f.CustomerName)-1], []byte(r.CustomerName))
	f.IdCardType = thost.TThostFtdcIdCardTypeType(r.IdCardType)
	copy(f.IdentifiedCardNo[:len(f.IdentifiedCardNo)-1], []byte(r.IdentifiedCardNo))
	f.CustType = thost.TThostFtdcCustTypeType(r.CustType)
	copy(f.BankAccount[:len(f.BankAccount)-1], []byte(r.BankAccount))
	copy(f.BankPassWord[:len(f.BankPassWord)-1], []byte(r.BankPassWord))
	copy(f.AccountID[:len(f.AccountID)-1], []byte(r.AccountID))
	copy(f.Password[:len(f.Password)-1], []byte(r.Password))
	f.InstallID = thost.TThostFtdcInstallIDType(r.InstallID)
	f.FutureSerial = thost.TThostFtdcFutureSerialType(r.FutureSerial)
	copy(f.UserID[:len(f.UserID)-1], []byte(r.UserID))
	f.VerifyCertNoFlag = thost.TThostFtdcYesNoIndicatorType(r.VerifyCertNoFlag)
	copy(f.CurrencyID[:len(f.CurrencyID)-1], []byte(r.CurrencyID))
	f.TradeAmount = thost.TThostFtdcTradeAmountType(r.TradeAmount)
	f.FutureFetchAmount = thost.TThostFtdcTradeAmountType(r.FutureFetchAmount)
	f.FeePayFlag = thost.TThostFtdcFeePayFlagType(r.FeePayFlag)
	f.CustFee = thost.TThostFtdcCustFeeType(r.CustFee)
	f.BrokerFee = thost.TThostFtdcFutureFeeType(r.BrokerFee)
	copy(f.Message[:len(f.Message)-1], []byte(r.Message))
	copy(f.Digest[:len(f.Digest)-1], []byte(r.Digest))
	f.BankAccType = thost.TThostFtdcBankAccTypeType(r.BankAccType)
	copy(f.DeviceID[:len(f.DeviceID)-1], []byte(r.DeviceID))
	f.BankSecuAccType = thost.TThostFtdcBankAccTypeType(r.BankSecuAccType)
	copy(f.BrokerIDByBank[:len(f.BrokerIDByBank)-1], []byte(r.BrokerIDByBank))
	copy(f.BankSecuAcc[:len(f.BankSecuAcc)-1], []byte(r.BankSecuAcc))
	f.BankPwdFlag = thost.TThostFtdcPwdFlagType(r.BankPwdFlag)
	f.SecuPwdFlag = thost.TThostFtdcPwdFlagType(r.SecuPwdFlag)
	copy(f.OperNo[:len(f.OperNo)-1], []byte(r.OperNo))
	f.RequestID = thost.TThostFtdcRequestIDType(r.RequestID)
	f.TID = thost.TThostFtdcTIDType(r.TID)
	f.TransferStatus = thost.TThostFtdcTransferStatusType(r.TransferStatus)
	copy(f.LongCustomerName[:len(f.LongCustomerName)-1], []byte(r.LongCustomerName))

	return f
}

// 查询合约保证金率
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

// 合约保证金率
func toCThostFtdcInstrumentMarginRateField(r *InstrumentMarginRateField) *thost.CThostFtdcInstrumentMarginRateField {

	f := &thost.CThostFtdcInstrumentMarginRateField{}

	f.InvestorRange = thost.TThostFtdcInvestorRangeType(r.InvestorRange)
	copy(f.BrokerID[:len(f.BrokerID)-1], []byte(r.BrokerID))
	copy(f.InvestorID[:len(f.InvestorID)-1], []byte(r.InvestorID))
	f.HedgeFlag = thost.TThostFtdcHedgeFlagType(r.HedgeFlag)
	f.LongMarginRatioByMoney = thost.TThostFtdcRatioType(r.LongMarginRatioByMoney)
	f.LongMarginRatioByVolume = thost.TThostFtdcMoneyType(r.LongMarginRatioByVolume)
	f.ShortMarginRatioByMoney = thost.TThostFtdcRatioType(r.ShortMarginRatioByMoney)
	f.ShortMarginRatioByVolume = thost.TThostFtdcMoneyType(r.ShortMarginRatioByVolume)
	f.IsRelative = thost.TThostFtdcBoolType(r.IsRelative)
	copy(f.ExchangeID[:len(f.ExchangeID)-1], []byte(r.ExchangeID))
	copy(f.InvestUnitID[:len(f.InvestUnitID)-1], []byte(r.InvestUnitID))
	copy(f.InstrumentID[:len(f.InstrumentID)-1], []byte(r.InstrumentID))

	return f
}

// 用户事件通知信息
func toCThostFtdcTradingNoticeInfoField(r *TradingNoticeInfoField) *thost.CThostFtdcTradingNoticeInfoField {

	f := &thost.CThostFtdcTradingNoticeInfoField{}

	copy(f.BrokerID[:len(f.BrokerID)-1], []byte(r.BrokerID))
	copy(f.InvestorID[:len(f.InvestorID)-1], []byte(r.InvestorID))
	copy(f.SendTime[:len(f.SendTime)-1], []byte(r.SendTime))
	copy(f.FieldContent[:len(f.FieldContent)-1], []byte(r.FieldContent))
	f.SequenceSeries = thost.TThostFtdcSequenceSeriesType(r.SequenceSeries)
	f.SequenceNo = thost.TThostFtdcSequenceNoType(r.SequenceNo)
	copy(f.InvestUnitID[:len(f.InvestUnitID)-1], []byte(r.InvestUnitID))

	return f
}

// 查询资金账户
func toCThostFtdcQryTradingAccountField(r *QryTradingAccountField) *thost.CThostFtdcQryTradingAccountField {

	f := &thost.CThostFtdcQryTradingAccountField{}

	copy(f.BrokerID[:len(f.BrokerID)-1], []byte(r.BrokerID))
	copy(f.InvestorID[:len(f.InvestorID)-1], []byte(r.InvestorID))
	copy(f.CurrencyID[:len(f.CurrencyID)-1], []byte(r.CurrencyID))
	f.BizType = thost.TThostFtdcBizTypeType(r.BizType)
	copy(f.AccountID[:len(f.AccountID)-1], []byte(r.AccountID))

	return f
}

// 客户端认证请求
func toCThostFtdcReqAuthenticateField(r *ReqAuthenticateField) *thost.CThostFtdcReqAuthenticateField {

	f := &thost.CThostFtdcReqAuthenticateField{}

	copy(f.BrokerID[:len(f.BrokerID)-1], []byte(r.BrokerID))
	copy(f.UserID[:len(f.UserID)-1], []byte(r.UserID))
	copy(f.UserProductInfo[:len(f.UserProductInfo)-1], []byte(r.UserProductInfo))
	copy(f.AuthCode[:len(f.AuthCode)-1], []byte(r.AuthCode))
	copy(f.AppID[:len(f.AppID)-1], []byte(r.AppID))

	return f
}

// 成交
func toCThostFtdcTradeField(r *TradeField) *thost.CThostFtdcTradeField {

	f := &thost.CThostFtdcTradeField{}

	copy(f.BrokerID[:len(f.BrokerID)-1], []byte(r.BrokerID))
	copy(f.InvestorID[:len(f.InvestorID)-1], []byte(r.InvestorID))

	copy(f.OrderRef[:len(f.OrderRef)-1], []byte(r.OrderRef))
	copy(f.UserID[:len(f.UserID)-1], []byte(r.UserID))
	copy(f.ExchangeID[:len(f.ExchangeID)-1], []byte(r.ExchangeID))
	copy(f.TradeID[:len(f.TradeID)-1], []byte(r.TradeID))
	f.Direction = thost.TThostFtdcDirectionType(r.Direction)
	copy(f.OrderSysID[:len(f.OrderSysID)-1], []byte(r.OrderSysID))
	copy(f.ParticipantID[:len(f.ParticipantID)-1], []byte(r.ParticipantID))
	copy(f.ClientID[:len(f.ClientID)-1], []byte(r.ClientID))
	f.TradingRole = thost.TThostFtdcTradingRoleType(r.TradingRole)

	f.OffsetFlag = thost.TThostFtdcOffsetFlagType(r.OffsetFlag)
	f.HedgeFlag = thost.TThostFtdcHedgeFlagType(r.HedgeFlag)
	f.Price = thost.TThostFtdcPriceType(r.Price)
	f.Volume = thost.TThostFtdcVolumeType(r.Volume)
	copy(f.TradeDate[:len(f.TradeDate)-1], []byte(r.TradeDate))
	copy(f.TradeTime[:len(f.TradeTime)-1], []byte(r.TradeTime))
	f.TradeType = thost.TThostFtdcTradeTypeType(r.TradeType)
	f.PriceSource = thost.TThostFtdcPriceSourceType(r.PriceSource)
	copy(f.TraderID[:len(f.TraderID)-1], []byte(r.TraderID))
	copy(f.OrderLocalID[:len(f.OrderLocalID)-1], []byte(r.OrderLocalID))
	copy(f.ClearingPartID[:len(f.ClearingPartID)-1], []byte(r.ClearingPartID))
	copy(f.BusinessUnit[:len(f.BusinessUnit)-1], []byte(r.BusinessUnit))
	f.SequenceNo = thost.TThostFtdcSequenceNoType(r.SequenceNo)
	copy(f.TradingDay[:len(f.TradingDay)-1], []byte(r.TradingDay))
	f.SettlementID = thost.TThostFtdcSettlementIDType(r.SettlementID)
	f.BrokerOrderSeq = thost.TThostFtdcSequenceNoType(r.BrokerOrderSeq)
	f.TradeSource = thost.TThostFtdcTradeSourceType(r.TradeSource)
	copy(f.InvestUnitID[:len(f.InvestUnitID)-1], []byte(r.InvestUnitID))
	copy(f.InstrumentID[:len(f.InstrumentID)-1], []byte(r.InstrumentID))
	copy(f.ExchangeInstID[:len(f.ExchangeInstID)-1], []byte(r.ExchangeInstID))

	return f
}

// 投资者结算结果
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

// 用户口令变更
func toCThostFtdcUserPasswordUpdateField(r *UserPasswordUpdateField) *thost.CThostFtdcUserPasswordUpdateField {

	f := &thost.CThostFtdcUserPasswordUpdateField{}

	copy(f.BrokerID[:len(f.BrokerID)-1], []byte(r.BrokerID))
	copy(f.UserID[:len(f.UserID)-1], []byte(r.UserID))
	copy(f.OldPassword[:len(f.OldPassword)-1], []byte(r.OldPassword))
	copy(f.NewPassword[:len(f.NewPassword)-1], []byte(r.NewPassword))

	return f
}

// 合约状态
func toCThostFtdcInstrumentStatusField(r *InstrumentStatusField) *thost.CThostFtdcInstrumentStatusField {

	f := &thost.CThostFtdcInstrumentStatusField{}

	copy(f.ExchangeID[:len(f.ExchangeID)-1], []byte(r.ExchangeID))

	copy(f.SettlementGroupID[:len(f.SettlementGroupID)-1], []byte(r.SettlementGroupID))

	f.InstrumentStatus = thost.TThostFtdcInstrumentStatusType(r.InstrumentStatus)
	f.TradingSegmentSN = thost.TThostFtdcTradingSegmentSNType(r.TradingSegmentSN)
	copy(f.EnterTime[:len(f.EnterTime)-1], []byte(r.EnterTime))
	f.EnterReason = thost.TThostFtdcInstStatusEnterReasonType(r.EnterReason)
	copy(f.ExchangeInstID[:len(f.ExchangeInstID)-1], []byte(r.ExchangeInstID))
	copy(f.InstrumentID[:len(f.InstrumentID)-1], []byte(r.InstrumentID))

	return f
}

//---------------------------------------------------------------------------

// 客户端认证响应
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

// 查询手续费率
func fromCThostFtdcQryInstrumentCommissionRateField(r *thost.CThostFtdcQryInstrumentCommissionRateField) *QryInstrumentCommissionRateField {

	f := &QryInstrumentCommissionRateField{

		BrokerID:   bytes2String(r.BrokerID[:]),
		InvestorID: bytes2String(r.InvestorID[:]),

		ExchangeID:   bytes2String(r.ExchangeID[:]),
		InvestUnitID: bytes2String(r.InvestUnitID[:]),
		InstrumentID: bytes2String(r.InstrumentID[:]),
	}

	return f
}

// 查询合约保证金率
func fromCThostFtdcQryInstrumentMarginRateField(r *thost.CThostFtdcQryInstrumentMarginRateField) *QryInstrumentMarginRateField {

	f := &QryInstrumentMarginRateField{

		BrokerID:   bytes2String(r.BrokerID[:]),
		InvestorID: bytes2String(r.InvestorID[:]),

		HedgeFlag:    byte(r.HedgeFlag),
		ExchangeID:   bytes2String(r.ExchangeID[:]),
		InvestUnitID: bytes2String(r.InvestUnitID[:]),
		InstrumentID: bytes2String(r.InstrumentID[:]),
	}

	return f
}

// 银期转账交易流水表
func fromCThostFtdcTransferSerialField(r *thost.CThostFtdcTransferSerialField) *TransferSerialField {

	f := &TransferSerialField{

		PlateSerial:      int(r.PlateSerial),
		TradeDate:        bytes2String(r.TradeDate[:]),
		TradingDay:       bytes2String(r.TradingDay[:]),
		TradeTime:        bytes2String(r.TradeTime[:]),
		TradeCode:        bytes2String(r.TradeCode[:]),
		SessionID:        int(r.SessionID),
		BankID:           bytes2String(r.BankID[:]),
		BankBranchID:     bytes2String(r.BankBranchID[:]),
		BankAccType:      byte(r.BankAccType),
		BankAccount:      bytes2String(r.BankAccount[:]),
		BankSerial:       bytes2String(r.BankSerial[:]),
		BrokerID:         bytes2String(r.BrokerID[:]),
		BrokerBranchID:   bytes2String(r.BrokerBranchID[:]),
		FutureAccType:    byte(r.FutureAccType),
		AccountID:        bytes2String(r.AccountID[:]),
		InvestorID:       bytes2String(r.InvestorID[:]),
		FutureSerial:     int(r.FutureSerial),
		IdCardType:       byte(r.IdCardType),
		IdentifiedCardNo: bytes2String(r.IdentifiedCardNo[:]),
		CurrencyID:       bytes2String(r.CurrencyID[:]),
		TradeAmount:      float64(r.TradeAmount),
		CustFee:          float64(r.CustFee),
		BrokerFee:        float64(r.BrokerFee),
		AvailabilityFlag: byte(r.AvailabilityFlag),
		OperatorCode:     bytes2String(r.OperatorCode[:]),
		BankNewAccount:   bytes2String(r.BankNewAccount[:]),
		ErrorID:          int(r.ErrorID),
		ErrorMsg:         bytes2String(r.ErrorMsg[:]),
	}

	return f
}

// 查询经纪公司交易参数
func fromCThostFtdcQryBrokerTradingParamsField(r *thost.CThostFtdcQryBrokerTradingParamsField) *QryBrokerTradingParamsField {

	f := &QryBrokerTradingParamsField{

		BrokerID:   bytes2String(r.BrokerID[:]),
		InvestorID: bytes2String(r.InvestorID[:]),
		CurrencyID: bytes2String(r.CurrencyID[:]),
		AccountID:  bytes2String(r.AccountID[:]),
	}

	return f
}

// 输入报单
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

// 请求查询转帐流水
func fromCThostFtdcQryTransferSerialField(r *thost.CThostFtdcQryTransferSerialField) *QryTransferSerialField {

	f := &QryTransferSerialField{

		BrokerID:   bytes2String(r.BrokerID[:]),
		AccountID:  bytes2String(r.AccountID[:]),
		BankID:     bytes2String(r.BankID[:]),
		CurrencyID: bytes2String(r.CurrencyID[:]),
	}

	return f
}

// 用户事件通知信息
func fromCThostFtdcTradingNoticeInfoField(r *thost.CThostFtdcTradingNoticeInfoField) *TradingNoticeInfoField {

	f := &TradingNoticeInfoField{

		BrokerID:       bytes2String(r.BrokerID[:]),
		InvestorID:     bytes2String(r.InvestorID[:]),
		SendTime:       bytes2String(r.SendTime[:]),
		FieldContent:   bytes2String(r.FieldContent[:]),
		SequenceSeries: int16(r.SequenceSeries),
		SequenceNo:     int(r.SequenceNo),
		InvestUnitID:   bytes2String(r.InvestUnitID[:]),
	}

	return f
}

// 输入报单操作
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

// 查询资金账户
func fromCThostFtdcQryTradingAccountField(r *thost.CThostFtdcQryTradingAccountField) *QryTradingAccountField {

	f := &QryTradingAccountField{

		BrokerID:   bytes2String(r.BrokerID[:]),
		InvestorID: bytes2String(r.InvestorID[:]),
		CurrencyID: bytes2String(r.CurrencyID[:]),
		BizType:    byte(r.BizType),
		AccountID:  bytes2String(r.AccountID[:]),
	}

	return f
}

// 用户口令变更
func fromCThostFtdcUserPasswordUpdateField(r *thost.CThostFtdcUserPasswordUpdateField) *UserPasswordUpdateField {

	f := &UserPasswordUpdateField{

		BrokerID:    bytes2String(r.BrokerID[:]),
		UserID:      bytes2String(r.UserID[:]),
		OldPassword: bytes2String(r.OldPassword[:]),
		NewPassword: bytes2String(r.NewPassword[:]),
	}

	return f
}

// 报单
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

// 响应信息
func fromCThostFtdcRspInfoField(r *thost.CThostFtdcRspInfoField) *RspInfoField {

	f := &RspInfoField{

		ErrorID:  int32(r.ErrorID),
		ErrorMsg: bytes2String(r.ErrorMsg[:]),
	}

	return f
}

// 合约保证金率
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

// 成交
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

// 查询投资者持仓
func fromCThostFtdcQryInvestorPositionField(r *thost.CThostFtdcQryInvestorPositionField) *QryInvestorPositionField {

	f := &QryInvestorPositionField{

		BrokerID:   bytes2String(r.BrokerID[:]),
		InvestorID: bytes2String(r.InvestorID[:]),

		ExchangeID:   bytes2String(r.ExchangeID[:]),
		InvestUnitID: bytes2String(r.InvestUnitID[:]),
		InstrumentID: bytes2String(r.InstrumentID[:]),
	}

	return f
}

// 客户端认证请求
func fromCThostFtdcReqAuthenticateField(r *thost.CThostFtdcReqAuthenticateField) *ReqAuthenticateField {

	f := &ReqAuthenticateField{

		BrokerID:        bytes2String(r.BrokerID[:]),
		UserID:          bytes2String(r.UserID[:]),
		UserProductInfo: bytes2String(r.UserProductInfo[:]),
		AuthCode:        bytes2String(r.AuthCode[:]),
		AppID:           bytes2String(r.AppID[:]),
	}

	return f
}

// 经纪公司交易参数
func fromCThostFtdcBrokerTradingParamsField(r *thost.CThostFtdcBrokerTradingParamsField) *BrokerTradingParamsField {

	f := &BrokerTradingParamsField{

		BrokerID:                bytes2String(r.BrokerID[:]),
		InvestorID:              bytes2String(r.InvestorID[:]),
		MarginPriceType:         byte(r.MarginPriceType),
		Algorithm:               byte(r.Algorithm),
		AvailIncludeCloseProfit: byte(r.AvailIncludeCloseProfit),
		CurrencyID:              bytes2String(r.CurrencyID[:]),
		OptionRoyaltyPriceType:  byte(r.OptionRoyaltyPriceType),
		AccountID:               bytes2String(r.AccountID[:]),
	}

	return f
}

// 请求查询银期签约关系
func fromCThostFtdcQryAccountregisterField(r *thost.CThostFtdcQryAccountregisterField) *QryAccountregisterField {

	f := &QryAccountregisterField{

		BrokerID:     bytes2String(r.BrokerID[:]),
		AccountID:    bytes2String(r.AccountID[:]),
		BankID:       bytes2String(r.BankID[:]),
		BankBranchID: bytes2String(r.BankBranchID[:]),
		CurrencyID:   bytes2String(r.CurrencyID[:]),
	}

	return f
}

// 查询签约银行响应
func fromCThostFtdcContractBankField(r *thost.CThostFtdcContractBankField) *ContractBankField {

	f := &ContractBankField{

		BrokerID:   bytes2String(r.BrokerID[:]),
		BankID:     bytes2String(r.BankID[:]),
		BankBrchID: bytes2String(r.BankBrchID[:]),
		BankName:   bytes2String(r.BankName[:]),
	}

	return f
}

// 投资者结算结果确认信息
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

// 客户开销户信息表
func fromCThostFtdcAccountregisterField(r *thost.CThostFtdcAccountregisterField) *AccountregisterField {

	f := &AccountregisterField{

		TradeDay:         bytes2String(r.TradeDay[:]),
		BankID:           bytes2String(r.BankID[:]),
		BankBranchID:     bytes2String(r.BankBranchID[:]),
		BankAccount:      bytes2String(r.BankAccount[:]),
		BrokerID:         bytes2String(r.BrokerID[:]),
		BrokerBranchID:   bytes2String(r.BrokerBranchID[:]),
		AccountID:        bytes2String(r.AccountID[:]),
		IdCardType:       byte(r.IdCardType),
		IdentifiedCardNo: bytes2String(r.IdentifiedCardNo[:]),
		CustomerName:     bytes2String(r.CustomerName[:]),
		CurrencyID:       bytes2String(r.CurrencyID[:]),
		OpenOrDestroy:    byte(r.OpenOrDestroy),
		RegDate:          bytes2String(r.RegDate[:]),
		OutDate:          bytes2String(r.OutDate[:]),
		TID:              int(r.TID),
		CustType:         byte(r.CustType),
		BankAccType:      byte(r.BankAccType),
		LongCustomerName: bytes2String(r.LongCustomerName[:]),
	}

	return f
}

// 用户登录应答
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

// 投资者结算结果
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

// 深度行情
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

// 查询签约银行请求
func fromCThostFtdcQryContractBankField(r *thost.CThostFtdcQryContractBankField) *QryContractBankField {

	f := &QryContractBankField{

		BrokerID:   bytes2String(r.BrokerID[:]),
		BankID:     bytes2String(r.BankID[:]),
		BankBrchID: bytes2String(r.BankBrchID[:]),
	}

	return f
}

// 资金账户
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

// 合约状态
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

// 银行发起银行资金转期货响应
func fromCThostFtdcRspTransferField(r *thost.CThostFtdcRspTransferField) *RspTransferField {

	f := &RspTransferField{

		TradeCode:         bytes2String(r.TradeCode[:]),
		BankID:            bytes2String(r.BankID[:]),
		BankBranchID:      bytes2String(r.BankBranchID[:]),
		BrokerID:          bytes2String(r.BrokerID[:]),
		BrokerBranchID:    bytes2String(r.BrokerBranchID[:]),
		TradeDate:         bytes2String(r.TradeDate[:]),
		TradeTime:         bytes2String(r.TradeTime[:]),
		BankSerial:        bytes2String(r.BankSerial[:]),
		TradingDay:        bytes2String(r.TradingDay[:]),
		PlateSerial:       int(r.PlateSerial),
		LastFragment:      byte(r.LastFragment),
		SessionID:         int(r.SessionID),
		CustomerName:      bytes2String(r.CustomerName[:]),
		IdCardType:        byte(r.IdCardType),
		IdentifiedCardNo:  bytes2String(r.IdentifiedCardNo[:]),
		CustType:          byte(r.CustType),
		BankAccount:       bytes2String(r.BankAccount[:]),
		BankPassWord:      bytes2String(r.BankPassWord[:]),
		AccountID:         bytes2String(r.AccountID[:]),
		Password:          bytes2String(r.Password[:]),
		InstallID:         int(r.InstallID),
		FutureSerial:      int(r.FutureSerial),
		UserID:            bytes2String(r.UserID[:]),
		VerifyCertNoFlag:  byte(r.VerifyCertNoFlag),
		CurrencyID:        bytes2String(r.CurrencyID[:]),
		TradeAmount:       float64(r.TradeAmount),
		FutureFetchAmount: float64(r.FutureFetchAmount),
		FeePayFlag:        byte(r.FeePayFlag),
		CustFee:           float64(r.CustFee),
		BrokerFee:         float64(r.BrokerFee),
		Message:           bytes2String(r.Message[:]),
		Digest:            bytes2String(r.Digest[:]),
		BankAccType:       byte(r.BankAccType),
		DeviceID:          bytes2String(r.DeviceID[:]),
		BankSecuAccType:   byte(r.BankSecuAccType),
		BrokerIDByBank:    bytes2String(r.BrokerIDByBank[:]),
		BankSecuAcc:       bytes2String(r.BankSecuAcc[:]),
		BankPwdFlag:       byte(r.BankPwdFlag),
		SecuPwdFlag:       byte(r.SecuPwdFlag),
		OperNo:            bytes2String(r.OperNo[:]),
		RequestID:         int(r.RequestID),
		TID:               int(r.TID),
		TransferStatus:    byte(r.TransferStatus),
		ErrorID:           int(r.ErrorID),
		ErrorMsg:          bytes2String(r.ErrorMsg[:]),
		LongCustomerName:  bytes2String(r.LongCustomerName[:]),
	}

	return f
}

// 报单操作
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

// 查询成交
func fromCThostFtdcQryTradeField(r *thost.CThostFtdcQryTradeField) *QryTradeField {

	f := &QryTradeField{

		BrokerID:   bytes2String(r.BrokerID[:]),
		InvestorID: bytes2String(r.InvestorID[:]),

		ExchangeID:     bytes2String(r.ExchangeID[:]),
		TradeID:        bytes2String(r.TradeID[:]),
		TradeTimeStart: bytes2String(r.TradeTimeStart[:]),
		TradeTimeEnd:   bytes2String(r.TradeTimeEnd[:]),
		InvestUnitID:   bytes2String(r.InvestUnitID[:]),
		InstrumentID:   bytes2String(r.InstrumentID[:]),
	}

	return f
}

// 用户登录请求
func fromCThostFtdcReqUserLoginField(r *thost.CThostFtdcReqUserLoginField) *ReqUserLoginField {

	f := &ReqUserLoginField{

		TradingDay:           bytes2String(r.TradingDay[:]),
		BrokerID:             bytes2String(r.BrokerID[:]),
		UserID:               bytes2String(r.UserID[:]),
		Password:             bytes2String(r.Password[:]),
		UserProductInfo:      bytes2String(r.UserProductInfo[:]),
		InterfaceProductInfo: bytes2String(r.InterfaceProductInfo[:]),
		ProtocolInfo:         bytes2String(r.ProtocolInfo[:]),
		MacAddress:           bytes2String(r.MacAddress[:]),
		OneTimePassword:      bytes2String(r.OneTimePassword[:]),

		LoginRemark:     bytes2String(r.LoginRemark[:]),
		ClientIPPort:    int(r.ClientIPPort),
		ClientIPAddress: bytes2String(r.ClientIPAddress[:]),
	}

	return f
}

// 查询报单
func fromCThostFtdcQryOrderField(r *thost.CThostFtdcQryOrderField) *QryOrderField {

	f := &QryOrderField{

		BrokerID:   bytes2String(r.BrokerID[:]),
		InvestorID: bytes2String(r.InvestorID[:]),

		ExchangeID:      bytes2String(r.ExchangeID[:]),
		OrderSysID:      bytes2String(r.OrderSysID[:]),
		InsertTimeStart: bytes2String(r.InsertTimeStart[:]),
		InsertTimeEnd:   bytes2String(r.InsertTimeEnd[:]),
		InvestUnitID:    bytes2String(r.InvestUnitID[:]),
		InstrumentID:    bytes2String(r.InstrumentID[:]),
	}

	return f
}

// 查询投资者结算结果
func fromCThostFtdcQrySettlementInfoField(r *thost.CThostFtdcQrySettlementInfoField) *QrySettlementInfoField {

	f := &QrySettlementInfoField{

		BrokerID:   bytes2String(r.BrokerID[:]),
		InvestorID: bytes2String(r.InvestorID[:]),
		TradingDay: bytes2String(r.TradingDay[:]),
		AccountID:  bytes2String(r.AccountID[:]),
		CurrencyID: bytes2String(r.CurrencyID[:]),
	}

	return f
}

// 资金账户口令变更域
func fromCThostFtdcTradingAccountPasswordUpdateField(r *thost.CThostFtdcTradingAccountPasswordUpdateField) *TradingAccountPasswordUpdateField {

	f := &TradingAccountPasswordUpdateField{

		BrokerID:    bytes2String(r.BrokerID[:]),
		AccountID:   bytes2String(r.AccountID[:]),
		OldPassword: bytes2String(r.OldPassword[:]),
		NewPassword: bytes2String(r.NewPassword[:]),
		CurrencyID:  bytes2String(r.CurrencyID[:]),
	}

	return f
}

// 合约
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

// 查询合约
func fromCThostFtdcQryInstrumentField(r *thost.CThostFtdcQryInstrumentField) *QryInstrumentField {

	f := &QryInstrumentField{

		ExchangeID: bytes2String(r.ExchangeID[:]),

		InstrumentID:   bytes2String(r.InstrumentID[:]),
		ExchangeInstID: bytes2String(r.ExchangeInstID[:]),
		ProductID:      bytes2String(r.ProductID[:]),
	}

	return f
}

// 投资者持仓
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

// 查询结算信息确认域
func fromCThostFtdcQrySettlementInfoConfirmField(r *thost.CThostFtdcQrySettlementInfoConfirmField) *QrySettlementInfoConfirmField {

	f := &QrySettlementInfoConfirmField{

		BrokerID:   bytes2String(r.BrokerID[:]),
		InvestorID: bytes2String(r.InvestorID[:]),
		AccountID:  bytes2String(r.AccountID[:]),
		CurrencyID: bytes2String(r.CurrencyID[:]),
	}

	return f
}

// 合约手续费率
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

// 转账请求
func fromCThostFtdcReqTransferField(r *thost.CThostFtdcReqTransferField) *ReqTransferField {

	f := &ReqTransferField{

		TradeCode:         bytes2String(r.TradeCode[:]),
		BankID:            bytes2String(r.BankID[:]),
		BankBranchID:      bytes2String(r.BankBranchID[:]),
		BrokerID:          bytes2String(r.BrokerID[:]),
		BrokerBranchID:    bytes2String(r.BrokerBranchID[:]),
		TradeDate:         bytes2String(r.TradeDate[:]),
		TradeTime:         bytes2String(r.TradeTime[:]),
		BankSerial:        bytes2String(r.BankSerial[:]),
		TradingDay:        bytes2String(r.TradingDay[:]),
		PlateSerial:       int(r.PlateSerial),
		LastFragment:      byte(r.LastFragment),
		SessionID:         int(r.SessionID),
		CustomerName:      bytes2String(r.CustomerName[:]),
		IdCardType:        byte(r.IdCardType),
		IdentifiedCardNo:  bytes2String(r.IdentifiedCardNo[:]),
		CustType:          byte(r.CustType),
		BankAccount:       bytes2String(r.BankAccount[:]),
		BankPassWord:      bytes2String(r.BankPassWord[:]),
		AccountID:         bytes2String(r.AccountID[:]),
		Password:          bytes2String(r.Password[:]),
		InstallID:         int(r.InstallID),
		FutureSerial:      int(r.FutureSerial),
		UserID:            bytes2String(r.UserID[:]),
		VerifyCertNoFlag:  byte(r.VerifyCertNoFlag),
		CurrencyID:        bytes2String(r.CurrencyID[:]),
		TradeAmount:       float64(r.TradeAmount),
		FutureFetchAmount: float64(r.FutureFetchAmount),
		FeePayFlag:        byte(r.FeePayFlag),
		CustFee:           float64(r.CustFee),
		BrokerFee:         float64(r.BrokerFee),
		Message:           bytes2String(r.Message[:]),
		Digest:            bytes2String(r.Digest[:]),
		BankAccType:       byte(r.BankAccType),
		DeviceID:          bytes2String(r.DeviceID[:]),
		BankSecuAccType:   byte(r.BankSecuAccType),
		BrokerIDByBank:    bytes2String(r.BrokerIDByBank[:]),
		BankSecuAcc:       bytes2String(r.BankSecuAcc[:]),
		BankPwdFlag:       byte(r.BankPwdFlag),
		SecuPwdFlag:       byte(r.SecuPwdFlag),
		OperNo:            bytes2String(r.OperNo[:]),
		RequestID:         int(r.RequestID),
		TID:               int(r.TID),
		TransferStatus:    byte(r.TransferStatus),
		LongCustomerName:  bytes2String(r.LongCustomerName[:]),
	}

	return f
}

//-------------------------------------------------------------------

func bytes2String(t []byte) string {
	var s []byte = t

	i := bytes.Index(t, []byte{0})
	if i >= 0 {
		s = t[:i]
	}
	msg, _ := simplifiedchinese.GB18030.NewDecoder().Bytes(s)
	return strings.Trim(string(msg), "\u0000")
}
