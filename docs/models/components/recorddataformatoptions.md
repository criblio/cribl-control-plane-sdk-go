# RecordDataFormatOptions

Format to use to serialize events before writing to the Event Hubs Kafka brokers

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
)

value := components.RecordDataFormatOptionsJSON

// Open enum: custom values can be created with a direct type cast
custom := components.RecordDataFormatOptions("custom_value")
```


## Values

| Name                          | Value                         |
| ----------------------------- | ----------------------------- |
| `RecordDataFormatOptionsJSON` | json                          |
| `RecordDataFormatOptionsRaw`  | raw                           |