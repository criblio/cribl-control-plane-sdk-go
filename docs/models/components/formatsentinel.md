# FormatSentinel

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
)

value := components.FormatSentinelNdjson

// Open enum: custom values can be created with a direct type cast
custom := components.FormatSentinel("custom_value")
```


## Values

| Name                      | Value                     |
| ------------------------- | ------------------------- |
| `FormatSentinelNdjson`    | ndjson                    |
| `FormatSentinelJSONArray` | json_array                |
| `FormatSentinelCustom`    | custom                    |
| `FormatSentinelAdvanced`  | advanced                  |