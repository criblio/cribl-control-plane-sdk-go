# RestAuthenticationLoginGetAuthTokenFromHeaderTrueAuthentication

Authentication method for Discover and Collect REST calls. You can specify API key–based authentication by adding the appropriate Collect headers.

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
)

value := components.RestAuthenticationLoginGetAuthTokenFromHeaderTrueAuthenticationNone

// Open enum: custom values can be created with a direct type cast
custom := components.RestAuthenticationLoginGetAuthTokenFromHeaderTrueAuthentication("custom_value")
```


## Values

| Name                                                                               | Value                                                                              |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `RestAuthenticationLoginGetAuthTokenFromHeaderTrueAuthenticationNone`              | none                                                                               |
| `RestAuthenticationLoginGetAuthTokenFromHeaderTrueAuthenticationBasic`             | basic                                                                              |
| `RestAuthenticationLoginGetAuthTokenFromHeaderTrueAuthenticationBasicSecret`       | basicSecret                                                                        |
| `RestAuthenticationLoginGetAuthTokenFromHeaderTrueAuthenticationLogin`             | login                                                                              |
| `RestAuthenticationLoginGetAuthTokenFromHeaderTrueAuthenticationLoginSecret`       | loginSecret                                                                        |
| `RestAuthenticationLoginGetAuthTokenFromHeaderTrueAuthenticationOauth`             | oauth                                                                              |
| `RestAuthenticationLoginGetAuthTokenFromHeaderTrueAuthenticationOauthSecret`       | oauthSecret                                                                        |
| `RestAuthenticationLoginGetAuthTokenFromHeaderTrueAuthenticationGoogleOauth`       | google_oauth                                                                       |
| `RestAuthenticationLoginGetAuthTokenFromHeaderTrueAuthenticationGoogleOauthSecret` | google_oauthSecret                                                                 |
| `RestAuthenticationLoginGetAuthTokenFromHeaderTrueAuthenticationHmac`              | hmac                                                                               |