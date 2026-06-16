# Audit

Audit record for the revert operation, including the commit hash and affected files.


## Fields

| Field                                                                               | Type                                                                                | Required                                                                            | Description                                                                         |
| ----------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------- |
| `Files`                                                                             | [*components.GitRevertResultFiles](../../models/components/gitrevertresultfiles.md) | :heavy_minus_sign:                                                                  | Files affected by the revert, grouped by change type.                               |
| `Group`                                                                             | `*string`                                                                           | :heavy_minus_sign:                                                                  | Worker Group the revert was applied to, if applicable.                              |
| `ID`                                                                                | `string`                                                                            | :heavy_check_mark:                                                                  | SHA-1 hash of the revert commit that was created.                                   |