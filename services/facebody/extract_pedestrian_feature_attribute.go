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

// ExtractPedestrianFeatureAttribute invokes the facebody.ExtractPedestrianFeatureAttribute API synchronously
func (client *Client) ExtractPedestrianFeatureAttribute(request *ExtractPedestrianFeatureAttributeRequest) (response *ExtractPedestrianFeatureAttributeResponse, err error) {
	response = CreateExtractPedestrianFeatureAttributeResponse()
	err = client.DoAction(request, response)
	return
}

// ExtractPedestrianFeatureAttributeWithChan invokes the facebody.ExtractPedestrianFeatureAttribute API asynchronously
func (client *Client) ExtractPedestrianFeatureAttributeWithChan(request *ExtractPedestrianFeatureAttributeRequest) (<-chan *ExtractPedestrianFeatureAttributeResponse, <-chan error) {
	responseChan := make(chan *ExtractPedestrianFeatureAttributeResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ExtractPedestrianFeatureAttribute(request)
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

// ExtractPedestrianFeatureAttributeWithCallback invokes the facebody.ExtractPedestrianFeatureAttribute API asynchronously
func (client *Client) ExtractPedestrianFeatureAttributeWithCallback(request *ExtractPedestrianFeatureAttributeRequest, callback func(response *ExtractPedestrianFeatureAttributeResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ExtractPedestrianFeatureAttributeResponse
		var err error
		defer close(result)
		response, err = client.ExtractPedestrianFeatureAttribute(request)
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

// ExtractPedestrianFeatureAttributeRequest is the request struct for api ExtractPedestrianFeatureAttribute
type ExtractPedestrianFeatureAttributeRequest struct {
	*requests.RpcRequest
	UrlList  *[]ExtractPedestrianFeatureAttributeUrlList `position:"Body" name:"UrlList"  type:"Repeated"`
	Mode     string                                      `position:"Body" name:"Mode"`
	ImageURL string                                      `position:"Body" name:"ImageURL"`
}

// ExtractPedestrianFeatureAttributeUrlList is a repeated param struct in ExtractPedestrianFeatureAttributeRequest
type ExtractPedestrianFeatureAttributeUrlList struct {
	Url string `name:"Url"`
}

// ExtractPedestrianFeatureAttributeResponse is the response struct for api ExtractPedestrianFeatureAttribute
type ExtractPedestrianFeatureAttributeResponse struct {
	*responses.BaseResponse
	RequestId string                                  `json:"RequestId" xml:"RequestId"`
	Data      DataInExtractPedestrianFeatureAttribute `json:"Data" xml:"Data"`
}

// CreateExtractPedestrianFeatureAttributeRequest creates a request to invoke ExtractPedestrianFeatureAttribute API
func CreateExtractPedestrianFeatureAttributeRequest() (request *ExtractPedestrianFeatureAttributeRequest) {
	request = &ExtractPedestrianFeatureAttributeRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("facebody", "2019-12-30", "ExtractPedestrianFeatureAttribute", "", "")
	request.Method = requests.POST
	return
}

// CreateExtractPedestrianFeatureAttributeResponse creates a response to parse from ExtractPedestrianFeatureAttribute response
func CreateExtractPedestrianFeatureAttributeResponse() (response *ExtractPedestrianFeatureAttributeResponse) {
	response = &ExtractPedestrianFeatureAttributeResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
