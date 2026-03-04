# CreateOutputDataFormatSumoLogic

Preserve the raw event format instead of JSONifying it

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/operations"
)

value := operations.CreateOutputDataFormatSumoLogicJSON

// Open enum: custom values can be created with a direct type cast
custom := operations.CreateOutputDataFormatSumoLogic("custom_value")
```


## Values

| Name                                  | Value                                 |
| ------------------------------------- | ------------------------------------- |
| `CreateOutputDataFormatSumoLogicJSON` | json                                  |
| `CreateOutputDataFormatSumoLogicRaw`  | raw                                   |