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

// PedestrianDetectAttribute invokes the facebody.PedestrianDetectAttribute API synchronously
func (client *Client) PedestrianDetectAttribute(request *PedestrianDetectAttributeRequest) (response *PedestrianDetectAttributeResponse, err error) {
	response = CreatePedestrianDetectAttributeResponse()
	err = client.DoAction(request, response)
	return
}

// PedestrianDetectAttributeWithChan invokes the facebody.PedestrianDetectAttribute API asynchronously
func (client *Client) PedestrianDetectAttributeWithChan(request *PedestrianDetectAttributeRequest) (<-chan *PedestrianDetectAttributeResponse, <-chan error) {
	responseChan := make(chan *PedestrianDetectAttributeResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.PedestrianDetectAttribute(request)
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

// PedestrianDetectAttributeWithCallback invokes the facebody.PedestrianDetectAttribute API asynchronously
func (client *Client) PedestrianDetectAttributeWithCallback(request *PedestrianDetectAttributeRequest, callback func(response *PedestrianDetectAttributeResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *PedestrianDetectAttributeResponse
		var err error
		defer close(result)
		response, err = client.PedestrianDetectAttribute(request)
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

// PedestrianDetectAttributeRequest is the request struct for api PedestrianDetectAttribute
type PedestrianDetectAttributeRequest struct {
	*requests.RpcRequest
	ImageURL string `position:"Body" name:"ImageURL"`
}

// PedestrianDetectAttributeResponse is the response struct for api PedestrianDetectAttribute
type PedestrianDetectAttributeResponse struct {
	*responses.BaseResponse
	RequestId string                          `json:"RequestId" xml:"RequestId"`
	Data      DataInPedestrianDetectAttribute `json:"Data" xml:"Data"`
}

// CreatePedestrianDetectAttributeRequest creates a request to invoke PedestrianDetectAttribute API
func CreatePedestrianDetectAttributeRequest() (request *PedestrianDetectAttributeRequest) {
	request = &PedestrianDetectAttributeRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("facebody", "2019-12-30", "PedestrianDetectAttribute", "", "")
	request.Method = requests.POST
	return
}

// CreatePedestrianDetectAttributeResponse creates a response to parse from PedestrianDetectAttribute response
func CreatePedestrianDetectAttributeResponse() (response *PedestrianDetectAttributeResponse) {
	response = &PedestrianDetectAttributeResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
