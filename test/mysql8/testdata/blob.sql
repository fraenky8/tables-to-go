DROP TABLE IF EXISTS blob_ref CASCADE;
CREATE TABLE blob_ref (
    blob_ref blob NOT NULL,
    KEY blob_ref_len_key (blob_ref(10))
);

DROP TABLE IF EXISTS blob_table;
CREATE TABLE blob_table (
    col blob,

    blob_def_const blob DEFAULT ('42'),
    blob_def_func blob DEFAULT (pi()),

    blob_ref blob REFERENCES blob_ref(blob_ref),

    blob_nn blob NOT NULL,
    blob_nn_check_cmp blob NOT NULL CHECK ( blob_nn_check_cmp = '42' ),
    blob_nn_check_fn blob NOT NULL CHECK ( length(blob_nn_check_fn) > 0 ),
    blob_nn_ref blob NOT NULL REFERENCES blob_ref(blob_ref),
    blob_nn_def_const blob NOT NULL DEFAULT ('42'),
    blob_nn_def_func blob NOT NULL DEFAULT (pi()),

    blob_check blob CHECK ( length(blob_check) > 0 ),
    blob_check_ref blob CHECK ( length(blob_check_ref) > 0 ) REFERENCES blob_ref(blob_ref),
    blob_check_def_const blob CHECK ( length(blob_check_def_const) > 0 ) DEFAULT ('42'),
    blob_check_def_func blob CHECK ( length(blob_check_def_func) > 0 ) DEFAULT (pi())
);
