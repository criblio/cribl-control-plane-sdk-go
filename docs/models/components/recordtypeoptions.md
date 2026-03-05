# RecordTypeOptions

DNS record type to resolve

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
)

value := components.RecordTypeOptionsSrv

// Open enum: custom values can be created with a direct type cast
custom := components.RecordTypeOptions("custom_value")
```


## Values

| Name                    | Value                   |
| ----------------------- | ----------------------- |
| `RecordTypeOptionsSrv`  | SRV                     |
| `RecordTypeOptionsA`    | A                       |
| `RecordTypeOptionsAaaa` | AAAA                    |