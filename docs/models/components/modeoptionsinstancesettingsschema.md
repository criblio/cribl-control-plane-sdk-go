# ModeOptionsInstanceSettingsSchema

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
)

value := components.ModeOptionsInstanceSettingsSchemaSingle

// Open enum: custom values can be created with a direct type cast
custom := components.ModeOptionsInstanceSettingsSchema("custom_value")
```


## Values

| Name                                                | Value                                               |
| --------------------------------------------------- | --------------------------------------------------- |
| `ModeOptionsInstanceSettingsSchemaSingle`           | single                                              |
| `ModeOptionsInstanceSettingsSchemaMaster`           | master                                              |
| `ModeOptionsInstanceSettingsSchemaWorker`           | worker                                              |
| `ModeOptionsInstanceSettingsSchemaEdge`             | edge                                                |
| `ModeOptionsInstanceSettingsSchemaManagedEdge`      | managed-edge                                        |
| `ModeOptionsInstanceSettingsSchemaOutpost`          | outpost                                             |
| `ModeOptionsInstanceSettingsSchemaSearchSupervisor` | search-supervisor                                   |