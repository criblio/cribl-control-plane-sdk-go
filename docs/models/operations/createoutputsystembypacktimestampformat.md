# CreateOutputSystemByPackTimestampFormat

Timestamp format to use when serializing event's time field

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/operations"
)

value := operations.CreateOutputSystemByPackTimestampFormatSyslog

// Open enum: custom values can be created with a direct type cast
custom := operations.CreateOutputSystemByPackTimestampFormat("custom_value")
```


## Values

| Name                                             | Value                                            |
| ------------------------------------------------ | ------------------------------------------------ |
| `CreateOutputSystemByPackTimestampFormatSyslog`  | syslog                                           |
| `CreateOutputSystemByPackTimestampFormatIso8601` | iso8601                                          |