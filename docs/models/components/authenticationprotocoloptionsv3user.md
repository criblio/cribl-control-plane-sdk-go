# AuthenticationProtocolOptionsV3User

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
)

value := components.AuthenticationProtocolOptionsV3UserNone

// Open enum: custom values can be created with a direct type cast
custom := components.AuthenticationProtocolOptionsV3User("custom_value")
```


## Values

| Name                                        | Value                                       |
| ------------------------------------------- | ------------------------------------------- |
| `AuthenticationProtocolOptionsV3UserNone`   | none                                        |
| `AuthenticationProtocolOptionsV3UserMd5`    | md5                                         |
| `AuthenticationProtocolOptionsV3UserSha`    | sha                                         |
| `AuthenticationProtocolOptionsV3UserSha224` | sha224                                      |
| `AuthenticationProtocolOptionsV3UserSha256` | sha256                                      |
| `AuthenticationProtocolOptionsV3UserSha384` | sha384                                      |
| `AuthenticationProtocolOptionsV3UserSha512` | sha512                                      |