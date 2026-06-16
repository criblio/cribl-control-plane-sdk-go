# DiffLineContext

Unchanged context line in a Git diff hunk.


## Fields

| Field                                                                            | Type                                                                             | Required                                                                         | Description                                                                      |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `Type`                                                                           | [components.DiffLineContextType](../../models/components/difflinecontexttype.md) | :heavy_check_mark:                                                               | Line change type. Always <code>context</code> for unchanged lines.               |
| `NewNumber`                                                                      | `int64`                                                                          | :heavy_check_mark:                                                               | Line number in the new file.                                                     |
| `OldNumber`                                                                      | `int64`                                                                          | :heavy_check_mark:                                                               | Line number in the original file.                                                |
| `Content`                                                                        | `string`                                                                         | :heavy_check_mark:                                                               | Full content of the line, including the diff prefix character.                   |