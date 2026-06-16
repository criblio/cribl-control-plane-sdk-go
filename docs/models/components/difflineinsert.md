# DiffLineInsert

Inserted line in a Git diff hunk.


## Fields

| Field                                                                          | Type                                                                           | Required                                                                       | Description                                                                    |
| ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ |
| `Type`                                                                         | [components.DiffLineInsertType](../../models/components/difflineinserttype.md) | :heavy_check_mark:                                                             | Line change type. Always <code>insert</code> for inserted lines.               |
| `NewNumber`                                                                    | `int64`                                                                        | :heavy_check_mark:                                                             | Line number in the new file.                                                   |
| `Content`                                                                      | `string`                                                                       | :heavy_check_mark:                                                             | Full content of the line, including the diff prefix character.                 |