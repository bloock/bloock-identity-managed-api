package request

type CreateSchemaRequest struct {
	IssuerDID   string
	Attributes  []AttributeSchema
	Title       string
	SchemaType  string
	Version     string
	Description string
}

type AttributeSchema struct {
	Key         string
	Title       string
	DataType    string
	Description string
	Required    bool
}
