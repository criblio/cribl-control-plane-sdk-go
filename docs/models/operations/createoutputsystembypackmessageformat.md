# CreateOutputSystemByPackMessageFormat

The syslog message format depending on the receiver's support

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/operations"
)

value := operations.CreateOutputSystemByPackMessageFormatRfc3164

// Open enum: custom values can be created with a direct type cast
custom := operations.CreateOutputSystemByPackMessageFormat("custom_value")
```


## Values

| Name                                           | Value                                          |
| ---------------------------------------------- | ---------------------------------------------- |
| `CreateOutputSystemByPackMessageFormatRfc3164` | rfc3164                                        |
| `CreateOutputSystemByPackMessageFormatRfc5424` | rfc5424                                        |