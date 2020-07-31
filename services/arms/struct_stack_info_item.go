package arms

//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.
//
// Code generated by Alibaba Cloud SDK Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

// StackInfoItem is a nested struct in arms response
type StackInfoItem struct {
	StartTime   int64   `json:"StartTime" xml:"StartTime"`
	Duration    int64   `json:"Duration" xml:"Duration"`
	RpcId       string  `json:"RpcId" xml:"RpcId"`
	ServiceName string  `json:"ServiceName" xml:"ServiceName"`
	Api         string  `json:"Api" xml:"Api"`
	Exception   string  `json:"Exception" xml:"Exception"`
	Line        string  `json:"Line" xml:"Line"`
	ExtInfo     ExtInfo `json:"ExtInfo" xml:"ExtInfo"`
}
