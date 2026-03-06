# SplunkAuthenticationTokenSecretAuthentication

Authentication method for Discover and Collect REST calls

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
)

value := components.SplunkAuthenticationTokenSecretAuthenticationNone

// Open enum: custom values can be created with a direct type cast
custom := components.SplunkAuthenticationTokenSecretAuthentication("custom_value")
```


## Values

| Name                                                       | Value                                                      |
| ---------------------------------------------------------- | ---------------------------------------------------------- |
| `SplunkAuthenticationTokenSecretAuthenticationNone`        | none                                                       |
| `SplunkAuthenticationTokenSecretAuthenticationBasic`       | basic                                                      |
| `SplunkAuthenticationTokenSecretAuthenticationBasicSecret` | basicSecret                                                |
| `SplunkAuthenticationTokenSecretAuthenticationToken`       | token                                                      |
| `SplunkAuthenticationTokenSecretAuthenticationTokenSecret` | tokenSecret                                                |