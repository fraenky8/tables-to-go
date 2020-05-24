DROP TABLE IF EXISTS longtext_ref CASCADE;
CREATE TABLE longtext_ref (
    longtext_ref longtext NOT NULL,
    KEY longtext_ref_len_key (longtext_ref(10))
);

DROP TABLE IF EXISTS longtext_table;
CREATE TABLE longtext_table (
    col longtext,

    longtext_def_const longtext DEFAULT ('42'),
    longtext_def_func longtext DEFAULT (pi()),

    longtext_ref longtext REFERENCES longtext_ref(longtext_ref),

    longtext_nn longtext NOT NULL,
    longtext_nn_check_cmp longtext NOT NULL CHECK ( longtext_nn_check_cmp = '42' ),
    longtext_nn_check_fn longtext NOT NULL CHECK ( length(longtext_nn_check_fn) > 0 ),
    longtext_nn_ref longtext NOT NULL REFERENCES longtext_ref(longtext_ref),
    longtext_nn_def_const longtext NOT NULL DEFAULT ('42'),
    longtext_nn_def_func longtext NOT NULL DEFAULT (pi()),

    longtext_check longtext CHECK ( length(longtext_check) > 0 ),
    longtext_check_ref longtext CHECK ( length(longtext_check_ref) > 0 ) REFERENCES longtext_ref(longtext_ref),
    longtext_check_def_const longtext CHECK ( length(longtext_check_def_const) > 0 ) DEFAULT ('42'),
    longtext_check_def_func longtext CHECK ( length(longtext_check_def_func) > 0 ) DEFAULT (pi())
);
