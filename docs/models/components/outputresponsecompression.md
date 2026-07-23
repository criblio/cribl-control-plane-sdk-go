# OutputResponseCompression

Compression type to use for records

## Example Usage

```go
import (
	"github.com/Cribl-Community/cribl-control-plane-sdk-go/models/components"
)

value := components.OutputResponseCompressionNone

// Open enum: custom values can be created with a direct type cast
custom := components.OutputResponseCompression("custom_value")
```


## Values

| Name                            | Value                           |
| ------------------------------- | ------------------------------- |
| `OutputResponseCompressionNone` | none                            |
| `OutputResponseCompressionGzip` | gzip                            |