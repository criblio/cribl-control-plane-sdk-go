# CreateOutputFormatSentinel

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/operations"
)

value := operations.CreateOutputFormatSentinelNdjson

// Open enum: custom values can be created with a direct type cast
custom := operations.CreateOutputFormatSentinel("custom_value")
```


## Values

| Name                                  | Value                                 |
| ------------------------------------- | ------------------------------------- |
| `CreateOutputFormatSentinelNdjson`    | ndjson                                |
| `CreateOutputFormatSentinelJSONArray` | json_array                            |
| `CreateOutputFormatSentinelCustom`    | custom                                |
| `CreateOutputFormatSentinelAdvanced`  | advanced                              |