# CreateOutputSystemByPackProtocolSyslog

The network protocol to use for sending out syslog messages

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/operations"
)

value := operations.CreateOutputSystemByPackProtocolSyslogTCP

// Open enum: custom values can be created with a direct type cast
custom := operations.CreateOutputSystemByPackProtocolSyslog("custom_value")
```


## Values

| Name                                        | Value                                       |
| ------------------------------------------- | ------------------------------------------- |
| `CreateOutputSystemByPackProtocolSyslogTCP` | tcp                                         |
| `CreateOutputSystemByPackProtocolSyslogUDP` | udp                                         |