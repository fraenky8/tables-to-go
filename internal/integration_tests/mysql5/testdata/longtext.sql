DROP TABLE IF EXISTS longtext_ref CASCADE;
CREATE TABLE longtext_ref
(
    longtext_ref longtext NOT NULL,
    KEY          longtext_ref_len_key (longtext_ref(10))
);

DROP TABLE IF EXISTS longtext_table;
CREATE TABLE longtext_table
(
    col                      longtext,

    longtext_def_const       longtext,
    longtext_def_func        longtext,

    longtext_ref             longtext REFERENCES longtext_ref (longtext_ref),

    longtext_nn              longtext NOT NULL,
    longtext_nn_check_cmp    longtext NOT NULL,
    longtext_nn_check_fn     longtext NOT NULL,
    longtext_nn_ref          longtext NOT NULL REFERENCES longtext_ref (longtext_ref),
    longtext_nn_def_const    longtext NOT NULL,
    longtext_nn_def_func     longtext NOT NULL,

    longtext_check           longtext,
    longtext_check_ref       longtext REFERENCES longtext_ref (longtext_ref),
    longtext_check_def_const longtext,
    longtext_check_def_func  longtext
);
