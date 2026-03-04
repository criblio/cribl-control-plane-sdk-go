# UDMType

Defines the specific format for UDM events sent to Google SecOps. This must match the type of UDM data being sent.

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
)

value := components.UDMTypeEntities

// Open enum: custom values can be created with a direct type cast
custom := components.UDMType("custom_value")
```


## Values

| Name              | Value             |
| ----------------- | ----------------- |
| `UDMTypeEntities` | entities          |
| `UDMTypeLogs`     | logs              |