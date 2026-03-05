# ObjectACLOptionsAuthenticatedreadBucketownerfullcontrol

Object ACL to assign to uploaded objects

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
)

value := components.ObjectACLOptionsAuthenticatedreadBucketownerfullcontrolPrivate

// Open enum: custom values can be created with a direct type cast
custom := components.ObjectACLOptionsAuthenticatedreadBucketownerfullcontrol("custom_value")
```


## Values

| Name                                                                            | Value                                                                           |
| ------------------------------------------------------------------------------- | ------------------------------------------------------------------------------- |
| `ObjectACLOptionsAuthenticatedreadBucketownerfullcontrolPrivate`                | private                                                                         |
| `ObjectACLOptionsAuthenticatedreadBucketownerfullcontrolBucketOwnerRead`        | bucket-owner-read                                                               |
| `ObjectACLOptionsAuthenticatedreadBucketownerfullcontrolBucketOwnerFullControl` | bucket-owner-full-control                                                       |
| `ObjectACLOptionsAuthenticatedreadBucketownerfullcontrolProjectPrivate`         | project-private                                                                 |
| `ObjectACLOptionsAuthenticatedreadBucketownerfullcontrolAuthenticatedRead`      | authenticated-read                                                              |
| `ObjectACLOptionsAuthenticatedreadBucketownerfullcontrolPublicRead`             | public-read                                                                     |