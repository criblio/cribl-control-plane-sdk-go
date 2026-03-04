# CreateInputSystemByPackCompression

Controls whether to support reading compressed data from a forwarder. Select 'Automatic' to match the forwarder's configuration, or 'Disabled' to reject compressed connections.

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/operations"
)

value := operations.CreateInputSystemByPackCompressionDisabled

// Open enum: custom values can be created with a direct type cast
custom := operations.CreateInputSystemByPackCompression("custom_value")
```


## Values

| Name                                         | Value                                        |
| -------------------------------------------- | -------------------------------------------- |
| `CreateInputSystemByPackCompressionDisabled` | disabled                                     |
| `CreateInputSystemByPackCompressionAuto`     | auto                                         |
| `CreateInputSystemByPackCompressionAlways`   | always                                       |