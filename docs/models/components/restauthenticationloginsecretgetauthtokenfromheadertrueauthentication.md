# RestAuthenticationLoginSecretGetAuthTokenFromHeaderTrueAuthentication

Authentication method for Discover and Collect REST calls. You can specify API key–based authentication by adding the appropriate Collect headers.

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
)

value := components.RestAuthenticationLoginSecretGetAuthTokenFromHeaderTrueAuthenticationNone

// Open enum: custom values can be created with a direct type cast
custom := components.RestAuthenticationLoginSecretGetAuthTokenFromHeaderTrueAuthentication("custom_value")
```


## Values

| Name                                                                                     | Value                                                                                    |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `RestAuthenticationLoginSecretGetAuthTokenFromHeaderTrueAuthenticationNone`              | none                                                                                     |
| `RestAuthenticationLoginSecretGetAuthTokenFromHeaderTrueAuthenticationBasic`             | basic                                                                                    |
| `RestAuthenticationLoginSecretGetAuthTokenFromHeaderTrueAuthenticationBasicSecret`       | basicSecret                                                                              |
| `RestAuthenticationLoginSecretGetAuthTokenFromHeaderTrueAuthenticationLogin`             | login                                                                                    |
| `RestAuthenticationLoginSecretGetAuthTokenFromHeaderTrueAuthenticationLoginSecret`       | loginSecret                                                                              |
| `RestAuthenticationLoginSecretGetAuthTokenFromHeaderTrueAuthenticationOauth`             | oauth                                                                                    |
| `RestAuthenticationLoginSecretGetAuthTokenFromHeaderTrueAuthenticationOauthSecret`       | oauthSecret                                                                              |
| `RestAuthenticationLoginSecretGetAuthTokenFromHeaderTrueAuthenticationGoogleOauth`       | google_oauth                                                                             |
| `RestAuthenticationLoginSecretGetAuthTokenFromHeaderTrueAuthenticationGoogleOauthSecret` | google_oauthSecret                                                                       |
| `RestAuthenticationLoginSecretGetAuthTokenFromHeaderTrueAuthenticationHmac`              | hmac                                                                                     |