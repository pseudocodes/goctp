//go:build darwin && !linux

package dc

/*
#cgo darwin LDFLAGS: -L. -L${SRCDIR}/../api/v6.6.9_MacOS_20220926 -lMacDataCollect -framework Cocoa -framework IOKit
#cgo darwin CPPFLAGS: -I. -I${SRCDIR}/../api/v6.6.9_MacOS_20220926

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
