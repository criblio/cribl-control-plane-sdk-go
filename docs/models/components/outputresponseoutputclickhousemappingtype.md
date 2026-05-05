# OutputResponseOutputClickHouseMappingType

How event fields are mapped to ClickHouse columns.

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
)

value := components.OutputResponseOutputClickHouseMappingTypeAutomatic

// Open enum: custom values can be created with a direct type cast
custom := components.OutputResponseOutputClickHouseMappingType("custom_value")
```


## Values

| Name                                                 | Value                                                |
| ---------------------------------------------------- | ---------------------------------------------------- |
| `OutputResponseOutputClickHouseMappingTypeAutomatic` | automatic                                            |
| `OutputResponseOutputClickHouseMappingTypeCustom`    | custom                                               |