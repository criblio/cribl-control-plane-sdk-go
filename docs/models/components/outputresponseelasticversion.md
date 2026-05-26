# OutputResponseElasticVersion

Optional Elasticsearch version, used to format events. If not specified, will auto-discover version.

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
)

value := components.OutputResponseElasticVersionAuto

// Open enum: custom values can be created with a direct type cast
custom := components.OutputResponseElasticVersion("custom_value")
```


## Values

| Name                                | Value                               |
| ----------------------------------- | ----------------------------------- |
| `OutputResponseElasticVersionAuto`  | auto                                |
| `OutputResponseElasticVersionSix`   | 6                                   |
| `OutputResponseElasticVersionSeven` | 7                                   |