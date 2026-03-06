# SaslMechanismOptionsSasl

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
)

value := components.SaslMechanismOptionsSaslPlain

// Open enum: custom values can be created with a direct type cast
custom := components.SaslMechanismOptionsSasl("custom_value")
```


## Values

| Name                                  | Value                                 |
| ------------------------------------- | ------------------------------------- |
| `SaslMechanismOptionsSaslPlain`       | plain                                 |
| `SaslMechanismOptionsSaslScramSha256` | scram-sha-256                         |
| `SaslMechanismOptionsSaslScramSha512` | scram-sha-512                         |
| `SaslMechanismOptionsSaslKerberos`    | kerberos                              |