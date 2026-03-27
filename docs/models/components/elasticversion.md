# ElasticVersion

Optional Elasticsearch version, used to format events. If not specified, will auto-discover version.

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
)

value := components.ElasticVersionAuto

// Open enum: custom values can be created with a direct type cast
custom := components.ElasticVersion("custom_value")
```


## Values

| Name                  | Value                 |
| --------------------- | --------------------- |
| `ElasticVersionAuto`  | auto                  |
| `ElasticVersionSix`   | 6                     |
| `ElasticVersionSeven` | 7                     |