package schema

type Field struct {
	DBName          string
	DBType          string
	PrimaryKey      bool
	AutoIncrement   bool
	HasDefaultValue bool
	DefaultValue    string
	NotNull         bool
	Unique          bool
	UniqueName      string
	UniqueIndex     bool
	UniqueIndexName string
	Comment         string
	GoType          FieldGoType
}

type FieldGoType string

const (
	Bool   FieldGoType = "bool"
	Int    FieldGoType = "int"
	Uint   FieldGoType = "uint"
	Float  FieldGoType = "float"
	String FieldGoType = "string"
	Time   FieldGoType = "time"
)
