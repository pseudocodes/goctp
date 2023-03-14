# GoCTP
Pure Golang bindings for CTP v2 Version without Swig

原生 Golang CTP 封装，最大限度的利用了 Go 本身的 FFI 特性



## Why not Swig
* 编译慢
* 编辑器提示不友好
* 封装接口太难看，不好用


## How to Build goctp
> Only support Linux X64 now.

```
go get github.com/pseudocodes/goctp

# make build
```

## How to Build example 

```
make example

> go build sample/simple_market/simple_market.go
> go build sample/simple_trader/simple_trader.go

```

## Sample 

### simple market quote 
```go

import "github.com/pseudocodes/goctp"


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
	log.Printf("tick {%+v}\n", quote)
}

func (s *baseSpi) OnRspError(rspInfo *goctp.RspInfoField, requestID int, isLast bool) {
	log.Printf("RspInfo: %+v\n", rspInfo)

}

```

### simple trader
```go
import "github.com/pseudocodes/goctp"

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

	s.tdspi.SetOnRspAuthenticate(s.OnRspAuthenticate)

	s.tdspi.SetOnRspUserLogin(s.OnRspUserLogin)
	
	s.tdspi.SetOnRspSettlementInfoConfirm(s.OnRspSettlementInfoConfirm)

	s.tdspi.SetOnRspQryInstrument(s.OnRspQryInstrument)

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

func main() {

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

```

## 免责条款声明

本项目明确拒绝对产品做任何明示或暗示的担保。由于软件系统开发本身的复杂性，无法保证本项目完全没有错误。如选择使用本项目表示您同意错误和/或遗漏的存在，在任何情况下本项目对于直接、间接、特殊的、偶然的、或间接产生的、使用或无法使用本项目进行交易和投资造成的盈亏、直接或间接引起的赔偿、损失、债务或是任何交易中止均不承担责任和义务。 

此声明永久有效