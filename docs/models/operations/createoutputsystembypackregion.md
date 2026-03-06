# CreateOutputSystemByPackRegion

The SentinelOne region to send events to. In most cases you can find the region by either looking at your SentinelOne URL or knowing what geographic region your SentinelOne instance is contained in.

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/operations"
)

value := operations.CreateOutputSystemByPackRegionUs

// Open enum: custom values can be created with a direct type cast
custom := operations.CreateOutputSystemByPackRegion("custom_value")
```


## Values

| Name                                   | Value                                  |
| -------------------------------------- | -------------------------------------- |
| `CreateOutputSystemByPackRegionUs`     | US                                     |
| `CreateOutputSystemByPackRegionCa`     | CA                                     |
| `CreateOutputSystemByPackRegionEmea`   | EMEA                                   |
| `CreateOutputSystemByPackRegionAp`     | AP                                     |
| `CreateOutputSystemByPackRegionAps`    | APS                                    |
| `CreateOutputSystemByPackRegionAu`     | AU                                     |
| `CreateOutputSystemByPackRegionCustom` | Custom                                 |