# SeverityDataset

Default value for event severity. If the `sev` or `__severity` fields are set on an event, the first one matching will override this value.

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
)

value := components.SeverityDatasetFinest

// Open enum: custom values can be created with a direct type cast
custom := components.SeverityDataset("custom_value")
```


## Values

| Name                     | Value                    |
| ------------------------ | ------------------------ |
| `SeverityDatasetFinest`  | finest                   |
| `SeverityDatasetFiner`   | finer                    |
| `SeverityDatasetFine`    | fine                     |
| `SeverityDatasetInfo`    | info                     |
| `SeverityDatasetWarning` | warning                  |
| `SeverityDatasetError`   | error                    |
| `SeverityDatasetFatal`   | fatal                    |