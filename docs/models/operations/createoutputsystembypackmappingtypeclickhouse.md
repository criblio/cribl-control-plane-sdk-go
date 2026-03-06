# CreateOutputSystemByPackMappingTypeClickHouse

How event fields are mapped to ClickHouse columns.

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/operations"
)

value := operations.CreateOutputSystemByPackMappingTypeClickHouseAutomatic

// Open enum: custom values can be created with a direct type cast
custom := operations.CreateOutputSystemByPackMappingTypeClickHouse("custom_value")
```


## Values

| Name                                                     | Value                                                    |
| -------------------------------------------------------- | -------------------------------------------------------- |
| `CreateOutputSystemByPackMappingTypeClickHouseAutomatic` | automatic                                                |
| `CreateOutputSystemByPackMappingTypeClickHouseCustom`    | custom                                                   |