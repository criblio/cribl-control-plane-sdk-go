# OutputSyslogMessageFormat

The syslog message format depending on the receiver's support

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
)

value := components.OutputSyslogMessageFormatRfc3164

// Open enum: custom values can be created with a direct type cast
custom := components.OutputSyslogMessageFormat("custom_value")
```


## Values

| Name                               | Value                              |
| ---------------------------------- | ---------------------------------- |
| `OutputSyslogMessageFormatRfc3164` | rfc3164                            |
| `OutputSyslogMessageFormatRfc5424` | rfc5424                            |