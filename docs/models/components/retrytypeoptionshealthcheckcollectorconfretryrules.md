# RetryTypeOptionsHealthCheckCollectorConfRetryRules

The algorithm to use when performing HTTP retries

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
)

value := components.RetryTypeOptionsHealthCheckCollectorConfRetryRulesNone

// Open enum: custom values can be created with a direct type cast
custom := components.RetryTypeOptionsHealthCheckCollectorConfRetryRules("custom_value")
```


## Values

| Name                                                        | Value                                                       |
| ----------------------------------------------------------- | ----------------------------------------------------------- |
| `RetryTypeOptionsHealthCheckCollectorConfRetryRulesNone`    | none                                                        |
| `RetryTypeOptionsHealthCheckCollectorConfRetryRulesBackoff` | backoff                                                     |
| `RetryTypeOptionsHealthCheckCollectorConfRetryRulesStatic`  | static                                                      |