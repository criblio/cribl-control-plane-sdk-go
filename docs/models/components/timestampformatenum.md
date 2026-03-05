# TimestampFormatEnum

Timestamp format to use when serializing event's time field

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
)

value := components.TimestampFormatEnumSyslog

// Open enum: custom values can be created with a direct type cast
custom := components.TimestampFormatEnum("custom_value")
```


## Values

| Name                         | Value                        |
| ---------------------------- | ---------------------------- |
| `TimestampFormatEnumSyslog`  | syslog                       |
| `TimestampFormatEnumIso8601` | iso8601                      |