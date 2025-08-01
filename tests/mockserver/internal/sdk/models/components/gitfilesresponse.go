// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

type GitFilesResponse struct {
	CommitMessage map[string]any `json:"commitMessage"`
	Count         float64        `json:"count"`
	Items         []GitFile      `json:"items"`
}

func (o *GitFilesResponse) GetCommitMessage() map[string]any {
	if o == nil {
		return map[string]any{}
	}
	return o.CommitMessage
}

func (o *GitFilesResponse) GetCount() float64 {
	if o == nil {
		return 0.0
	}
	return o.Count
}

func (o *GitFilesResponse) GetItems() []GitFile {
	if o == nil {
		return []GitFile{}
	}
	return o.Items
}
