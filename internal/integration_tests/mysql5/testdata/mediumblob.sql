DROP TABLE IF EXISTS mediumblob_ref CASCADE;
CREATE TABLE mediumblob_ref
(
    mediumblob_ref mediumblob NOT NULL,
    KEY            mediumblob_ref_len_key (mediumblob_ref(10))
);

DROP TABLE IF EXISTS mediumblob_table;
CREATE TABLE mediumblob_table
(
    col                        mediumblob,

    mediumblob_def_const       mediumblob,
    mediumblob_def_func        mediumblob,

    mediumblob_ref             mediumblob REFERENCES mediumblob_ref (mediumblob_ref),

    mediumblob_nn              mediumblob NOT NULL,
    mediumblob_nn_check_cmp    mediumblob NOT NULL,
    mediumblob_nn_check_fn     mediumblob NOT NULL,
    mediumblob_nn_ref          mediumblob NOT NULL REFERENCES mediumblob_ref (mediumblob_ref),
    mediumblob_nn_def_const    mediumblob NOT NULL,
    mediumblob_nn_def_func     mediumblob NOT NULL,

    mediumblob_check           mediumblob,
    mediumblob_check_ref       mediumblob REFERENCES mediumblob_ref (mediumblob_ref),
    mediumblob_check_def_const mediumblob,
    mediumblob_check_def_func  mediumblob
);
