# UpdatePipelinesByPackAndIDRequest


## Fields

| Field                                                                | Type                                                                 | Required                                                             | Description                                                          |
| -------------------------------------------------------------------- | -------------------------------------------------------------------- | -------------------------------------------------------------------- | -------------------------------------------------------------------- |
| `ID`                                                                 | *string*                                                             | :heavy_check_mark:                                                   | The <code>id</code> of the Pipeline to update.                       |
| `Pack`                                                               | *string*                                                             | :heavy_check_mark:                                                   | The <code>id</code> of the Pack to update.                           |
| `Pipeline`                                                           | [components.PipelineInput](../../models/components/pipelineinput.md) | :heavy_check_mark:                                                   | Pipeline object                                                      |