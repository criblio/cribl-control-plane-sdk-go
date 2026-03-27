# RestAuthenticationGoogleOauthSecretAuthentication

Authentication method for Discover and Collect REST calls. You can specify API key–based authentication by adding the appropriate Collect headers.

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
)

value := components.RestAuthenticationGoogleOauthSecretAuthenticationNone

// Open enum: custom values can be created with a direct type cast
custom := components.RestAuthenticationGoogleOauthSecretAuthentication("custom_value")
```


## Values

| Name                                                                 | Value                                                                |
| -------------------------------------------------------------------- | -------------------------------------------------------------------- |
| `RestAuthenticationGoogleOauthSecretAuthenticationNone`              | none                                                                 |
| `RestAuthenticationGoogleOauthSecretAuthenticationBasic`             | basic                                                                |
| `RestAuthenticationGoogleOauthSecretAuthenticationBasicSecret`       | basicSecret                                                          |
| `RestAuthenticationGoogleOauthSecretAuthenticationLogin`             | login                                                                |
| `RestAuthenticationGoogleOauthSecretAuthenticationLoginSecret`       | loginSecret                                                          |
| `RestAuthenticationGoogleOauthSecretAuthenticationOauth`             | oauth                                                                |
| `RestAuthenticationGoogleOauthSecretAuthenticationOauthSecret`       | oauthSecret                                                          |
| `RestAuthenticationGoogleOauthSecretAuthenticationGoogleOauth`       | google_oauth                                                         |
| `RestAuthenticationGoogleOauthSecretAuthenticationGoogleOauthSecret` | google_oauthSecret                                                   |
| `RestAuthenticationGoogleOauthSecretAuthenticationHmac`              | hmac                                                                 |