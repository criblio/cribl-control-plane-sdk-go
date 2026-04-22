# CreateInputAuthenticationMethodMicrosoftGraph

Select authentication method.

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/operations"
)

value := operations.CreateInputAuthenticationMethodMicrosoftGraphOauth

// Open enum: custom values can be created with a direct type cast
custom := operations.CreateInputAuthenticationMethodMicrosoftGraph("custom_value")
```


## Values

| Name                                                       | Value                                                      |
| ---------------------------------------------------------- | ---------------------------------------------------------- |
| `CreateInputAuthenticationMethodMicrosoftGraphOauth`       | oauth                                                      |
| `CreateInputAuthenticationMethodMicrosoftGraphOauthSecret` | oauthSecret                                                |
| `CreateInputAuthenticationMethodMicrosoftGraphOauthCert`   | oauthCert                                                  |