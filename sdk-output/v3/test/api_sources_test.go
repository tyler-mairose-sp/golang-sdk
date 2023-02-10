/*
IdentityNow V3 API

Testing SourcesApiService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package v3

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func Test_v3_SourcesApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test SourcesApiService CreateProvisioningPolicy", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var sourceId string

		resp, httpRes, err := apiClient.SourcesApi.CreateProvisioningPolicy(context.Background(), sourceId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SourcesApiService CreateSource", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.SourcesApi.CreateSource(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SourcesApiService CreateSourceSchema", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var sourceId string

		resp, httpRes, err := apiClient.SourcesApi.CreateSourceSchema(context.Background(), sourceId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SourcesApiService DeleteProvisioningPolicy", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var sourceId string
		var usageType UsageType

		resp, httpRes, err := apiClient.SourcesApi.DeleteProvisioningPolicy(context.Background(), sourceId, usageType).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SourcesApiService DeleteSource", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		resp, httpRes, err := apiClient.SourcesApi.DeleteSource(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SourcesApiService DeleteSourceSchema", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var sourceId string
		var schemaId string

		resp, httpRes, err := apiClient.SourcesApi.DeleteSourceSchema(context.Background(), sourceId, schemaId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SourcesApiService DownloadSourceAccountsSchema", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		resp, httpRes, err := apiClient.SourcesApi.DownloadSourceAccountsSchema(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SourcesApiService DownloadSourceEntitlementsSchema", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		resp, httpRes, err := apiClient.SourcesApi.DownloadSourceEntitlementsSchema(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SourcesApiService GetProvisioningPolicy", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var sourceId string
		var usageType UsageType

		resp, httpRes, err := apiClient.SourcesApi.GetProvisioningPolicy(context.Background(), sourceId, usageType).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SourcesApiService GetSource", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		resp, httpRes, err := apiClient.SourcesApi.GetSource(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SourcesApiService GetSourceHealth", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var sourceId string

		resp, httpRes, err := apiClient.SourcesApi.GetSourceHealth(context.Background(), sourceId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SourcesApiService GetSourceSchema", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var sourceId string
		var schemaId string

		resp, httpRes, err := apiClient.SourcesApi.GetSourceSchema(context.Background(), sourceId, schemaId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SourcesApiService ListProvisioningPolicies", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var sourceId string

		resp, httpRes, err := apiClient.SourcesApi.ListProvisioningPolicies(context.Background(), sourceId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SourcesApiService ListSourceSchemas", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var sourceId string

		resp, httpRes, err := apiClient.SourcesApi.ListSourceSchemas(context.Background(), sourceId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SourcesApiService ListSources", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.SourcesApi.ListSources(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SourcesApiService PutProvisioningPolicy", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var sourceId string
		var usageType UsageType

		resp, httpRes, err := apiClient.SourcesApi.PutProvisioningPolicy(context.Background(), sourceId, usageType).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SourcesApiService PutSource", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		resp, httpRes, err := apiClient.SourcesApi.PutSource(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SourcesApiService PutSourceSchema", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var sourceId string
		var schemaId string

		resp, httpRes, err := apiClient.SourcesApi.PutSourceSchema(context.Background(), sourceId, schemaId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SourcesApiService UpdateProvisioningPoliciesInBulk", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var sourceId string

		resp, httpRes, err := apiClient.SourcesApi.UpdateProvisioningPoliciesInBulk(context.Background(), sourceId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SourcesApiService UpdateProvisioningPolicy", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var sourceId string
		var usageType UsageType

		resp, httpRes, err := apiClient.SourcesApi.UpdateProvisioningPolicy(context.Background(), sourceId, usageType).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SourcesApiService UpdateSource", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		resp, httpRes, err := apiClient.SourcesApi.UpdateSource(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SourcesApiService UpdateSourceSchema", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var sourceId string
		var schemaId string

		resp, httpRes, err := apiClient.SourcesApi.UpdateSourceSchema(context.Background(), sourceId, schemaId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SourcesApiService UploadSourceAccountsSchema", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		resp, httpRes, err := apiClient.SourcesApi.UploadSourceAccountsSchema(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SourcesApiService UploadSourceConnectorFile", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var sourceId string

		resp, httpRes, err := apiClient.SourcesApi.UploadSourceConnectorFile(context.Background(), sourceId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SourcesApiService UploadSourceEntitlementsSchema", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		resp, httpRes, err := apiClient.SourcesApi.UploadSourceEntitlementsSchema(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
