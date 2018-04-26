package cdn

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

// ClearUserDomainBlackList invokes the cdn.ClearUserDomainBlackList API synchronously
// api document: https://help.aliyun.com/api/cdn/clearuserdomainblacklist.html
func (client *Client) ClearUserDomainBlackList(request *ClearUserDomainBlackListRequest) (response *ClearUserDomainBlackListResponse, err error) {
	response = CreateClearUserDomainBlackListResponse()
	err = client.DoAction(request, response)
	return
}

// ClearUserDomainBlackListWithChan invokes the cdn.ClearUserDomainBlackList API asynchronously
// api document: https://help.aliyun.com/api/cdn/clearuserdomainblacklist.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ClearUserDomainBlackListWithChan(request *ClearUserDomainBlackListRequest) (<-chan *ClearUserDomainBlackListResponse, <-chan error) {
	responseChan := make(chan *ClearUserDomainBlackListResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ClearUserDomainBlackList(request)
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

// ClearUserDomainBlackListWithCallback invokes the cdn.ClearUserDomainBlackList API asynchronously
// api document: https://help.aliyun.com/api/cdn/clearuserdomainblacklist.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ClearUserDomainBlackListWithCallback(request *ClearUserDomainBlackListRequest, callback func(response *ClearUserDomainBlackListResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ClearUserDomainBlackListResponse
		var err error
		defer close(result)
		response, err = client.ClearUserDomainBlackList(request)
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

// ClearUserDomainBlackListRequest is the request struct for api ClearUserDomainBlackList
type ClearUserDomainBlackListRequest struct {
	*requests.RpcRequest
	OwnerId       requests.Integer `position:"Query" name:"OwnerId"`
	OwnerAccount  string           `position:"Query" name:"OwnerAccount"`
	SecurityToken string           `position:"Query" name:"SecurityToken"`
	DomainName    string           `position:"Query" name:"DomainName"`
}

// ClearUserDomainBlackListResponse is the response struct for api ClearUserDomainBlackList
type ClearUserDomainBlackListResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateClearUserDomainBlackListRequest creates a request to invoke ClearUserDomainBlackList API
func CreateClearUserDomainBlackListRequest() (request *ClearUserDomainBlackListRequest) {
	request = &ClearUserDomainBlackListRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cdn", "2014-11-11", "ClearUserDomainBlackList", "", "")
	return
}

// CreateClearUserDomainBlackListResponse creates a response to parse from ClearUserDomainBlackList response
func CreateClearUserDomainBlackListResponse() (response *ClearUserDomainBlackListResponse) {
	response = &ClearUserDomainBlackListResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}