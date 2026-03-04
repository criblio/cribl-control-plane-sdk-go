# EventFormat

Format of individual events

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
)

value := components.EventFormatJSON

// Open enum: custom values can be created with a direct type cast
custom := components.EventFormat("custom_value")
```


## Values

| Name              | Value             |
| ----------------- | ----------------- |
| `EventFormatJSON` | json              |
| `EventFormatXML`  | xml               |