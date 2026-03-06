# RestAuthenticationOauthSecretAuthentication

Authentication method for Discover and Collect REST calls. You can specify API key–based authentication by adding the appropriate Collect headers.

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
)

value := components.RestAuthenticationOauthSecretAuthenticationNone

// Open enum: custom values can be created with a direct type cast
custom := components.RestAuthenticationOauthSecretAuthentication("custom_value")
```


## Values

| Name                                                           | Value                                                          |
| -------------------------------------------------------------- | -------------------------------------------------------------- |
| `RestAuthenticationOauthSecretAuthenticationNone`              | none                                                           |
| `RestAuthenticationOauthSecretAuthenticationBasic`             | basic                                                          |
| `RestAuthenticationOauthSecretAuthenticationBasicSecret`       | basicSecret                                                    |
| `RestAuthenticationOauthSecretAuthenticationLogin`             | login                                                          |
| `RestAuthenticationOauthSecretAuthenticationLoginSecret`       | loginSecret                                                    |
| `RestAuthenticationOauthSecretAuthenticationOauth`             | oauth                                                          |
| `RestAuthenticationOauthSecretAuthenticationOauthSecret`       | oauthSecret                                                    |
| `RestAuthenticationOauthSecretAuthenticationGoogleOauth`       | google_oauth                                                   |
| `RestAuthenticationOauthSecretAuthenticationGoogleOauthSecret` | google_oauthSecret                                             |
| `RestAuthenticationOauthSecretAuthenticationHmac`              | hmac                                                           |