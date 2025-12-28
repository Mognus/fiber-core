package crud

// CRUDProvider is the interface that all modules must implement
// to be manageable through the admin panel
type CRUDProvider interface {
	// Model Info
	GetModelName() string // e.g. "users", "todos"
	GetSchema() Schema    // Describes fields for frontend

	// CRUD Operations
	List(filters map[string]string, page, limit int) (ListResponse, error)
	Get(id string) (any, error)
	Create(data map[string]any) (any, error)
	Update(id string, data map[string]any) (any, error)
	Delete(id string) error
}

// Schema describes the structure of a model for the frontend
type Schema struct {
	Name        string   `json:"name"`        // "users"
	DisplayName string   `json:"displayName"` // "Users"
	Fields      []Field  `json:"fields"`      // List of fields
	Filterable  []string `json:"filterable"`  // Fields that can be filtered
	Searchable  []string `json:"searchable"`  // Fields that are searchable
}

// Field describes a single field in the model
type Field struct {
	Name       string   `json:"name"`       // "email"
	Type       string   `json:"type"`       // "string", "number", "boolean", "date", "enum"
	Label      string   `json:"label"`      // "Email Address"
	Required   bool     `json:"required"`   // Is this field required?
	Readonly   bool     `json:"readonly"`   // Is this field readonly?
	Editable   bool     `json:"editable"`   // Is this field editable after creation?
	EnumValues []string `json:"enumValues"` // Possible values for enum fields
}

// ListResponse is the response format for list operations
type ListResponse struct {
	Items []any `json:"items"` // The data items
	Total int64 `json:"total"` // Total count of items
	Page  int   `json:"page"`  // Current page
	Limit int   `json:"limit"` // Items per page
}
