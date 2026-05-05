# InputResponseInputElasticAuthenticationMethod

Enter credentials directly, or select a stored secret

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
)

value := components.InputResponseInputElasticAuthenticationMethodNone

// Open enum: custom values can be created with a direct type cast
custom := components.InputResponseInputElasticAuthenticationMethod("custom_value")
```


## Values

| Name                                                  | Value                                                 |
| ----------------------------------------------------- | ----------------------------------------------------- |
| `InputResponseInputElasticAuthenticationMethodNone`   | none                                                  |
| `InputResponseInputElasticAuthenticationMethodManual` | manual                                                |
| `InputResponseInputElasticAuthenticationMethodSecret` | secret                                                |