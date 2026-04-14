DROP TABLE IF EXISTS tinyblob_ref CASCADE;
CREATE TABLE tinyblob_ref
(
    tinyblob_ref tinyblob NOT NULL,
    KEY          tinyblob_ref_len_key (tinyblob_ref(10))
);

DROP TABLE IF EXISTS tinyblob_table;
CREATE TABLE tinyblob_table
(
    col                      tinyblob,

    tinyblob_def_const       tinyblob,
    tinyblob_def_func        tinyblob,

    tinyblob_ref             tinyblob REFERENCES tinyblob_ref (tinyblob_ref),

    tinyblob_nn              tinyblob NOT NULL,
    tinyblob_nn_check_cmp    tinyblob NOT NULL,
    tinyblob_nn_check_fn     tinyblob NOT NULL,
    tinyblob_nn_ref          tinyblob NOT NULL REFERENCES tinyblob_ref (tinyblob_ref),
    tinyblob_nn_def_const    tinyblob NOT NULL,
    tinyblob_nn_def_func     tinyblob NOT NULL,

    tinyblob_check           tinyblob,
    tinyblob_check_ref       tinyblob REFERENCES tinyblob_ref (tinyblob_ref),
    tinyblob_check_def_const tinyblob,
    tinyblob_check_def_func  tinyblob
);
