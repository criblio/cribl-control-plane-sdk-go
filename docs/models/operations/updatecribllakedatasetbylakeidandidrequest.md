# UpdateCriblLakeDatasetByLakeIDAndIDRequest


## Fields

| Field                                                                      | Type                                                                       | Required                                                                   | Description                                                                |
| -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- |
| `LakeID`                                                                   | *string*                                                                   | :heavy_check_mark:                                                         | lake id that contains the Datasets                                         |
| `ID`                                                                       | *string*                                                                   | :heavy_check_mark:                                                         | dataset id to update                                                       |
| `CriblLakeDataset`                                                         | [components.CriblLakeDataset](../../models/components/cribllakedataset.md) | :heavy_check_mark:                                                         | CriblLakeDataset object                                                    |