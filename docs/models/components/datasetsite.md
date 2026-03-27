# DataSetSite

DataSet site to which events should be sent

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
)

value := components.DataSetSiteUs

// Open enum: custom values can be created with a direct type cast
custom := components.DataSetSite("custom_value")
```


## Values

| Name                | Value               |
| ------------------- | ------------------- |
| `DataSetSiteUs`     | us                  |
| `DataSetSiteEu`     | eu                  |
| `DataSetSiteCustom` | custom              |