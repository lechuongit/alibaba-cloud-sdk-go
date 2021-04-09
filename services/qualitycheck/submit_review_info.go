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

// SubmitReviewInfo invokes the qualitycheck.SubmitReviewInfo API synchronously
func (client *Client) SubmitReviewInfo(request *SubmitReviewInfoRequest) (response *SubmitReviewInfoResponse, err error) {
	response = CreateSubmitReviewInfoResponse()
	err = client.DoAction(request, response)
	return
}

// SubmitReviewInfoWithChan invokes the qualitycheck.SubmitReviewInfo API asynchronously
func (client *Client) SubmitReviewInfoWithChan(request *SubmitReviewInfoRequest) (<-chan *SubmitReviewInfoResponse, <-chan error) {
	responseChan := make(chan *SubmitReviewInfoResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.SubmitReviewInfo(request)
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

// SubmitReviewInfoWithCallback invokes the qualitycheck.SubmitReviewInfo API asynchronously
func (client *Client) SubmitReviewInfoWithCallback(request *SubmitReviewInfoRequest, callback func(response *SubmitReviewInfoResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *SubmitReviewInfoResponse
		var err error
		defer close(result)
		response, err = client.SubmitReviewInfo(request)
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

// SubmitReviewInfoRequest is the request struct for api SubmitReviewInfo
type SubmitReviewInfoRequest struct {
	*requests.RpcRequest
	ResourceOwnerId requests.Integer `position:"Query" name:"ResourceOwnerId"`
	JsonStr         string           `position:"Query" name:"JsonStr"`
}

// SubmitReviewInfoResponse is the response struct for api SubmitReviewInfo
type SubmitReviewInfoResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Success   bool   `json:"Success" xml:"Success"`
	Code      string `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
	Data      string `json:"Data" xml:"Data"`
}

// CreateSubmitReviewInfoRequest creates a request to invoke SubmitReviewInfo API
func CreateSubmitReviewInfoRequest() (request *SubmitReviewInfoRequest) {
	request = &SubmitReviewInfoRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Qualitycheck", "2019-01-15", "SubmitReviewInfo", "", "")
	request.Method = requests.POST
	return
}

// CreateSubmitReviewInfoResponse creates a response to parse from SubmitReviewInfo response
func CreateSubmitReviewInfoResponse() (response *SubmitReviewInfoResponse) {
	response = &SubmitReviewInfoResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
