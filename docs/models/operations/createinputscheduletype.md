# CreateInputScheduleType

Select a schedule type; either an interval (in seconds) or a cron-style schedule.

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/operations"
)

value := operations.CreateInputScheduleTypeInterval

// Open enum: custom values can be created with a direct type cast
custom := operations.CreateInputScheduleType("custom_value")
```


## Values

| Name                                  | Value                                 |
| ------------------------------------- | ------------------------------------- |
| `CreateInputScheduleTypeInterval`     | interval                              |
| `CreateInputScheduleTypeCronSchedule` | cronSchedule                          |