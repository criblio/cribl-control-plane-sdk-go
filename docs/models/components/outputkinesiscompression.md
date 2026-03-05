# OutputKinesisCompression

Compression type to use for records

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
)

value := components.OutputKinesisCompressionNone

// Open enum: custom values can be created with a direct type cast
custom := components.OutputKinesisCompression("custom_value")
```


## Values

| Name                           | Value                          |
| ------------------------------ | ------------------------------ |
| `OutputKinesisCompressionNone` | none                           |
| `OutputKinesisCompressionGzip` | gzip                           |