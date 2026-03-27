# RestCollectMethodOtherAuthentication

Authentication method for Discover and Collect REST calls. You can specify API key–based authentication by adding the appropriate Collect headers.

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
)

value := components.RestCollectMethodOtherAuthenticationNone

// Open enum: custom values can be created with a direct type cast
custom := components.RestCollectMethodOtherAuthentication("custom_value")
```


## Values

| Name                                                    | Value                                                   |
| ------------------------------------------------------- | ------------------------------------------------------- |
| `RestCollectMethodOtherAuthenticationNone`              | none                                                    |
| `RestCollectMethodOtherAuthenticationBasic`             | basic                                                   |
| `RestCollectMethodOtherAuthenticationBasicSecret`       | basicSecret                                             |
| `RestCollectMethodOtherAuthenticationLogin`             | login                                                   |
| `RestCollectMethodOtherAuthenticationLoginSecret`       | loginSecret                                             |
| `RestCollectMethodOtherAuthenticationOauth`             | oauth                                                   |
| `RestCollectMethodOtherAuthenticationOauthSecret`       | oauthSecret                                             |
| `RestCollectMethodOtherAuthenticationGoogleOauth`       | google_oauth                                            |
| `RestCollectMethodOtherAuthenticationGoogleOauthSecret` | google_oauthSecret                                      |
| `RestCollectMethodOtherAuthenticationHmac`              | hmac                                                    |