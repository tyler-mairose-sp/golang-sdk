/*
IdentityNow Beta API

Testing IAICommonAccessApiService

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

func Test_beta_IAICommonAccessApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test IAICommonAccessApiService CreateCommonAccess", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.IAICommonAccessApi.CreateCommonAccess(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test IAICommonAccessApiService GetCommonAccess", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.IAICommonAccessApi.GetCommonAccess(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test IAICommonAccessApiService UpdateCommonAccessStatusInBulk", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.IAICommonAccessApi.UpdateCommonAccessStatusInBulk(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
