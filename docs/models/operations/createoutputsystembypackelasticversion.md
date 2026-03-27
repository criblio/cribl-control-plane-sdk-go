# CreateOutputSystemByPackElasticVersion

Optional Elasticsearch version, used to format events. If not specified, will auto-discover version.

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/operations"
)

value := operations.CreateOutputSystemByPackElasticVersionAuto

// Open enum: custom values can be created with a direct type cast
custom := operations.CreateOutputSystemByPackElasticVersion("custom_value")
```


## Values

| Name                                          | Value                                         |
| --------------------------------------------- | --------------------------------------------- |
| `CreateOutputSystemByPackElasticVersionAuto`  | auto                                          |
| `CreateOutputSystemByPackElasticVersionSix`   | 6                                             |
| `CreateOutputSystemByPackElasticVersionSeven` | 7                                             |