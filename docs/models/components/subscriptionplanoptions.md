# SubscriptionPlanOptions

Office 365 subscription plan for your organization, typically Office 365 Enterprise

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
)

value := components.SubscriptionPlanOptionsEnterpriseGcc

// Open enum: custom values can be created with a direct type cast
custom := components.SubscriptionPlanOptions("custom_value")
```


## Values

| Name                                   | Value                                  |
| -------------------------------------- | -------------------------------------- |
| `SubscriptionPlanOptionsEnterpriseGcc` | enterprise_gcc                         |
| `SubscriptionPlanOptionsGcc`           | gcc                                    |
| `SubscriptionPlanOptionsGccHigh`       | gcc_high                               |
| `SubscriptionPlanOptionsDod`           | dod                                    |