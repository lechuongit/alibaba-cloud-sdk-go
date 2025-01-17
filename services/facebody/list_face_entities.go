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

// ListFaceEntities invokes the facebody.ListFaceEntities API synchronously
func (client *Client) ListFaceEntities(request *ListFaceEntitiesRequest) (response *ListFaceEntitiesResponse, err error) {
	response = CreateListFaceEntitiesResponse()
	err = client.DoAction(request, response)
	return
}

// ListFaceEntitiesWithChan invokes the facebody.ListFaceEntities API asynchronously
func (client *Client) ListFaceEntitiesWithChan(request *ListFaceEntitiesRequest) (<-chan *ListFaceEntitiesResponse, <-chan error) {
	responseChan := make(chan *ListFaceEntitiesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListFaceEntities(request)
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

// ListFaceEntitiesWithCallback invokes the facebody.ListFaceEntities API asynchronously
func (client *Client) ListFaceEntitiesWithCallback(request *ListFaceEntitiesRequest, callback func(response *ListFaceEntitiesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListFaceEntitiesResponse
		var err error
		defer close(result)
		response, err = client.ListFaceEntities(request)
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

// ListFaceEntitiesRequest is the request struct for api ListFaceEntities
type ListFaceEntitiesRequest struct {
	*requests.RpcRequest
	EntityIdPrefix string           `position:"Body" name:"EntityIdPrefix"`
	Limit          requests.Integer `position:"Body" name:"Limit"`
	Order          string           `position:"Body" name:"Order"`
	Offset         requests.Integer `position:"Body" name:"Offset"`
	Token          string           `position:"Body" name:"Token"`
	Labels         string           `position:"Body" name:"Labels"`
	DbName         string           `position:"Body" name:"DbName"`
}

// ListFaceEntitiesResponse is the response struct for api ListFaceEntities
type ListFaceEntitiesResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Data      Data   `json:"Data" xml:"Data"`
}

// CreateListFaceEntitiesRequest creates a request to invoke ListFaceEntities API
func CreateListFaceEntitiesRequest() (request *ListFaceEntitiesRequest) {
	request = &ListFaceEntitiesRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("facebody", "2019-12-30", "ListFaceEntities", "", "")
	request.Method = requests.POST
	return
}

// CreateListFaceEntitiesResponse creates a response to parse from ListFaceEntities response
func CreateListFaceEntitiesResponse() (response *ListFaceEntitiesResponse) {
	response = &ListFaceEntitiesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
