# AuthenticationTypeServicenowTable

ServiceNow Table API authentication method

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
)

value := components.AuthenticationTypeServicenowTableNone

// Open enum: custom values can be created with a direct type cast
custom := components.AuthenticationTypeServicenowTable("custom_value")
```


## Values

| Name                                           | Value                                          |
| ---------------------------------------------- | ---------------------------------------------- |
| `AuthenticationTypeServicenowTableNone`        | none                                           |
| `AuthenticationTypeServicenowTableBasicSecret` | basicSecret                                    |
| `AuthenticationTypeServicenowTableOauthSecret` | oauthSecret                                    |