#ifndef _DATA_COLLECT_WRAP_H_
#define _DATA_COLLECT_WRAP_H_

struct DataCollectSystemInfo {
    int Length;
    char SystemInfo[273];
};

#ifdef __cplusplus
extern "C" {
#endif
int __wrap_ctp_get_system_info(struct DataCollectSystemInfo* pSystemInfo);

int __wrap_ctp_get_system_info_unaes_encode(struct DataCollectSystemInfo* pSystemInfo);

const char* __wrap_ctp_get_data_collect_api_version(void);

#ifdef __cplusplus
}
#endif

#endif