# OutputDatasetSeverity

Default value for event severity. If the `sev` or `__severity` fields are set on an event, the first one matching will override this value.

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
)

value := components.OutputDatasetSeverityFinest

// Open enum: custom values can be created with a direct type cast
custom := components.OutputDatasetSeverity("custom_value")
```


## Values

| Name                           | Value                          |
| ------------------------------ | ------------------------------ |
| `OutputDatasetSeverityFinest`  | finest                         |
| `OutputDatasetSeverityFiner`   | finer                          |
| `OutputDatasetSeverityFine`    | fine                           |
| `OutputDatasetSeverityInfo`    | info                           |
| `OutputDatasetSeverityWarning` | warning                        |
| `OutputDatasetSeverityError`   | error                          |
| `OutputDatasetSeverityFatal`   | fatal                          |