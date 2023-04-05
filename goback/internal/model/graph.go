package model

type CellData struct {
	Description string `json:"Description,omitempty"`
	LeadTime    uint64 `json:"LeadTime,omitempty"`
	Performers  string `json:"Performers,omitempty"`
	Energy      uint64 `json:"Energy,omitempty"`
	Status      string `json:"Status,omitempty"`
}

type Cell struct {
	Id       string         `json:"id,omitempty"`
	Shape    string         `json:"shape,omitempty"`
	Data     CellData       `json:"data,omitempty"`
	Attrs    map[string]any `json:"attrs,omitempty"`
	ZIndex   any            `json:"zIndex,omitempty"`
	Visible  bool           `json:"visible,omitempty"`
	Size     map[string]any `json:"size,omitempty"`
	Position map[string]any `json:"position,omitempty"`
	Ports    map[string]any `json:"ports,omitempty"`

	// for edge
	Source any `json:"source,omitempty"`
	Target any `json:"target,omitempty"`
}

type Graph struct {
	Length uint8  `json:"length,omitempty"`
	Cells  []Cell `json:"cells,omitempty"`
}
