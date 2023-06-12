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

/*
#cgo linux LDFLAGS: -fPIC -L. -L${SRCDIR}/api/v6.6.9_CP_20220809_api_traderapi_se_linux64 -Wl,-rpath=${SRCDIR}/api/v6.6.9_CP_20220809_api_traderapi_se_linux64  -lthostmduserapi_se -lthosttraderapi_se -lstdc++
#cgo linux CPPFLAGS: -fPIC -I. -I/${SRCDIR}/api/v6.6.9_CP_20220809_api_traderapi_se_linux64/

#cgo darwin LDFLAGS: -L. -L${SRCDIR}/api/v6.6.9_MacOS_20220926 -lthostmduserapi_se -lthosttraderapi_se -lssl -lcrypto -lcomunicationkey -lMacDataCollect -framework Cocoa -framework IOKit
#cgo darwin CPPFLAGS: -I. -I${SRCDIR}/api/v6.6.9_MacOS_20220926

*/
import "C"
