# CreateOutputCompression

Compression type to use for records

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/operations"
)

value := operations.CreateOutputCompressionNone

// Open enum: custom values can be created with a direct type cast
custom := operations.CreateOutputCompression("custom_value")
```


## Values

| Name                          | Value                         |
| ----------------------------- | ----------------------------- |
| `CreateOutputCompressionNone` | none                          |
| `CreateOutputCompressionGzip` | gzip                          |