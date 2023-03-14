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

package main

import (
	"log"
	"os"

	"github.com/gookit/goutil/dump"
	"github.com/pseudocodes/goctp"
)

var SimnowEnv map[string]map[string]string = map[string]map[string]string{
	"td": {
		"7x24":      "tcp://180.168.146.187:10130",
		"telesim1":  "tcp://180.168.146.187:10201",
		"telesim2":  "tcp://180.168.146.187:10202",
		"moblesim3": "tcp://218.202.237.33:10203",
	},
	"md": {
		"7x24":      "tcp://180.168.146.187:10131",
		"telesim1":  "tcp://180.168.146.187:10211",
		"telesim2":  "tcp://180.168.146.187:10212",
		"moblesim3": "tcp://218.202.237.33:10213",
	},
}

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

type baseSpi struct {
	tdspi *goctp.TraderSpiLite
	tdapi *goctp.TraderApiLite
}

func CreateBaseSpi() *baseSpi {
	s := &baseSpi{
		// tdapi: tdapi,
		tdspi: &goctp.TraderSpiLite{},
	}

	s.tdspi.SetOnFrontConnected(s.OnFrontConnected)

	s.tdspi.SetOnFrontDisconnected(s.OnFrontDisconnected)

	// s.tdspi.SetOnHeartBeatWarning(s.OnHeartBeatWarning)

	s.tdspi.SetOnRspAuthenticate(s.OnRspAuthenticate)

	s.tdspi.SetOnRspUserLogin(s.OnRspUserLogin)

	// s.tdspi.SetOnRspOrderInsert(s.OnRspOrderInsert)

	// s.tdspi.SetOnRspOrderAction(s.OnRspOrderAction)

	s.tdspi.SetOnRspSettlementInfoConfirm(s.OnRspSettlementInfoConfirm)

	// s.tdspi.SetOnRspQryOrder(s.OnRspQryOrder)

	// s.tdspi.SetOnRspQryInvestorPosition(s.OnRspQryInvestorPosition)

	// s.tdspi.SetOnRspQryTradingAccount(s.OnRspQryTradingAccount)

	// s.tdspi.SetOnRspQryInstrumentMarginRate(s.OnRspQryInstrumentMarginRate)

	// s.tdspi.SetOnRspQryInstrumentCommissionRate(s.OnRspQryInstrumentCommissionRate)

	s.tdspi.SetOnRspQryInstrument(s.OnRspQryInstrument)

	// s.tdspi.SetOnRspQrySettlementInfo(s.OnRspQrySettlementInfo)

	// s.tdspi.SetOnRspError(s.OnRspError)

	// s.tdspi.SetOnRtnOrder(s.OnRtnOrder)

	// s.tdspi.SetOnRtnTrade(s.OnRtnTrade)

	// s.tdspi.SetOnErrRtnOrderInsert(s.OnErrRtnOrderInsert)

	// s.tdspi.SetOnErrRtnOrderAction(s.OnErrRtnOrderAction)

	s.tdspi.SetOnRtnInstrumentStatus(s.OnRtnInstrumentStatus)
	return s
}

func (s *baseSpi) OnFrontConnected() {
	var ret int
	log.Printf("OnFrontConnected\n")

	ret = s.tdapi.ReqAuthenticate(&goctp.ReqAuthenticateField{
		BrokerID: "9999",
		UserID:   os.Getenv("SIMNOW_USER_ID"), // <- 环境变量设置
		AuthCode: "0000000000000000",
		AppID:    "simnow_client_test",
	}, 1)

	// req := &goctp.ReqUserLoginField{
	// 	BrokerID: "9999",
	// 	UserID:   os.Getenv("SIMNOW_USER_ID"), // <- 环境变量设置
	// 	Password: os.Getenv("SIMNOW_USER_PASSWORD"),
	// }
	// dump.Println(req)
	// ret = s.tdapi.ReqUserLogin(req, 1)

	log.Printf("user auth: %v\n", ret)
}

func (s *baseSpi) OnRspAuthenticate(f *goctp.RspAuthenticateField, r *goctp.RspInfoField, nRequestID int, bIsLast bool) {
	dump.Println(r)
	dump.Println(f)
	req := &goctp.ReqUserLoginField{
		BrokerID: "9999",
		UserID:   os.Getenv("SIMNOW_USER_ID"), // <- 环境变量设置
		Password: os.Getenv("SIMNOW_USER_PASSWORD"),
	}
	dump.Println(req)
	ret := s.tdapi.ReqUserLogin(req, 1)
	log.Printf("user login: %v\n", ret)
}

func (s *baseSpi) OnRspSettlementInfoConfirm(f *goctp.SettlementInfoConfirmField, r *goctp.RspInfoField, nRequestID int, bIsLast bool) {
	dump.Println(r)
	dump.Println(f)
}

func (s *baseSpi) OnFrontDisconnected(nReason int) {
	log.Printf("OnFrontDissconnected: %v\n", nReason)
}

func (s *baseSpi) OnRspUserLogin(f *goctp.RspUserLoginField, r *goctp.RspInfoField, nRequestID int, bIsLast bool) {
	dump.Println(r)
	dump.Println(f)
	// req := &goctp.QryInstrumentField{}
	// ret := s.tdapi.ReqQryInstrument(req, 3)
	// log.Printf("user qry ins: %v\n", ret)
}

func (s *baseSpi) OnRspQryInstrument(pInstrument *goctp.InstrumentField, pRspInfo *goctp.RspInfoField, nRequestID int, bIsLast bool) {
	dump.Print(pRspInfo, nRequestID, bIsLast)
	dump.Println(pInstrument.InstrumentName)

}

//合约交易状态通知
func (s *baseSpi) OnRtnInstrumentStatus(pInstrumentStatus *goctp.InstrumentStatusField) {
	dump.P(pInstrumentStatus)
}

func sample1() {

	tdapi := goctp.CreateTraderApiLite(goctp.TraderFlowPath("./data/"))
	baseSpi := CreateBaseSpi()
	baseSpi.tdapi = tdapi
	log.Printf("baseSpi %+v\n", baseSpi)
	tdapi.RegisterSpi(baseSpi.tdspi)
	tdapi.RegisterFront(SimnowEnv["td"]["7x24"])

	tdapi.Init()

	println(tdapi.GetTradingDay())
	println(tdapi.GetApiVersion())
	tdapi.Join()
}

func main() {
	sample1()
}
