# CreateOutputEndpointGoogleCloudObservability

Fixed Google Cloud Observability gRPC endpoint. All three signals share this transport; the OTLP service path determines whether the call lands on traces, metrics, or logs.

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/operations"
)

value := operations.CreateOutputEndpointGoogleCloudObservabilityTelemetryGoogleapisCom443

// Open enum: custom values can be created with a direct type cast
custom := operations.CreateOutputEndpointGoogleCloudObservability("custom_value")
```


## Values

| Name                                                                    | Value                                                                   |
| ----------------------------------------------------------------------- | ----------------------------------------------------------------------- |
| `CreateOutputEndpointGoogleCloudObservabilityTelemetryGoogleapisCom443` | telemetry.googleapis.com:443                                            |