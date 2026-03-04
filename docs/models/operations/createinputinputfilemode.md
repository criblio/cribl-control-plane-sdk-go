# CreateInputInputFileMode

Choose how to discover files to monitor

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/operations"
)

value := operations.CreateInputInputFileModeManual

// Open enum: custom values can be created with a direct type cast
custom := operations.CreateInputInputFileMode("custom_value")
```


## Values

| Name                             | Value                            |
| -------------------------------- | -------------------------------- |
| `CreateInputInputFileModeManual` | manual                           |
| `CreateInputInputFileModeAuto`   | auto                             |