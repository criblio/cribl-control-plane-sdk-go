# CreateOutputTimestampFormat

Timestamp format to use when serializing event's time field

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/operations"
)

value := operations.CreateOutputTimestampFormatSyslog

// Open enum: custom values can be created with a direct type cast
custom := operations.CreateOutputTimestampFormat("custom_value")
```


## Values

| Name                                 | Value                                |
| ------------------------------------ | ------------------------------------ |
| `CreateOutputTimestampFormatSyslog`  | syslog                               |
| `CreateOutputTimestampFormatIso8601` | iso8601                              |