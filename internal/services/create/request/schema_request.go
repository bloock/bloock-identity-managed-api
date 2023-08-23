package request

type CreateSchemaRequest struct {
	Attributes  []AttributeSchema
	DisplayName string
	SchemaType  string
	Version     string
	Description string
}

type AttributeSchema struct {
	Id          string
	Name        string
	DataType    string
	Description string
	Required    bool
}
