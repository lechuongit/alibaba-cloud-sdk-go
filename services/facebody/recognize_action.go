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

// RecognizeAction invokes the facebody.RecognizeAction API synchronously
func (client *Client) RecognizeAction(request *RecognizeActionRequest) (response *RecognizeActionResponse, err error) {
	response = CreateRecognizeActionResponse()
	err = client.DoAction(request, response)
	return
}

// RecognizeActionWithChan invokes the facebody.RecognizeAction API asynchronously
func (client *Client) RecognizeActionWithChan(request *RecognizeActionRequest) (<-chan *RecognizeActionResponse, <-chan error) {
	responseChan := make(chan *RecognizeActionResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.RecognizeAction(request)
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

// RecognizeActionWithCallback invokes the facebody.RecognizeAction API asynchronously
func (client *Client) RecognizeActionWithCallback(request *RecognizeActionRequest, callback func(response *RecognizeActionResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *RecognizeActionResponse
		var err error
		defer close(result)
		response, err = client.RecognizeAction(request)
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

// RecognizeActionRequest is the request struct for api RecognizeAction
type RecognizeActionRequest struct {
	*requests.RpcRequest
	Type      requests.Integer          `position:"Body" name:"Type"`
	VideoData string                    `position:"Body" name:"VideoData"`
	URLList   *[]RecognizeActionURLList `position:"Body" name:"URLList"  type:"Repeated"`
	VideoUrl  string                    `position:"Body" name:"VideoUrl"`
}

// RecognizeActionURLList is a repeated param struct in RecognizeActionRequest
type RecognizeActionURLList struct {
	ImageData string `name:"imageData"`
	URL       string `name:"URL"`
}

// RecognizeActionResponse is the response struct for api RecognizeAction
type RecognizeActionResponse struct {
	*responses.BaseResponse
	RequestId string                `json:"RequestId" xml:"RequestId"`
	Data      DataInRecognizeAction `json:"Data" xml:"Data"`
}

// CreateRecognizeActionRequest creates a request to invoke RecognizeAction API
func CreateRecognizeActionRequest() (request *RecognizeActionRequest) {
	request = &RecognizeActionRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("facebody", "2019-12-30", "RecognizeAction", "", "")
	request.Method = requests.POST
	return
}

// CreateRecognizeActionResponse creates a response to parse from RecognizeAction response
func CreateRecognizeActionResponse() (response *RecognizeActionResponse) {
	response = &RecognizeActionResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
