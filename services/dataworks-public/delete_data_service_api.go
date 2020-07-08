package dataworks_public

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

// DeleteDataServiceApi invokes the dataworks_public.DeleteDataServiceApi API synchronously
// api document: https://help.aliyun.com/api/dataworks-public/deletedataserviceapi.html
func (client *Client) DeleteDataServiceApi(request *DeleteDataServiceApiRequest) (response *DeleteDataServiceApiResponse, err error) {
	response = CreateDeleteDataServiceApiResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteDataServiceApiWithChan invokes the dataworks_public.DeleteDataServiceApi API asynchronously
// api document: https://help.aliyun.com/api/dataworks-public/deletedataserviceapi.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DeleteDataServiceApiWithChan(request *DeleteDataServiceApiRequest) (<-chan *DeleteDataServiceApiResponse, <-chan error) {
	responseChan := make(chan *DeleteDataServiceApiResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteDataServiceApi(request)
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

// DeleteDataServiceApiWithCallback invokes the dataworks_public.DeleteDataServiceApi API asynchronously
// api document: https://help.aliyun.com/api/dataworks-public/deletedataserviceapi.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DeleteDataServiceApiWithCallback(request *DeleteDataServiceApiRequest, callback func(response *DeleteDataServiceApiResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteDataServiceApiResponse
		var err error
		defer close(result)
		response, err = client.DeleteDataServiceApi(request)
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

// DeleteDataServiceApiRequest is the request struct for api DeleteDataServiceApi
type DeleteDataServiceApiRequest struct {
	*requests.RpcRequest
	TenantId  requests.Integer `position:"Body" name:"TenantId"`
	ProjectId requests.Integer `position:"Body" name:"ProjectId"`
	ApiId     requests.Integer `position:"Body" name:"ApiId"`
}

// DeleteDataServiceApiResponse is the response struct for api DeleteDataServiceApi
type DeleteDataServiceApiResponse struct {
	*responses.BaseResponse
	Data           bool   `json:"Data" xml:"Data"`
	ErrorCode      string `json:"ErrorCode" xml:"ErrorCode"`
	ErrorMessage   string `json:"ErrorMessage" xml:"ErrorMessage"`
	HttpStatusCode int    `json:"HttpStatusCode" xml:"HttpStatusCode"`
	RequestId      string `json:"RequestId" xml:"RequestId"`
	Success        bool   `json:"Success" xml:"Success"`
}

// CreateDeleteDataServiceApiRequest creates a request to invoke DeleteDataServiceApi API
func CreateDeleteDataServiceApiRequest() (request *DeleteDataServiceApiRequest) {
	request = &DeleteDataServiceApiRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("dataworks-public", "2020-05-18", "DeleteDataServiceApi", "dide", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDeleteDataServiceApiResponse creates a response to parse from DeleteDataServiceApi response
func CreateDeleteDataServiceApiResponse() (response *DeleteDataServiceApiResponse) {
	response = &DeleteDataServiceApiResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}