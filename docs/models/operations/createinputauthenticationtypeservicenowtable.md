# CreateInputAuthenticationTypeServicenowTable

ServiceNow Table API authentication method

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/operations"
)

value := operations.CreateInputAuthenticationTypeServicenowTableNone

// Open enum: custom values can be created with a direct type cast
custom := operations.CreateInputAuthenticationTypeServicenowTable("custom_value")
```


## Values

| Name                                                      | Value                                                     |
| --------------------------------------------------------- | --------------------------------------------------------- |
| `CreateInputAuthenticationTypeServicenowTableNone`        | none                                                      |
| `CreateInputAuthenticationTypeServicenowTableBasic`       | basic                                                     |
| `CreateInputAuthenticationTypeServicenowTableBasicSecret` | basicSecret                                               |
| `CreateInputAuthenticationTypeServicenowTableOauth`       | oauth                                                     |
| `CreateInputAuthenticationTypeServicenowTableOauthSecret` | oauthSecret                                               |