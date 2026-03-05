# CreateOutputOauthTypeAuthenticationMethod

The type of OAuth 2.0 client credentials grant flow to use

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/operations"
)

value := operations.CreateOutputOauthTypeAuthenticationMethodClientSecret

// Open enum: custom values can be created with a direct type cast
custom := operations.CreateOutputOauthTypeAuthenticationMethod("custom_value")
```


## Values

| Name                                                        | Value                                                       |
| ----------------------------------------------------------- | ----------------------------------------------------------- |
| `CreateOutputOauthTypeAuthenticationMethodClientSecret`     | clientSecret                                                |
| `CreateOutputOauthTypeAuthenticationMethodClientTextSecret` | clientTextSecret                                            |
| `CreateOutputOauthTypeAuthenticationMethodCertificate`      | certificate                                                 |