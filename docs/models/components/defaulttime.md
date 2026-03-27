# DefaultTime

How to set the time field if no timestamp is found

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
)

value := components.DefaultTimeNow

// Open enum: custom values can be created with a direct type cast
custom := components.DefaultTime("custom_value")
```


## Values

| Name              | Value             |
| ----------------- | ----------------- |
| `DefaultTimeNow`  | now               |
| `DefaultTimeLast` | last              |
| `DefaultTimeNone` | none              |