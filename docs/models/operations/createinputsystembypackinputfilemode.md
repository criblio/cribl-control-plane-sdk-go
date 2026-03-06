# CreateInputSystemByPackInputFileMode

Choose how to discover files to monitor

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/operations"
)

value := operations.CreateInputSystemByPackInputFileModeManual

// Open enum: custom values can be created with a direct type cast
custom := operations.CreateInputSystemByPackInputFileMode("custom_value")
```


## Values

| Name                                         | Value                                        |
| -------------------------------------------- | -------------------------------------------- |
| `CreateInputSystemByPackInputFileModeManual` | manual                                       |
| `CreateInputSystemByPackInputFileModeAuto`   | auto                                         |