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

// DetectLivingFace invokes the facebody.DetectLivingFace API synchronously
func (client *Client) DetectLivingFace(request *DetectLivingFaceRequest) (response *DetectLivingFaceResponse, err error) {
	response = CreateDetectLivingFaceResponse()
	err = client.DoAction(request, response)
	return
}

// DetectLivingFaceWithChan invokes the facebody.DetectLivingFace API asynchronously
func (client *Client) DetectLivingFaceWithChan(request *DetectLivingFaceRequest) (<-chan *DetectLivingFaceResponse, <-chan error) {
	responseChan := make(chan *DetectLivingFaceResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DetectLivingFace(request)
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

// DetectLivingFaceWithCallback invokes the facebody.DetectLivingFace API asynchronously
func (client *Client) DetectLivingFaceWithCallback(request *DetectLivingFaceRequest, callback func(response *DetectLivingFaceResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DetectLivingFaceResponse
		var err error
		defer close(result)
		response, err = client.DetectLivingFace(request)
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

// DetectLivingFaceRequest is the request struct for api DetectLivingFace
type DetectLivingFaceRequest struct {
	*requests.RpcRequest
	Tasks *[]DetectLivingFaceTasks `position:"Body" name:"Tasks"  type:"Repeated"`
}

// DetectLivingFaceTasks is a repeated param struct in DetectLivingFaceRequest
type DetectLivingFaceTasks struct {
	ImageURL  string `name:"ImageURL"`
	ImageData string `name:"ImageData"`
}

// DetectLivingFaceResponse is the response struct for api DetectLivingFace
type DetectLivingFaceResponse struct {
	*responses.BaseResponse
	RequestId string                 `json:"RequestId" xml:"RequestId"`
	Data      DataInDetectLivingFace `json:"Data" xml:"Data"`
}

// CreateDetectLivingFaceRequest creates a request to invoke DetectLivingFace API
func CreateDetectLivingFaceRequest() (request *DetectLivingFaceRequest) {
	request = &DetectLivingFaceRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("facebody", "2019-12-30", "DetectLivingFace", "", "")
	request.Method = requests.POST
	return
}

// CreateDetectLivingFaceResponse creates a response to parse from DetectLivingFace response
func CreateDetectLivingFaceResponse() (response *DetectLivingFaceResponse) {
	response = &DetectLivingFaceResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
