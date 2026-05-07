# DataFormatSumoLogic

Preserve the raw event format instead of JSONifying it

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
)

value := components.DataFormatSumoLogicJSON

// Open enum: custom values can be created with a direct type cast
custom := components.DataFormatSumoLogic("custom_value")
```


## Values

| Name                      | Value                     |
| ------------------------- | ------------------------- |
| `DataFormatSumoLogicJSON` | json                      |
| `DataFormatSumoLogicRaw`  | raw                       |