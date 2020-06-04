/*
 * 3DS OUTSCALE API
 *
 * Welcome to the 3DS OUTSCALE's API documentation.<br /><br />  The 3DS OUTSCALE API enables you to manage your resources in the 3DS OUTSCALE Cloud. This documentation describes the different actions available along with code examples.<br /><br />  Note that the 3DS OUTSCALE Cloud is compatible with Amazon Web Services (AWS) APIs, but some resources have different names in AWS than in the 3DS OUTSCALE API. You can find a list of the differences [here](https://wiki.outscale.net/display/EN/3DS+OUTSCALE+APIs+Reference).<br /><br />  You can also manage your resources using the [Cockpit](https://wiki.outscale.net/display/EN/About+Cockpit) web interface.
 *
 * API version: 1.1
 * Contact: support@outscale.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package oscgo

import (
	"bytes"
	"encoding/json"
)

// ImageExportTask Information about the OMI export task.
type ImageExportTask struct {
	// If the OMI export task fails, an error message appears.
	Comment *string `json:"Comment,omitempty"`
	// The ID of the OMI to be exported.
	ImageId   *string    `json:"ImageId,omitempty"`
	OsuExport *OsuExport `json:"OsuExport,omitempty"`
	// The progress of the OMI export task, as a percentage.
	Progress *int64 `json:"Progress,omitempty"`
	// The state of the OMI export task (`pending/queued` \\| `pending` \\| `completed` \\| `failed` \\| `cancelled`).
	State *string `json:"State,omitempty"`
	// One or more tags associated with the image export task.
	Tags *[]ResourceTag `json:"Tags,omitempty"`
	// The ID of the OMI export task.
	TaskId *string `json:"TaskId,omitempty"`
}

// GetComment returns the Comment field value if set, zero value otherwise.
func (o *ImageExportTask) GetComment() string {
	if o == nil || o.Comment == nil {
		var ret string
		return ret
	}
	return *o.Comment
}

// GetCommentOk returns a tuple with the Comment field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *ImageExportTask) GetCommentOk() (string, bool) {
	if o == nil || o.Comment == nil {
		var ret string
		return ret, false
	}
	return *o.Comment, true
}

// HasComment returns a boolean if a field has been set.
func (o *ImageExportTask) HasComment() bool {
	if o != nil && o.Comment != nil {
		return true
	}

	return false
}

// SetComment gets a reference to the given string and assigns it to the Comment field.
func (o *ImageExportTask) SetComment(v string) {
	o.Comment = &v
}

// GetImageId returns the ImageId field value if set, zero value otherwise.
func (o *ImageExportTask) GetImageId() string {
	if o == nil || o.ImageId == nil {
		var ret string
		return ret
	}
	return *o.ImageId
}

// GetImageIdOk returns a tuple with the ImageId field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *ImageExportTask) GetImageIdOk() (string, bool) {
	if o == nil || o.ImageId == nil {
		var ret string
		return ret, false
	}
	return *o.ImageId, true
}

// HasImageId returns a boolean if a field has been set.
func (o *ImageExportTask) HasImageId() bool {
	if o != nil && o.ImageId != nil {
		return true
	}

	return false
}

// SetImageId gets a reference to the given string and assigns it to the ImageId field.
func (o *ImageExportTask) SetImageId(v string) {
	o.ImageId = &v
}

// GetOsuExport returns the OsuExport field value if set, zero value otherwise.
func (o *ImageExportTask) GetOsuExport() OsuExport {
	if o == nil || o.OsuExport == nil {
		var ret OsuExport
		return ret
	}
	return *o.OsuExport
}

// GetOsuExportOk returns a tuple with the OsuExport field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *ImageExportTask) GetOsuExportOk() (OsuExport, bool) {
	if o == nil || o.OsuExport == nil {
		var ret OsuExport
		return ret, false
	}
	return *o.OsuExport, true
}

// HasOsuExport returns a boolean if a field has been set.
func (o *ImageExportTask) HasOsuExport() bool {
	if o != nil && o.OsuExport != nil {
		return true
	}

	return false
}

// SetOsuExport gets a reference to the given OsuExport and assigns it to the OsuExport field.
func (o *ImageExportTask) SetOsuExport(v OsuExport) {
	o.OsuExport = &v
}

// GetProgress returns the Progress field value if set, zero value otherwise.
func (o *ImageExportTask) GetProgress() int64 {
	if o == nil || o.Progress == nil {
		var ret int64
		return ret
	}
	return *o.Progress
}

// GetProgressOk returns a tuple with the Progress field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *ImageExportTask) GetProgressOk() (int64, bool) {
	if o == nil || o.Progress == nil {
		var ret int64
		return ret, false
	}
	return *o.Progress, true
}

// HasProgress returns a boolean if a field has been set.
func (o *ImageExportTask) HasProgress() bool {
	if o != nil && o.Progress != nil {
		return true
	}

	return false
}

// SetProgress gets a reference to the given int64 and assigns it to the Progress field.
func (o *ImageExportTask) SetProgress(v int64) {
	o.Progress = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *ImageExportTask) GetState() string {
	if o == nil || o.State == nil {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *ImageExportTask) GetStateOk() (string, bool) {
	if o == nil || o.State == nil {
		var ret string
		return ret, false
	}
	return *o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *ImageExportTask) HasState() bool {
	if o != nil && o.State != nil {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *ImageExportTask) SetState(v string) {
	o.State = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *ImageExportTask) GetTags() []ResourceTag {
	if o == nil || o.Tags == nil {
		var ret []ResourceTag
		return ret
	}
	return *o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *ImageExportTask) GetTagsOk() ([]ResourceTag, bool) {
	if o == nil || o.Tags == nil {
		var ret []ResourceTag
		return ret, false
	}
	return *o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *ImageExportTask) HasTags() bool {
	if o != nil && o.Tags != nil {
		return true
	}

	return false
}

// SetTags gets a reference to the given []ResourceTag and assigns it to the Tags field.
func (o *ImageExportTask) SetTags(v []ResourceTag) {
	o.Tags = &v
}

// GetTaskId returns the TaskId field value if set, zero value otherwise.
func (o *ImageExportTask) GetTaskId() string {
	if o == nil || o.TaskId == nil {
		var ret string
		return ret
	}
	return *o.TaskId
}

// GetTaskIdOk returns a tuple with the TaskId field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *ImageExportTask) GetTaskIdOk() (string, bool) {
	if o == nil || o.TaskId == nil {
		var ret string
		return ret, false
	}
	return *o.TaskId, true
}

// HasTaskId returns a boolean if a field has been set.
func (o *ImageExportTask) HasTaskId() bool {
	if o != nil && o.TaskId != nil {
		return true
	}

	return false
}

// SetTaskId gets a reference to the given string and assigns it to the TaskId field.
func (o *ImageExportTask) SetTaskId(v string) {
	o.TaskId = &v
}

type NullableImageExportTask struct {
	Value        ImageExportTask
	ExplicitNull bool
}

func (v NullableImageExportTask) MarshalJSON() ([]byte, error) {
	switch {
	case v.ExplicitNull:
		return []byte("null"), nil
	default:
		return json.Marshal(v.Value)
	}
}

func (v *NullableImageExportTask) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}
