# InputResponseRecordDataFormat

Format of data inside the Kinesis Stream records. Gzip compression is automatically detected.

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
)

value := components.InputResponseRecordDataFormatCribl

// Open enum: custom values can be created with a direct type cast
custom := components.InputResponseRecordDataFormat("custom_value")
```


## Values

| Name                                      | Value                                     |
| ----------------------------------------- | ----------------------------------------- |
| `InputResponseRecordDataFormatCribl`      | cribl                                     |
| `InputResponseRecordDataFormatNdjson`     | ndjson                                    |
| `InputResponseRecordDataFormatCloudwatch` | cloudwatch                                |
| `InputResponseRecordDataFormatLine`       | line                                      |