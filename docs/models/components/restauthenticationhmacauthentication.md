# RestAuthenticationHmacAuthentication

Authentication method for Discover and Collect REST calls. You can specify API key–based authentication by adding the appropriate Collect headers.

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
)

value := components.RestAuthenticationHmacAuthenticationNone

// Open enum: custom values can be created with a direct type cast
custom := components.RestAuthenticationHmacAuthentication("custom_value")
```


## Values

| Name                                                    | Value                                                   |
| ------------------------------------------------------- | ------------------------------------------------------- |
| `RestAuthenticationHmacAuthenticationNone`              | none                                                    |
| `RestAuthenticationHmacAuthenticationBasic`             | basic                                                   |
| `RestAuthenticationHmacAuthenticationBasicSecret`       | basicSecret                                             |
| `RestAuthenticationHmacAuthenticationLogin`             | login                                                   |
| `RestAuthenticationHmacAuthenticationLoginSecret`       | loginSecret                                             |
| `RestAuthenticationHmacAuthenticationOauth`             | oauth                                                   |
| `RestAuthenticationHmacAuthenticationOauthSecret`       | oauthSecret                                             |
| `RestAuthenticationHmacAuthenticationGoogleOauth`       | google_oauth                                            |
| `RestAuthenticationHmacAuthenticationGoogleOauthSecret` | google_oauthSecret                                      |
| `RestAuthenticationHmacAuthenticationHmac`              | hmac                                                    |