# RbacResource

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
)

value := components.RbacResourceGroups

// Open enum: custom values can be created with a direct type cast
custom := components.RbacResource("custom_value")
```


## Values

| Name                            | Value                           |
| ------------------------------- | ------------------------------- |
| `RbacResourceGroups`            | groups                          |
| `RbacResourceDatasets`          | datasets                        |
| `RbacResourceDatasetProviders`  | dataset-providers               |
| `RbacResourceProjects`          | projects                        |
| `RbacResourceDashboards`        | dashboards                      |
| `RbacResourceMacros`            | macros                          |
| `RbacResourceNotebooks`         | notebooks                       |
| `RbacResourceNotebookTemplates` | notebook-templates              |