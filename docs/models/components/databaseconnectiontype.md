# DatabaseConnectionType

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
)

value := components.DatabaseConnectionTypeMysql

// Open enum: custom values can be created with a direct type cast
custom := components.DatabaseConnectionType("custom_value")
```


## Values

| Name                              | Value                             |
| --------------------------------- | --------------------------------- |
| `DatabaseConnectionTypeMysql`     | mysql                             |
| `DatabaseConnectionTypeOracle`    | oracle                            |
| `DatabaseConnectionTypePostgres`  | postgres                          |
| `DatabaseConnectionTypeSqlserver` | sqlserver                         |