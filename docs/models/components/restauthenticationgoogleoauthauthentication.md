# RestAuthenticationGoogleOauthAuthentication

Authentication method for Discover and Collect REST calls. You can specify API key–based authentication by adding the appropriate Collect headers.

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
)

value := components.RestAuthenticationGoogleOauthAuthenticationNone

// Open enum: custom values can be created with a direct type cast
custom := components.RestAuthenticationGoogleOauthAuthentication("custom_value")
```


## Values

| Name                                                           | Value                                                          |
| -------------------------------------------------------------- | -------------------------------------------------------------- |
| `RestAuthenticationGoogleOauthAuthenticationNone`              | none                                                           |
| `RestAuthenticationGoogleOauthAuthenticationBasic`             | basic                                                          |
| `RestAuthenticationGoogleOauthAuthenticationBasicSecret`       | basicSecret                                                    |
| `RestAuthenticationGoogleOauthAuthenticationLogin`             | login                                                          |
| `RestAuthenticationGoogleOauthAuthenticationLoginSecret`       | loginSecret                                                    |
| `RestAuthenticationGoogleOauthAuthenticationOauth`             | oauth                                                          |
| `RestAuthenticationGoogleOauthAuthenticationOauthSecret`       | oauthSecret                                                    |
| `RestAuthenticationGoogleOauthAuthenticationGoogleOauth`       | google_oauth                                                   |
| `RestAuthenticationGoogleOauthAuthenticationGoogleOauthSecret` | google_oauthSecret                                             |
| `RestAuthenticationGoogleOauthAuthenticationHmac`              | hmac                                                           |