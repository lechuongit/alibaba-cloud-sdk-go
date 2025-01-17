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

// DescribeAsyncTaskStatus invokes the waf_openapi.DescribeAsyncTaskStatus API synchronously
func (client *Client) DescribeAsyncTaskStatus(request *DescribeAsyncTaskStatusRequest) (response *DescribeAsyncTaskStatusResponse, err error) {
	response = CreateDescribeAsyncTaskStatusResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeAsyncTaskStatusWithChan invokes the waf_openapi.DescribeAsyncTaskStatus API asynchronously
func (client *Client) DescribeAsyncTaskStatusWithChan(request *DescribeAsyncTaskStatusRequest) (<-chan *DescribeAsyncTaskStatusResponse, <-chan error) {
	responseChan := make(chan *DescribeAsyncTaskStatusResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeAsyncTaskStatus(request)
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

// DescribeAsyncTaskStatusWithCallback invokes the waf_openapi.DescribeAsyncTaskStatus API asynchronously
func (client *Client) DescribeAsyncTaskStatusWithCallback(request *DescribeAsyncTaskStatusRequest, callback func(response *DescribeAsyncTaskStatusResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeAsyncTaskStatusResponse
		var err error
		defer close(result)
		response, err = client.DescribeAsyncTaskStatus(request)
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

// DescribeAsyncTaskStatusRequest is the request struct for api DescribeAsyncTaskStatus
type DescribeAsyncTaskStatusRequest struct {
	*requests.RpcRequest
	ResourceGroupId string `position:"Query" name:"ResourceGroupId"`
	InstanceId      string `position:"Query" name:"InstanceId"`
	SourceIp        string `position:"Query" name:"SourceIp"`
	WafRequestId    string `position:"Query" name:"WafRequestId"`
	Lang            string `position:"Query" name:"Lang"`
	Region          string `position:"Query" name:"Region"`
}

// DescribeAsyncTaskStatusResponse is the response struct for api DescribeAsyncTaskStatus
type DescribeAsyncTaskStatusResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Result    Result `json:"Result" xml:"Result"`
}

// CreateDescribeAsyncTaskStatusRequest creates a request to invoke DescribeAsyncTaskStatus API
func CreateDescribeAsyncTaskStatusRequest() (request *DescribeAsyncTaskStatusRequest) {
	request = &DescribeAsyncTaskStatusRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("waf-openapi", "2018-01-17", "DescribeAsyncTaskStatus", "waf", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeAsyncTaskStatusResponse creates a response to parse from DescribeAsyncTaskStatus response
func CreateDescribeAsyncTaskStatusResponse() (response *DescribeAsyncTaskStatusResponse) {
	response = &DescribeAsyncTaskStatusResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
