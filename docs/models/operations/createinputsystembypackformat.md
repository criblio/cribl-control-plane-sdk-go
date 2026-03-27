# CreateInputSystemByPackFormat

Content format in which the endpoint should deliver events

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/operations"
)

value := operations.CreateInputSystemByPackFormatRaw

// Open enum: custom values can be created with a direct type cast
custom := operations.CreateInputSystemByPackFormat("custom_value")
```


## Values

| Name                                        | Value                                       |
| ------------------------------------------- | ------------------------------------------- |
| `CreateInputSystemByPackFormatRaw`          | Raw                                         |
| `CreateInputSystemByPackFormatRenderedText` | RenderedText                                |