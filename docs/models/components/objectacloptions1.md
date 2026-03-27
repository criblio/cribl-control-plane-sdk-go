# ObjectACLOptions1

Object ACL to assign to uploaded objects

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
)

value := components.ObjectACLOptions1Private

// Open enum: custom values can be created with a direct type cast
custom := components.ObjectACLOptions1("custom_value")
```


## Values

| Name                                      | Value                                     |
| ----------------------------------------- | ----------------------------------------- |
| `ObjectACLOptions1Private`                | private                                   |
| `ObjectACLOptions1BucketOwnerRead`        | bucket-owner-read                         |
| `ObjectACLOptions1BucketOwnerFullControl` | bucket-owner-full-control                 |
| `ObjectACLOptions1ProjectPrivate`         | project-private                           |
| `ObjectACLOptions1AuthenticatedRead`      | authenticated-read                        |
| `ObjectACLOptions1PublicRead`             | public-read                               |