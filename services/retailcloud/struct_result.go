package retailcloud

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

// Result is a nested struct in retailcloud response
type Result struct {
	Nonsense                int                  `json:"Nonsense" xml:"Nonsense"`
	BusinessCode            string               `json:"BusinessCode" xml:"BusinessCode"`
	Headless                bool                 `json:"Headless" xml:"Headless"`
	AppEnvId                int64                `json:"AppEnvId" xml:"AppEnvId"`
	StickySession           int                  `json:"StickySession" xml:"StickySession"`
	SlbAPId                 int64                `json:"SlbAPId" xml:"SlbAPId"`
	LabelValue              string               `json:"LabelValue" xml:"LabelValue"`
	EnvType                 int                  `json:"EnvType" xml:"EnvType"`
	ListenerPort            int                  `json:"ListenerPort" xml:"ListenerPort"`
	ServiceType             string               `json:"ServiceType" xml:"ServiceType"`
	ClusterId               string               `json:"ClusterId" xml:"ClusterId"`
	SchemaId                int64                `json:"SchemaId" xml:"SchemaId"`
	RealServerPort          int                  `json:"RealServerPort" xml:"RealServerPort"`
	Id                      int64                `json:"Id" xml:"Id"`
	DeployOrderId           int64                `json:"DeployOrderId" xml:"DeployOrderId"`
	TotalBackupSize         int64                `json:"TotalBackupSize" xml:"TotalBackupSize"`
	Success                 bool                 `json:"Success" xml:"Success"`
	Admitted                bool                 `json:"Admitted" xml:"Admitted"`
	ServiceId               int64                `json:"ServiceId" xml:"ServiceId"`
	EnvTypeName             string               `json:"EnvTypeName" xml:"EnvTypeName"`
	PersistentVolumeClaimId int64                `json:"PersistentVolumeClaimId" xml:"PersistentVolumeClaimId"`
	SlbId                   string               `json:"SlbId" xml:"SlbId"`
	Name                    string               `json:"Name" xml:"Name"`
	Region                  string               `json:"Region" xml:"Region"`
	BizTitle                string               `json:"BizTitle" xml:"BizTitle"`
	ResourceDef             string               `json:"ResourceDef" xml:"ResourceDef"`
	Title                   string               `json:"Title" xml:"Title"`
	EnvId                   int64                `json:"EnvId" xml:"EnvId"`
	BizName                 string               `json:"BizName" xml:"BizName"`
	Language                string               `json:"Language" xml:"Language"`
	LabelKey                string               `json:"LabelKey" xml:"LabelKey"`
	DeployType              string               `json:"DeployType" xml:"DeployType"`
	AppId                   int64                `json:"AppId" xml:"AppId"`
	CookieTimeout           int                  `json:"CookieTimeout" xml:"CookieTimeout"`
	Replicas                int                  `json:"Replicas" xml:"Replicas"`
	SslCertId               string               `json:"SslCertId" xml:"SslCertId"`
	K8sServiceId            string               `json:"K8sServiceId" xml:"K8sServiceId"`
	PersistentVolumeId      int64                `json:"PersistentVolumeId" xml:"PersistentVolumeId"`
	TotalRecordCount        string               `json:"TotalRecordCount" xml:"TotalRecordCount"`
	Description             string               `json:"Description" xml:"Description"`
	DeployOrderName         string               `json:"DeployOrderName" xml:"DeployOrderName"`
	PageNumber              string               `json:"PageNumber" xml:"PageNumber"`
	EstablishedTimeout      int                  `json:"EstablishedTimeout" xml:"EstablishedTimeout"`
	Protocol                string               `json:"Protocol" xml:"Protocol"`
	EnvName                 string               `json:"EnvName" xml:"EnvName"`
	AppSchemaId             int64                `json:"AppSchemaId" xml:"AppSchemaId"`
	NetworkMode             string               `json:"NetworkMode" xml:"NetworkMode"`
	RequestId               string               `json:"RequestId" xml:"RequestId"`
	OperatingSystem         string               `json:"OperatingSystem" xml:"OperatingSystem"`
	ClusterInstanceId       string               `json:"ClusterInstanceId" xml:"ClusterInstanceId"`
	SlbIp                   string               `json:"SlbIp" xml:"SlbIp"`
	PageRecordCount         string               `json:"PageRecordCount" xml:"PageRecordCount"`
	NodeWorkLoadList        []string             `json:"NodeWorkLoadList" xml:"NodeWorkLoadList"`
	InstanceInfo            InstanceInfo         `json:"InstanceInfo" xml:"InstanceInfo"`
	WorkLoad                WorkLoad             `json:"WorkLoad" xml:"WorkLoad"`
	BasicInfo               BasicInfo            `json:"BasicInfo" xml:"BasicInfo"`
	NetInfo                 NetInfo              `json:"NetInfo" xml:"NetInfo"`
	PodEvents               []PodEvent           `json:"PodEvents" xml:"PodEvents"`
	PortMappings            []ServicePortMapping `json:"PortMappings" xml:"PortMappings"`
	Accounts                []AccountsItem       `json:"Accounts" xml:"Accounts"`
	DeployStepList          []DeployLogStepRC    `json:"DeployStepList" xml:"DeployStepList"`
	Items                   []Backup             `json:"Items" xml:"Items"`
	Databases               []DatabasesItem      `json:"Databases" xml:"Databases"`
}
