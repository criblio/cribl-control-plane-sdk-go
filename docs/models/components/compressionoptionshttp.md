# CompressionOptionsHTTP

Data compression format to apply to HTTP content before it is delivered

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
)

value := components.CompressionOptionsHTTPNone

// Open enum: custom values can be created with a direct type cast
custom := components.CompressionOptionsHTTP("custom_value")
```


## Values

| Name                         | Value                        |
| ---------------------------- | ---------------------------- |
| `CompressionOptionsHTTPNone` | none                         |
| `CompressionOptionsHTTPGzip` | gzip                         |