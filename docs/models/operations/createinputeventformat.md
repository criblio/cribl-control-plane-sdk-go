# CreateInputEventFormat

Format of individual events

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/operations"
)

value := operations.CreateInputEventFormatJSON

// Open enum: custom values can be created with a direct type cast
custom := operations.CreateInputEventFormat("custom_value")
```


## Values

| Name                         | Value                        |
| ---------------------------- | ---------------------------- |
| `CreateInputEventFormatJSON` | json                         |
| `CreateInputEventFormatXML`  | xml                          |