# FunctionConfSchemaHandlebarsTemplateDefinition


## Fields

| Field                                                                | Type                                                                 | Required                                                             | Description                                                          |
| -------------------------------------------------------------------- | -------------------------------------------------------------------- | -------------------------------------------------------------------- | -------------------------------------------------------------------- |
| `ID`                                                                 | *string*                                                             | :heavy_check_mark:                                                   | Unique identifier for this template                                  |
| `Content`                                                            | *string*                                                             | :heavy_check_mark:                                                   | Handlebars template string                                           |
| `Description`                                                        | **string*                                                            | :heavy_minus_sign:                                                   | Optional description of what this template is used for               |
| `Type`                                                               | **string*                                                            | :heavy_minus_sign:                                                   | Type categorization for the template (e.g., Universal, Email, Slack) |