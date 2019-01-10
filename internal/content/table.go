package content

var _ Content = (*Table)(nil)

type Table struct {
	Type         string        `json:"type,omitempty"`
	Title        string        `json:"title,omitempty"`
	Columns      []TableColumn `json:"columns,omitempty"`
	Rows         []TableRow    `json:"rows"`
	EmptyContent string        `json:"empty_content,omitempty"`
}

func NewTable(title, emptyContent string) Table {
	return Table{
		Type:         "table",
		Title:        title,
		EmptyContent: emptyContent,
	}
}

func (t *Table) IsEmpty() bool {
	return len(t.Rows) == 0
}

func (t *Table) ColumnNames() []string {
	var names []string
	for _, col := range t.Columns {
		names = append(names, col.Name)
	}

	return names
}

func (t *Table) AddRow(row TableRow) {
	t.Rows = append(t.Rows, row)
}

func (t *Table) ViewComponent() ViewComponent {
	tc := TableConfig{
		Columns:      t.Columns,
		Rows:         t.Rows,
		EmptyContent: t.EmptyContent,
	}

	return ViewComponent{
		Metadata: Metadata{
			Type:  "table",
			Title: t.Title,
		},
		Config: tc,
	}
}

type TableConfig struct {
	Columns      []TableColumn `json:"columns,omitempty"`
	Rows         []TableRow    `json:"rows,omitempty"`
	EmptyContent string        `json:"emptyContent,omitempty"`
}

type TableColumn struct {
	Name     string `json:"name,omitempty"`
	Accessor string `json:"accessor,omitempty"`
}

type TableRow map[string]Text
