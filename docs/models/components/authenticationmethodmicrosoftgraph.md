# AuthenticationMethodMicrosoftGraph

Select authentication method.

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
)

value := components.AuthenticationMethodMicrosoftGraphOauth

// Open enum: custom values can be created with a direct type cast
custom := components.AuthenticationMethodMicrosoftGraph("custom_value")
```


## Values

| Name                                            | Value                                           |
| ----------------------------------------------- | ----------------------------------------------- |
| `AuthenticationMethodMicrosoftGraphOauth`       | oauth                                           |
| `AuthenticationMethodMicrosoftGraphOauthSecret` | oauthSecret                                     |
| `AuthenticationMethodMicrosoftGraphOauthCert`   | oauthCert                                       |