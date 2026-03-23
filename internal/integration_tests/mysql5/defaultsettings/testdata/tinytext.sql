DROP TABLE IF EXISTS tinytext_ref CASCADE;
CREATE TABLE tinytext_ref
(
    tinytext_ref tinytext NOT NULL,
    KEY          tinytext_ref_len_key (tinytext_ref(10))
);

DROP TABLE IF EXISTS tinytext_table;
CREATE TABLE tinytext_table
(
    col                      tinytext,

    tinytext_def_const       tinytext,
    tinytext_def_func        tinytext,

    tinytext_ref             tinytext REFERENCES tinytext_ref (tinytext_ref),

    tinytext_nn              tinytext NOT NULL,
    tinytext_nn_check_cmp    tinytext NOT NULL,
    tinytext_nn_check_fn     tinytext NOT NULL,
    tinytext_nn_ref          tinytext NOT NULL REFERENCES tinytext_ref (tinytext_ref),
    tinytext_nn_def_const    tinytext NOT NULL,
    tinytext_nn_def_func     tinytext NOT NULL,

    tinytext_check           tinytext,
    tinytext_check_ref       tinytext REFERENCES tinytext_ref (tinytext_ref),
    tinytext_check_def_const tinytext,
    tinytext_check_def_func  tinytext
);
