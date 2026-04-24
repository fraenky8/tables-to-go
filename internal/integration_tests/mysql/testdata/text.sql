DROP TABLE IF EXISTS text_table;
CREATE TABLE text_table
(
    col                  text,

    text_def_const       text                                            DEFAULT ('42'),
    text_def_func        text                                            DEFAULT (pi()),

    text_nn              text NOT NULL,
    text_nn_check_cmp    text NOT NULL CHECK ( text_nn_check_cmp = '42' ),
    text_nn_check_fn     text NOT NULL CHECK ( length(text_nn_check_fn) > 0 ),
    text_nn_def_const    text NOT NULL                                   DEFAULT ('42'),
    text_nn_def_func     text NOT NULL                                   DEFAULT (pi()),

    text_check           text CHECK ( length(text_check) > 0 ),
    text_check_def_const text CHECK ( length(text_check_def_const) > 0 ) DEFAULT ('42'),
    text_check_def_func  text CHECK ( length(text_check_def_func) > 0 )  DEFAULT (pi())
);
