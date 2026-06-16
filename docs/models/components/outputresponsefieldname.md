# OutputResponseFieldName

Name of the metadata field.

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
)

value := components.OutputResponseFieldNameService

// Open enum: custom values can be created with a direct type cast
custom := components.OutputResponseFieldName("custom_value")
```


## Values

| Name                               | Value                              |
| ---------------------------------- | ---------------------------------- |
| `OutputResponseFieldNameService`   | service                            |
| `OutputResponseFieldNameHostname`  | hostname                           |
| `OutputResponseFieldNameTimestamp` | timestamp                          |
| `OutputResponseFieldNameAuditID`   | auditId                            |