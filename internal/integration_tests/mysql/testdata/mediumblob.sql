DROP TABLE IF EXISTS mediumblob_table;
CREATE TABLE mediumblob_table
(
    col                        mediumblob,

    mediumblob_def_const       mediumblob                                                  DEFAULT ('42'),
    mediumblob_def_func        mediumblob                                                  DEFAULT (pi()),

    mediumblob_nn              mediumblob NOT NULL,
    mediumblob_nn_check_cmp    mediumblob NOT NULL CHECK ( mediumblob_nn_check_cmp = '42' ),
    mediumblob_nn_check_fn     mediumblob NOT NULL CHECK ( length(mediumblob_nn_check_fn) > 0 ),
    mediumblob_nn_def_const    mediumblob NOT NULL                                         DEFAULT ('42'),
    mediumblob_nn_def_func     mediumblob NOT NULL                                         DEFAULT (pi()),

    mediumblob_check           mediumblob CHECK ( length(mediumblob_check) > 0 ),
    mediumblob_check_def_const mediumblob CHECK ( length(mediumblob_check_def_const) > 0 ) DEFAULT ('42'),
    mediumblob_check_def_func  mediumblob CHECK ( length(mediumblob_check_def_func) > 0 )  DEFAULT (pi())
);
