# CompressionOptions2

Data compression format to apply to HTTP content before it is delivered

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
)

value := components.CompressionOptions2None

// Open enum: custom values can be created with a direct type cast
custom := components.CompressionOptions2("custom_value")
```


## Values

| Name                      | Value                     |
| ------------------------- | ------------------------- |
| `CompressionOptions2None` | none                      |
| `CompressionOptions2Gzip` | gzip                      |