# OutputResponseOutputDatasetSeverity

Default value for event severity. If the `sev` or `__severity` fields are set on an event, the first one matching will override this value.

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
)

value := components.OutputResponseOutputDatasetSeverityFinest

// Open enum: custom values can be created with a direct type cast
custom := components.OutputResponseOutputDatasetSeverity("custom_value")
```


## Values

| Name                                         | Value                                        |
| -------------------------------------------- | -------------------------------------------- |
| `OutputResponseOutputDatasetSeverityFinest`  | finest                                       |
| `OutputResponseOutputDatasetSeverityFiner`   | finer                                        |
| `OutputResponseOutputDatasetSeverityFine`    | fine                                         |
| `OutputResponseOutputDatasetSeverityInfo`    | info                                         |
| `OutputResponseOutputDatasetSeverityWarning` | warning                                      |
| `OutputResponseOutputDatasetSeverityError`   | error                                        |
| `OutputResponseOutputDatasetSeverityFatal`   | fatal                                        |