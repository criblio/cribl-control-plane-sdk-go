# Region

The SentinelOne region to send events to. In most cases you can find the region by either looking at your SentinelOne URL or knowing what geographic region your SentinelOne instance is contained in.

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
)

value := components.RegionUs

// Open enum: custom values can be created with a direct type cast
custom := components.Region("custom_value")
```


## Values

| Name           | Value          |
| -------------- | -------------- |
| `RegionUs`     | US             |
| `RegionCa`     | CA             |
| `RegionEmea`   | EMEA           |
| `RegionAp`     | AP             |
| `RegionAps`    | APS            |
| `RegionAu`     | AU             |
| `RegionCustom` | Custom         |