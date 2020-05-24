DROP TABLE IF EXISTS mediumtext_ref CASCADE;
CREATE TABLE mediumtext_ref (
    mediumtext_ref mediumtext NOT NULL,
    KEY mediumtext_ref_len_key (mediumtext_ref(10))
);

DROP TABLE IF EXISTS mediumtext_table;
CREATE TABLE mediumtext_table (
    col mediumtext,

    mediumtext_def_const mediumtext DEFAULT ('42'),
    mediumtext_def_func mediumtext DEFAULT (pi()),

    mediumtext_ref mediumtext REFERENCES mediumtext_ref(mediumtext_ref),

    mediumtext_nn mediumtext NOT NULL,
    mediumtext_nn_check_cmp mediumtext NOT NULL CHECK ( mediumtext_nn_check_cmp = '42' ),
    mediumtext_nn_check_fn mediumtext NOT NULL CHECK ( length(mediumtext_nn_check_fn) > 0 ),
    mediumtext_nn_ref mediumtext NOT NULL REFERENCES mediumtext_ref(mediumtext_ref),
    mediumtext_nn_def_const mediumtext NOT NULL DEFAULT ('42'),
    mediumtext_nn_def_func mediumtext NOT NULL DEFAULT (pi()),

    mediumtext_check mediumtext CHECK ( length(mediumtext_check) > 0 ),
    mediumtext_check_ref mediumtext CHECK ( length(mediumtext_check_ref) > 0 ) REFERENCES mediumtext_ref(mediumtext_ref),
    mediumtext_check_def_const mediumtext CHECK ( length(mediumtext_check_def_const) > 0 ) DEFAULT ('42'),
    mediumtext_check_def_func mediumtext CHECK ( length(mediumtext_check_def_func) > 0 ) DEFAULT (pi())
);
