# OutputResponseLogLocationType

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
)

value := components.OutputResponseLogLocationTypeProject

// Open enum: custom values can be created with a direct type cast
custom := components.OutputResponseLogLocationType("custom_value")
```


## Values

| Name                                          | Value                                         |
| --------------------------------------------- | --------------------------------------------- |
| `OutputResponseLogLocationTypeProject`        | project                                       |
| `OutputResponseLogLocationTypeOrganization`   | organization                                  |
| `OutputResponseLogLocationTypeBillingAccount` | billingAccount                                |
| `OutputResponseLogLocationTypeFolder`         | folder                                        |