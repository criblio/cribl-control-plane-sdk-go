# SecureVersion

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
)

value := components.SecureVersionTlSv13

// Open enum: custom values can be created with a direct type cast
custom := components.SecureVersion("custom_value")
```


## Values

| Name                  | Value                 |
| --------------------- | --------------------- |
| `SecureVersionTlSv13` | TLSv1.3               |
| `SecureVersionTlSv12` | TLSv1.2               |
| `SecureVersionTlSv11` | TLSv1.1               |
| `SecureVersionTlSv1`  | TLSv1                 |