DROP TABLE IF EXISTS text_ref CASCADE;
CREATE TABLE text_ref
(
    text_ref text NOT NULL,
    KEY      text_ref_len_key (text_ref(10))
);

DROP TABLE IF EXISTS text_table;
CREATE TABLE text_table
(
    col                  text,

    text_def_const       text,
    text_def_func        text,

    text_ref             text REFERENCES text_ref (text_ref),

    text_nn              text NOT NULL,
    text_nn_check_cmp    text NOT NULL,
    text_nn_check_fn     text NOT NULL,
    text_nn_ref          text NOT NULL REFERENCES text_ref (text_ref),
    text_nn_def_const    text NOT NULL,
    text_nn_def_func     text NOT NULL,

    text_check           text,
    text_check_ref       text REFERENCES text_ref (text_ref),
    text_check_def_const text,
    text_check_def_func  text
);
