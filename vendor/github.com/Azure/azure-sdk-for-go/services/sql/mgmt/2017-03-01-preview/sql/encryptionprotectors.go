package sql

// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"net/http"
)

// EncryptionProtectorsClient is the the Azure SQL Database management API provides a RESTful set of web services that
// interact with Azure SQL Database services to manage your databases. The API enables you to create, retrieve, update,
// and delete databases.
type EncryptionProtectorsClient struct {
	ManagementClient
}

// NewEncryptionProtectorsClient creates an instance of the EncryptionProtectorsClient client.
func NewEncryptionProtectorsClient(subscriptionID string) EncryptionProtectorsClient {
	return NewEncryptionProtectorsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewEncryptionProtectorsClientWithBaseURI creates an instance of the EncryptionProtectorsClient client.
func NewEncryptionProtectorsClientWithBaseURI(baseURI string, subscriptionID string) EncryptionProtectorsClient {
	return EncryptionProtectorsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// CreateOrUpdate updates an existing encryption protector. This method may poll for completion. Polling can be
// canceled by passing the cancel channel argument. The channel will be used to cancel polling and any outstanding HTTP
// requests.
//
// resourceGroupName is the name of the resource group that contains the resource. You can obtain this value from the
// Azure Resource Manager API or the portal. serverName is the name of the server. encryptionProtectorName is the name
// of the encryption protector to be updated. parameters is the requested encryption protector resource state.
func (client EncryptionProtectorsClient) CreateOrUpdate(resourceGroupName string, serverName string, encryptionProtectorName string, parameters EncryptionProtector, cancel <-chan struct{}) (<-chan EncryptionProtector, <-chan error) {
	resultChan := make(chan EncryptionProtector, 1)
	errChan := make(chan error, 1)
	go func() {
		var err error
		var result EncryptionProtector
		defer func() {
			if err != nil {
				errChan <- err
			}
			resultChan <- result
			close(resultChan)
			close(errChan)
		}()
		req, err := client.CreateOrUpdatePreparer(resourceGroupName, serverName, encryptionProtectorName, parameters, cancel)
		if err != nil {
			err = autorest.NewErrorWithError(err, "sql.EncryptionProtectorsClient", "CreateOrUpdate", nil, "Failure preparing request")
			return
		}

		resp, err := client.CreateOrUpdateSender(req)
		if err != nil {
			result.Response = autorest.Response{Response: resp}
			err = autorest.NewErrorWithError(err, "sql.EncryptionProtectorsClient", "CreateOrUpdate", resp, "Failure sending request")
			return
		}

		result, err = client.CreateOrUpdateResponder(resp)
		if err != nil {
			err = autorest.NewErrorWithError(err, "sql.EncryptionProtectorsClient", "CreateOrUpdate", resp, "Failure responding to request")
		}
	}()
	return resultChan, errChan
}

// CreateOrUpdatePreparer prepares the CreateOrUpdate request.
func (client EncryptionProtectorsClient) CreateOrUpdatePreparer(resourceGroupName string, serverName string, encryptionProtectorName string, parameters EncryptionProtector, cancel <-chan struct{}) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"encryptionProtectorName": autorest.Encode("path", encryptionProtectorName),
		"resourceGroupName":       autorest.Encode("path", resourceGroupName),
		"serverName":              autorest.Encode("path", serverName),
		"subscriptionId":          autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2015-05-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsJSON(),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/encryptionProtector/{encryptionProtectorName}", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{Cancel: cancel})
}

// CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (client EncryptionProtectorsClient) CreateOrUpdateSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client,
		req,
		azure.DoRetryWithRegistration(client.Client),
		azure.DoPollForAsynchronous(client.PollingDelay))
}

// CreateOrUpdateResponder handles the response to the CreateOrUpdate request. The method always
// closes the http.Response Body.
func (client EncryptionProtectorsClient) CreateOrUpdateResponder(resp *http.Response) (result EncryptionProtector, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Get gets a server encryption protector.
//
// resourceGroupName is the name of the resource group that contains the resource. You can obtain this value from the
// Azure Resource Manager API or the portal. serverName is the name of the server. encryptionProtectorName is the name
// of the encryption protector to be retrieved.
func (client EncryptionProtectorsClient) Get(resourceGroupName string, serverName string, encryptionProtectorName string) (result EncryptionProtector, err error) {
	req, err := client.GetPreparer(resourceGroupName, serverName, encryptionProtectorName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.EncryptionProtectorsClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "sql.EncryptionProtectorsClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.EncryptionProtectorsClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client EncryptionProtectorsClient) GetPreparer(resourceGroupName string, serverName string, encryptionProtectorName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"encryptionProtectorName": autorest.Encode("path", encryptionProtectorName),
		"resourceGroupName":       autorest.Encode("path", resourceGroupName),
		"serverName":              autorest.Encode("path", serverName),
		"subscriptionId":          autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2015-05-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/encryptionProtector/{encryptionProtectorName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client EncryptionProtectorsClient) GetSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client,
		req,
		azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client EncryptionProtectorsClient) GetResponder(resp *http.Response) (result EncryptionProtector, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListByServer gets a list of server encryption protectors
//
// resourceGroupName is the name of the resource group that contains the resource. You can obtain this value from the
// Azure Resource Manager API or the portal. serverName is the name of the server.
func (client EncryptionProtectorsClient) ListByServer(resourceGroupName string, serverName string) (result EncryptionProtectorListResult, err error) {
	req, err := client.ListByServerPreparer(resourceGroupName, serverName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.EncryptionProtectorsClient", "ListByServer", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByServerSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "sql.EncryptionProtectorsClient", "ListByServer", resp, "Failure sending request")
		return
	}

	result, err = client.ListByServerResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.EncryptionProtectorsClient", "ListByServer", resp, "Failure responding to request")
	}

	return
}

// ListByServerPreparer prepares the ListByServer request.
func (client EncryptionProtectorsClient) ListByServerPreparer(resourceGroupName string, serverName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"serverName":        autorest.Encode("path", serverName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2015-05-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/encryptionProtector", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// ListByServerSender sends the ListByServer request. The method will close the
// http.Response Body if it receives an error.
func (client EncryptionProtectorsClient) ListByServerSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client,
		req,
		azure.DoRetryWithRegistration(client.Client))
}

// ListByServerResponder handles the response to the ListByServer request. The method always
// closes the http.Response Body.
func (client EncryptionProtectorsClient) ListByServerResponder(resp *http.Response) (result EncryptionProtectorListResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListByServerNextResults retrieves the next set of results, if any.
func (client EncryptionProtectorsClient) ListByServerNextResults(lastResults EncryptionProtectorListResult) (result EncryptionProtectorListResult, err error) {
	req, err := lastResults.EncryptionProtectorListResultPreparer()
	if err != nil {
		return result, autorest.NewErrorWithError(err, "sql.EncryptionProtectorsClient", "ListByServer", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}

	resp, err := client.ListByServerSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "sql.EncryptionProtectorsClient", "ListByServer", resp, "Failure sending next results request")
	}

	result, err = client.ListByServerResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.EncryptionProtectorsClient", "ListByServer", resp, "Failure responding to next results request")
	}

	return
}

// ListByServerComplete gets all elements from the list without paging.
func (client EncryptionProtectorsClient) ListByServerComplete(resourceGroupName string, serverName string, cancel <-chan struct{}) (<-chan EncryptionProtector, <-chan error) {
	resultChan := make(chan EncryptionProtector)
	errChan := make(chan error, 1)
	go func() {
		defer func() {
			close(resultChan)
			close(errChan)
		}()
		list, err := client.ListByServer(resourceGroupName, serverName)
		if err != nil {
			errChan <- err
			return
		}
		if list.Value != nil {
			for _, item := range *list.Value {
				select {
				case <-cancel:
					return
				case resultChan <- item:
					// Intentionally left blank
				}
			}
		}
		for list.NextLink != nil {
			list, err = client.ListByServerNextResults(list)
			if err != nil {
				errChan <- err
				return
			}
			if list.Value != nil {
				for _, item := range *list.Value {
					select {
					case <-cancel:
						return
					case resultChan <- item:
						// Intentionally left blank
					}
				}
			}
		}
	}()
	return resultChan, errChan
}
