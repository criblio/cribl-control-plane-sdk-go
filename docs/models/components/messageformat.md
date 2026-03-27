# MessageFormat

The syslog message format depending on the receiver's support

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
)

value := components.MessageFormatRfc3164

// Open enum: custom values can be created with a direct type cast
custom := components.MessageFormat("custom_value")
```


## Values

| Name                   | Value                  |
| ---------------------- | ---------------------- |
| `MessageFormatRfc3164` | rfc3164                |
| `MessageFormatRfc5424` | rfc5424                |