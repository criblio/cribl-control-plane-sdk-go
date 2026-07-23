# UpgradeOptionsSystemSettingsConfSystem

Upgrade permission policy: <code>api</code> to allow upgrades from the UI or API or <code>false</code> to disable.

## Example Usage

```go
import (
	"github.com/Cribl-Community/cribl-control-plane-sdk-go/models/components"
)

value := components.UpgradeOptionsSystemSettingsConfSystemAPI

// Open enum: custom values can be created with a direct type cast
custom := components.UpgradeOptionsSystemSettingsConfSystem("custom_value")
```


## Values

| Name                                          | Value                                         |
| --------------------------------------------- | --------------------------------------------- |
| `UpgradeOptionsSystemSettingsConfSystemAPI`   | api                                           |
| `UpgradeOptionsSystemSettingsConfSystemFalse` | false                                         |