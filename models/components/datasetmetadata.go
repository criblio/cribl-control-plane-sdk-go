// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

import (
	"encoding/json"
	"fmt"
)

type ScanMode string

const (
	ScanModeDetailed ScanMode = "detailed"
	ScanModeQuick    ScanMode = "quick"
)

func (e ScanMode) ToPointer() *ScanMode {
	return &e
}
func (e *ScanMode) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "detailed":
		fallthrough
	case "quick":
		*e = ScanMode(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ScanMode: %v", v)
	}
}

type DatasetMetadata struct {
	Earliest           string                  `json:"earliest"`
	EnableAcceleration bool                    `json:"enableAcceleration"`
	FieldList          []string                `json:"fieldList"`
	LatestRunInfo      *DatasetMetadataRunInfo `json:"latestRunInfo,omitempty"`
	ScanMode           ScanMode                `json:"scanMode"`
}

func (o *DatasetMetadata) GetEarliest() string {
	if o == nil {
		return ""
	}
	return o.Earliest
}

func (o *DatasetMetadata) GetEnableAcceleration() bool {
	if o == nil {
		return false
	}
	return o.EnableAcceleration
}

func (o *DatasetMetadata) GetFieldList() []string {
	if o == nil {
		return []string{}
	}
	return o.FieldList
}

func (o *DatasetMetadata) GetLatestRunInfo() *DatasetMetadataRunInfo {
	if o == nil {
		return nil
	}
	return o.LatestRunInfo
}

func (o *DatasetMetadata) GetScanMode() ScanMode {
	if o == nil {
		return ScanMode("")
	}
	return o.ScanMode
}
