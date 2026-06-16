# DiffLineDelete

Deleted line in a Git diff hunk.


## Fields

| Field                                                                          | Type                                                                           | Required                                                                       | Description                                                                    |
| ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ |
| `Type`                                                                         | [components.DiffLineDeleteType](../../models/components/difflinedeletetype.md) | :heavy_check_mark:                                                             | Line change type. Always <code>delete</code> for deleted lines.                |
| `OldNumber`                                                                    | `int64`                                                                        | :heavy_check_mark:                                                             | Line number in the original file.                                              |
| `Content`                                                                      | `string`                                                                       | :heavy_check_mark:                                                             | Full content of the line, including the diff prefix character.                 |