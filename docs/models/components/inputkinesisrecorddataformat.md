# InputKinesisRecordDataFormat

Format of data inside the Kinesis Stream records. Gzip compression is automatically detected.

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
)

value := components.InputKinesisRecordDataFormatCribl

// Open enum: custom values can be created with a direct type cast
custom := components.InputKinesisRecordDataFormat("custom_value")
```


## Values

| Name                                     | Value                                    |
| ---------------------------------------- | ---------------------------------------- |
| `InputKinesisRecordDataFormatCribl`      | cribl                                    |
| `InputKinesisRecordDataFormatNdjson`     | ndjson                                   |
| `InputKinesisRecordDataFormatCloudwatch` | cloudwatch                               |
| `InputKinesisRecordDataFormatLine`       | line                                     |