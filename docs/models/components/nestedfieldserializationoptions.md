# NestedFieldSerializationOptions

How to serialize nested fields into index-time fields

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
)

value := components.NestedFieldSerializationOptionsJSON

// Open enum: custom values can be created with a direct type cast
custom := components.NestedFieldSerializationOptions("custom_value")
```


## Values

| Name                                  | Value                                 |
| ------------------------------------- | ------------------------------------- |
| `NestedFieldSerializationOptionsJSON` | json                                  |
| `NestedFieldSerializationOptionsNone` | none                                  |