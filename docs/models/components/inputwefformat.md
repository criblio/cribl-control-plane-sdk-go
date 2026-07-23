# InputWefFormat

Content format in which the endpoint should deliver events

## Example Usage

```go
import (
	"github.com/Cribl-Community/cribl-control-plane-sdk-go/models/components"
)

value := components.InputWefFormatRaw

// Open enum: custom values can be created with a direct type cast
custom := components.InputWefFormat("custom_value")
```


## Values

| Name                         | Value                        |
| ---------------------------- | ---------------------------- |
| `InputWefFormatRaw`          | Raw                          |
| `InputWefFormatRenderedText` | RenderedText                 |