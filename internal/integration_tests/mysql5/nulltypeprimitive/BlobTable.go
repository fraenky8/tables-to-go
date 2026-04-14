package dto

type BlobTable struct {
	Col               *string `db:"col"`
	BlobDefConst      *string `db:"blob_def_const"`
	BlobDefFunc       *string `db:"blob_def_func"`
	BlobRef           *string `db:"blob_ref"`
	BlobNn            string  `db:"blob_nn"`
	BlobNnCheckCmp    string  `db:"blob_nn_check_cmp"`
	BlobNnCheckFn     string  `db:"blob_nn_check_fn"`
	BlobNnRef         string  `db:"blob_nn_ref"`
	BlobNnDefConst    string  `db:"blob_nn_def_const"`
	BlobNnDefFunc     string  `db:"blob_nn_def_func"`
	BlobCheck         *string `db:"blob_check"`
	BlobCheckRef      *string `db:"blob_check_ref"`
	BlobCheckDefConst *string `db:"blob_check_def_const"`
	BlobCheckDefFunc  *string `db:"blob_check_def_func"`
}
