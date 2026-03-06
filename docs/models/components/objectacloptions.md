# ObjectACLOptions

Object ACL to assign to uploaded objects

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
)

value := components.ObjectACLOptionsPrivate

// Open enum: custom values can be created with a direct type cast
custom := components.ObjectACLOptions("custom_value")
```


## Values

| Name                                     | Value                                    |
| ---------------------------------------- | ---------------------------------------- |
| `ObjectACLOptionsPrivate`                | private                                  |
| `ObjectACLOptionsPublicRead`             | public-read                              |
| `ObjectACLOptionsPublicReadWrite`        | public-read-write                        |
| `ObjectACLOptionsAuthenticatedRead`      | authenticated-read                       |
| `ObjectACLOptionsAwsExecRead`            | aws-exec-read                            |
| `ObjectACLOptionsBucketOwnerRead`        | bucket-owner-read                        |
| `ObjectACLOptionsBucketOwnerFullControl` | bucket-owner-full-control                |