package ens

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

// DescribeLoadBalancerUDPListenerAttribute invokes the ens.DescribeLoadBalancerUDPListenerAttribute API synchronously
func (client *Client) DescribeLoadBalancerUDPListenerAttribute(request *DescribeLoadBalancerUDPListenerAttributeRequest) (response *DescribeLoadBalancerUDPListenerAttributeResponse, err error) {
	response = CreateDescribeLoadBalancerUDPListenerAttributeResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeLoadBalancerUDPListenerAttributeWithChan invokes the ens.DescribeLoadBalancerUDPListenerAttribute API asynchronously
func (client *Client) DescribeLoadBalancerUDPListenerAttributeWithChan(request *DescribeLoadBalancerUDPListenerAttributeRequest) (<-chan *DescribeLoadBalancerUDPListenerAttributeResponse, <-chan error) {
	responseChan := make(chan *DescribeLoadBalancerUDPListenerAttributeResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeLoadBalancerUDPListenerAttribute(request)
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

// DescribeLoadBalancerUDPListenerAttributeWithCallback invokes the ens.DescribeLoadBalancerUDPListenerAttribute API asynchronously
func (client *Client) DescribeLoadBalancerUDPListenerAttributeWithCallback(request *DescribeLoadBalancerUDPListenerAttributeRequest, callback func(response *DescribeLoadBalancerUDPListenerAttributeResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeLoadBalancerUDPListenerAttributeResponse
		var err error
		defer close(result)
		response, err = client.DescribeLoadBalancerUDPListenerAttribute(request)
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

// DescribeLoadBalancerUDPListenerAttributeRequest is the request struct for api DescribeLoadBalancerUDPListenerAttribute
type DescribeLoadBalancerUDPListenerAttributeRequest struct {
	*requests.RpcRequest
	Protocol       string           `position:"Query" name:"Protocol"`
	ListenerPort   requests.Integer `position:"Query" name:"ListenerPort"`
	LoadBalancerId string           `position:"Query" name:"LoadBalancerId"`
}

// DescribeLoadBalancerUDPListenerAttributeResponse is the response struct for api DescribeLoadBalancerUDPListenerAttribute
type DescribeLoadBalancerUDPListenerAttributeResponse struct {
	*responses.BaseResponse
	RequestId                 string `json:"RequestId" xml:"RequestId"`
	ListenerPort              int    `json:"ListenerPort" xml:"ListenerPort"`
	Status                    string `json:"Status" xml:"Status"`
	Bandwidth                 int    `json:"Bandwidth" xml:"Bandwidth"`
	Scheduler                 string `json:"Scheduler" xml:"Scheduler"`
	HealthCheck               string `json:"HealthCheck" xml:"HealthCheck"`
	HealthyThreshold          int    `json:"HealthyThreshold" xml:"HealthyThreshold"`
	UnhealthyThreshold        int    `json:"UnhealthyThreshold" xml:"UnhealthyThreshold"`
	HealthCheckConnectTimeout int    `json:"HealthCheckConnectTimeout" xml:"HealthCheckConnectTimeout"`
	HealthCheckInterval       int    `json:"HealthCheckInterval" xml:"HealthCheckInterval"`
	Description               string `json:"Description" xml:"Description"`
	BackendServerPort         int    `json:"BackendServerPort" xml:"BackendServerPort"`
	HealthCheckConnectPort    int    `json:"HealthCheckConnectPort" xml:"HealthCheckConnectPort"`
	HealthCheckReq            string `json:"HealthCheckReq" xml:"HealthCheckReq"`
	HealthCheckExp            string `json:"HealthCheckExp" xml:"HealthCheckExp"`
}

// CreateDescribeLoadBalancerUDPListenerAttributeRequest creates a request to invoke DescribeLoadBalancerUDPListenerAttribute API
func CreateDescribeLoadBalancerUDPListenerAttributeRequest() (request *DescribeLoadBalancerUDPListenerAttributeRequest) {
	request = &DescribeLoadBalancerUDPListenerAttributeRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ens", "2017-11-10", "DescribeLoadBalancerUDPListenerAttribute", "ens", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeLoadBalancerUDPListenerAttributeResponse creates a response to parse from DescribeLoadBalancerUDPListenerAttribute response
func CreateDescribeLoadBalancerUDPListenerAttributeResponse() (response *DescribeLoadBalancerUDPListenerAttributeResponse) {
	response = &DescribeLoadBalancerUDPListenerAttributeResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
