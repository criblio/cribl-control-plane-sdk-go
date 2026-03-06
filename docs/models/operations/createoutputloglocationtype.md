# CreateOutputLogLocationType

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/operations"
)

value := operations.CreateOutputLogLocationTypeProject

// Open enum: custom values can be created with a direct type cast
custom := operations.CreateOutputLogLocationType("custom_value")
```


## Values

| Name                                        | Value                                       |
| ------------------------------------------- | ------------------------------------------- |
| `CreateOutputLogLocationTypeProject`        | project                                     |
| `CreateOutputLogLocationTypeOrganization`   | organization                                |
| `CreateOutputLogLocationTypeBillingAccount` | billingAccount                              |
| `CreateOutputLogLocationTypeFolder`         | folder                                      |