# TimestampFormat

Timestamp format to use when serializing event's time field

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
)

value := components.TimestampFormatSyslog

// Open enum: custom values can be created with a direct type cast
custom := components.TimestampFormat("custom_value")
```


## Values

| Name                     | Value                    |
| ------------------------ | ------------------------ |
| `TimestampFormatSyslog`  | syslog                   |
| `TimestampFormatIso8601` | iso8601                  |