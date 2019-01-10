package component

import "encoding/json"

// List contains other ViewComponents
type List struct {
	Metadata Metadata   `json:"metadata"`
	Config   ListConfig `json:"config"`
}

// ListConfig is the contents of a List
type ListConfig struct {
	Items []ViewComponent `json:"items"`
}

// NewList creates a list component
func NewList(title string, items []ViewComponent) *List {
	return &List{
		Metadata: Metadata{
			Type: "list",
		},
		Config: ListConfig{
			Items: items,
		},
	}
}

// GetMetadata accesses the components metadata. Implements ViewComponent.
func (t *List) GetMetadata() Metadata {
	return t.Metadata
}

// IsEmpty specifes whether the component is considered empty. Implements ViewComponent.
func (t *List) IsEmpty() bool {
	return len(t.Config.Items) == 0
}

// Add adds additional items to the tail of the list.
func (t *List) Add(items ...ViewComponent) {
	t.Config.Items = append(t.Config.Items, items...)
}

type listMarshal List

// MarshalJSON implements json.Marshaler
func (t *List) MarshalJSON() ([]byte, error) {
	m := listMarshal(*t)
	m.Metadata.Type = "list"
	return json.Marshal(&m)
}
