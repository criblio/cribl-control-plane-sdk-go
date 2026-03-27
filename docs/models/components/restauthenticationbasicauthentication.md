# RestAuthenticationBasicAuthentication

Authentication method for Discover and Collect REST calls. You can specify API key–based authentication by adding the appropriate Collect headers.

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
)

value := components.RestAuthenticationBasicAuthenticationNone

// Open enum: custom values can be created with a direct type cast
custom := components.RestAuthenticationBasicAuthentication("custom_value")
```


## Values

| Name                                                     | Value                                                    |
| -------------------------------------------------------- | -------------------------------------------------------- |
| `RestAuthenticationBasicAuthenticationNone`              | none                                                     |
| `RestAuthenticationBasicAuthenticationBasic`             | basic                                                    |
| `RestAuthenticationBasicAuthenticationBasicSecret`       | basicSecret                                              |
| `RestAuthenticationBasicAuthenticationLogin`             | login                                                    |
| `RestAuthenticationBasicAuthenticationLoginSecret`       | loginSecret                                              |
| `RestAuthenticationBasicAuthenticationOauth`             | oauth                                                    |
| `RestAuthenticationBasicAuthenticationOauthSecret`       | oauthSecret                                              |
| `RestAuthenticationBasicAuthenticationGoogleOauth`       | google_oauth                                             |
| `RestAuthenticationBasicAuthenticationGoogleOauthSecret` | google_oauthSecret                                       |
| `RestAuthenticationBasicAuthenticationHmac`              | hmac                                                     |