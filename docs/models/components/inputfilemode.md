# InputFileMode

Choose how to discover files to monitor

## Example Usage

```go
import (
	"github.com/Cribl-Community/cribl-control-plane-sdk-go/models/components"
)

value := components.InputFileModeManual

// Open enum: custom values can be created with a direct type cast
custom := components.InputFileMode("custom_value")
```


## Values

| Name                  | Value                 |
| --------------------- | --------------------- |
| `InputFileModeManual` | manual                |
| `InputFileModeAuto`   | auto                  |