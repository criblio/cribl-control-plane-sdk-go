# RegionOptions

Which New Relic region endpoint to use.

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
)

value := components.RegionOptionsUs

// Open enum: custom values can be created with a direct type cast
custom := components.RegionOptions("custom_value")
```


## Values

| Name                  | Value                 |
| --------------------- | --------------------- |
| `RegionOptionsUs`     | US                    |
| `RegionOptionsEu`     | EU                    |
| `RegionOptionsCustom` | Custom                |