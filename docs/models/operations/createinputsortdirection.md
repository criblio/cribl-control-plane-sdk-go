# CreateInputSortDirection

Used only when Sort by field is set.

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/operations"
)

value := operations.CreateInputSortDirectionAsc

// Open enum: custom values can be created with a direct type cast
custom := operations.CreateInputSortDirection("custom_value")
```


## Values

| Name                           | Value                          |
| ------------------------------ | ------------------------------ |
| `CreateInputSortDirectionAsc`  | asc                            |
| `CreateInputSortDirectionDesc` | desc                           |