# LogLevelOptionsRunnableJobCollectionScheduleRun

Level at which to set task logging

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
)

value := components.LogLevelOptionsRunnableJobCollectionScheduleRunError

// Open enum: custom values can be created with a direct type cast
custom := components.LogLevelOptionsRunnableJobCollectionScheduleRun("custom_value")
```


## Values

| Name                                                   | Value                                                  |
| ------------------------------------------------------ | ------------------------------------------------------ |
| `LogLevelOptionsRunnableJobCollectionScheduleRunError` | error                                                  |
| `LogLevelOptionsRunnableJobCollectionScheduleRunWarn`  | warn                                                   |
| `LogLevelOptionsRunnableJobCollectionScheduleRunInfo`  | info                                                   |
| `LogLevelOptionsRunnableJobCollectionScheduleRunDebug` | debug                                                  |
| `LogLevelOptionsRunnableJobCollectionScheduleRunSilly` | silly                                                  |