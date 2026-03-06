# RecordDataFormat

Format of data inside the Kinesis Stream records. Gzip compression is automatically detected.

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
)

value := components.RecordDataFormatCribl

// Open enum: custom values can be created with a direct type cast
custom := components.RecordDataFormat("custom_value")
```


## Values

| Name                         | Value                        |
| ---------------------------- | ---------------------------- |
| `RecordDataFormatCribl`      | cribl                        |
| `RecordDataFormatNdjson`     | ndjson                       |
| `RecordDataFormatCloudwatch` | cloudwatch                   |
| `RecordDataFormatLine`       | line                         |