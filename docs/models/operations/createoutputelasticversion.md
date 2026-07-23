# CreateOutputElasticVersion

Optional Elasticsearch version, used to format events. If not specified, will auto-discover version.

## Example Usage

```go
import (
	"github.com/Cribl-Community/cribl-control-plane-sdk-go/models/operations"
)

value := operations.CreateOutputElasticVersionAuto

// Open enum: custom values can be created with a direct type cast
custom := operations.CreateOutputElasticVersion("custom_value")
```


## Values

| Name                              | Value                             |
| --------------------------------- | --------------------------------- |
| `CreateOutputElasticVersionAuto`  | auto                              |
| `CreateOutputElasticVersionSix`   | 6                                 |
| `CreateOutputElasticVersionSeven` | 7                                 |