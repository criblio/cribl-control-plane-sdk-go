# OutputLocalSearchStorageFormat

Data format to use when sending data. Defaults to JSON Compact.

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
)

value := components.OutputLocalSearchStorageFormatJSONCompactEachRowWithNames

// Open enum: custom values can be created with a direct type cast
custom := components.OutputLocalSearchStorageFormat("custom_value")
```


## Values

| Name                                                        | Value                                                       |
| ----------------------------------------------------------- | ----------------------------------------------------------- |
| `OutputLocalSearchStorageFormatJSONCompactEachRowWithNames` | json-compact-each-row-with-names                            |
| `OutputLocalSearchStorageFormatJSONEachRow`                 | json-each-row                                               |