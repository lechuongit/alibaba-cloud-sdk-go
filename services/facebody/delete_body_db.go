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

// DeleteBodyDb invokes the facebody.DeleteBodyDb API synchronously
func (client *Client) DeleteBodyDb(request *DeleteBodyDbRequest) (response *DeleteBodyDbResponse, err error) {
	response = CreateDeleteBodyDbResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteBodyDbWithChan invokes the facebody.DeleteBodyDb API asynchronously
func (client *Client) DeleteBodyDbWithChan(request *DeleteBodyDbRequest) (<-chan *DeleteBodyDbResponse, <-chan error) {
	responseChan := make(chan *DeleteBodyDbResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteBodyDb(request)
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

// DeleteBodyDbWithCallback invokes the facebody.DeleteBodyDb API asynchronously
func (client *Client) DeleteBodyDbWithCallback(request *DeleteBodyDbRequest, callback func(response *DeleteBodyDbResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteBodyDbResponse
		var err error
		defer close(result)
		response, err = client.DeleteBodyDb(request)
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

// DeleteBodyDbRequest is the request struct for api DeleteBodyDb
type DeleteBodyDbRequest struct {
	*requests.RpcRequest
	Id requests.Integer `position:"Body" name:"Id"`
}

// DeleteBodyDbResponse is the response struct for api DeleteBodyDb
type DeleteBodyDbResponse struct {
	*responses.BaseResponse
	Message   string `json:"Message" xml:"Message"`
	RequestId string `json:"RequestId" xml:"RequestId"`
	Code      string `json:"Code" xml:"Code"`
}

// CreateDeleteBodyDbRequest creates a request to invoke DeleteBodyDb API
func CreateDeleteBodyDbRequest() (request *DeleteBodyDbRequest) {
	request = &DeleteBodyDbRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("facebody", "2019-12-30", "DeleteBodyDb", "", "")
	request.Method = requests.POST
	return
}

// CreateDeleteBodyDbResponse creates a response to parse from DeleteBodyDb response
func CreateDeleteBodyDbResponse() (response *DeleteBodyDbResponse) {
	response = &DeleteBodyDbResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
