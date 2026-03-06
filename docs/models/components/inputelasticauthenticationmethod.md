# InputElasticAuthenticationMethod

Enter credentials directly, or select a stored secret

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
)

value := components.InputElasticAuthenticationMethodNone

// Open enum: custom values can be created with a direct type cast
custom := components.InputElasticAuthenticationMethod("custom_value")
```


## Values

| Name                                     | Value                                    |
| ---------------------------------------- | ---------------------------------------- |
| `InputElasticAuthenticationMethodNone`   | none                                     |
| `InputElasticAuthenticationMethodManual` | manual                                   |
| `InputElasticAuthenticationMethodSecret` | secret                                   |