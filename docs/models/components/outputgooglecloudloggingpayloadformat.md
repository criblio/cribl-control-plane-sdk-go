# OutputGoogleCloudLoggingPayloadFormat

Format to use when sending payload. Defaults to Text.

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
)

value := components.OutputGoogleCloudLoggingPayloadFormatText

// Open enum: custom values can be created with a direct type cast
custom := components.OutputGoogleCloudLoggingPayloadFormat("custom_value")
```


## Values

| Name                                        | Value                                       |
| ------------------------------------------- | ------------------------------------------- |
| `OutputGoogleCloudLoggingPayloadFormatText` | text                                        |
| `OutputGoogleCloudLoggingPayloadFormatJSON` | json                                        |