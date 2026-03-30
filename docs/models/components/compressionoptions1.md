# CompressionOptions1

Codec to use to compress the data before sending

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
)

value := components.CompressionOptions1None

// Open enum: custom values can be created with a direct type cast
custom := components.CompressionOptions1("custom_value")
```


## Values

| Name                      | Value                     |
| ------------------------- | ------------------------- |
| `CompressionOptions1None` | none                      |
| `CompressionOptions1Gzip` | gzip                      |