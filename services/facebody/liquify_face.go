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

// LiquifyFace invokes the facebody.LiquifyFace API synchronously
func (client *Client) LiquifyFace(request *LiquifyFaceRequest) (response *LiquifyFaceResponse, err error) {
	response = CreateLiquifyFaceResponse()
	err = client.DoAction(request, response)
	return
}

// LiquifyFaceWithChan invokes the facebody.LiquifyFace API asynchronously
func (client *Client) LiquifyFaceWithChan(request *LiquifyFaceRequest) (<-chan *LiquifyFaceResponse, <-chan error) {
	responseChan := make(chan *LiquifyFaceResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.LiquifyFace(request)
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

// LiquifyFaceWithCallback invokes the facebody.LiquifyFace API asynchronously
func (client *Client) LiquifyFaceWithCallback(request *LiquifyFaceRequest, callback func(response *LiquifyFaceResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *LiquifyFaceResponse
		var err error
		defer close(result)
		response, err = client.LiquifyFace(request)
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

// LiquifyFaceRequest is the request struct for api LiquifyFace
type LiquifyFaceRequest struct {
	*requests.RpcRequest
	SlimDegree requests.Float `position:"Body" name:"SlimDegree"`
	ImageURL   string         `position:"Body" name:"ImageURL"`
}

// LiquifyFaceResponse is the response struct for api LiquifyFace
type LiquifyFaceResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Code      string `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
	Data      Data   `json:"Data" xml:"Data"`
}

// CreateLiquifyFaceRequest creates a request to invoke LiquifyFace API
func CreateLiquifyFaceRequest() (request *LiquifyFaceRequest) {
	request = &LiquifyFaceRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("facebody", "2019-12-30", "LiquifyFace", "", "")
	request.Method = requests.POST
	return
}

// CreateLiquifyFaceResponse creates a response to parse from LiquifyFace response
func CreateLiquifyFaceResponse() (response *LiquifyFaceResponse) {
	response = &LiquifyFaceResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
