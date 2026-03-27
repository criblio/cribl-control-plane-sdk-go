# RestAuthenticationLoginSecretGetAuthTokenFromHeaderFalseAuthentication

Authentication method for Discover and Collect REST calls. You can specify API key–based authentication by adding the appropriate Collect headers.

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
)

value := components.RestAuthenticationLoginSecretGetAuthTokenFromHeaderFalseAuthenticationNone

// Open enum: custom values can be created with a direct type cast
custom := components.RestAuthenticationLoginSecretGetAuthTokenFromHeaderFalseAuthentication("custom_value")
```


## Values

| Name                                                                                      | Value                                                                                     |
| ----------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------- |
| `RestAuthenticationLoginSecretGetAuthTokenFromHeaderFalseAuthenticationNone`              | none                                                                                      |
| `RestAuthenticationLoginSecretGetAuthTokenFromHeaderFalseAuthenticationBasic`             | basic                                                                                     |
| `RestAuthenticationLoginSecretGetAuthTokenFromHeaderFalseAuthenticationBasicSecret`       | basicSecret                                                                               |
| `RestAuthenticationLoginSecretGetAuthTokenFromHeaderFalseAuthenticationLogin`             | login                                                                                     |
| `RestAuthenticationLoginSecretGetAuthTokenFromHeaderFalseAuthenticationLoginSecret`       | loginSecret                                                                               |
| `RestAuthenticationLoginSecretGetAuthTokenFromHeaderFalseAuthenticationOauth`             | oauth                                                                                     |
| `RestAuthenticationLoginSecretGetAuthTokenFromHeaderFalseAuthenticationOauthSecret`       | oauthSecret                                                                               |
| `RestAuthenticationLoginSecretGetAuthTokenFromHeaderFalseAuthenticationGoogleOauth`       | google_oauth                                                                              |
| `RestAuthenticationLoginSecretGetAuthTokenFromHeaderFalseAuthenticationGoogleOauthSecret` | google_oauthSecret                                                                        |
| `RestAuthenticationLoginSecretGetAuthTokenFromHeaderFalseAuthenticationHmac`              | hmac                                                                                      |