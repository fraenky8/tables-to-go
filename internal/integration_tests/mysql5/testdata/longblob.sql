DROP TABLE IF EXISTS longblob_ref CASCADE;
CREATE TABLE longblob_ref
(
    longblob_ref longblob NOT NULL,
    KEY          longblob_ref_len_key (longblob_ref(10))
);

DROP TABLE IF EXISTS longblob_table;
CREATE TABLE longblob_table
(
    col                      longblob,

    longblob_def_const       longblob,
    longblob_def_func        longblob,

    longblob_ref             longblob REFERENCES longblob_ref (longblob_ref),

    longblob_nn              longblob NOT NULL,
    longblob_nn_check_cmp    longblob NOT NULL,
    longblob_nn_check_fn     longblob NOT NULL,
    longblob_nn_ref          longblob NOT NULL REFERENCES longblob_ref (longblob_ref),
    longblob_nn_def_const    longblob NOT NULL,
    longblob_nn_def_func     longblob NOT NULL,

    longblob_check           longblob,
    longblob_check_ref       longblob REFERENCES longblob_ref (longblob_ref),
    longblob_check_def_const longblob,
    longblob_check_def_func  longblob
);
