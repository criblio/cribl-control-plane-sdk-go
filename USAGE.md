<!-- Start SDK Example Usage [usage] -->
```go
package main

import (
	"context"
	criblcontrolplanesdkgo "github.com/criblio/cribl-control-plane-sdk-go"
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
	"log"
	"os"
)

func main() {
	ctx := context.Background()

	s := criblcontrolplanesdkgo.New(
		"https://api.example.com",
		criblcontrolplanesdkgo.WithSecurity(components.Security{
			BearerAuth: criblcontrolplanesdkgo.Pointer(os.Getenv("CRIBLCONTROLPLANE_BEARER_AUTH")),
		}),
	)

	res, err := s.LakeDatasets.Create(ctx, "<id>", components.CriblLakeDataset{
		AcceleratedFields: []string{
			"<value 1>",
			"<value 2>",
		},
		BucketName: criblcontrolplanesdkgo.Pointer("<value>"),
		CacheConnection: &components.CacheConnection{
			AcceleratedFields: []string{
				"<value 1>",
				"<value 2>",
			},
			BackfillStatus:          components.CacheConnectionBackfillStatusPending.ToPointer(),
			CacheRef:                "<value>",
			CreatedAt:               7795.06,
			LakehouseConnectionType: components.LakehouseConnectionTypeCache.ToPointer(),
			MigrationQueryID:        criblcontrolplanesdkgo.Pointer("<id>"),
			RetentionInDays:         1466.58,
		},
		DeletionStartedAt: criblcontrolplanesdkgo.Pointer[float64](8310.58),
		Description:       criblcontrolplanesdkgo.Pointer("pleased toothbrush long brush smooth swiftly rightfully phooey chapel"),
		Format:            components.CriblLakeDatasetFormatDdss.ToPointer(),
		HTTPDAUsed:        criblcontrolplanesdkgo.Pointer(true),
		ID:                "<id>",
		Metrics: &components.LakeDatasetMetrics{
			CurrentSizeBytes: 6170.04,
			MetricsDate:      "<value>",
		},
		RetentionPeriodInDays: criblcontrolplanesdkgo.Pointer[float64](456.37),
		SearchConfig: &components.LakeDatasetSearchConfig{
			Datatypes: []string{
				"<value 1>",
			},
			Metadata: &components.DatasetMetadata{
				Earliest:           "<value>",
				EnableAcceleration: true,
				FieldList: []string{
					"<value 1>",
					"<value 2>",
				},
				LatestRunInfo: &components.DatasetMetadataRunInfo{
					EarliestScannedTime: criblcontrolplanesdkgo.Pointer[float64](4334.7),
					FinishedAt:          criblcontrolplanesdkgo.Pointer[float64](6811.22),
					LatestScannedTime:   criblcontrolplanesdkgo.Pointer[float64](5303.3),
					ObjectCount:         criblcontrolplanesdkgo.Pointer[float64](9489.04),
				},
				ScanMode: components.ScanModeDetailed,
			},
		},
		StorageLocationID: criblcontrolplanesdkgo.Pointer("<id>"),
		ViewName:          criblcontrolplanesdkgo.Pointer("<value>"),
	})
	if err != nil {
		log.Fatal(err)
	}
	if res.CountedListCriblLakeDataset != nil {
		// handle response
	}
}

```
<!-- End SDK Example Usage [usage] -->