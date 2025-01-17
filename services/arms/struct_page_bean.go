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

// PageBean is a nested struct in arms response
type PageBean struct {
	Total          int               `json:"Total" xml:"Total"`
	TotalCount     int               `json:"TotalCount" xml:"TotalCount"`
	PageNumber     int               `json:"PageNumber" xml:"PageNumber"`
	PageSize       int               `json:"PageSize" xml:"PageSize"`
	Event          []EventItem       `json:"Event" xml:"Event"`
	TraceApps      []TraceApp        `json:"TraceApps" xml:"TraceApps"`
	Contacts       []Contact         `json:"Contacts" xml:"Contacts"`
	AlertRules     []AlertRuleEntity `json:"AlertRules" xml:"AlertRules"`
	RetcodeApps    []RetcodeApp      `json:"RetcodeApps" xml:"RetcodeApps"`
	AlarmHistories []AlarmHistory    `json:"AlarmHistories" xml:"AlarmHistories"`
	TraceInfos     []TraceInfo       `json:"TraceInfos" xml:"TraceInfos"`
}
