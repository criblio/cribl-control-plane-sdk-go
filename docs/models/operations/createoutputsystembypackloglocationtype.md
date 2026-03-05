# CreateOutputSystemByPackLogLocationType

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/operations"
)

value := operations.CreateOutputSystemByPackLogLocationTypeProject

// Open enum: custom values can be created with a direct type cast
custom := operations.CreateOutputSystemByPackLogLocationType("custom_value")
```


## Values

| Name                                                    | Value                                                   |
| ------------------------------------------------------- | ------------------------------------------------------- |
| `CreateOutputSystemByPackLogLocationTypeProject`        | project                                                 |
| `CreateOutputSystemByPackLogLocationTypeOrganization`   | organization                                            |
| `CreateOutputSystemByPackLogLocationTypeBillingAccount` | billingAccount                                          |
| `CreateOutputSystemByPackLogLocationTypeFolder`         | folder                                                  |