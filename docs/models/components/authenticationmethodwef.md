# AuthenticationMethodWef

How to authenticate incoming client connections

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
)

value := components.AuthenticationMethodWefClientCert

// Open enum: custom values can be created with a direct type cast
custom := components.AuthenticationMethodWef("custom_value")
```


## Values

| Name                                | Value                               |
| ----------------------------------- | ----------------------------------- |
| `AuthenticationMethodWefClientCert` | clientCert                          |
| `AuthenticationMethodWefKerberos`   | kerberos                            |