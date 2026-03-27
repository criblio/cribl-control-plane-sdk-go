# MethodOptions

The method to use when sending events

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
)

value := components.MethodOptionsPost

// Open enum: custom values can be created with a direct type cast
custom := components.MethodOptions("custom_value")
```


## Values

| Name                 | Value                |
| -------------------- | -------------------- |
| `MethodOptionsPost`  | POST                 |
| `MethodOptionsPut`   | PUT                  |
| `MethodOptionsPatch` | PATCH                |