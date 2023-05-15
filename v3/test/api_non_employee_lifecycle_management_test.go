/*
IdentityNow V3 API

Testing NonEmployeeLifecycleManagementApiService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package v3

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/sailpoint-oss/golang-sdk"
)

func Test_v3_NonEmployeeLifecycleManagementApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test NonEmployeeLifecycleManagementApiService ApproveNonEmployeeRequest", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		resp, httpRes, err := apiClient.NonEmployeeLifecycleManagementApi.ApproveNonEmployeeRequest(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test NonEmployeeLifecycleManagementApiService CreateNonEmployeeRecord", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.NonEmployeeLifecycleManagementApi.CreateNonEmployeeRecord(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test NonEmployeeLifecycleManagementApiService CreateNonEmployeeRequest", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.NonEmployeeLifecycleManagementApi.CreateNonEmployeeRequest(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test NonEmployeeLifecycleManagementApiService CreateNonEmployeeSource", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.NonEmployeeLifecycleManagementApi.CreateNonEmployeeSource(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test NonEmployeeLifecycleManagementApiService CreateNonEmployeeSourceSchemaAttributes", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var sourceId string

		resp, httpRes, err := apiClient.NonEmployeeLifecycleManagementApi.CreateNonEmployeeSourceSchemaAttributes(context.Background(), sourceId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test NonEmployeeLifecycleManagementApiService DeleteNonEmployeeRecord", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		httpRes, err := apiClient.NonEmployeeLifecycleManagementApi.DeleteNonEmployeeRecord(context.Background(), id).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test NonEmployeeLifecycleManagementApiService DeleteNonEmployeeRecordsInBulk", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		httpRes, err := apiClient.NonEmployeeLifecycleManagementApi.DeleteNonEmployeeRecordsInBulk(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test NonEmployeeLifecycleManagementApiService DeleteNonEmployeeRequest", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		httpRes, err := apiClient.NonEmployeeLifecycleManagementApi.DeleteNonEmployeeRequest(context.Background(), id).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test NonEmployeeLifecycleManagementApiService DeleteNonEmployeeSchemaAttribute", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var attributeId string
		var sourceId string

		httpRes, err := apiClient.NonEmployeeLifecycleManagementApi.DeleteNonEmployeeSchemaAttribute(context.Background(), attributeId, sourceId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test NonEmployeeLifecycleManagementApiService DeleteNonEmployeeSource", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var sourceId string

		httpRes, err := apiClient.NonEmployeeLifecycleManagementApi.DeleteNonEmployeeSource(context.Background(), sourceId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test NonEmployeeLifecycleManagementApiService DeleteNonEmployeeSourceSchemaAttributes", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var sourceId string

		httpRes, err := apiClient.NonEmployeeLifecycleManagementApi.DeleteNonEmployeeSourceSchemaAttributes(context.Background(), sourceId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test NonEmployeeLifecycleManagementApiService ExportNonEmployeeRecords", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		httpRes, err := apiClient.NonEmployeeLifecycleManagementApi.ExportNonEmployeeRecords(context.Background(), id).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test NonEmployeeLifecycleManagementApiService ExportNonEmployeeSourceSchemaTemplate", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		httpRes, err := apiClient.NonEmployeeLifecycleManagementApi.ExportNonEmployeeSourceSchemaTemplate(context.Background(), id).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test NonEmployeeLifecycleManagementApiService GetNonEmployeeApproval", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		resp, httpRes, err := apiClient.NonEmployeeLifecycleManagementApi.GetNonEmployeeApproval(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test NonEmployeeLifecycleManagementApiService GetNonEmployeeApprovalSummary", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var requestedFor string

		resp, httpRes, err := apiClient.NonEmployeeLifecycleManagementApi.GetNonEmployeeApprovalSummary(context.Background(), requestedFor).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test NonEmployeeLifecycleManagementApiService GetNonEmployeeBulkUploadStatus", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		resp, httpRes, err := apiClient.NonEmployeeLifecycleManagementApi.GetNonEmployeeBulkUploadStatus(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test NonEmployeeLifecycleManagementApiService GetNonEmployeeRecord", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		resp, httpRes, err := apiClient.NonEmployeeLifecycleManagementApi.GetNonEmployeeRecord(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test NonEmployeeLifecycleManagementApiService GetNonEmployeeRequest", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		resp, httpRes, err := apiClient.NonEmployeeLifecycleManagementApi.GetNonEmployeeRequest(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test NonEmployeeLifecycleManagementApiService GetNonEmployeeRequestSummary", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var requestedFor string

		resp, httpRes, err := apiClient.NonEmployeeLifecycleManagementApi.GetNonEmployeeRequestSummary(context.Background(), requestedFor).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test NonEmployeeLifecycleManagementApiService GetNonEmployeeSchemaAttribute", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var attributeId string
		var sourceId string

		resp, httpRes, err := apiClient.NonEmployeeLifecycleManagementApi.GetNonEmployeeSchemaAttribute(context.Background(), attributeId, sourceId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test NonEmployeeLifecycleManagementApiService GetNonEmployeeSource", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var sourceId string

		resp, httpRes, err := apiClient.NonEmployeeLifecycleManagementApi.GetNonEmployeeSource(context.Background(), sourceId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test NonEmployeeLifecycleManagementApiService GetNonEmployeeSourceSchemaAttributes", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var sourceId string

		resp, httpRes, err := apiClient.NonEmployeeLifecycleManagementApi.GetNonEmployeeSourceSchemaAttributes(context.Background(), sourceId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test NonEmployeeLifecycleManagementApiService ImportNonEmployeeRecordsInBulk", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		resp, httpRes, err := apiClient.NonEmployeeLifecycleManagementApi.ImportNonEmployeeRecordsInBulk(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test NonEmployeeLifecycleManagementApiService ListNonEmployeeApprovals", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.NonEmployeeLifecycleManagementApi.ListNonEmployeeApprovals(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test NonEmployeeLifecycleManagementApiService ListNonEmployeeRecords", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.NonEmployeeLifecycleManagementApi.ListNonEmployeeRecords(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test NonEmployeeLifecycleManagementApiService ListNonEmployeeRequests", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.NonEmployeeLifecycleManagementApi.ListNonEmployeeRequests(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test NonEmployeeLifecycleManagementApiService ListNonEmployeeSources", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.NonEmployeeLifecycleManagementApi.ListNonEmployeeSources(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test NonEmployeeLifecycleManagementApiService PatchNonEmployeeRecord", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		resp, httpRes, err := apiClient.NonEmployeeLifecycleManagementApi.PatchNonEmployeeRecord(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test NonEmployeeLifecycleManagementApiService PatchNonEmployeeSchemaAttribute", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var attributeId string
		var sourceId string

		resp, httpRes, err := apiClient.NonEmployeeLifecycleManagementApi.PatchNonEmployeeSchemaAttribute(context.Background(), attributeId, sourceId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test NonEmployeeLifecycleManagementApiService PatchNonEmployeeSource", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var sourceId string

		resp, httpRes, err := apiClient.NonEmployeeLifecycleManagementApi.PatchNonEmployeeSource(context.Background(), sourceId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test NonEmployeeLifecycleManagementApiService RejectNonEmployeeRequest", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		resp, httpRes, err := apiClient.NonEmployeeLifecycleManagementApi.RejectNonEmployeeRequest(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test NonEmployeeLifecycleManagementApiService UpdateNonEmployeeRecord", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		resp, httpRes, err := apiClient.NonEmployeeLifecycleManagementApi.UpdateNonEmployeeRecord(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
