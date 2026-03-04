# CreateOutputPayloadFormat

Format to use when sending payload. Defaults to Text.

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/operations"
)

value := operations.CreateOutputPayloadFormatText

// Open enum: custom values can be created with a direct type cast
custom := operations.CreateOutputPayloadFormat("custom_value")
```


## Values

| Name                            | Value                           |
| ------------------------------- | ------------------------------- |
| `CreateOutputPayloadFormatText` | text                            |
| `CreateOutputPayloadFormatJSON` | json                            |