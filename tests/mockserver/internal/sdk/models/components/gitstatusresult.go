// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

type File struct {
	Index      string `json:"index"`
	Path       string `json:"path"`
	WorkingDir string `json:"working_dir"`
}

func (o *File) GetIndex() string {
	if o == nil {
		return ""
	}
	return o.Index
}

func (o *File) GetPath() string {
	if o == nil {
		return ""
	}
	return o.Path
}

func (o *File) GetWorkingDir() string {
	if o == nil {
		return ""
	}
	return o.WorkingDir
}

type Renamed struct {
	From string `json:"from"`
	To   string `json:"to"`
}

func (o *Renamed) GetFrom() string {
	if o == nil {
		return ""
	}
	return o.From
}

func (o *Renamed) GetTo() string {
	if o == nil {
		return ""
	}
	return o.To
}

type GitStatusResult struct {
	Ahead      float64   `json:"ahead"`
	Behind     float64   `json:"behind"`
	Conflicted []string  `json:"conflicted"`
	Created    []string  `json:"created"`
	Current    string    `json:"current"`
	Deleted    []string  `json:"deleted"`
	Files      []File    `json:"files"`
	Modified   []string  `json:"modified"`
	NotAdded   []string  `json:"not_added"`
	Renamed    []Renamed `json:"renamed"`
	Staged     []string  `json:"staged"`
	Tracking   string    `json:"tracking"`
}

func (o *GitStatusResult) GetAhead() float64 {
	if o == nil {
		return 0.0
	}
	return o.Ahead
}

func (o *GitStatusResult) GetBehind() float64 {
	if o == nil {
		return 0.0
	}
	return o.Behind
}

func (o *GitStatusResult) GetConflicted() []string {
	if o == nil {
		return []string{}
	}
	return o.Conflicted
}

func (o *GitStatusResult) GetCreated() []string {
	if o == nil {
		return []string{}
	}
	return o.Created
}

func (o *GitStatusResult) GetCurrent() string {
	if o == nil {
		return ""
	}
	return o.Current
}

func (o *GitStatusResult) GetDeleted() []string {
	if o == nil {
		return []string{}
	}
	return o.Deleted
}

func (o *GitStatusResult) GetFiles() []File {
	if o == nil {
		return []File{}
	}
	return o.Files
}

func (o *GitStatusResult) GetModified() []string {
	if o == nil {
		return []string{}
	}
	return o.Modified
}

func (o *GitStatusResult) GetNotAdded() []string {
	if o == nil {
		return []string{}
	}
	return o.NotAdded
}

func (o *GitStatusResult) GetRenamed() []Renamed {
	if o == nil {
		return []Renamed{}
	}
	return o.Renamed
}

func (o *GitStatusResult) GetStaged() []string {
	if o == nil {
		return []string{}
	}
	return o.Staged
}

func (o *GitStatusResult) GetTracking() string {
	if o == nil {
		return ""
	}
	return o.Tracking
}
