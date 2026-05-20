# ProtocolSyslog

The network protocol to use for sending out syslog messages

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
)

value := components.ProtocolSyslogTCP

// Open enum: custom values can be created with a direct type cast
custom := components.ProtocolSyslog("custom_value")
```


## Values

| Name                | Value               |
| ------------------- | ------------------- |
| `ProtocolSyslogTCP` | tcp                 |
| `ProtocolSyslogUDP` | udp                 |