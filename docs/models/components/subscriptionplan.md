# SubscriptionPlan

Microsoft 365 subscription plan for your organization, typically Microsoft 365 Enterprise

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
)

value := components.SubscriptionPlanEnterpriseGcc

// Open enum: custom values can be created with a direct type cast
custom := components.SubscriptionPlan("custom_value")
```


## Values

| Name                            | Value                           |
| ------------------------------- | ------------------------------- |
| `SubscriptionPlanEnterpriseGcc` | enterprise_gcc                  |
| `SubscriptionPlanGcc`           | gcc                             |
| `SubscriptionPlanGccHigh`       | gcc_high                        |
| `SubscriptionPlanDod`           | dod                             |
| `SubscriptionPlanChina`         | china                           |