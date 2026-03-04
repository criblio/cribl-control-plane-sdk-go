# CompressionOptions

Controls whether the sender should send compressed data to the server. Select 'Disabled' to reject compressed connections or 'Always' to ignore server's configuration and send compressed data.

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
)

value := components.CompressionOptionsDisabled

// Open enum: custom values can be created with a direct type cast
custom := components.CompressionOptions("custom_value")
```


## Values

| Name                         | Value                        |
| ---------------------------- | ---------------------------- |
| `CompressionOptionsDisabled` | disabled                     |
| `CompressionOptionsAuto`     | auto                         |
| `CompressionOptionsAlways`   | always                       |