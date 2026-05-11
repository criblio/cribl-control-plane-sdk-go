# InputCollectionOriginDataSourceDiscoveryWithDestinationArnConstraint

Read-only metadata that records how the Source was created. Preserved on update when omitted from the request body. Cannot be set on create.


## Fields

| Field                                                                      | Type                                                                       | Required                                                                   | Description                                                                |
| -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- |
| `Origin`                                                                   | [*components.Origin](../../models/components/origin.md)                    | :heavy_minus_sign:                                                         | Feature that created the Source.                                           |
| `DestinationArn`                                                           | `*string`                                                                  | :heavy_minus_sign:                                                         | ARN of the S3 bucket or Firehose delivery stream configured as the Source. |
| `SourceArn`                                                                | `*string`                                                                  | :heavy_minus_sign:                                                         | ARN of the AWS resource that produces the logs.                            |