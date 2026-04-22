# Role

Leader Node role: <code>primary</code> or <code>standby</code>.

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
)

value := components.RolePrimary

// Open enum: custom values can be created with a direct type cast
custom := components.Role("custom_value")
```


## Values

| Name          | Value         |
| ------------- | ------------- |
| `RolePrimary` | primary       |
| `RoleStandby` | standby       |