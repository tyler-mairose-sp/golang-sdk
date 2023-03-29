/*
IdentityNow cc (private) APIs

Testing SystemApiService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package cc

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/sailpoint-oss/golang-sdk"
)

func Test_cc_SystemApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test SystemApiService RefreshIdentities", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		httpRes, err := apiClient.SystemApi.RefreshIdentities(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
