# OutputResponseRegion

The SentinelOne region to send events to. In most cases you can find the region by either looking at your SentinelOne URL or knowing what geographic region your SentinelOne instance is contained in.

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
)

value := components.OutputResponseRegionUs

// Open enum: custom values can be created with a direct type cast
custom := components.OutputResponseRegion("custom_value")
```


## Values

| Name                         | Value                        |
| ---------------------------- | ---------------------------- |
| `OutputResponseRegionUs`     | US                           |
| `OutputResponseRegionCa`     | CA                           |
| `OutputResponseRegionEmea`   | EMEA                         |
| `OutputResponseRegionAp`     | AP                           |
| `OutputResponseRegionAps`    | APS                          |
| `OutputResponseRegionAu`     | AU                           |
| `OutputResponseRegionCustom` | Custom                       |