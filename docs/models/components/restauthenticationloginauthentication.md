# RestAuthenticationLoginAuthentication

Authentication method for Discover and Collect REST calls. You can specify API key–based authentication by adding the appropriate Collect headers.

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
)

value := components.RestAuthenticationLoginAuthenticationNone

// Open enum: custom values can be created with a direct type cast
custom := components.RestAuthenticationLoginAuthentication("custom_value")
```


## Values

| Name                                                     | Value                                                    |
| -------------------------------------------------------- | -------------------------------------------------------- |
| `RestAuthenticationLoginAuthenticationNone`              | none                                                     |
| `RestAuthenticationLoginAuthenticationBasic`             | basic                                                    |
| `RestAuthenticationLoginAuthenticationBasicSecret`       | basicSecret                                              |
| `RestAuthenticationLoginAuthenticationLogin`             | login                                                    |
| `RestAuthenticationLoginAuthenticationLoginSecret`       | loginSecret                                              |
| `RestAuthenticationLoginAuthenticationOauth`             | oauth                                                    |
| `RestAuthenticationLoginAuthenticationOauthSecret`       | oauthSecret                                              |
| `RestAuthenticationLoginAuthenticationGoogleOauth`       | google_oauth                                             |
| `RestAuthenticationLoginAuthenticationGoogleOauthSecret` | google_oauthSecret                                       |
| `RestAuthenticationLoginAuthenticationHmac`              | hmac                                                     |