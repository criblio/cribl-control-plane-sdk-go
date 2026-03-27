# RestAuthenticationBasicSecretAuthentication

Authentication method for Discover and Collect REST calls. You can specify API key–based authentication by adding the appropriate Collect headers.

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
)

value := components.RestAuthenticationBasicSecretAuthenticationNone

// Open enum: custom values can be created with a direct type cast
custom := components.RestAuthenticationBasicSecretAuthentication("custom_value")
```


## Values

| Name                                                           | Value                                                          |
| -------------------------------------------------------------- | -------------------------------------------------------------- |
| `RestAuthenticationBasicSecretAuthenticationNone`              | none                                                           |
| `RestAuthenticationBasicSecretAuthenticationBasic`             | basic                                                          |
| `RestAuthenticationBasicSecretAuthenticationBasicSecret`       | basicSecret                                                    |
| `RestAuthenticationBasicSecretAuthenticationLogin`             | login                                                          |
| `RestAuthenticationBasicSecretAuthenticationLoginSecret`       | loginSecret                                                    |
| `RestAuthenticationBasicSecretAuthenticationOauth`             | oauth                                                          |
| `RestAuthenticationBasicSecretAuthenticationOauthSecret`       | oauthSecret                                                    |
| `RestAuthenticationBasicSecretAuthenticationGoogleOauth`       | google_oauth                                                   |
| `RestAuthenticationBasicSecretAuthenticationGoogleOauthSecret` | google_oauthSecret                                             |
| `RestAuthenticationBasicSecretAuthenticationHmac`              | hmac                                                           |