# MappingTypeLocalSearchStorage

How event fields are mapped to columns.

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
)

value := components.MappingTypeLocalSearchStorageAutomatic

// Open enum: custom values can be created with a direct type cast
custom := components.MappingTypeLocalSearchStorage("custom_value")
```


## Values

| Name                                     | Value                                    |
| ---------------------------------------- | ---------------------------------------- |
| `MappingTypeLocalSearchStorageAutomatic` | automatic                                |
| `MappingTypeLocalSearchStorageCustom`    | custom                                   |