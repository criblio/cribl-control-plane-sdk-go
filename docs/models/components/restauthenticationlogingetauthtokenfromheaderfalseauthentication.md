# RestAuthenticationLoginGetAuthTokenFromHeaderFalseAuthentication

Authentication method for Discover and Collect REST calls. You can specify API key–based authentication by adding the appropriate Collect headers.

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
)

value := components.RestAuthenticationLoginGetAuthTokenFromHeaderFalseAuthenticationNone

// Open enum: custom values can be created with a direct type cast
custom := components.RestAuthenticationLoginGetAuthTokenFromHeaderFalseAuthentication("custom_value")
```


## Values

| Name                                                                                | Value                                                                               |
| ----------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------- |
| `RestAuthenticationLoginGetAuthTokenFromHeaderFalseAuthenticationNone`              | none                                                                                |
| `RestAuthenticationLoginGetAuthTokenFromHeaderFalseAuthenticationBasic`             | basic                                                                               |
| `RestAuthenticationLoginGetAuthTokenFromHeaderFalseAuthenticationBasicSecret`       | basicSecret                                                                         |
| `RestAuthenticationLoginGetAuthTokenFromHeaderFalseAuthenticationLogin`             | login                                                                               |
| `RestAuthenticationLoginGetAuthTokenFromHeaderFalseAuthenticationLoginSecret`       | loginSecret                                                                         |
| `RestAuthenticationLoginGetAuthTokenFromHeaderFalseAuthenticationOauth`             | oauth                                                                               |
| `RestAuthenticationLoginGetAuthTokenFromHeaderFalseAuthenticationOauthSecret`       | oauthSecret                                                                         |
| `RestAuthenticationLoginGetAuthTokenFromHeaderFalseAuthenticationGoogleOauth`       | google_oauth                                                                        |
| `RestAuthenticationLoginGetAuthTokenFromHeaderFalseAuthenticationGoogleOauthSecret` | google_oauthSecret                                                                  |
| `RestAuthenticationLoginGetAuthTokenFromHeaderFalseAuthenticationHmac`              | hmac                                                                                |