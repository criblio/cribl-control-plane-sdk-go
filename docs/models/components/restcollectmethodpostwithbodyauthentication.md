# RestCollectMethodPostWithBodyAuthentication

Authentication method for Discover and Collect REST calls. You can specify API key–based authentication by adding the appropriate Collect headers.

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
)

value := components.RestCollectMethodPostWithBodyAuthenticationNone

// Open enum: custom values can be created with a direct type cast
custom := components.RestCollectMethodPostWithBodyAuthentication("custom_value")
```


## Values

| Name                                                           | Value                                                          |
| -------------------------------------------------------------- | -------------------------------------------------------------- |
| `RestCollectMethodPostWithBodyAuthenticationNone`              | none                                                           |
| `RestCollectMethodPostWithBodyAuthenticationBasic`             | basic                                                          |
| `RestCollectMethodPostWithBodyAuthenticationBasicSecret`       | basicSecret                                                    |
| `RestCollectMethodPostWithBodyAuthenticationLogin`             | login                                                          |
| `RestCollectMethodPostWithBodyAuthenticationLoginSecret`       | loginSecret                                                    |
| `RestCollectMethodPostWithBodyAuthenticationOauth`             | oauth                                                          |
| `RestCollectMethodPostWithBodyAuthenticationOauthSecret`       | oauthSecret                                                    |
| `RestCollectMethodPostWithBodyAuthenticationGoogleOauth`       | google_oauth                                                   |
| `RestCollectMethodPostWithBodyAuthenticationGoogleOauthSecret` | google_oauthSecret                                             |
| `RestCollectMethodPostWithBodyAuthenticationHmac`              | hmac                                                           |