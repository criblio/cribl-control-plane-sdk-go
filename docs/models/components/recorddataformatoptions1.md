# RecordDataFormatOptions1

Format to use to serialize events before writing to Kafka.

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
)

value := components.RecordDataFormatOptions1JSON

// Open enum: custom values can be created with a direct type cast
custom := components.RecordDataFormatOptions1("custom_value")
```


## Values

| Name                               | Value                              |
| ---------------------------------- | ---------------------------------- |
| `RecordDataFormatOptions1JSON`     | json                               |
| `RecordDataFormatOptions1Raw`      | raw                                |
| `RecordDataFormatOptions1Protobuf` | protobuf                           |