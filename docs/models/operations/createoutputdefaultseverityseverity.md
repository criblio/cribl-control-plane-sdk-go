# CreateOutputDefaultSeveritySeverity

Default value for event severity. If the `sev` or `__severity` fields are set on an event, the first one matching will override this value.

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/operations"
)

value := operations.CreateOutputDefaultSeveritySeverityFinest

// Open enum: custom values can be created with a direct type cast
custom := operations.CreateOutputDefaultSeveritySeverity("custom_value")
```


## Values

| Name                                         | Value                                        |
| -------------------------------------------- | -------------------------------------------- |
| `CreateOutputDefaultSeveritySeverityFinest`  | finest                                       |
| `CreateOutputDefaultSeveritySeverityFiner`   | finer                                        |
| `CreateOutputDefaultSeveritySeverityFine`    | fine                                         |
| `CreateOutputDefaultSeveritySeverityInfo`    | info                                         |
| `CreateOutputDefaultSeveritySeverityWarning` | warning                                      |
| `CreateOutputDefaultSeveritySeverityError`   | error                                        |
| `CreateOutputDefaultSeveritySeverityFatal`   | fatal                                        |