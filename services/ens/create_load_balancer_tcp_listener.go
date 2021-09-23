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

// CreateLoadBalancerTCPListener invokes the ens.CreateLoadBalancerTCPListener API synchronously
func (client *Client) CreateLoadBalancerTCPListener(request *CreateLoadBalancerTCPListenerRequest) (response *CreateLoadBalancerTCPListenerResponse, err error) {
	response = CreateCreateLoadBalancerTCPListenerResponse()
	err = client.DoAction(request, response)
	return
}

// CreateLoadBalancerTCPListenerWithChan invokes the ens.CreateLoadBalancerTCPListener API asynchronously
func (client *Client) CreateLoadBalancerTCPListenerWithChan(request *CreateLoadBalancerTCPListenerRequest) (<-chan *CreateLoadBalancerTCPListenerResponse, <-chan error) {
	responseChan := make(chan *CreateLoadBalancerTCPListenerResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateLoadBalancerTCPListener(request)
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

// CreateLoadBalancerTCPListenerWithCallback invokes the ens.CreateLoadBalancerTCPListener API asynchronously
func (client *Client) CreateLoadBalancerTCPListenerWithCallback(request *CreateLoadBalancerTCPListenerRequest, callback func(response *CreateLoadBalancerTCPListenerResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateLoadBalancerTCPListenerResponse
		var err error
		defer close(result)
		response, err = client.CreateLoadBalancerTCPListener(request)
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

// CreateLoadBalancerTCPListenerRequest is the request struct for api CreateLoadBalancerTCPListener
type CreateLoadBalancerTCPListenerRequest struct {
	*requests.RpcRequest
	HealthCheckURI            string           `position:"Query" name:"HealthCheckURI"`
	Protocol                  string           `position:"Query" name:"Protocol"`
	EstablishedTimeout        requests.Integer `position:"Query" name:"EstablishedTimeout"`
	PersistenceTimeout        requests.Integer `position:"Query" name:"PersistenceTimeout"`
	HealthCheckDomain         string           `position:"Query" name:"HealthCheckDomain"`
	LoadBalancerId            string           `position:"Query" name:"LoadBalancerId"`
	HealthCheckInterval       requests.Integer `position:"Query" name:"HealthCheckInterval"`
	BackendServerPort         requests.Integer `position:"Query" name:"BackendServerPort"`
	HealthCheckConnectTimeout requests.Integer `position:"Query" name:"HealthCheckConnectTimeout"`
	Description               string           `position:"Query" name:"Description"`
	UnhealthyThreshold        requests.Integer `position:"Query" name:"UnhealthyThreshold"`
	HealthyThreshold          requests.Integer `position:"Query" name:"HealthyThreshold"`
	Scheduler                 string           `position:"Query" name:"Scheduler"`
	ListenerPort              requests.Integer `position:"Query" name:"ListenerPort"`
	HealthCheckType           string           `position:"Query" name:"HealthCheckType"`
	HealthCheckHttpCode       string           `position:"Query" name:"HealthCheckHttpCode"`
	HealthCheckConnectPort    requests.Integer `position:"Query" name:"HealthCheckConnectPort"`
}

// CreateLoadBalancerTCPListenerResponse is the response struct for api CreateLoadBalancerTCPListener
type CreateLoadBalancerTCPListenerResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateCreateLoadBalancerTCPListenerRequest creates a request to invoke CreateLoadBalancerTCPListener API
func CreateCreateLoadBalancerTCPListenerRequest() (request *CreateLoadBalancerTCPListenerRequest) {
	request = &CreateLoadBalancerTCPListenerRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ens", "2017-11-10", "CreateLoadBalancerTCPListener", "ens", "openAPI")
	request.Method = requests.POST
	return
}

// CreateCreateLoadBalancerTCPListenerResponse creates a response to parse from CreateLoadBalancerTCPListener response
func CreateCreateLoadBalancerTCPListenerResponse() (response *CreateLoadBalancerTCPListenerResponse) {
	response = &CreateLoadBalancerTCPListenerResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
