# CreateOutputDataSetSite

DataSet site to which events should be sent

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/operations"
)

value := operations.CreateOutputDataSetSiteUs

// Open enum: custom values can be created with a direct type cast
custom := operations.CreateOutputDataSetSite("custom_value")
```


## Values

| Name                            | Value                           |
| ------------------------------- | ------------------------------- |
| `CreateOutputDataSetSiteUs`     | us                              |
| `CreateOutputDataSetSiteEu`     | eu                              |
| `CreateOutputDataSetSiteCustom` | custom                          |