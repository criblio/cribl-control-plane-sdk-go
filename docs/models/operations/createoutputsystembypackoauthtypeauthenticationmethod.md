# CreateOutputSystemByPackOauthTypeAuthenticationMethod

The type of OAuth 2.0 client credentials grant flow to use

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/operations"
)

value := operations.CreateOutputSystemByPackOauthTypeAuthenticationMethodClientSecret

// Open enum: custom values can be created with a direct type cast
custom := operations.CreateOutputSystemByPackOauthTypeAuthenticationMethod("custom_value")
```


## Values

| Name                                                                    | Value                                                                   |
| ----------------------------------------------------------------------- | ----------------------------------------------------------------------- |
| `CreateOutputSystemByPackOauthTypeAuthenticationMethodClientSecret`     | clientSecret                                                            |
| `CreateOutputSystemByPackOauthTypeAuthenticationMethodClientTextSecret` | clientTextSecret                                                        |
| `CreateOutputSystemByPackOauthTypeAuthenticationMethodCertificate`      | certificate                                                             |