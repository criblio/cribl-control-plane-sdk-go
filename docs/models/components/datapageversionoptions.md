# DataPageVersionOptions

Serialization format of data pages. Note that some reader implementations use Data page V2's attributes to work more efficiently, while others ignore it.

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
)

value := components.DataPageVersionOptionsDataPageV1

// Open enum: custom values can be created with a direct type cast
custom := components.DataPageVersionOptions("custom_value")
```


## Values

| Name                               | Value                              |
| ---------------------------------- | ---------------------------------- |
| `DataPageVersionOptionsDataPageV1` | DATA_PAGE_V1                       |
| `DataPageVersionOptionsDataPageV2` | DATA_PAGE_V2                       |