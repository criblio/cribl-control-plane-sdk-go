# AcknowledgmentsOptions1

Control the number of required acknowledgments.

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
)

value := components.AcknowledgmentsOptions1Leader

// Open enum: custom values can be created with a direct type cast
custom := components.AcknowledgmentsOptions1(999)
```


## Values

| Name                            | Value                           |
| ------------------------------- | ------------------------------- |
| `AcknowledgmentsOptions1Leader` | 1                               |
| `AcknowledgmentsOptions1None`   | 0                               |
| `AcknowledgmentsOptions1All`    | -1                              |