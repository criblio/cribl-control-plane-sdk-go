# CreateInputSystemByPackAuthenticationTypeSplunkSearch

Splunk Search authentication type

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/operations"
)

value := operations.CreateInputSystemByPackAuthenticationTypeSplunkSearchNone

// Open enum: custom values can be created with a direct type cast
custom := operations.CreateInputSystemByPackAuthenticationTypeSplunkSearch("custom_value")
```


## Values

| Name                                                                     | Value                                                                    |
| ------------------------------------------------------------------------ | ------------------------------------------------------------------------ |
| `CreateInputSystemByPackAuthenticationTypeSplunkSearchNone`              | none                                                                     |
| `CreateInputSystemByPackAuthenticationTypeSplunkSearchBasic`             | basic                                                                    |
| `CreateInputSystemByPackAuthenticationTypeSplunkSearchCredentialsSecret` | credentialsSecret                                                        |
| `CreateInputSystemByPackAuthenticationTypeSplunkSearchToken`             | token                                                                    |
| `CreateInputSystemByPackAuthenticationTypeSplunkSearchTextSecret`        | textSecret                                                               |