# AuthenticationMethodElastic

Enter credentials directly, or select a stored secret

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
)

value := components.AuthenticationMethodElasticNone

// Open enum: custom values can be created with a direct type cast
custom := components.AuthenticationMethodElastic("custom_value")
```


## Values

| Name                                | Value                               |
| ----------------------------------- | ----------------------------------- |
| `AuthenticationMethodElasticNone`   | none                                |
| `AuthenticationMethodElasticManual` | manual                              |
| `AuthenticationMethodElasticSecret` | secret                              |