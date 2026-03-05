# OutputSumoLogicDataFormat

Preserve the raw event format instead of JSONifying it

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
)

value := components.OutputSumoLogicDataFormatJSON

// Open enum: custom values can be created with a direct type cast
custom := components.OutputSumoLogicDataFormat("custom_value")
```


## Values

| Name                            | Value                           |
| ------------------------------- | ------------------------------- |
| `OutputSumoLogicDataFormatJSON` | json                            |
| `OutputSumoLogicDataFormatRaw`  | raw                             |