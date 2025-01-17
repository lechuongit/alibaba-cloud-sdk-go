package facebody

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

// ListBodyDbs invokes the facebody.ListBodyDbs API synchronously
func (client *Client) ListBodyDbs(request *ListBodyDbsRequest) (response *ListBodyDbsResponse, err error) {
	response = CreateListBodyDbsResponse()
	err = client.DoAction(request, response)
	return
}

// ListBodyDbsWithChan invokes the facebody.ListBodyDbs API asynchronously
func (client *Client) ListBodyDbsWithChan(request *ListBodyDbsRequest) (<-chan *ListBodyDbsResponse, <-chan error) {
	responseChan := make(chan *ListBodyDbsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListBodyDbs(request)
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

// ListBodyDbsWithCallback invokes the facebody.ListBodyDbs API asynchronously
func (client *Client) ListBodyDbsWithCallback(request *ListBodyDbsRequest, callback func(response *ListBodyDbsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListBodyDbsResponse
		var err error
		defer close(result)
		response, err = client.ListBodyDbs(request)
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

// ListBodyDbsRequest is the request struct for api ListBodyDbs
type ListBodyDbsRequest struct {
	*requests.RpcRequest
	Limit  requests.Integer `position:"Query" name:"Limit"`
	Offset requests.Integer `position:"Query" name:"Offset"`
}

// ListBodyDbsResponse is the response struct for api ListBodyDbs
type ListBodyDbsResponse struct {
	*responses.BaseResponse
	Message   string `json:"Message" xml:"Message"`
	RequestId string `json:"RequestId" xml:"RequestId"`
	Code      string `json:"Code" xml:"Code"`
	Data      Data   `json:"Data" xml:"Data"`
}

// CreateListBodyDbsRequest creates a request to invoke ListBodyDbs API
func CreateListBodyDbsRequest() (request *ListBodyDbsRequest) {
	request = &ListBodyDbsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("facebody", "2019-12-30", "ListBodyDbs", "", "")
	request.Method = requests.GET
	return
}

// CreateListBodyDbsResponse creates a response to parse from ListBodyDbs response
func CreateListBodyDbsResponse() (response *ListBodyDbsResponse) {
	response = &ListBodyDbsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
