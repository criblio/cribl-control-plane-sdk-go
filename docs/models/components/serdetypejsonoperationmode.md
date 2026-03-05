# SerdeTypeJSONOperationMode

Extract creates new fields. Reserialize extracts and filters fields, and then reserializes.

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
)

value := components.SerdeTypeJSONOperationModeExtract

// Open enum: custom values can be created with a direct type cast
custom := components.SerdeTypeJSONOperationMode("custom_value")
```


## Values

| Name                                    | Value                                   |
| --------------------------------------- | --------------------------------------- |
| `SerdeTypeJSONOperationModeExtract`     | extract                                 |
| `SerdeTypeJSONOperationModeReserialize` | reserialize                             |