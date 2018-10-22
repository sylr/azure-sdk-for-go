package servicefabric

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
	"context"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/autorest/validation"
	"net/http"
)

// MeshNetworkClient is the service Fabric REST Client APIs allows management of Service Fabric clusters, applications
// and services.
type MeshNetworkClient struct {
	BaseClient
}

// NewMeshNetworkClient creates an instance of the MeshNetworkClient client.
func NewMeshNetworkClient() MeshNetworkClient {
	return NewMeshNetworkClientWithBaseURI(DefaultBaseURI)
}

// NewMeshNetworkClientWithBaseURI creates an instance of the MeshNetworkClient client.
func NewMeshNetworkClientWithBaseURI(baseURI string) MeshNetworkClient {
	return MeshNetworkClient{NewWithBaseURI(baseURI)}
}

// CreateOrUpdate creates a Network resource with the specified name, description and properties. If Network resource
// with the same name exists, then it is updated with the specified description and properties. Network resource
// provides connectivity between application services.
// Parameters:
// networkResourceName - the identity of the network.
// networkResourceDescription - description for creating a Network resource.
func (client MeshNetworkClient) CreateOrUpdate(ctx context.Context, networkResourceName string, networkResourceDescription NetworkResourceDescription) (result NetworkResourceDescription, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: networkResourceDescription,
			Constraints: []validation.Constraint{{Target: "networkResourceDescription.Name", Name: validation.Null, Rule: true, Chain: nil},
				{Target: "networkResourceDescription.Properties", Name: validation.Null, Rule: true, Chain: nil}}}}); err != nil {
		return result, validation.NewError("servicefabric.MeshNetworkClient", "CreateOrUpdate", err.Error())
	}

	req, err := client.CreateOrUpdatePreparer(ctx, networkResourceName, networkResourceDescription)
	if err != nil {
		err = autorest.NewErrorWithError(err, "servicefabric.MeshNetworkClient", "CreateOrUpdate", nil, "Failure preparing request")
		return
	}

	resp, err := client.CreateOrUpdateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "servicefabric.MeshNetworkClient", "CreateOrUpdate", resp, "Failure sending request")
		return
	}

	result, err = client.CreateOrUpdateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "servicefabric.MeshNetworkClient", "CreateOrUpdate", resp, "Failure responding to request")
	}

	return
}

// CreateOrUpdatePreparer prepares the CreateOrUpdate request.
func (client MeshNetworkClient) CreateOrUpdatePreparer(ctx context.Context, networkResourceName string, networkResourceDescription NetworkResourceDescription) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"networkResourceName": networkResourceName,
	}

	const APIVersion = "6.4-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/Resources/Networks/{networkResourceName}", pathParameters),
		autorest.WithJSON(networkResourceDescription),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (client MeshNetworkClient) CreateOrUpdateSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// CreateOrUpdateResponder handles the response to the CreateOrUpdate request. The method always
// closes the http.Response Body.
func (client MeshNetworkClient) CreateOrUpdateResponder(resp *http.Response) (result NetworkResourceDescription, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated, http.StatusAccepted),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete deletes the Network resource identified by the name.
// Parameters:
// networkResourceName - the identity of the network.
func (client MeshNetworkClient) Delete(ctx context.Context, networkResourceName string) (result autorest.Response, err error) {
	req, err := client.DeletePreparer(ctx, networkResourceName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "servicefabric.MeshNetworkClient", "Delete", nil, "Failure preparing request")
		return
	}

	resp, err := client.DeleteSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "servicefabric.MeshNetworkClient", "Delete", resp, "Failure sending request")
		return
	}

	result, err = client.DeleteResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "servicefabric.MeshNetworkClient", "Delete", resp, "Failure responding to request")
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client MeshNetworkClient) DeletePreparer(ctx context.Context, networkResourceName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"networkResourceName": networkResourceName,
	}

	const APIVersion = "6.4-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/Resources/Networks/{networkResourceName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client MeshNetworkClient) DeleteSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client MeshNetworkClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get gets the information about the Network resource with the given name. The information include the description and
// other properties of the Network.
// Parameters:
// networkResourceName - the identity of the network.
func (client MeshNetworkClient) Get(ctx context.Context, networkResourceName string) (result NetworkResourceDescription, err error) {
	req, err := client.GetPreparer(ctx, networkResourceName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "servicefabric.MeshNetworkClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "servicefabric.MeshNetworkClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "servicefabric.MeshNetworkClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client MeshNetworkClient) GetPreparer(ctx context.Context, networkResourceName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"networkResourceName": networkResourceName,
	}

	const APIVersion = "6.4-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/Resources/Networks/{networkResourceName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client MeshNetworkClient) GetSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client MeshNetworkClient) GetResponder(resp *http.Response) (result NetworkResourceDescription, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List gets the information about all network resources in a given resource group. The information include the
// description and other properties of the Network.
func (client MeshNetworkClient) List(ctx context.Context) (result PagedNetworkResourceDescriptionList, err error) {
	req, err := client.ListPreparer(ctx)
	if err != nil {
		err = autorest.NewErrorWithError(err, "servicefabric.MeshNetworkClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "servicefabric.MeshNetworkClient", "List", resp, "Failure sending request")
		return
	}

	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "servicefabric.MeshNetworkClient", "List", resp, "Failure responding to request")
	}

	return
}

// ListPreparer prepares the List request.
func (client MeshNetworkClient) ListPreparer(ctx context.Context) (*http.Request, error) {
	const APIVersion = "6.4-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/Resources/Networks"),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client MeshNetworkClient) ListSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client MeshNetworkClient) ListResponder(resp *http.Response) (result PagedNetworkResourceDescriptionList, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
