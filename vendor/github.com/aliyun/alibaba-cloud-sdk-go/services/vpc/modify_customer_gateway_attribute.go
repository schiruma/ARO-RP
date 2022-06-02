package vpc

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

// ModifyCustomerGatewayAttribute invokes the vpc.ModifyCustomerGatewayAttribute API synchronously
func (client *Client) ModifyCustomerGatewayAttribute(request *ModifyCustomerGatewayAttributeRequest) (response *ModifyCustomerGatewayAttributeResponse, err error) {
	response = CreateModifyCustomerGatewayAttributeResponse()
	err = client.DoAction(request, response)
	return
}

// ModifyCustomerGatewayAttributeWithChan invokes the vpc.ModifyCustomerGatewayAttribute API asynchronously
func (client *Client) ModifyCustomerGatewayAttributeWithChan(request *ModifyCustomerGatewayAttributeRequest) (<-chan *ModifyCustomerGatewayAttributeResponse, <-chan error) {
	responseChan := make(chan *ModifyCustomerGatewayAttributeResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifyCustomerGatewayAttribute(request)
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

// ModifyCustomerGatewayAttributeWithCallback invokes the vpc.ModifyCustomerGatewayAttribute API asynchronously
func (client *Client) ModifyCustomerGatewayAttributeWithCallback(request *ModifyCustomerGatewayAttributeRequest, callback func(response *ModifyCustomerGatewayAttributeResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifyCustomerGatewayAttributeResponse
		var err error
		defer close(result)
		response, err = client.ModifyCustomerGatewayAttribute(request)
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

// ModifyCustomerGatewayAttributeRequest is the request struct for api ModifyCustomerGatewayAttribute
type ModifyCustomerGatewayAttributeRequest struct {
	*requests.RpcRequest
	AuthKey              string           `position:"Query" name:"AuthKey"`
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ClientToken          string           `position:"Query" name:"ClientToken"`
	Description          string           `position:"Query" name:"Description"`
	CustomerGatewayId    string           `position:"Query" name:"CustomerGatewayId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	Name                 string           `position:"Query" name:"Name"`
}

// ModifyCustomerGatewayAttributeResponse is the response struct for api ModifyCustomerGatewayAttribute
type ModifyCustomerGatewayAttributeResponse struct {
	*responses.BaseResponse
	RequestId         string `json:"RequestId" xml:"RequestId"`
	IpAddress         string `json:"IpAddress" xml:"IpAddress"`
	Description       string `json:"Description" xml:"Description"`
	CustomerGatewayId string `json:"CustomerGatewayId" xml:"CustomerGatewayId"`
	CreateTime        int64  `json:"CreateTime" xml:"CreateTime"`
	Name              string `json:"Name" xml:"Name"`
}

// CreateModifyCustomerGatewayAttributeRequest creates a request to invoke ModifyCustomerGatewayAttribute API
func CreateModifyCustomerGatewayAttributeRequest() (request *ModifyCustomerGatewayAttributeRequest) {
	request = &ModifyCustomerGatewayAttributeRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Vpc", "2016-04-28", "ModifyCustomerGatewayAttribute", "vpc", "openAPI")
	request.Method = requests.POST
	return
}

// CreateModifyCustomerGatewayAttributeResponse creates a response to parse from ModifyCustomerGatewayAttribute response
func CreateModifyCustomerGatewayAttributeResponse() (response *ModifyCustomerGatewayAttributeResponse) {
	response = &ModifyCustomerGatewayAttributeResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}