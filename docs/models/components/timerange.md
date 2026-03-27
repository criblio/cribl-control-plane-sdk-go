# TimeRange

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
)

value := components.TimeRangeAbsolute

// Open enum: custom values can be created with a direct type cast
custom := components.TimeRange("custom_value")
```


## Values

| Name                | Value               |
| ------------------- | ------------------- |
| `TimeRangeAbsolute` | absolute            |
| `TimeRangeRelative` | relative            |