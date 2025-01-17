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

// FaceBeauty invokes the facebody.FaceBeauty API synchronously
func (client *Client) FaceBeauty(request *FaceBeautyRequest) (response *FaceBeautyResponse, err error) {
	response = CreateFaceBeautyResponse()
	err = client.DoAction(request, response)
	return
}

// FaceBeautyWithChan invokes the facebody.FaceBeauty API asynchronously
func (client *Client) FaceBeautyWithChan(request *FaceBeautyRequest) (<-chan *FaceBeautyResponse, <-chan error) {
	responseChan := make(chan *FaceBeautyResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.FaceBeauty(request)
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

// FaceBeautyWithCallback invokes the facebody.FaceBeauty API asynchronously
func (client *Client) FaceBeautyWithCallback(request *FaceBeautyRequest, callback func(response *FaceBeautyResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *FaceBeautyResponse
		var err error
		defer close(result)
		response, err = client.FaceBeauty(request)
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

// FaceBeautyRequest is the request struct for api FaceBeauty
type FaceBeautyRequest struct {
	*requests.RpcRequest
	White    requests.Float `position:"Body" name:"White"`
	Smooth   requests.Float `position:"Body" name:"Smooth"`
	Sharp    requests.Float `position:"Body" name:"Sharp"`
	ImageURL string         `position:"Body" name:"ImageURL"`
}

// FaceBeautyResponse is the response struct for api FaceBeauty
type FaceBeautyResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Data      Data   `json:"Data" xml:"Data"`
}

// CreateFaceBeautyRequest creates a request to invoke FaceBeauty API
func CreateFaceBeautyRequest() (request *FaceBeautyRequest) {
	request = &FaceBeautyRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("facebody", "2019-12-30", "FaceBeauty", "", "")
	request.Method = requests.POST
	return
}

// CreateFaceBeautyResponse creates a response to parse from FaceBeauty response
func CreateFaceBeautyResponse() (response *FaceBeautyResponse) {
	response = &FaceBeautyResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
