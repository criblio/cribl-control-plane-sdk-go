# CompressionOptionsPq

Codec to use to compress the persisted data

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
)

value := components.CompressionOptionsPqNone

// Open enum: custom values can be created with a direct type cast
custom := components.CompressionOptionsPq("custom_value")
```


## Values

| Name                       | Value                      |
| -------------------------- | -------------------------- |
| `CompressionOptionsPqNone` | none                       |
| `CompressionOptionsPqGzip` | gzip                       |