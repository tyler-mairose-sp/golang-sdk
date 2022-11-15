/*
IdentityNow Beta API

Testing NonEmployeeLifecycleManagementApiService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package sailpointsdk

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func Test_sailpointsdk_NonEmployeeLifecycleManagementApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test NonEmployeeLifecycleManagementApiService CreateSchemaAttribute", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var sourceId string

		resp, httpRes, err := apiClient.NonEmployeeLifecycleManagementApi.CreateSchemaAttribute(context.Background(), sourceId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test NonEmployeeLifecycleManagementApiService DeleteSchemaAttribute", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var attributeId string
		var sourceId string

		resp, httpRes, err := apiClient.NonEmployeeLifecycleManagementApi.DeleteSchemaAttribute(context.Background(), attributeId, sourceId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test NonEmployeeLifecycleManagementApiService DeleteSchemaAttributes", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var sourceId string

		resp, httpRes, err := apiClient.NonEmployeeLifecycleManagementApi.DeleteSchemaAttributes(context.Background(), sourceId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test NonEmployeeLifecycleManagementApiService GetSchemaAttribute", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var attributeId string
		var sourceId string

		resp, httpRes, err := apiClient.NonEmployeeLifecycleManagementApi.GetSchemaAttribute(context.Background(), attributeId, sourceId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test NonEmployeeLifecycleManagementApiService GetSchemaAttributes", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var sourceId string

		resp, httpRes, err := apiClient.NonEmployeeLifecycleManagementApi.GetSchemaAttributes(context.Background(), sourceId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test NonEmployeeLifecycleManagementApiService NonEmployeeApprovalGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		resp, httpRes, err := apiClient.NonEmployeeLifecycleManagementApi.NonEmployeeApprovalGet(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test NonEmployeeLifecycleManagementApiService NonEmployeeApprovalList", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.NonEmployeeLifecycleManagementApi.NonEmployeeApprovalList(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test NonEmployeeLifecycleManagementApiService NonEmployeeApprovalSummary", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var requestedFor string

		resp, httpRes, err := apiClient.NonEmployeeLifecycleManagementApi.NonEmployeeApprovalSummary(context.Background(), requestedFor).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test NonEmployeeLifecycleManagementApiService NonEmployeeApproveRequest", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		resp, httpRes, err := apiClient.NonEmployeeLifecycleManagementApi.NonEmployeeApproveRequest(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test NonEmployeeLifecycleManagementApiService NonEmployeeBulkUploadStatus", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		resp, httpRes, err := apiClient.NonEmployeeLifecycleManagementApi.NonEmployeeBulkUploadStatus(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test NonEmployeeLifecycleManagementApiService NonEmployeeExportSourceSchemaTemplate", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		resp, httpRes, err := apiClient.NonEmployeeLifecycleManagementApi.NonEmployeeExportSourceSchemaTemplate(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test NonEmployeeLifecycleManagementApiService NonEmployeeRecordBulkDelete", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.NonEmployeeLifecycleManagementApi.NonEmployeeRecordBulkDelete(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test NonEmployeeLifecycleManagementApiService NonEmployeeRecordCreation", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.NonEmployeeLifecycleManagementApi.NonEmployeeRecordCreation(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test NonEmployeeLifecycleManagementApiService NonEmployeeRecordDelete", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		resp, httpRes, err := apiClient.NonEmployeeLifecycleManagementApi.NonEmployeeRecordDelete(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test NonEmployeeLifecycleManagementApiService NonEmployeeRecordGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		resp, httpRes, err := apiClient.NonEmployeeLifecycleManagementApi.NonEmployeeRecordGet(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test NonEmployeeLifecycleManagementApiService NonEmployeeRecordList", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.NonEmployeeLifecycleManagementApi.NonEmployeeRecordList(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test NonEmployeeLifecycleManagementApiService NonEmployeeRecordPatch", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		resp, httpRes, err := apiClient.NonEmployeeLifecycleManagementApi.NonEmployeeRecordPatch(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test NonEmployeeLifecycleManagementApiService NonEmployeeRecordUpdate", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		resp, httpRes, err := apiClient.NonEmployeeLifecycleManagementApi.NonEmployeeRecordUpdate(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test NonEmployeeLifecycleManagementApiService NonEmployeeRecordsBulkUpload", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		resp, httpRes, err := apiClient.NonEmployeeLifecycleManagementApi.NonEmployeeRecordsBulkUpload(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test NonEmployeeLifecycleManagementApiService NonEmployeeRecordsExport", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		resp, httpRes, err := apiClient.NonEmployeeLifecycleManagementApi.NonEmployeeRecordsExport(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test NonEmployeeLifecycleManagementApiService NonEmployeeRejectRequest", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		resp, httpRes, err := apiClient.NonEmployeeLifecycleManagementApi.NonEmployeeRejectRequest(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test NonEmployeeLifecycleManagementApiService NonEmployeeRequestCreation", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.NonEmployeeLifecycleManagementApi.NonEmployeeRequestCreation(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test NonEmployeeLifecycleManagementApiService NonEmployeeRequestDeletion", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		resp, httpRes, err := apiClient.NonEmployeeLifecycleManagementApi.NonEmployeeRequestDeletion(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test NonEmployeeLifecycleManagementApiService NonEmployeeRequestGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		resp, httpRes, err := apiClient.NonEmployeeLifecycleManagementApi.NonEmployeeRequestGet(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test NonEmployeeLifecycleManagementApiService NonEmployeeRequestList", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.NonEmployeeLifecycleManagementApi.NonEmployeeRequestList(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test NonEmployeeLifecycleManagementApiService NonEmployeeRequestSummaryGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var requestedFor string

		resp, httpRes, err := apiClient.NonEmployeeLifecycleManagementApi.NonEmployeeRequestSummaryGet(context.Background(), requestedFor).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test NonEmployeeLifecycleManagementApiService NonEmployeeSourceDelete", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var sourceId string

		resp, httpRes, err := apiClient.NonEmployeeLifecycleManagementApi.NonEmployeeSourceDelete(context.Background(), sourceId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test NonEmployeeLifecycleManagementApiService NonEmployeeSourceGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var sourceId string

		resp, httpRes, err := apiClient.NonEmployeeLifecycleManagementApi.NonEmployeeSourceGet(context.Background(), sourceId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test NonEmployeeLifecycleManagementApiService NonEmployeeSourcePatch", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var sourceId string

		resp, httpRes, err := apiClient.NonEmployeeLifecycleManagementApi.NonEmployeeSourcePatch(context.Background(), sourceId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test NonEmployeeLifecycleManagementApiService NonEmployeeSourcesCreation", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.NonEmployeeLifecycleManagementApi.NonEmployeeSourcesCreation(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test NonEmployeeLifecycleManagementApiService NonEmployeeSourcesList", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.NonEmployeeLifecycleManagementApi.NonEmployeeSourcesList(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test NonEmployeeLifecycleManagementApiService PatchSchemaAttribute", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var attributeId string
		var sourceId string

		resp, httpRes, err := apiClient.NonEmployeeLifecycleManagementApi.PatchSchemaAttribute(context.Background(), attributeId, sourceId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
