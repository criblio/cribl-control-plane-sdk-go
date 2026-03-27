# CreateOutputSystemByPackDefaultSeveritySeverity

Default value for event severity. If the `sev` or `__severity` fields are set on an event, the first one matching will override this value.

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/operations"
)

value := operations.CreateOutputSystemByPackDefaultSeveritySeverityFinest

// Open enum: custom values can be created with a direct type cast
custom := operations.CreateOutputSystemByPackDefaultSeveritySeverity("custom_value")
```


## Values

| Name                                                     | Value                                                    |
| -------------------------------------------------------- | -------------------------------------------------------- |
| `CreateOutputSystemByPackDefaultSeveritySeverityFinest`  | finest                                                   |
| `CreateOutputSystemByPackDefaultSeveritySeverityFiner`   | finer                                                    |
| `CreateOutputSystemByPackDefaultSeveritySeverityFine`    | fine                                                     |
| `CreateOutputSystemByPackDefaultSeveritySeverityInfo`    | info                                                     |
| `CreateOutputSystemByPackDefaultSeveritySeverityWarning` | warning                                                  |
| `CreateOutputSystemByPackDefaultSeveritySeverityError`   | error                                                    |
| `CreateOutputSystemByPackDefaultSeveritySeverityFatal`   | fatal                                                    |