# OutputElasticElasticVersion

Optional Elasticsearch version, used to format events. If not specified, will auto-discover version.

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
)

value := components.OutputElasticElasticVersionAuto

// Open enum: custom values can be created with a direct type cast
custom := components.OutputElasticElasticVersion("custom_value")
```


## Values

| Name                               | Value                              |
| ---------------------------------- | ---------------------------------- |
| `OutputElasticElasticVersionAuto`  | auto                               |
| `OutputElasticElasticVersionSix`   | 6                                  |
| `OutputElasticElasticVersionSeven` | 7                                  |