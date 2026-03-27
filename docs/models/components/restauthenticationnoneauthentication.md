# RestAuthenticationNoneAuthentication

Authentication method for Discover and Collect REST calls. You can specify API key–based authentication by adding the appropriate Collect headers.

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
)

value := components.RestAuthenticationNoneAuthenticationNone

// Open enum: custom values can be created with a direct type cast
custom := components.RestAuthenticationNoneAuthentication("custom_value")
```


## Values

| Name                                                    | Value                                                   |
| ------------------------------------------------------- | ------------------------------------------------------- |
| `RestAuthenticationNoneAuthenticationNone`              | none                                                    |
| `RestAuthenticationNoneAuthenticationBasic`             | basic                                                   |
| `RestAuthenticationNoneAuthenticationBasicSecret`       | basicSecret                                             |
| `RestAuthenticationNoneAuthenticationLogin`             | login                                                   |
| `RestAuthenticationNoneAuthenticationLoginSecret`       | loginSecret                                             |
| `RestAuthenticationNoneAuthenticationOauth`             | oauth                                                   |
| `RestAuthenticationNoneAuthenticationOauthSecret`       | oauthSecret                                             |
| `RestAuthenticationNoneAuthenticationGoogleOauth`       | google_oauth                                            |
| `RestAuthenticationNoneAuthenticationGoogleOauthSecret` | google_oauthSecret                                      |
| `RestAuthenticationNoneAuthenticationHmac`              | hmac                                                    |