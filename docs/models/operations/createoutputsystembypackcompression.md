# CreateOutputSystemByPackCompression

Compression type to use for records

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/operations"
)

value := operations.CreateOutputSystemByPackCompressionNone

// Open enum: custom values can be created with a direct type cast
custom := operations.CreateOutputSystemByPackCompression("custom_value")
```


## Values

| Name                                      | Value                                     |
| ----------------------------------------- | ----------------------------------------- |
| `CreateOutputSystemByPackCompressionNone` | none                                      |
| `CreateOutputSystemByPackCompressionGzip` | gzip                                      |