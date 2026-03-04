# RunType

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
)

value := components.RunTypeAdhoc

// Open enum: custom values can be created with a direct type cast
custom := components.RunType("custom_value")
```


## Values

| Name               | Value              |
| ------------------ | ------------------ |
| `RunTypeAdhoc`     | adhoc              |
| `RunTypeScheduled` | scheduled          |
| `RunTypeSystem`    | system             |