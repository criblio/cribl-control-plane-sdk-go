# RestAuthenticationLoginSecretAuthentication

Authentication method for Discover and Collect REST calls. You can specify API key–based authentication by adding the appropriate Collect headers.

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
)

value := components.RestAuthenticationLoginSecretAuthenticationNone

// Open enum: custom values can be created with a direct type cast
custom := components.RestAuthenticationLoginSecretAuthentication("custom_value")
```


## Values

| Name                                                           | Value                                                          |
| -------------------------------------------------------------- | -------------------------------------------------------------- |
| `RestAuthenticationLoginSecretAuthenticationNone`              | none                                                           |
| `RestAuthenticationLoginSecretAuthenticationBasic`             | basic                                                          |
| `RestAuthenticationLoginSecretAuthenticationBasicSecret`       | basicSecret                                                    |
| `RestAuthenticationLoginSecretAuthenticationLogin`             | login                                                          |
| `RestAuthenticationLoginSecretAuthenticationLoginSecret`       | loginSecret                                                    |
| `RestAuthenticationLoginSecretAuthenticationOauth`             | oauth                                                          |
| `RestAuthenticationLoginSecretAuthenticationOauthSecret`       | oauthSecret                                                    |
| `RestAuthenticationLoginSecretAuthenticationGoogleOauth`       | google_oauth                                                   |
| `RestAuthenticationLoginSecretAuthenticationGoogleOauthSecret` | google_oauthSecret                                             |
| `RestAuthenticationLoginSecretAuthenticationHmac`              | hmac                                                           |