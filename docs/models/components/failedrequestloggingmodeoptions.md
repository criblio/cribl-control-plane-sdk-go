# FailedRequestLoggingModeOptions

Data to log when a request fails. All headers are redacted by default, unless listed as safe headers below.

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
)

value := components.FailedRequestLoggingModeOptionsPayload

// Open enum: custom values can be created with a direct type cast
custom := components.FailedRequestLoggingModeOptions("custom_value")
```


## Values

| Name                                               | Value                                              |
| -------------------------------------------------- | -------------------------------------------------- |
| `FailedRequestLoggingModeOptionsPayload`           | payload                                            |
| `FailedRequestLoggingModeOptionsPayloadAndHeaders` | payloadAndHeaders                                  |
| `FailedRequestLoggingModeOptionsNone`              | none                                               |