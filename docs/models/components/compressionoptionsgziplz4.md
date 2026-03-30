# CompressionOptionsGzipLz4

Codec to use to compress the data before sending to Kafka

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
)

value := components.CompressionOptionsGzipLz4None

// Open enum: custom values can be created with a direct type cast
custom := components.CompressionOptionsGzipLz4("custom_value")
```


## Values

| Name                              | Value                             |
| --------------------------------- | --------------------------------- |
| `CompressionOptionsGzipLz4None`   | none                              |
| `CompressionOptionsGzipLz4Gzip`   | gzip                              |
| `CompressionOptionsGzipLz4Snappy` | snappy                            |
| `CompressionOptionsGzipLz4Lz4`    | lz4                               |
| `CompressionOptionsGzipLz4Zstd`   | zstd                              |