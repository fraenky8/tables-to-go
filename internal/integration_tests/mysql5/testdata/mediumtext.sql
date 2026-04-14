DROP TABLE IF EXISTS mediumtext_ref CASCADE;
CREATE TABLE mediumtext_ref
(
    mediumtext_ref mediumtext NOT NULL,
    KEY            mediumtext_ref_len_key (mediumtext_ref(10))
);

DROP TABLE IF EXISTS mediumtext_table;
CREATE TABLE mediumtext_table
(
    col                        mediumtext,

    mediumtext_def_const       mediumtext,
    mediumtext_def_func        mediumtext,

    mediumtext_ref             mediumtext REFERENCES mediumtext_ref (mediumtext_ref),

    mediumtext_nn              mediumtext NOT NULL,
    mediumtext_nn_check_cmp    mediumtext NOT NULL,
    mediumtext_nn_check_fn     mediumtext NOT NULL,
    mediumtext_nn_ref          mediumtext NOT NULL REFERENCES mediumtext_ref (mediumtext_ref),
    mediumtext_nn_def_const    mediumtext NOT NULL,
    mediumtext_nn_def_func     mediumtext NOT NULL,

    mediumtext_check           mediumtext,
    mediumtext_check_ref       mediumtext REFERENCES mediumtext_ref (mediumtext_ref),
    mediumtext_check_def_const mediumtext,
    mediumtext_check_def_func  mediumtext
);
