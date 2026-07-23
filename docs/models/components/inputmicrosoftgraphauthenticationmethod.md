# InputMicrosoftGraphAuthenticationMethod

Select authentication method.

## Example Usage

```go
import (
	"github.com/Cribl-Community/cribl-control-plane-sdk-go/models/components"
)

value := components.InputMicrosoftGraphAuthenticationMethodOauth

// Open enum: custom values can be created with a direct type cast
custom := components.InputMicrosoftGraphAuthenticationMethod("custom_value")
```


## Values

| Name                                                 | Value                                                |
| ---------------------------------------------------- | ---------------------------------------------------- |
| `InputMicrosoftGraphAuthenticationMethodOauth`       | oauth                                                |
| `InputMicrosoftGraphAuthenticationMethodOauthSecret` | oauthSecret                                          |
| `InputMicrosoftGraphAuthenticationMethodOauthCert`   | oauthCert                                            |