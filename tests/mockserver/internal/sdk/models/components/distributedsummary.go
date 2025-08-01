// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

type DistributedSummaryGroups struct {
	Count        float64 `json:"count"`
	Destinations float64 `json:"destinations"`
	Pipelines    float64 `json:"pipelines"`
	Routes       float64 `json:"routes"`
	Sources      float64 `json:"sources"`
}

func (o *DistributedSummaryGroups) GetCount() float64 {
	if o == nil {
		return 0.0
	}
	return o.Count
}

func (o *DistributedSummaryGroups) GetDestinations() float64 {
	if o == nil {
		return 0.0
	}
	return o.Destinations
}

func (o *DistributedSummaryGroups) GetPipelines() float64 {
	if o == nil {
		return 0.0
	}
	return o.Pipelines
}

func (o *DistributedSummaryGroups) GetRoutes() float64 {
	if o == nil {
		return 0.0
	}
	return o.Routes
}

func (o *DistributedSummaryGroups) GetSources() float64 {
	if o == nil {
		return 0.0
	}
	return o.Sources
}

type DistributedSummaryWorkers struct {
	Alive             float64 `json:"alive"`
	ConfVersions      float64 `json:"confVersions"`
	Count             float64 `json:"count"`
	DisconnectedCount float64 `json:"disconnectedCount"`
	Groups            float64 `json:"groups"`
	SoftwareVersions  float64 `json:"softwareVersions"`
	Unhealthy         float64 `json:"unhealthy"`
}

func (o *DistributedSummaryWorkers) GetAlive() float64 {
	if o == nil {
		return 0.0
	}
	return o.Alive
}

func (o *DistributedSummaryWorkers) GetConfVersions() float64 {
	if o == nil {
		return 0.0
	}
	return o.ConfVersions
}

func (o *DistributedSummaryWorkers) GetCount() float64 {
	if o == nil {
		return 0.0
	}
	return o.Count
}

func (o *DistributedSummaryWorkers) GetDisconnectedCount() float64 {
	if o == nil {
		return 0.0
	}
	return o.DisconnectedCount
}

func (o *DistributedSummaryWorkers) GetGroups() float64 {
	if o == nil {
		return 0.0
	}
	return o.Groups
}

func (o *DistributedSummaryWorkers) GetSoftwareVersions() float64 {
	if o == nil {
		return 0.0
	}
	return o.SoftwareVersions
}

func (o *DistributedSummaryWorkers) GetUnhealthy() float64 {
	if o == nil {
		return 0.0
	}
	return o.Unhealthy
}

type DistributedSummary struct {
	Groups  DistributedSummaryGroups  `json:"groups"`
	Workers DistributedSummaryWorkers `json:"workers"`
}

func (o *DistributedSummary) GetGroups() DistributedSummaryGroups {
	if o == nil {
		return DistributedSummaryGroups{}
	}
	return o.Groups
}

func (o *DistributedSummary) GetWorkers() DistributedSummaryWorkers {
	if o == nil {
		return DistributedSummaryWorkers{}
	}
	return o.Workers
}
