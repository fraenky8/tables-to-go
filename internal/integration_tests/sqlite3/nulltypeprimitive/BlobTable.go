package dto

type BlobTable struct {
	B            *string `db:"b"`
	BlobNn       string  `db:"blob_nn"`
	BlobUnique   *string `db:"blob_unique"`
	BlobCheck    *string `db:"blob_check"`
	BlobRef      *string `db:"blob_ref"`
	BlobDefConst *string `db:"blob_def_const"`
	BlobPk       *string `db:"blob_pk"`
}
