package waf_openapi

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

// CreatePostpaidInstance invokes the waf_openapi.CreatePostpaidInstance API synchronously
func (client *Client) CreatePostpaidInstance(request *CreatePostpaidInstanceRequest) (response *CreatePostpaidInstanceResponse, err error) {
	response = CreateCreatePostpaidInstanceResponse()
	err = client.DoAction(request, response)
	return
}

// CreatePostpaidInstanceWithChan invokes the waf_openapi.CreatePostpaidInstance API asynchronously
func (client *Client) CreatePostpaidInstanceWithChan(request *CreatePostpaidInstanceRequest) (<-chan *CreatePostpaidInstanceResponse, <-chan error) {
	responseChan := make(chan *CreatePostpaidInstanceResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreatePostpaidInstance(request)
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

// CreatePostpaidInstanceWithCallback invokes the waf_openapi.CreatePostpaidInstance API asynchronously
func (client *Client) CreatePostpaidInstanceWithCallback(request *CreatePostpaidInstanceRequest, callback func(response *CreatePostpaidInstanceResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreatePostpaidInstanceResponse
		var err error
		defer close(result)
		response, err = client.CreatePostpaidInstance(request)
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

// CreatePostpaidInstanceRequest is the request struct for api CreatePostpaidInstance
type CreatePostpaidInstanceRequest struct {
	*requests.RpcRequest
	IsAutoRenew       requests.Boolean `position:"Query" name:"IsAutoRenew"`
	ClientToken       string           `position:"Query" name:"ClientToken"`
	PackageCode       string           `position:"Query" name:"PackageCode"`
	AutoRenewDuration requests.Integer `position:"Query" name:"AutoRenewDuration"`
	Duration          requests.Integer `position:"Query" name:"Duration"`
	SourceIp          string           `position:"Query" name:"SourceIp"`
	AliUid            requests.Integer `position:"Query" name:"AliUid"`
	Lang              string           `position:"Query" name:"Lang"`
	OwnerId           requests.Integer `position:"Query" name:"OwnerId"`
	PricingCycle      string           `position:"Query" name:"PricingCycle"`
}

// CreatePostpaidInstanceResponse is the response struct for api CreatePostpaidInstance
type CreatePostpaidInstanceResponse struct {
	*responses.BaseResponse
	OrderId    string `json:"OrderId" xml:"OrderId"`
	InstanceId string `json:"InstanceId" xml:"InstanceId"`
	RequestId  string `json:"RequestId" xml:"RequestId"`
}

// CreateCreatePostpaidInstanceRequest creates a request to invoke CreatePostpaidInstance API
func CreateCreatePostpaidInstanceRequest() (request *CreatePostpaidInstanceRequest) {
	request = &CreatePostpaidInstanceRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("waf-openapi", "2019-09-10", "CreatePostpaidInstance", "waf", "openAPI")
	request.Method = requests.POST
	return
}

// CreateCreatePostpaidInstanceResponse creates a response to parse from CreatePostpaidInstance response
func CreateCreatePostpaidInstanceResponse() (response *CreatePostpaidInstanceResponse) {
	response = &CreatePostpaidInstanceResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}