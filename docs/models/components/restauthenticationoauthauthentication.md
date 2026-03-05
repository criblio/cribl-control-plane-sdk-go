# RestAuthenticationOauthAuthentication

Authentication method for Discover and Collect REST calls. You can specify API key–based authentication by adding the appropriate Collect headers.

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
)

value := components.RestAuthenticationOauthAuthenticationNone

// Open enum: custom values can be created with a direct type cast
custom := components.RestAuthenticationOauthAuthentication("custom_value")
```


## Values

| Name                                                     | Value                                                    |
| -------------------------------------------------------- | -------------------------------------------------------- |
| `RestAuthenticationOauthAuthenticationNone`              | none                                                     |
| `RestAuthenticationOauthAuthenticationBasic`             | basic                                                    |
| `RestAuthenticationOauthAuthenticationBasicSecret`       | basicSecret                                              |
| `RestAuthenticationOauthAuthenticationLogin`             | login                                                    |
| `RestAuthenticationOauthAuthenticationLoginSecret`       | loginSecret                                              |
| `RestAuthenticationOauthAuthenticationOauth`             | oauth                                                    |
| `RestAuthenticationOauthAuthenticationOauthSecret`       | oauthSecret                                              |
| `RestAuthenticationOauthAuthenticationGoogleOauth`       | google_oauth                                             |
| `RestAuthenticationOauthAuthenticationGoogleOauthSecret` | google_oauthSecret                                       |
| `RestAuthenticationOauthAuthenticationHmac`              | hmac                                                     |