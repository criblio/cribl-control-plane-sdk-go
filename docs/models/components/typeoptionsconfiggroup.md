# TypeOptionsConfigGroup

Explicit type of the Worker Group, Outpost Group, or Edge Fleet.

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
)

value := components.TypeOptionsConfigGroupEdge

// Open enum: custom values can be created with a direct type cast
custom := components.TypeOptionsConfigGroup("custom_value")
```


## Values

| Name                                | Value                               |
| ----------------------------------- | ----------------------------------- |
| `TypeOptionsConfigGroupEdge`        | edge                                |
| `TypeOptionsConfigGroupLakeAccess`  | lake_access                         |
| `TypeOptionsConfigGroupLocalSearch` | local_search                        |
| `TypeOptionsConfigGroupOutpost`     | outpost                             |
| `TypeOptionsConfigGroupSearch`      | search                              |
| `TypeOptionsConfigGroupStream`      | stream                              |