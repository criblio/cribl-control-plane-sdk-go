# InputWefAuthenticationMethod

How to authenticate incoming client connections

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
)

value := components.InputWefAuthenticationMethodClientCert

// Open enum: custom values can be created with a direct type cast
custom := components.InputWefAuthenticationMethod("custom_value")
```


## Values

| Name                                     | Value                                    |
| ---------------------------------------- | ---------------------------------------- |
| `InputWefAuthenticationMethodClientCert` | clientCert                               |
| `InputWefAuthenticationMethodKerberos`   | kerberos                                 |