# OutputResponseOutputLocalSearchStorageFormat

Data format to use when sending data. Defaults to JSON Compact.

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
)

value := components.OutputResponseOutputLocalSearchStorageFormatJSONCompactEachRowWithNames

// Open enum: custom values can be created with a direct type cast
custom := components.OutputResponseOutputLocalSearchStorageFormat("custom_value")
```


## Values

| Name                                                                      | Value                                                                     |
| ------------------------------------------------------------------------- | ------------------------------------------------------------------------- |
| `OutputResponseOutputLocalSearchStorageFormatJSONCompactEachRowWithNames` | json-compact-each-row-with-names                                          |
| `OutputResponseOutputLocalSearchStorageFormatJSONEachRow`                 | json-each-row                                                             |