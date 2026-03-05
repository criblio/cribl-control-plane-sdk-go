# CreateOutputMessageFormat

The syslog message format depending on the receiver's support

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/operations"
)

value := operations.CreateOutputMessageFormatRfc3164

// Open enum: custom values can be created with a direct type cast
custom := operations.CreateOutputMessageFormat("custom_value")
```


## Values

| Name                               | Value                              |
| ---------------------------------- | ---------------------------------- |
| `CreateOutputMessageFormatRfc3164` | rfc3164                            |
| `CreateOutputMessageFormatRfc5424` | rfc5424                            |