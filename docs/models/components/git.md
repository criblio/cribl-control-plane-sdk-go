# Git

Git status of the Worker Group, Outpost Group, or Edge Fleet configuration. Automatically populated and returned in responses.


## Fields

| Field                                                         | Type                                                          | Required                                                      | Description                                                   |
| ------------------------------------------------------------- | ------------------------------------------------------------- | ------------------------------------------------------------- | ------------------------------------------------------------- |
| `Commit`                                                      | `*string`                                                     | :heavy_minus_sign:                                            | Commit hash of the currently committed configuration version. |
| `LocalChanges`                                                | `*int64`                                                      | :heavy_minus_sign:                                            | Number of local configuration changes not yet committed.      |
| `Log`                                                         | [][components.Commit](../../models/components/commit.md)      | :heavy_minus_sign:                                            | List of recent configuration commits.                         |