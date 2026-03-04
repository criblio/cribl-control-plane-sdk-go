# CompressionOptions3

Codec to use to compress the data before sending to Kafka

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
)

value := components.CompressionOptions3None

// Open enum: custom values can be created with a direct type cast
custom := components.CompressionOptions3("custom_value")
```


## Values

| Name                        | Value                       |
| --------------------------- | --------------------------- |
| `CompressionOptions3None`   | none                        |
| `CompressionOptions3Gzip`   | gzip                        |
| `CompressionOptions3Snappy` | snappy                      |
| `CompressionOptions3Lz4`    | lz4                         |
| `CompressionOptions3Zstd`   | zstd                        |