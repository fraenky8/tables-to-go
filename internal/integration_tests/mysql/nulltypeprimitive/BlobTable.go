package dto

type BlobTable struct {
	Col               *string `db:"col"`
	BlobDefConst      *string `db:"blob_def_const"`
	BlobDefFunc       *string `db:"blob_def_func"`
	BlobNn            string  `db:"blob_nn"`
	BlobNnCheckCmp    string  `db:"blob_nn_check_cmp"`
	BlobNnCheckFn     string  `db:"blob_nn_check_fn"`
	BlobNnDefConst    string  `db:"blob_nn_def_const"`
	BlobNnDefFunc     string  `db:"blob_nn_def_func"`
	BlobCheck         *string `db:"blob_check"`
	BlobCheckDefConst *string `db:"blob_check_def_const"`
	BlobCheckDefFunc  *string `db:"blob_check_def_func"`
}
