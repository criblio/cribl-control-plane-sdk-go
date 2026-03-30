# CreateInputSystemByPackAuthenticationTypeServicenowTable

ServiceNow Table API authentication method

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/operations"
)

value := operations.CreateInputSystemByPackAuthenticationTypeServicenowTableNone

// Open enum: custom values can be created with a direct type cast
custom := operations.CreateInputSystemByPackAuthenticationTypeServicenowTable("custom_value")
```


## Values

| Name                                                                  | Value                                                                 |
| --------------------------------------------------------------------- | --------------------------------------------------------------------- |
| `CreateInputSystemByPackAuthenticationTypeServicenowTableNone`        | none                                                                  |
| `CreateInputSystemByPackAuthenticationTypeServicenowTableBasic`       | basic                                                                 |
| `CreateInputSystemByPackAuthenticationTypeServicenowTableBasicSecret` | basicSecret                                                           |
| `CreateInputSystemByPackAuthenticationTypeServicenowTableOauth`       | oauth                                                                 |
| `CreateInputSystemByPackAuthenticationTypeServicenowTableOauthSecret` | oauthSecret                                                           |