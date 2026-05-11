# AuthenticationTypeSplunkSearch

Splunk Search authentication type

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
)

value := components.AuthenticationTypeSplunkSearchNone

// Open enum: custom values can be created with a direct type cast
custom := components.AuthenticationTypeSplunkSearch("custom_value")
```


## Values

| Name                                              | Value                                             |
| ------------------------------------------------- | ------------------------------------------------- |
| `AuthenticationTypeSplunkSearchNone`              | none                                              |
| `AuthenticationTypeSplunkSearchBasic`             | basic                                             |
| `AuthenticationTypeSplunkSearchCredentialsSecret` | credentialsSecret                                 |
| `AuthenticationTypeSplunkSearchToken`             | token                                             |
| `AuthenticationTypeSplunkSearchTextSecret`        | textSecret                                        |