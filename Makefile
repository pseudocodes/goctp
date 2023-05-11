
build:
	go build

v6.6.9:
	@sed -i '16c #cgo linux LDFLAGS: -fPIC -L. -L$${SRCDIR}/api/v6.6.9_20220914_api_tradeapi_se_linux64 -Wl,-rpath=$${SRCDIR}/api/v6.6.9_20220914_api_tradeapi_se_linux64  -lthostmduserapi_se -lthosttraderapi_se -lstdc++' libctp.go
	@sed -i '17c #cgo linux CPPFLAGS: -fPIC -I. -I$${SRCDIR}/api/v6.6.9_20220914_api_tradeapi_se_linux64' libctp.go

v6.6.7:
	@sed -i '16c #cgo linux LDFLAGS: -fPIC -L. -L$${SRCDIR}/api/v6.6.7_20220613_api_tradeapi_linux64 -Wl,-rpath=$${SRCDIR}/api/v6.6.7_20220613_api_tradeapi_linux64  -lthostmduserapi_se -lthosttraderapi_se -lstdc++' libctp.go
	@sed -i '17c #cgo linux CPPFLAGS: -fPIC -I. -I$${SRCDIR}/api/v6.6.7_20220613_api_tradeapi_linux64' libctp.go

v6.6.1_P1:
	@sed -i '16c #cgo linux LDFLAGS: -fPIC -L. -L$${SRCDIR}/api/v6.6.1_P1_20210406_api_tradeapi_se_linux64 -Wl,-rpath=$${SRCDIR}/api/v6.6.1_P1_20210406_api_tradeapi_se_linux64  -lthostmduserapi_se -lthosttraderapi_se -lstdc++' libctp.go
	@sed -i '17c #cgo linux CPPFLAGS: -fPIC -I. -I/$${SRCDIR}/api/v6.6.1_P1_20210406_api_tradeapi_se_linux64/' libctp.go
	

example:
	go build -linkshared sample/market/goctp_md_example.go
	go build -linkshared sample/trader/goctp_trader_example.go
