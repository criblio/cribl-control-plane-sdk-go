# CreateOutputMappingType

How event fields are mapped to columns.

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/operations"
)

value := operations.CreateOutputMappingTypeAutomatic

// Open enum: custom values can be created with a direct type cast
custom := operations.CreateOutputMappingType("custom_value")
```


## Values

| Name                               | Value                              |
| ---------------------------------- | ---------------------------------- |
| `CreateOutputMappingTypeAutomatic` | automatic                          |
| `CreateOutputMappingTypeCustom`    | custom                             |