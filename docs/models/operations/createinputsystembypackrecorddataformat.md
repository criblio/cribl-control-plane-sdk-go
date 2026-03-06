# CreateInputSystemByPackRecordDataFormat

Format of data inside the Kinesis Stream records. Gzip compression is automatically detected.

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/operations"
)

value := operations.CreateInputSystemByPackRecordDataFormatCribl

// Open enum: custom values can be created with a direct type cast
custom := operations.CreateInputSystemByPackRecordDataFormat("custom_value")
```


## Values

| Name                                                | Value                                               |
| --------------------------------------------------- | --------------------------------------------------- |
| `CreateInputSystemByPackRecordDataFormatCribl`      | cribl                                               |
| `CreateInputSystemByPackRecordDataFormatNdjson`     | ndjson                                              |
| `CreateInputSystemByPackRecordDataFormatCloudwatch` | cloudwatch                                          |
| `CreateInputSystemByPackRecordDataFormatLine`       | line                                                |