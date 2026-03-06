# InputSplunkCompression

Controls whether to support reading compressed data from a forwarder. Select 'Automatic' to match the forwarder's configuration, or 'Disabled' to reject compressed connections.

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
)

value := components.InputSplunkCompressionDisabled

// Open enum: custom values can be created with a direct type cast
custom := components.InputSplunkCompression("custom_value")
```


## Values

| Name                             | Value                            |
| -------------------------------- | -------------------------------- |
| `InputSplunkCompressionDisabled` | disabled                         |
| `InputSplunkCompressionAuto`     | auto                             |
| `InputSplunkCompressionAlways`   | always                           |