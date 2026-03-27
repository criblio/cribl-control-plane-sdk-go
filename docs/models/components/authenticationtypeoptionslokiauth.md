# AuthenticationTypeOptionsLokiAuth

Loki logs authentication type

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
)

value := components.AuthenticationTypeOptionsLokiAuthNone

// Open enum: custom values can be created with a direct type cast
custom := components.AuthenticationTypeOptionsLokiAuth("custom_value")
```


## Values

| Name                                                 | Value                                                |
| ---------------------------------------------------- | ---------------------------------------------------- |
| `AuthenticationTypeOptionsLokiAuthNone`              | none                                                 |
| `AuthenticationTypeOptionsLokiAuthBasic`             | basic                                                |
| `AuthenticationTypeOptionsLokiAuthCredentialsSecret` | credentialsSecret                                    |
| `AuthenticationTypeOptionsLokiAuthToken`             | token                                                |
| `AuthenticationTypeOptionsLokiAuthTextSecret`        | textSecret                                           |