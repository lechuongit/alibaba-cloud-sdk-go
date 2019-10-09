package green

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

// DescribeViewContent invokes the green.DescribeViewContent API synchronously
// api document: https://help.aliyun.com/api/green/describeviewcontent.html
func (client *Client) DescribeViewContent(request *DescribeViewContentRequest) (response *DescribeViewContentResponse, err error) {
	response = CreateDescribeViewContentResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeViewContentWithChan invokes the green.DescribeViewContent API asynchronously
// api document: https://help.aliyun.com/api/green/describeviewcontent.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeViewContentWithChan(request *DescribeViewContentRequest) (<-chan *DescribeViewContentResponse, <-chan error) {
	responseChan := make(chan *DescribeViewContentResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeViewContent(request)
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

// DescribeViewContentWithCallback invokes the green.DescribeViewContent API asynchronously
// api document: https://help.aliyun.com/api/green/describeviewcontent.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeViewContentWithCallback(request *DescribeViewContentRequest, callback func(response *DescribeViewContentResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeViewContentResponse
		var err error
		defer close(result)
		response, err = client.DescribeViewContent(request)
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

// DescribeViewContentRequest is the request struct for api DescribeViewContent
type DescribeViewContentRequest struct {
	*requests.RpcRequest
	ImageId      string           `position:"Query" name:"ImageId"`
	StartDate    string           `position:"Query" name:"StartDate"`
	Scene        string           `position:"Query" name:"Scene"`
	SourceIp     string           `position:"Query" name:"SourceIp"`
	LibType      string           `position:"Query" name:"LibType"`
	AuditResult  string           `position:"Query" name:"AuditResult"`
	PageSize     requests.Integer `position:"Query" name:"PageSize"`
	TaskId       string           `position:"Query" name:"TaskId"`
	TotalCount   requests.Integer `position:"Query" name:"TotalCount"`
	KeywordId    string           `position:"Query" name:"KeywordId"`
	Suggestion   string           `position:"Query" name:"Suggestion"`
	CurrentPage  requests.Integer `position:"Query" name:"CurrentPage"`
	Label        string           `position:"Query" name:"Label"`
	ResourceType string           `position:"Query" name:"ResourceType"`
	BizType      string           `position:"Query" name:"BizType"`
	EndDate      string           `position:"Query" name:"EndDate"`
	DataId       string           `position:"Query" name:"DataId"`
}

// DescribeViewContentResponse is the response struct for api DescribeViewContent
type DescribeViewContentResponse struct {
	*responses.BaseResponse
	RequestId       string        `json:"RequestId" xml:"RequestId"`
	PageSize        int           `json:"PageSize" xml:"PageSize"`
	CurrentPage     int           `json:"CurrentPage" xml:"CurrentPage"`
	TotalCount      int           `json:"TotalCount" xml:"TotalCount"`
	ViewContentList []ViewContent `json:"ViewContentList" xml:"ViewContentList"`
}

// CreateDescribeViewContentRequest creates a request to invoke DescribeViewContent API
func CreateDescribeViewContentRequest() (request *DescribeViewContentRequest) {
	request = &DescribeViewContentRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Green", "2017-08-23", "DescribeViewContent", "green", "openAPI")
	return
}

// CreateDescribeViewContentResponse creates a response to parse from DescribeViewContent response
func CreateDescribeViewContentResponse() (response *DescribeViewContentResponse) {
	response = &DescribeViewContentResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}