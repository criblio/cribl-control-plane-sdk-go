# CreateInputSystemByPackScheduleType

Select a schedule type; either an interval (in seconds) or a cron-style schedule.

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/operations"
)

value := operations.CreateInputSystemByPackScheduleTypeInterval

// Open enum: custom values can be created with a direct type cast
custom := operations.CreateInputSystemByPackScheduleType("custom_value")
```


## Values

| Name                                              | Value                                             |
| ------------------------------------------------- | ------------------------------------------------- |
| `CreateInputSystemByPackScheduleTypeInterval`     | interval                                          |
| `CreateInputSystemByPackScheduleTypeCronSchedule` | cronSchedule                                      |