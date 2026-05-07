# OutputResponseTimestampFormat

Timestamp format to use when serializing event's time field

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
)

value := components.OutputResponseTimestampFormatSyslog

// Open enum: custom values can be created with a direct type cast
custom := components.OutputResponseTimestampFormat("custom_value")
```


## Values

| Name                                   | Value                                  |
| -------------------------------------- | -------------------------------------- |
| `OutputResponseTimestampFormatSyslog`  | syslog                                 |
| `OutputResponseTimestampFormatIso8601` | iso8601                                |