# CreateOutputMappingTypeClickHouse

How event fields are mapped to ClickHouse columns.

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/operations"
)

value := operations.CreateOutputMappingTypeClickHouseAutomatic

// Open enum: custom values can be created with a direct type cast
custom := operations.CreateOutputMappingTypeClickHouse("custom_value")
```


## Values

| Name                                         | Value                                        |
| -------------------------------------------- | -------------------------------------------- |
| `CreateOutputMappingTypeClickHouseAutomatic` | automatic                                    |
| `CreateOutputMappingTypeClickHouseCustom`    | custom                                       |