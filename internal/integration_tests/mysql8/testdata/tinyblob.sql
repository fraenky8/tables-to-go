DROP TABLE IF EXISTS tinyblob_ref CASCADE;
CREATE TABLE tinyblob_ref (
    tinyblob_ref tinyblob NOT NULL,
    KEY tinyblob_ref_len_key (tinyblob_ref(10))
);

DROP TABLE IF EXISTS tinyblob_table;
CREATE TABLE tinyblob_table (
    col tinyblob,

    tinyblob_def_const tinyblob DEFAULT ('42'),
    tinyblob_def_func tinyblob DEFAULT (pi()),

    tinyblob_ref tinyblob REFERENCES tinyblob_ref(tinyblob_ref),

    tinyblob_nn tinyblob NOT NULL,
    tinyblob_nn_check_cmp tinyblob NOT NULL CHECK ( tinyblob_nn_check_cmp = '42' ),
    tinyblob_nn_check_fn tinyblob NOT NULL CHECK ( length(tinyblob_nn_check_fn) > 0 ),
    tinyblob_nn_ref tinyblob NOT NULL REFERENCES tinyblob_ref(tinyblob_ref),
    tinyblob_nn_def_const tinyblob NOT NULL DEFAULT ('42'),
    tinyblob_nn_def_func tinyblob NOT NULL DEFAULT (pi()),

    tinyblob_check tinyblob CHECK ( length(tinyblob_check) > 0 ),
    tinyblob_check_ref tinyblob CHECK ( length(tinyblob_check_ref) > 0 ) REFERENCES tinyblob_ref(tinyblob_ref),
    tinyblob_check_def_const tinyblob CHECK ( length(tinyblob_check_def_const) > 0 ) DEFAULT ('42'),
    tinyblob_check_def_func tinyblob CHECK ( length(tinyblob_check_def_func) > 0 ) DEFAULT (pi())
);
