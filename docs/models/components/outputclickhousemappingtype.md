# OutputClickHouseMappingType

How event fields are mapped to ClickHouse columns.

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
)

value := components.OutputClickHouseMappingTypeAutomatic

// Open enum: custom values can be created with a direct type cast
custom := components.OutputClickHouseMappingType("custom_value")
```


## Values

| Name                                   | Value                                  |
| -------------------------------------- | -------------------------------------- |
| `OutputClickHouseMappingTypeAutomatic` | automatic                              |
| `OutputClickHouseMappingTypeCustom`    | custom                                 |