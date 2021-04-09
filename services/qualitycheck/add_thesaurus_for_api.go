package qualitycheck

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

// AddThesaurusForApi invokes the qualitycheck.AddThesaurusForApi API synchronously
func (client *Client) AddThesaurusForApi(request *AddThesaurusForApiRequest) (response *AddThesaurusForApiResponse, err error) {
	response = CreateAddThesaurusForApiResponse()
	err = client.DoAction(request, response)
	return
}

// AddThesaurusForApiWithChan invokes the qualitycheck.AddThesaurusForApi API asynchronously
func (client *Client) AddThesaurusForApiWithChan(request *AddThesaurusForApiRequest) (<-chan *AddThesaurusForApiResponse, <-chan error) {
	responseChan := make(chan *AddThesaurusForApiResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.AddThesaurusForApi(request)
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

// AddThesaurusForApiWithCallback invokes the qualitycheck.AddThesaurusForApi API asynchronously
func (client *Client) AddThesaurusForApiWithCallback(request *AddThesaurusForApiRequest, callback func(response *AddThesaurusForApiResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *AddThesaurusForApiResponse
		var err error
		defer close(result)
		response, err = client.AddThesaurusForApi(request)
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

// AddThesaurusForApiRequest is the request struct for api AddThesaurusForApi
type AddThesaurusForApiRequest struct {
	*requests.RpcRequest
	ResourceOwnerId requests.Integer `position:"Query" name:"ResourceOwnerId"`
	JsonStr         string           `position:"Query" name:"JsonStr"`
}

// AddThesaurusForApiResponse is the response struct for api AddThesaurusForApi
type AddThesaurusForApiResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Success   bool   `json:"Success" xml:"Success"`
	Code      string `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
	Data      int64  `json:"Data" xml:"Data"`
}

// CreateAddThesaurusForApiRequest creates a request to invoke AddThesaurusForApi API
func CreateAddThesaurusForApiRequest() (request *AddThesaurusForApiRequest) {
	request = &AddThesaurusForApiRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Qualitycheck", "2019-01-15", "AddThesaurusForApi", "", "")
	request.Method = requests.POST
	return
}

// CreateAddThesaurusForApiResponse creates a response to parse from AddThesaurusForApi response
func CreateAddThesaurusForApiResponse() (response *AddThesaurusForApiResponse) {
	response = &AddThesaurusForApiResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
