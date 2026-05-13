# MappingTypeClickHouse

How event fields are mapped to ClickHouse columns

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
)

value := components.MappingTypeClickHouseAutomatic

// Open enum: custom values can be created with a direct type cast
custom := components.MappingTypeClickHouse("custom_value")
```


## Values

| Name                             | Value                            |
| -------------------------------- | -------------------------------- |
| `MappingTypeClickHouseAutomatic` | automatic                        |
| `MappingTypeClickHouseCustom`    | custom                           |