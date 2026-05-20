# OutputSentinelOneAiSiemRegion

The SentinelOne region to send events to. In most cases you can find the region by either looking at your SentinelOne URL or knowing what geographic region your SentinelOne instance is contained in.

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
)

value := components.OutputSentinelOneAiSiemRegionUs

// Open enum: custom values can be created with a direct type cast
custom := components.OutputSentinelOneAiSiemRegion("custom_value")
```


## Values

| Name                                  | Value                                 |
| ------------------------------------- | ------------------------------------- |
| `OutputSentinelOneAiSiemRegionUs`     | US                                    |
| `OutputSentinelOneAiSiemRegionCa`     | CA                                    |
| `OutputSentinelOneAiSiemRegionEmea`   | EMEA                                  |
| `OutputSentinelOneAiSiemRegionAp`     | AP                                    |
| `OutputSentinelOneAiSiemRegionAps`    | APS                                   |
| `OutputSentinelOneAiSiemRegionAu`     | AU                                    |
| `OutputSentinelOneAiSiemRegionCustom` | Custom                                |