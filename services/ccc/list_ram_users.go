package ccc

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

// ListRamUsers invokes the ccc.ListRamUsers API synchronously
func (client *Client) ListRamUsers(request *ListRamUsersRequest) (response *ListRamUsersResponse, err error) {
	response = CreateListRamUsersResponse()
	err = client.DoAction(request, response)
	return
}

// ListRamUsersWithChan invokes the ccc.ListRamUsers API asynchronously
func (client *Client) ListRamUsersWithChan(request *ListRamUsersRequest) (<-chan *ListRamUsersResponse, <-chan error) {
	responseChan := make(chan *ListRamUsersResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListRamUsers(request)
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

// ListRamUsersWithCallback invokes the ccc.ListRamUsers API asynchronously
func (client *Client) ListRamUsersWithCallback(request *ListRamUsersRequest, callback func(response *ListRamUsersResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListRamUsersResponse
		var err error
		defer close(result)
		response, err = client.ListRamUsers(request)
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

// ListRamUsersRequest is the request struct for api ListRamUsers
type ListRamUsersRequest struct {
	*requests.RpcRequest
	PageNumber    requests.Integer `position:"Query" name:"PageNumber"`
	SearchPattern string           `position:"Query" name:"SearchPattern"`
	InstanceId    string           `position:"Query" name:"InstanceId"`
	PageSize      requests.Integer `position:"Query" name:"PageSize"`
}

// ListRamUsersResponse is the response struct for api ListRamUsers
type ListRamUsersResponse struct {
	*responses.BaseResponse
	Code           string             `json:"Code" xml:"Code"`
	HttpStatusCode int                `json:"HttpStatusCode" xml:"HttpStatusCode"`
	Message        string             `json:"Message" xml:"Message"`
	RequestId      string             `json:"RequestId" xml:"RequestId"`
	Params         []string           `json:"Params" xml:"Params"`
	Data           DataInListRamUsers `json:"Data" xml:"Data"`
}

// CreateListRamUsersRequest creates a request to invoke ListRamUsers API
func CreateListRamUsersRequest() (request *ListRamUsersRequest) {
	request = &ListRamUsersRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CCC", "2020-07-01", "ListRamUsers", "CCC", "openAPI")
	request.Method = requests.POST
	return
}

// CreateListRamUsersResponse creates a response to parse from ListRamUsers response
func CreateListRamUsersResponse() (response *ListRamUsersResponse) {
	response = &ListRamUsersResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
