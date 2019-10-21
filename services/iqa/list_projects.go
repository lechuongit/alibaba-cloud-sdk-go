package iqa

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

// ListProjects invokes the iqa.ListProjects API synchronously
// api document: https://help.aliyun.com/api/iqa/listprojects.html
func (client *Client) ListProjects(request *ListProjectsRequest) (response *ListProjectsResponse, err error) {
	response = CreateListProjectsResponse()
	err = client.DoAction(request, response)
	return
}

// ListProjectsWithChan invokes the iqa.ListProjects API asynchronously
// api document: https://help.aliyun.com/api/iqa/listprojects.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ListProjectsWithChan(request *ListProjectsRequest) (<-chan *ListProjectsResponse, <-chan error) {
	responseChan := make(chan *ListProjectsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListProjects(request)
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

// ListProjectsWithCallback invokes the iqa.ListProjects API asynchronously
// api document: https://help.aliyun.com/api/iqa/listprojects.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ListProjectsWithCallback(request *ListProjectsRequest, callback func(response *ListProjectsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListProjectsResponse
		var err error
		defer close(result)
		response, err = client.ListProjects(request)
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

// ListProjectsRequest is the request struct for api ListProjects
type ListProjectsRequest struct {
	*requests.RpcRequest
	FilterParam string           `position:"Query" name:"FilterParam"`
	PageNumber  requests.Integer `position:"Query" name:"PageNumber"`
	PageSize    requests.Integer `position:"Query" name:"PageSize"`
	ProjectType string           `position:"Query" name:"ProjectType"`
}

// ListProjectsResponse is the response struct for api ListProjects
type ListProjectsResponse struct {
	*responses.BaseResponse
	RequestId  string    `json:"RequestId" xml:"RequestId"`
	TotalCount int       `json:"TotalCount" xml:"TotalCount"`
	PageNumber int       `json:"PageNumber" xml:"PageNumber"`
	PageSize   int       `json:"PageSize" xml:"PageSize"`
	Projects   []Project `json:"Projects" xml:"Projects"`
}

// CreateListProjectsRequest creates a request to invoke ListProjects API
func CreateListProjectsRequest() (request *ListProjectsRequest) {
	request = &ListProjectsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("iqa", "2019-08-13", "ListProjects", "iqa", "openAPI")
	return
}

// CreateListProjectsResponse creates a response to parse from ListProjects response
func CreateListProjectsResponse() (response *ListProjectsResponse) {
	response = &ListProjectsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
