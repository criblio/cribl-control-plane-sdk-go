# RestCollectMethodGetAuthentication

Authentication method for Discover and Collect REST calls. You can specify API key–based authentication by adding the appropriate Collect headers.

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
)

value := components.RestCollectMethodGetAuthenticationNone

// Open enum: custom values can be created with a direct type cast
custom := components.RestCollectMethodGetAuthentication("custom_value")
```


## Values

| Name                                                  | Value                                                 |
| ----------------------------------------------------- | ----------------------------------------------------- |
| `RestCollectMethodGetAuthenticationNone`              | none                                                  |
| `RestCollectMethodGetAuthenticationBasic`             | basic                                                 |
| `RestCollectMethodGetAuthenticationBasicSecret`       | basicSecret                                           |
| `RestCollectMethodGetAuthenticationLogin`             | login                                                 |
| `RestCollectMethodGetAuthenticationLoginSecret`       | loginSecret                                           |
| `RestCollectMethodGetAuthenticationOauth`             | oauth                                                 |
| `RestCollectMethodGetAuthenticationOauthSecret`       | oauthSecret                                           |
| `RestCollectMethodGetAuthenticationGoogleOauth`       | google_oauth                                          |
| `RestCollectMethodGetAuthenticationGoogleOauthSecret` | google_oauthSecret                                    |
| `RestCollectMethodGetAuthenticationHmac`              | hmac                                                  |