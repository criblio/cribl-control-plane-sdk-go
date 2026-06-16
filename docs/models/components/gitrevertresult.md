# GitRevertResult


## Fields

| Field                                                                                     | Type                                                                                      | Required                                                                                  | Description                                                                               |
| ----------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------- |
| `Audit`                                                                                   | [components.Audit](../../models/components/audit.md)                                      | :heavy_check_mark:                                                                        | Audit record for the revert operation, including the commit hash and affected files.      |
| `Reverted`                                                                                | `bool`                                                                                    | :heavy_check_mark:                                                                        | If <code>true</code>, the revert was applied successfully. Otherwise, <code>false</code>. |