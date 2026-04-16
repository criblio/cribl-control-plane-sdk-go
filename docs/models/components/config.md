# Config

Configuration bundle and policy revision metadata for this node.


## Fields

| Field                                                             | Type                                                              | Required                                                          | Description                                                       |
| ----------------------------------------------------------------- | ----------------------------------------------------------------- | ----------------------------------------------------------------- | ----------------------------------------------------------------- |
| `FeaturesRev`                                                     | `*string`                                                         | :heavy_minus_sign:                                                | Feature flags or feature revision string for this bundle.         |
| `HbPeriodSeconds`                                                 | `*int64`                                                          | :heavy_minus_sign:                                                | Worker-to-Leader heartbeat interval, in seconds.                  |
| `LogStreamEnv`                                                    | `*string`                                                         | :heavy_minus_sign:                                                | GitOps or LogStream environment label associated with the bundle. |
| `PolicyRev`                                                       | `*string`                                                         | :heavy_minus_sign:                                                | Current policies revision string.                                 |
| `Version`                                                         | `*string`                                                         | :heavy_minus_sign:                                                | Configuration bundle version.                                     |