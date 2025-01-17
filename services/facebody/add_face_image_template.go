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

// AddFaceImageTemplate invokes the facebody.AddFaceImageTemplate API synchronously
func (client *Client) AddFaceImageTemplate(request *AddFaceImageTemplateRequest) (response *AddFaceImageTemplateResponse, err error) {
	response = CreateAddFaceImageTemplateResponse()
	err = client.DoAction(request, response)
	return
}

// AddFaceImageTemplateWithChan invokes the facebody.AddFaceImageTemplate API asynchronously
func (client *Client) AddFaceImageTemplateWithChan(request *AddFaceImageTemplateRequest) (<-chan *AddFaceImageTemplateResponse, <-chan error) {
	responseChan := make(chan *AddFaceImageTemplateResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.AddFaceImageTemplate(request)
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

// AddFaceImageTemplateWithCallback invokes the facebody.AddFaceImageTemplate API asynchronously
func (client *Client) AddFaceImageTemplateWithCallback(request *AddFaceImageTemplateRequest, callback func(response *AddFaceImageTemplateResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *AddFaceImageTemplateResponse
		var err error
		defer close(result)
		response, err = client.AddFaceImageTemplate(request)
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

// AddFaceImageTemplateRequest is the request struct for api AddFaceImageTemplate
type AddFaceImageTemplateRequest struct {
	*requests.RpcRequest
	UserId   string `position:"Body" name:"UserId"`
	ImageURL string `position:"Body" name:"ImageURL"`
}

// AddFaceImageTemplateResponse is the response struct for api AddFaceImageTemplate
type AddFaceImageTemplateResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Code      string `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
	Data      Data   `json:"Data" xml:"Data"`
}

// CreateAddFaceImageTemplateRequest creates a request to invoke AddFaceImageTemplate API
func CreateAddFaceImageTemplateRequest() (request *AddFaceImageTemplateRequest) {
	request = &AddFaceImageTemplateRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("facebody", "2019-12-30", "AddFaceImageTemplate", "", "")
	request.Method = requests.POST
	return
}

// CreateAddFaceImageTemplateResponse creates a response to parse from AddFaceImageTemplate response
func CreateAddFaceImageTemplateResponse() (response *AddFaceImageTemplateResponse) {
	response = &AddFaceImageTemplateResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
