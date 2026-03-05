# OutputSyslogProtocol

The network protocol to use for sending out syslog messages

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
)

value := components.OutputSyslogProtocolTCP

// Open enum: custom values can be created with a direct type cast
custom := components.OutputSyslogProtocol("custom_value")
```


## Values

| Name                      | Value                     |
| ------------------------- | ------------------------- |
| `OutputSyslogProtocolTCP` | tcp                       |
| `OutputSyslogProtocolUDP` | udp                       |