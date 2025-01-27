/*
IdentityNow Beta API

Testing RoleInsightsApiService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package beta

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/sailpoint-oss/golang-sdk"
)

func Test_beta_RoleInsightsApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test RoleInsightsApiService CreateRoleInsightRequests", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.RoleInsightsApi.CreateRoleInsightRequests(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RoleInsightsApiService DownloadRoleInsightsEntitlementsChanges", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var insightId string

		resp, httpRes, err := apiClient.RoleInsightsApi.DownloadRoleInsightsEntitlementsChanges(context.Background(), insightId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RoleInsightsApiService GetEntitlementChangesIdentities", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var insightId string
		var entitlementId string

		resp, httpRes, err := apiClient.RoleInsightsApi.GetEntitlementChangesIdentities(context.Background(), insightId, entitlementId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RoleInsightsApiService GetRoleInsight", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var insightId string

		resp, httpRes, err := apiClient.RoleInsightsApi.GetRoleInsight(context.Background(), insightId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RoleInsightsApiService GetRoleInsights", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.RoleInsightsApi.GetRoleInsights(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RoleInsightsApiService GetRoleInsightsCurrentEntitlements", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var insightId string

		resp, httpRes, err := apiClient.RoleInsightsApi.GetRoleInsightsCurrentEntitlements(context.Background(), insightId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RoleInsightsApiService GetRoleInsightsEntitlementsChanges", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var insightId string

		resp, httpRes, err := apiClient.RoleInsightsApi.GetRoleInsightsEntitlementsChanges(context.Background(), insightId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RoleInsightsApiService GetRoleInsightsRequests", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		resp, httpRes, err := apiClient.RoleInsightsApi.GetRoleInsightsRequests(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RoleInsightsApiService GetRoleInsightsSummary", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.RoleInsightsApi.GetRoleInsightsSummary(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
