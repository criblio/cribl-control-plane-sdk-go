# CreateOutputRegion

The SentinelOne region to send events to. In most cases you can find the region by either looking at your SentinelOne URL or knowing what geographic region your SentinelOne instance is contained in.

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/operations"
)

value := operations.CreateOutputRegionUs

// Open enum: custom values can be created with a direct type cast
custom := operations.CreateOutputRegion("custom_value")
```


## Values

| Name                       | Value                      |
| -------------------------- | -------------------------- |
| `CreateOutputRegionUs`     | US                         |
| `CreateOutputRegionCa`     | CA                         |
| `CreateOutputRegionEmea`   | EMEA                       |
| `CreateOutputRegionAp`     | AP                         |
| `CreateOutputRegionAps`    | APS                        |
| `CreateOutputRegionAu`     | AU                         |
| `CreateOutputRegionCustom` | Custom                     |