# LogLocationType

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
)

value := components.LogLocationTypeProject

// Open enum: custom values can be created with a direct type cast
custom := components.LogLocationType("custom_value")
```


## Values

| Name                            | Value                           |
| ------------------------------- | ------------------------------- |
| `LogLocationTypeProject`        | project                         |
| `LogLocationTypeOrganization`   | organization                    |
| `LogLocationTypeBillingAccount` | billingAccount                  |
| `LogLocationTypeFolder`         | folder                          |