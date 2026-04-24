DROP TABLE IF EXISTS blob_table;
CREATE TABLE blob_table
(
    col                  blob,

    blob_def_const       blob                                            DEFAULT ('42'),
    blob_def_func        blob                                            DEFAULT (pi()),

    blob_nn              blob NOT NULL,
    blob_nn_check_cmp    blob NOT NULL CHECK ( blob_nn_check_cmp = '42' ),
    blob_nn_check_fn     blob NOT NULL CHECK ( length(blob_nn_check_fn) > 0 ),
    blob_nn_def_const    blob NOT NULL                                   DEFAULT ('42'),
    blob_nn_def_func     blob NOT NULL                                   DEFAULT (pi()),

    blob_check           blob CHECK ( length(blob_check) > 0 ),
    blob_check_def_const blob CHECK ( length(blob_check_def_const) > 0 ) DEFAULT ('42'),
    blob_check_def_func  blob CHECK ( length(blob_check_def_func) > 0 )  DEFAULT (pi())
);
