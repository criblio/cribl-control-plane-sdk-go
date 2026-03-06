# CreateOutputSystemByPackFormatLocalSearchStorage

Data format to use when sending data. Defaults to JSON Compact.

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/operations"
)

value := operations.CreateOutputSystemByPackFormatLocalSearchStorageJSONCompactEachRowWithNames

// Open enum: custom values can be created with a direct type cast
custom := operations.CreateOutputSystemByPackFormatLocalSearchStorage("custom_value")
```


## Values

| Name                                                                          | Value                                                                         |
| ----------------------------------------------------------------------------- | ----------------------------------------------------------------------------- |
| `CreateOutputSystemByPackFormatLocalSearchStorageJSONCompactEachRowWithNames` | json-compact-each-row-with-names                                              |
| `CreateOutputSystemByPackFormatLocalSearchStorageJSONEachRow`                 | json-each-row                                                                 |