DROP TABLE IF EXISTS tinytext_ref CASCADE;
CREATE TABLE tinytext_ref (
    tinytext_ref tinytext NOT NULL,
    KEY tinytext_ref_len_key (tinytext_ref(10))
);

DROP TABLE IF EXISTS tinytext_table;
CREATE TABLE tinytext_table (
    col tinytext,

    tinytext_def_const tinytext DEFAULT ('42'),
    tinytext_def_func tinytext DEFAULT (pi()),

    tinytext_ref tinytext REFERENCES tinytext_ref(tinytext_ref),

    tinytext_nn tinytext NOT NULL,
    tinytext_nn_check_cmp tinytext NOT NULL CHECK ( tinytext_nn_check_cmp = '42' ),
    tinytext_nn_check_fn tinytext NOT NULL CHECK ( length(tinytext_nn_check_fn) > 0 ),
    tinytext_nn_ref tinytext NOT NULL REFERENCES tinytext_ref(tinytext_ref),
    tinytext_nn_def_const tinytext NOT NULL DEFAULT ('42'),
    tinytext_nn_def_func tinytext NOT NULL DEFAULT (pi()),

    tinytext_check tinytext CHECK ( length(tinytext_check) > 0 ),
    tinytext_check_ref tinytext CHECK ( length(tinytext_check_ref) > 0 ) REFERENCES tinytext_ref(tinytext_ref),
    tinytext_check_def_const tinytext CHECK ( length(tinytext_check_def_const) > 0 ) DEFAULT ('42'),
    tinytext_check_def_func tinytext CHECK ( length(tinytext_check_def_func) > 0 ) DEFAULT (pi())
);
