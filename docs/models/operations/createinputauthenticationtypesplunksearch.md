# CreateInputAuthenticationTypeSplunkSearch

Splunk Search authentication type

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/operations"
)

value := operations.CreateInputAuthenticationTypeSplunkSearchNone

// Open enum: custom values can be created with a direct type cast
custom := operations.CreateInputAuthenticationTypeSplunkSearch("custom_value")
```


## Values

| Name                                                         | Value                                                        |
| ------------------------------------------------------------ | ------------------------------------------------------------ |
| `CreateInputAuthenticationTypeSplunkSearchNone`              | none                                                         |
| `CreateInputAuthenticationTypeSplunkSearchBasic`             | basic                                                        |
| `CreateInputAuthenticationTypeSplunkSearchCredentialsSecret` | credentialsSecret                                            |
| `CreateInputAuthenticationTypeSplunkSearchToken`             | token                                                        |
| `CreateInputAuthenticationTypeSplunkSearchTextSecret`        | textSecret                                                   |