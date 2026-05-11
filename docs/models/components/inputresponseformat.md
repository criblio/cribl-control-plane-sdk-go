# InputResponseFormat

Content format in which the endpoint should deliver events

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
)

value := components.InputResponseFormatRaw

// Open enum: custom values can be created with a direct type cast
custom := components.InputResponseFormat("custom_value")
```


## Values

| Name                              | Value                             |
| --------------------------------- | --------------------------------- |
| `InputResponseFormatRaw`          | Raw                               |
| `InputResponseFormatRenderedText` | RenderedText                      |