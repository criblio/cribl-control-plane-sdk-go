# RecordDataFormatOptionsJSONProtobuf

Format to use to serialize events before writing to Kafka.

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
)

value := components.RecordDataFormatOptionsJSONProtobufJSON

// Open enum: custom values can be created with a direct type cast
custom := components.RecordDataFormatOptionsJSONProtobuf("custom_value")
```


## Values

| Name                                          | Value                                         |
| --------------------------------------------- | --------------------------------------------- |
| `RecordDataFormatOptionsJSONProtobufJSON`     | json                                          |
| `RecordDataFormatOptionsJSONProtobufRaw`      | raw                                           |
| `RecordDataFormatOptionsJSONProtobufProtobuf` | protobuf                                      |