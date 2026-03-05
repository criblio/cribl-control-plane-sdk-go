# OutputSentinelFormat

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
)

value := components.OutputSentinelFormatNdjson

// Open enum: custom values can be created with a direct type cast
custom := components.OutputSentinelFormat("custom_value")
```


## Values

| Name                            | Value                           |
| ------------------------------- | ------------------------------- |
| `OutputSentinelFormatNdjson`    | ndjson                          |
| `OutputSentinelFormatJSONArray` | json_array                      |
| `OutputSentinelFormatCustom`    | custom                          |
| `OutputSentinelFormatAdvanced`  | advanced                        |