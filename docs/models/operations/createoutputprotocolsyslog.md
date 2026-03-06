# CreateOutputProtocolSyslog

The network protocol to use for sending out syslog messages

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/operations"
)

value := operations.CreateOutputProtocolSyslogTCP

// Open enum: custom values can be created with a direct type cast
custom := operations.CreateOutputProtocolSyslog("custom_value")
```


## Values

| Name                            | Value                           |
| ------------------------------- | ------------------------------- |
| `CreateOutputProtocolSyslogTCP` | tcp                             |
| `CreateOutputProtocolSyslogUDP` | udp                             |