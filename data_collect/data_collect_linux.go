//go:build !darwin && linux

package dc

/*
#cgo linux LDFLAGS: -L. -L${SRCDIR}/../api/v6.7.0_20230209_api_traderapi_se_linux64 -lLinuxDataCollect
#cgo linux CPPFLAGS: -I. -I${SRCDIR}/../api/v6.7.0_20230209_api_traderapi_se_linux64

#include "data_collect_wrap.h"
*/
import "C"
import "unsafe"

// CTP_GetSystemInfo 中继模式使用方法
func CTP_GetSystemInfo(data *DataCollectSystemInfo) int {
	var arg0 = data
	r := C.__wrap_ctp_get_system_info((*C.struct_DataCollectSystemInfo)(unsafe.Pointer(arg0)))
	return int(r)
}

// CTP_GetSystemInfoUnAesEncode 直连模式使用方法
func CTP_GetSystemInfoUnAesEncode(data *DataCollectSystemInfo) int {
	var arg0 = data
	r := C.__wrap_ctp_get_system_info_unaes_encode((*C.struct_DataCollectSystemInfo)(unsafe.Pointer(arg0)))
	return int(r)
}

// 查询采集库版本信息
func CTP_GetDataCollectApiVersion() string {
	r := C.__wrap_ctp_get_data_collect_api_version()
	result := C.GoString(r)
	return result
}
