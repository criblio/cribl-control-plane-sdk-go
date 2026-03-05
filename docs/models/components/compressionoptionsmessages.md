# CompressionOptionsMessages

Type of compression to apply to messages sent to the OpenTelemetry endpoint

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
)

value := components.CompressionOptionsMessagesNone

// Open enum: custom values can be created with a direct type cast
custom := components.CompressionOptionsMessages("custom_value")
```


## Values

| Name                             | Value                            |
| -------------------------------- | -------------------------------- |
| `CompressionOptionsMessagesNone` | none                             |
| `CompressionOptionsMessagesGzip` | gzip                             |