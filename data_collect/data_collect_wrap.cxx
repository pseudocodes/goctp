#include <stdio.h>
#include <stdlib.h>
#include <string.h>

#include "data_collect_wrap.h"

#include "DataCollect.h"

int __wrap_ctp_get_system_info(DataCollectSystemInfo* pSystemInfo)
{
    char buffer[273];
    int nLen = 0;
    memset(buffer, 0, sizeof(buffer));
    int ret = CTP_GetSystemInfo(buffer, nLen);
    pSystemInfo->Length = nLen;
    memcpy(pSystemInfo->SystemInfo, buffer, nLen);
    return ret;
}

int __wrap_ctp_get_system_info_unaes_encode(DataCollectSystemInfo* pSystemInfo)
{
    char buffer[273];
    int nLen = 0;
    memset(buffer, 0, sizeof(buffer));
    int ret = CTP_GetSystemInfoUnAesEncode(buffer, nLen);
    pSystemInfo->Length = nLen;
    memcpy(pSystemInfo->SystemInfo, buffer, nLen);
    return ret;
}

const char* __wrap_ctp_get_data_collect_api_version(void)
{
    return CTP_GetDataCollectApiVersion();
}