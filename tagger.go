package tablestogo

// Tagger interface for types of struct-tages
type Tagger interface {
	GenerateTag(db Database, column Column) string
}
