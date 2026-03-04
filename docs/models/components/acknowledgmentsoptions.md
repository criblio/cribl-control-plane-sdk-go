# AcknowledgmentsOptions

Control the number of required acknowledgments

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
)

value := components.AcknowledgmentsOptionsLeader

// Open enum: custom values can be created with a direct type cast
custom := components.AcknowledgmentsOptions(999)
```


## Values

| Name                           | Value                          |
| ------------------------------ | ------------------------------ |
| `AcknowledgmentsOptionsLeader` | 1                              |
| `AcknowledgmentsOptionsNone`   | 0                              |
| `AcknowledgmentsOptionsAll`    | -1                             |