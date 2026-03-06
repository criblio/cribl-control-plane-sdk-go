# ModeOptionsHost

Select level of detail for host metrics

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
)

value := components.ModeOptionsHostBasic

// Open enum: custom values can be created with a direct type cast
custom := components.ModeOptionsHost("custom_value")
```


## Values

| Name                      | Value                     |
| ------------------------- | ------------------------- |
| `ModeOptionsHostBasic`    | basic                     |
| `ModeOptionsHostAll`      | all                       |
| `ModeOptionsHostCustom`   | custom                    |
| `ModeOptionsHostDisabled` | disabled                  |