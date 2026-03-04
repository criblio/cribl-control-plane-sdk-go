# MessageFormatOptions

Format to use when sending logs to Loki (Protobuf or JSON)

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
)

value := components.MessageFormatOptionsProtobuf

// Open enum: custom values can be created with a direct type cast
custom := components.MessageFormatOptions("custom_value")
```


## Values

| Name                           | Value                          |
| ------------------------------ | ------------------------------ |
| `MessageFormatOptionsProtobuf` | protobuf                       |
| `MessageFormatOptionsJSON`     | json                           |