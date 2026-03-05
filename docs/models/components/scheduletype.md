# ScheduleType

Select a schedule type; either an interval (in seconds) or a cron-style schedule.

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
)

value := components.ScheduleTypeInterval

// Open enum: custom values can be created with a direct type cast
custom := components.ScheduleType("custom_value")
```


## Values

| Name                       | Value                      |
| -------------------------- | -------------------------- |
| `ScheduleTypeInterval`     | interval                   |
| `ScheduleTypeCronSchedule` | cronSchedule               |