# SplunkAuthenticationBasicAuthentication

Authentication method for Discover and Collect REST calls

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
)

value := components.SplunkAuthenticationBasicAuthenticationNone

// Open enum: custom values can be created with a direct type cast
custom := components.SplunkAuthenticationBasicAuthentication("custom_value")
```


## Values

| Name                                                 | Value                                                |
| ---------------------------------------------------- | ---------------------------------------------------- |
| `SplunkAuthenticationBasicAuthenticationNone`        | none                                                 |
| `SplunkAuthenticationBasicAuthenticationBasic`       | basic                                                |
| `SplunkAuthenticationBasicAuthenticationBasicSecret` | basicSecret                                          |
| `SplunkAuthenticationBasicAuthenticationToken`       | token                                                |
| `SplunkAuthenticationBasicAuthenticationTokenSecret` | tokenSecret                                          |