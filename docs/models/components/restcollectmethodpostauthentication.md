# RestCollectMethodPostAuthentication

Authentication method for Discover and Collect REST calls. You can specify API key–based authentication by adding the appropriate Collect headers.

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
)

value := components.RestCollectMethodPostAuthenticationNone

// Open enum: custom values can be created with a direct type cast
custom := components.RestCollectMethodPostAuthentication("custom_value")
```


## Values

| Name                                                   | Value                                                  |
| ------------------------------------------------------ | ------------------------------------------------------ |
| `RestCollectMethodPostAuthenticationNone`              | none                                                   |
| `RestCollectMethodPostAuthenticationBasic`             | basic                                                  |
| `RestCollectMethodPostAuthenticationBasicSecret`       | basicSecret                                            |
| `RestCollectMethodPostAuthenticationLogin`             | login                                                  |
| `RestCollectMethodPostAuthenticationLoginSecret`       | loginSecret                                            |
| `RestCollectMethodPostAuthenticationOauth`             | oauth                                                  |
| `RestCollectMethodPostAuthenticationOauthSecret`       | oauthSecret                                            |
| `RestCollectMethodPostAuthenticationGoogleOauth`       | google_oauth                                           |
| `RestCollectMethodPostAuthenticationGoogleOauthSecret` | google_oauthSecret                                     |
| `RestCollectMethodPostAuthenticationHmac`              | hmac                                                   |