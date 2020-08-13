package dds

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

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

// CheckRecoveryCondition invokes the dds.CheckRecoveryCondition API synchronously
// api document: https://help.aliyun.com/api/dds/checkrecoverycondition.html
func (client *Client) CheckRecoveryCondition(request *CheckRecoveryConditionRequest) (response *CheckRecoveryConditionResponse, err error) {
	response = CreateCheckRecoveryConditionResponse()
	err = client.DoAction(request, response)
	return
}

// CheckRecoveryConditionWithChan invokes the dds.CheckRecoveryCondition API asynchronously
// api document: https://help.aliyun.com/api/dds/checkrecoverycondition.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CheckRecoveryConditionWithChan(request *CheckRecoveryConditionRequest) (<-chan *CheckRecoveryConditionResponse, <-chan error) {
	responseChan := make(chan *CheckRecoveryConditionResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CheckRecoveryCondition(request)
		if err != nil {
			errChan <- err
		} else {
			responseChan <- response
		}
	})
	if err != nil {
		errChan <- err
		close(responseChan)
		close(errChan)
	}
	return responseChan, errChan
}

// CheckRecoveryConditionWithCallback invokes the dds.CheckRecoveryCondition API asynchronously
// api document: https://help.aliyun.com/api/dds/checkrecoverycondition.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CheckRecoveryConditionWithCallback(request *CheckRecoveryConditionRequest, callback func(response *CheckRecoveryConditionResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CheckRecoveryConditionResponse
		var err error
		defer close(result)
		response, err = client.CheckRecoveryCondition(request)
		callback(response, err)
		result <- 1
	})
	if err != nil {
		defer close(result)
		callback(nil, err)
		result <- 0
	}
	return result
}

// CheckRecoveryConditionRequest is the request struct for api CheckRecoveryCondition
type CheckRecoveryConditionRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ResourceGroupId      string           `position:"Query" name:"ResourceGroupId"`
	DatabaseNames        string           `position:"Query" name:"DatabaseNames"`
	SecurityToken        string           `position:"Query" name:"SecurityToken"`
	RestoreTime          string           `position:"Query" name:"RestoreTime"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	BackupId             string           `position:"Query" name:"BackupId"`
	SourceDBInstance     string           `position:"Query" name:"SourceDBInstance"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
}

// CheckRecoveryConditionResponse is the response struct for api CheckRecoveryCondition
type CheckRecoveryConditionResponse struct {
	*responses.BaseResponse
	RequestId      string `json:"RequestId" xml:"RequestId"`
	DBInstanceName string `json:"DBInstanceName" xml:"DBInstanceName"`
	IsValid        bool   `json:"IsValid" xml:"IsValid"`
}

// CreateCheckRecoveryConditionRequest creates a request to invoke CheckRecoveryCondition API
func CreateCheckRecoveryConditionRequest() (request *CheckRecoveryConditionRequest) {
	request = &CheckRecoveryConditionRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Dds", "2015-12-01", "CheckRecoveryCondition", "Dds", "openAPI")
	request.Method = requests.POST
	return
}

// CreateCheckRecoveryConditionResponse creates a response to parse from CheckRecoveryCondition response
func CreateCheckRecoveryConditionResponse() (response *CheckRecoveryConditionResponse) {
	response = &CheckRecoveryConditionResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
