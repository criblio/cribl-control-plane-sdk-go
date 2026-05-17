# MappingTypeOptions

How event fields are mapped to ClickHouse columns

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
)

value := components.MappingTypeOptionsAutomatic

// Open enum: custom values can be created with a direct type cast
custom := components.MappingTypeOptions("custom_value")
```


## Values

| Name                          | Value                         |
| ----------------------------- | ----------------------------- |
| `MappingTypeOptionsAutomatic` | automatic                     |
| `MappingTypeOptionsCustom`    | custom                        |