# SerdeTypeAutoOperationMode

Extract creates new fields. Reserialize extracts and filters fields, and then reserializes.

## Example Usage

```go
import (
	"github.com/Cribl-Community/cribl-control-plane-sdk-go/models/components"
)

value := components.SerdeTypeAutoOperationModeExtract

// Open enum: custom values can be created with a direct type cast
custom := components.SerdeTypeAutoOperationMode("custom_value")
```


## Values

| Name                                    | Value                                   |
| --------------------------------------- | --------------------------------------- |
| `SerdeTypeAutoOperationModeExtract`     | extract                                 |
| `SerdeTypeAutoOperationModeReserialize` | reserialize                             |