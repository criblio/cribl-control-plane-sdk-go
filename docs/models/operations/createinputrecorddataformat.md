# CreateInputRecordDataFormat

Format of data inside the Kinesis Stream records. Gzip compression is automatically detected.

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/operations"
)

value := operations.CreateInputRecordDataFormatCribl

// Open enum: custom values can be created with a direct type cast
custom := operations.CreateInputRecordDataFormat("custom_value")
```


## Values

| Name                                    | Value                                   |
| --------------------------------------- | --------------------------------------- |
| `CreateInputRecordDataFormatCribl`      | cribl                                   |
| `CreateInputRecordDataFormatNdjson`     | ndjson                                  |
| `CreateInputRecordDataFormatCloudwatch` | cloudwatch                              |
| `CreateInputRecordDataFormatLine`       | line                                    |