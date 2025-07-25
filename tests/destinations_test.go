// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package tests

import (
	"context"
	criblcontrolplanesdkgo "github.com/criblio/cribl-control-plane-sdk-go"
	"github.com/criblio/cribl-control-plane-sdk-go/internal/utils"
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
	"github.com/criblio/cribl-control-plane-sdk-go/models/operations"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestDestinations_ListOutput(t *testing.T) {
	ctx := context.Background()

	testHTTPClient := createTestHTTPClient("listOutput")

	s := criblcontrolplanesdkgo.New(
		utils.GetEnv("TEST_SERVER_URL", "http://localhost:18080"),
		criblcontrolplanesdkgo.WithClient(testHTTPClient),
		criblcontrolplanesdkgo.WithSecurity(components.Security{
			BearerAuth: criblcontrolplanesdkgo.String(utils.GetEnv("CRIBLCONTROLPLANE_BEARER_AUTH", "value")),
		}),
	)

	res, err := s.Destinations.ListDestination(ctx)
	require.NoError(t, err)
	assert.Equal(t, 200, res.HTTPMeta.Response.StatusCode)
	assert.NotNil(t, res.Object)
	assert.Equal(t, &operations.ListOutputResponseBody{}, res.Object)

}

func TestDestinations_CreateOutput(t *testing.T) {
	ctx := context.Background()

	testHTTPClient := createTestHTTPClient("createOutput")

	s := criblcontrolplanesdkgo.New(
		utils.GetEnv("TEST_SERVER_URL", "http://localhost:18080"),
		criblcontrolplanesdkgo.WithClient(testHTTPClient),
		criblcontrolplanesdkgo.WithSecurity(components.Security{
			BearerAuth: criblcontrolplanesdkgo.String(utils.GetEnv("CRIBLCONTROLPLANE_BEARER_AUTH", "value")),
		}),
	)

	res, err := s.Destinations.CreateDestination(ctx, operations.CreateCreateOutputRequestOutputElasticCloud(
		operations.OutputElasticCloud{
			ID:    "<id>",
			URL:   "https://probable-rationale.com/",
			Index: "<value>",
		},
	))
	require.NoError(t, err)
	assert.Equal(t, 200, res.HTTPMeta.Response.StatusCode)
	assert.NotNil(t, res.Object)
	assert.Equal(t, &operations.CreateOutputResponseBody{}, res.Object)

}

func TestDestinations_GetOutputByID(t *testing.T) {
	ctx := context.Background()

	testHTTPClient := createTestHTTPClient("getOutputById")

	s := criblcontrolplanesdkgo.New(
		utils.GetEnv("TEST_SERVER_URL", "http://localhost:18080"),
		criblcontrolplanesdkgo.WithClient(testHTTPClient),
		criblcontrolplanesdkgo.WithSecurity(components.Security{
			BearerAuth: criblcontrolplanesdkgo.String(utils.GetEnv("CRIBLCONTROLPLANE_BEARER_AUTH", "value")),
		}),
	)

	res, err := s.Destinations.GetDestinationByID(ctx, "<id>")
	require.NoError(t, err)
	assert.Equal(t, 200, res.HTTPMeta.Response.StatusCode)
	assert.NotNil(t, res.Object)
	assert.Equal(t, &operations.GetOutputByIDResponseBody{}, res.Object)

}

func TestDestinations_UpdateOutputByID(t *testing.T) {
	ctx := context.Background()

	testHTTPClient := createTestHTTPClient("updateOutputById")

	s := criblcontrolplanesdkgo.New(
		utils.GetEnv("TEST_SERVER_URL", "http://localhost:18080"),
		criblcontrolplanesdkgo.WithClient(testHTTPClient),
		criblcontrolplanesdkgo.WithSecurity(components.Security{
			BearerAuth: criblcontrolplanesdkgo.String(utils.GetEnv("CRIBLCONTROLPLANE_BEARER_AUTH", "value")),
		}),
	)

	res, err := s.Destinations.UpdateDestinationByID(ctx, "<id>", components.CreateOutputOutputSignalfx(
		components.OutputSignalfx{
			Type: components.OutputSignalfxTypeSignalfx,
		},
	))
	require.NoError(t, err)
	assert.Equal(t, 200, res.HTTPMeta.Response.StatusCode)
	assert.NotNil(t, res.Object)
	assert.Equal(t, &operations.UpdateOutputByIDResponseBody{}, res.Object)

}

func TestDestinations_DeleteOutputByID(t *testing.T) {
	ctx := context.Background()

	testHTTPClient := createTestHTTPClient("deleteOutputById")

	s := criblcontrolplanesdkgo.New(
		utils.GetEnv("TEST_SERVER_URL", "http://localhost:18080"),
		criblcontrolplanesdkgo.WithClient(testHTTPClient),
		criblcontrolplanesdkgo.WithSecurity(components.Security{
			BearerAuth: criblcontrolplanesdkgo.String(utils.GetEnv("CRIBLCONTROLPLANE_BEARER_AUTH", "value")),
		}),
	)

	res, err := s.Destinations.DeleteDestinationByID(ctx, "<id>")
	require.NoError(t, err)
	assert.Equal(t, 200, res.HTTPMeta.Response.StatusCode)
	assert.NotNil(t, res.Object)
	assert.Equal(t, &operations.DeleteOutputByIDResponseBody{}, res.Object)

}

func TestDestinations_DeleteOutputPqByID(t *testing.T) {
	ctx := context.Background()

	testHTTPClient := createTestHTTPClient("deleteOutputPqById")

	s := criblcontrolplanesdkgo.New(
		utils.GetEnv("TEST_SERVER_URL", "http://localhost:18080"),
		criblcontrolplanesdkgo.WithClient(testHTTPClient),
		criblcontrolplanesdkgo.WithSecurity(components.Security{
			BearerAuth: criblcontrolplanesdkgo.String(utils.GetEnv("CRIBLCONTROLPLANE_BEARER_AUTH", "value")),
		}),
	)

	res, err := s.Destinations.DeleteDestinationPqByID(ctx, "<id>")
	require.NoError(t, err)
	assert.Equal(t, 200, res.HTTPMeta.Response.StatusCode)
	assert.NotNil(t, res.Object)
	assert.Equal(t, &operations.DeleteOutputPqByIDResponseBody{}, res.Object)

}

func TestDestinations_GetOutputPqByID(t *testing.T) {
	ctx := context.Background()

	testHTTPClient := createTestHTTPClient("getOutputPqById")

	s := criblcontrolplanesdkgo.New(
		utils.GetEnv("TEST_SERVER_URL", "http://localhost:18080"),
		criblcontrolplanesdkgo.WithClient(testHTTPClient),
		criblcontrolplanesdkgo.WithSecurity(components.Security{
			BearerAuth: criblcontrolplanesdkgo.String(utils.GetEnv("CRIBLCONTROLPLANE_BEARER_AUTH", "value")),
		}),
	)

	res, err := s.Destinations.GetDestinationPqByID(ctx, "<id>")
	require.NoError(t, err)
	assert.Equal(t, 200, res.HTTPMeta.Response.StatusCode)
	assert.NotNil(t, res.Object)
	assert.Equal(t, &operations.GetOutputPqByIDResponseBody{}, res.Object)

}

func TestDestinations_GetOutputSamplesByID(t *testing.T) {
	ctx := context.Background()

	testHTTPClient := createTestHTTPClient("getOutputSamplesById")

	s := criblcontrolplanesdkgo.New(
		utils.GetEnv("TEST_SERVER_URL", "http://localhost:18080"),
		criblcontrolplanesdkgo.WithClient(testHTTPClient),
		criblcontrolplanesdkgo.WithSecurity(components.Security{
			BearerAuth: criblcontrolplanesdkgo.String(utils.GetEnv("CRIBLCONTROLPLANE_BEARER_AUTH", "value")),
		}),
	)

	res, err := s.Destinations.GetDestinationSamplesByID(ctx, "<id>")
	require.NoError(t, err)
	assert.Equal(t, 200, res.HTTPMeta.Response.StatusCode)
	assert.NotNil(t, res.Object)
	assert.Equal(t, &operations.GetOutputSamplesByIDResponseBody{}, res.Object)

}

func TestDestinations_CreateOutputTestByID(t *testing.T) {
	ctx := context.Background()

	testHTTPClient := createTestHTTPClient("createOutputTestById")

	s := criblcontrolplanesdkgo.New(
		utils.GetEnv("TEST_SERVER_URL", "http://localhost:18080"),
		criblcontrolplanesdkgo.WithClient(testHTTPClient),
		criblcontrolplanesdkgo.WithSecurity(components.Security{
			BearerAuth: criblcontrolplanesdkgo.String(utils.GetEnv("CRIBLCONTROLPLANE_BEARER_AUTH", "value")),
		}),
	)

	res, err := s.Destinations.CreateDestinationTestByID(ctx, "<id>", components.OutputTestRequest{
		Events: []components.CriblEvent{
			components.CriblEvent{
				Raw: "<value>",
			},
		},
	})
	require.NoError(t, err)
	assert.Equal(t, 200, res.HTTPMeta.Response.StatusCode)
	assert.NotNil(t, res.Object)
	assert.Equal(t, &operations.CreateOutputTestByIDResponseBody{}, res.Object)

}
