# RulesetType

Type of ruleset to apply: dataset or datatype.

## Example Usage

```go
import (
	"github.com/Cribl-Community/cribl-control-plane-sdk-go/models/components"
)

value := components.RulesetTypeDataset

// Open enum: custom values can be created with a direct type cast
custom := components.RulesetType("custom_value")
```


## Values

| Name                  | Value                 |
| --------------------- | --------------------- |
| `RulesetTypeDataset`  | dataset               |
| `RulesetTypeDatatype` | datatype              |