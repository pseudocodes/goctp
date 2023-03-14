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

/*
Simnow是上期技术提供的CTP程序测试、模拟、学习的模拟平台。

7x24环境-电信：
交易前置: tcp://180.168.146.187:10130
行情前置: tcp://180.168.146.187:10131

仿真环境1-电信：交易时段同实盘
交易前置: tcp://180.168.146.187:10201
行情前置: tcp://180.168.146.187:10211

仿真环境2-电信：交易时段同实盘
交易前置: tcp://180.168.146.187:10202
行情前置: tcp://180.168.146.187:10212

仿真环境3-移动：交易时段同实盘
交易前置: tcp://218.202.237.33:10203
行情前置: tcp://218.202.237.33:10213
*/
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
	mdapi *goctp.MdApiLite
	mdspi *goctp.MdSpiLite
}

func CreateBaseSpi() *baseSpi {
	s := &baseSpi{
		mdspi: &goctp.MdSpiLite{},
	}
	s.mdspi.SetOnFrontConnected(s.OnFrontConnected)

	s.mdspi.SetOnFrontDisconnected(s.OnFrontDisconnected)

	s.mdspi.SetOnHeartBeatWarning(s.OnHeartBeatWarning)

	s.mdspi.SetOnRspUserLogin(s.OnRspUserLogin)

	s.mdspi.SetOnRspError(s.OnRspError)

	s.mdspi.SetOnRtnDepthMarketData(s.OnRtnDepthMarketData)
	return s
}

func (s *baseSpi) OnFrontConnected() {
	log.Printf("OnFrontConnected\n")

	ret := s.mdapi.ReqUserLogin(&goctp.ReqUserLoginField{
		BrokerID: "9999",
		UserID:   os.Getenv("SIMNOW_USER_ID"), // <- 环境变量设置
		Password: os.Getenv("SIMNOW_USER_PASSWORD"),
	}, 1)

	log.Printf("user log: %v\n", ret)
}

func (s *baseSpi) OnHeartBeatWarning(timelapse int) {
	log.Printf("OnHeartBeatWarning: %v\n", timelapse)
}

func (s *baseSpi) OnFrontDisconnected(nReason int) {
	log.Printf("OnFrontDisconnected: %v\n", nReason)
}

func (s *baseSpi) OnRspUserLogin(pRspUserLogin *goctp.RspUserLoginField, rspInfo *goctp.RspInfoField, nRequestID int, bIsLast bool) {
	log.Printf("RspUserLogin: %+v\nRspInfo: %+v\n", pRspUserLogin, rspInfo)
	s.mdapi.SubscribeMarketData("ag2306")
}

func (s *baseSpi) OnRspSubMarketData(instrumentID string, rspInfo *goctp.RspInfoField, requestID int, isLast bool) {
	log.Printf("instrumentID: %+v\n  RspInfo: %+v\n", instrumentID, rspInfo)
}

func (s *baseSpi) OnRtnDepthMarketData(quote *goctp.DepthMarketDataField) {
	// log.Printf("tick {%+v}\n", quote)
	dump.Println(quote)
}

func (s *baseSpi) OnRspError(rspInfo *goctp.RspInfoField, requestID int, isLast bool) {
	log.Printf("RspInfo: %+v\n", rspInfo)

}

func sample1() {

	mdapi := goctp.CreateMdApiLite(goctp.MdFlowPath("./data/"), goctp.MdUsingUDP(false), goctp.MdMultiCast(false))
	baseSpi := CreateBaseSpi()
	baseSpi.mdapi = mdapi
	mdapi.RegisterSpi(baseSpi.mdspi)
	mdapi.RegisterFront(SimnowEnv["md"]["telesim1"])
	// mdapi.RegisterFront(SimnowEnv["md"]["md"])
	// mdapi.RegisterFront("tcp://0.0.0.0:9091")

	// mdapi.RegisterNameServer("tcp://localhost:9091")
	mdapi.Init()

	println(mdapi.GetTradingDay())
	println(mdapi.GetTradingDay())

	mdapi.Join()
}

func main() {
	sample1()
}
