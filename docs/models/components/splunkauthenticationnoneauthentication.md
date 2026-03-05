# SplunkAuthenticationNoneAuthentication

Authentication method for Discover and Collect REST calls

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
)

value := components.SplunkAuthenticationNoneAuthenticationNone

// Open enum: custom values can be created with a direct type cast
custom := components.SplunkAuthenticationNoneAuthentication("custom_value")
```


## Values

| Name                                                | Value                                               |
| --------------------------------------------------- | --------------------------------------------------- |
| `SplunkAuthenticationNoneAuthenticationNone`        | none                                                |
| `SplunkAuthenticationNoneAuthenticationBasic`       | basic                                               |
| `SplunkAuthenticationNoneAuthenticationBasicSecret` | basicSecret                                         |
| `SplunkAuthenticationNoneAuthenticationToken`       | token                                               |
| `SplunkAuthenticationNoneAuthenticationTokenSecret` | tokenSecret                                         |