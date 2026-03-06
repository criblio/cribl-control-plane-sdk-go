# CompressionOptions4

Type of compression to apply to messages sent to the OpenTelemetry endpoint

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
)

value := components.CompressionOptions4None

// Open enum: custom values can be created with a direct type cast
custom := components.CompressionOptions4("custom_value")
```


## Values

| Name                         | Value                        |
| ---------------------------- | ---------------------------- |
| `CompressionOptions4None`    | none                         |
| `CompressionOptions4Deflate` | deflate                      |
| `CompressionOptions4Gzip`    | gzip                         |