package ecs

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

// DescribeTasks invokes the ecs.DescribeTasks API synchronously
// api document: https://help.aliyun.com/api/ecs/describetasks.html
func (client *Client) DescribeTasks(request *DescribeTasksRequest) (response *DescribeTasksResponse, err error) {
	response = CreateDescribeTasksResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeTasksWithChan invokes the ecs.DescribeTasks API asynchronously
// api document: https://help.aliyun.com/api/ecs/describetasks.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeTasksWithChan(request *DescribeTasksRequest) (<-chan *DescribeTasksResponse, <-chan error) {
	responseChan := make(chan *DescribeTasksResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeTasks(request)
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

// DescribeTasksWithCallback invokes the ecs.DescribeTasks API asynchronously
// api document: https://help.aliyun.com/api/ecs/describetasks.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeTasksWithCallback(request *DescribeTasksRequest, callback func(response *DescribeTasksResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeTasksResponse
		var err error
		defer close(result)
		response, err = client.DescribeTasks(request)
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

// DescribeTasksRequest is the request struct for api DescribeTasks
type DescribeTasksRequest struct {
	*requests.RpcRequest
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	PageNumber           requests.Integer `position:"Query" name:"PageNumber"`
	PageSize             requests.Integer `position:"Query" name:"PageSize"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	TaskIds              string           `position:"Query" name:"TaskIds"`
	TaskAction           string           `position:"Query" name:"TaskAction"`
	TaskStatus           string           `position:"Query" name:"TaskStatus"`
	StartTime            string           `position:"Query" name:"StartTime"`
	EndTime              string           `position:"Query" name:"EndTime"`
}

// DescribeTasksResponse is the response struct for api DescribeTasks
type DescribeTasksResponse struct {
	*responses.BaseResponse
	RequestId  string  `json:"RequestId" xml:"RequestId"`
	RegionId   string  `json:"RegionId" xml:"RegionId"`
	TotalCount int     `json:"TotalCount" xml:"TotalCount"`
	PageNumber int     `json:"PageNumber" xml:"PageNumber"`
	PageSize   int     `json:"PageSize" xml:"PageSize"`
	TaskSet    TaskSet `json:"TaskSet" xml:"TaskSet"`
}

// CreateDescribeTasksRequest creates a request to invoke DescribeTasks API
func CreateDescribeTasksRequest() (request *DescribeTasksRequest) {
	request = &DescribeTasksRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ecs", "2014-05-26", "DescribeTasks", "ecs", "openAPI")
	return
}

// CreateDescribeTasksResponse creates a response to parse from DescribeTasks response
func CreateDescribeTasksResponse() (response *DescribeTasksResponse) {
	response = &DescribeTasksResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}