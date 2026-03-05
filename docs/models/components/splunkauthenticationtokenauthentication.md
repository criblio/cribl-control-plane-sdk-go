# SplunkAuthenticationTokenAuthentication

Authentication method for Discover and Collect REST calls

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
)

value := components.SplunkAuthenticationTokenAuthenticationNone

// Open enum: custom values can be created with a direct type cast
custom := components.SplunkAuthenticationTokenAuthentication("custom_value")
```


## Values

| Name                                                 | Value                                                |
| ---------------------------------------------------- | ---------------------------------------------------- |
| `SplunkAuthenticationTokenAuthenticationNone`        | none                                                 |
| `SplunkAuthenticationTokenAuthenticationBasic`       | basic                                                |
| `SplunkAuthenticationTokenAuthenticationBasicSecret` | basicSecret                                          |
| `SplunkAuthenticationTokenAuthenticationToken`       | token                                                |
| `SplunkAuthenticationTokenAuthenticationTokenSecret` | tokenSecret                                          |