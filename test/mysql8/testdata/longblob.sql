DROP TABLE IF EXISTS longblob_ref CASCADE;
CREATE TABLE longblob_ref (
    longblob_ref longblob NOT NULL,
    KEY longblob_ref_len_key (longblob_ref(10))
);

DROP TABLE IF EXISTS longblob_table;
CREATE TABLE longblob_table (
    col longblob,

    longblob_def_const longblob DEFAULT ('42'),
    longblob_def_func longblob DEFAULT (pi()),

    longblob_ref longblob REFERENCES longblob_ref(longblob_ref),

    longblob_nn longblob NOT NULL,
    longblob_nn_check_cmp longblob NOT NULL CHECK ( longblob_nn_check_cmp = '42' ),
    longblob_nn_check_fn longblob NOT NULL CHECK ( length(longblob_nn_check_fn) > 0 ),
    longblob_nn_ref longblob NOT NULL REFERENCES longblob_ref(longblob_ref),
    longblob_nn_def_const longblob NOT NULL DEFAULT ('42'),
    longblob_nn_def_func longblob NOT NULL DEFAULT (pi()),

    longblob_check longblob CHECK ( length(longblob_check) > 0 ),
    longblob_check_ref longblob CHECK ( length(longblob_check_ref) > 0 ) REFERENCES longblob_ref(longblob_ref),
    longblob_check_def_const longblob CHECK ( length(longblob_check_def_const) > 0 ) DEFAULT ('42'),
    longblob_check_def_func longblob CHECK ( length(longblob_check_def_func) > 0 ) DEFAULT (pi())
);
