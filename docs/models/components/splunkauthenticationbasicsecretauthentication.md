# SplunkAuthenticationBasicSecretAuthentication

Authentication method for Discover and Collect REST calls

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
)

value := components.SplunkAuthenticationBasicSecretAuthenticationNone

// Open enum: custom values can be created with a direct type cast
custom := components.SplunkAuthenticationBasicSecretAuthentication("custom_value")
```


## Values

| Name                                                       | Value                                                      |
| ---------------------------------------------------------- | ---------------------------------------------------------- |
| `SplunkAuthenticationBasicSecretAuthenticationNone`        | none                                                       |
| `SplunkAuthenticationBasicSecretAuthenticationBasic`       | basic                                                      |
| `SplunkAuthenticationBasicSecretAuthenticationBasicSecret` | basicSecret                                                |
| `SplunkAuthenticationBasicSecretAuthenticationToken`       | token                                                      |
| `SplunkAuthenticationBasicSecretAuthenticationTokenSecret` | tokenSecret                                                |