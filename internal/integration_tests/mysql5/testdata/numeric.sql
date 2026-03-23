DROP TABLE IF EXISTS numeric_ref CASCADE;
CREATE TABLE numeric_ref
(
    numeric_ref numeric UNIQUE
);

DROP TABLE IF EXISTS numeric_table;
CREATE TABLE numeric_table
(
    col                            numeric,

    numeric_nn                     numeric NOT NULL,
    numeric_nn_unique              numeric NOT NULL UNIQUE,
    numeric_nn_check               numeric NOT NULL,
    numeric_nn_ref                 numeric NOT NULL REFERENCES numeric_ref (numeric_ref),
    numeric_nn_def_const           numeric NOT NULL DEFAULT 42,
    numeric_nn_def_func            numeric NOT NULL DEFAULT 42,
    numeric_nn_unique_check        numeric NOT NULL UNIQUE,

    numeric_unique                 numeric UNIQUE,
    numeric_unique_check           numeric UNIQUE,
    numeric_unique_ref             numeric UNIQUE REFERENCES numeric_ref (numeric_ref),
    numeric_unique_def_const       numeric UNIQUE   DEFAULT 42,
    numeric_unique_def_func        numeric UNIQUE   DEFAULT 42,

    numeric_check                  numeric,
    numeric_check_ref              numeric REFERENCES numeric_ref (numeric_ref),
    numeric_check_def_const        numeric          DEFAULT 42,
    numeric_check_def_func         numeric          DEFAULT 42,

    numeric_ref                    numeric REFERENCES numeric_ref (numeric_ref),
    numeric_ref_unique_check       numeric UNIQUE REFERENCES numeric_ref (numeric_ref),

    numeric_def_const              numeric          DEFAULT 42,
    numeric_def_const_unique_check numeric UNIQUE   DEFAULT 42,

    numeric_def_func               numeric          DEFAULT 42,
    numeric_def_func_unique_check  numeric UNIQUE   DEFAULT 42
);

DROP TABLE IF EXISTS numeric_pk;
CREATE TABLE numeric_pk
(
    numeric_pk numeric PRIMARY KEY
);

DROP TABLE IF EXISTS numeric_pk_ref;
CREATE TABLE numeric_pk_ref
(
    numeric_pk_ref numeric PRIMARY KEY REFERENCES numeric_ref (numeric_ref)
);

DROP TABLE IF EXISTS numeric_pk_def_const;
CREATE TABLE numeric_pk_def_const
(
    numeric_pk_def_const numeric PRIMARY KEY DEFAULT 42
);

DROP TABLE IF EXISTS numeric_pk_def_func;
CREATE TABLE numeric_pk_def_func
(
    numeric_pk_def_func numeric PRIMARY KEY DEFAULT 42
);

DROP TABLE IF EXISTS numeric_nn_pk;
CREATE TABLE numeric_nn_pk
(
    numeric_nn_pk numeric NOT NULL PRIMARY KEY
);

DROP TABLE IF EXISTS numeric_nn_unique_check_pk;
CREATE TABLE numeric_nn_unique_check_pk
(
    numeric_nn_unique_check_pk numeric PRIMARY KEY NOT NULL UNIQUE
);

DROP TABLE IF EXISTS numeric_nn_unique_check_pk_ref;
CREATE TABLE numeric_nn_unique_check_pk_ref
(
    numeric_nn_unique_check_pk_ref numeric PRIMARY KEY NOT NULL UNIQUE REFERENCES numeric_ref (numeric_ref)
);

DROP TABLE IF EXISTS numeric_unique_pk;
CREATE TABLE numeric_unique_pk
(
    numeric_unique_pk numeric PRIMARY KEY UNIQUE
);

DROP TABLE IF EXISTS numeric_unique_check_pk;
CREATE TABLE numeric_unique_check_pk
(
    numeric_unique_check_pk numeric PRIMARY KEY UNIQUE
);

DROP TABLE IF EXISTS numeric_unique_check_pk_ref;
CREATE TABLE numeric_unique_check_pk_ref
(
    numeric_unique_check_pk_ref numeric PRIMARY KEY UNIQUE REFERENCES numeric_ref (numeric_ref)
);

DROP TABLE IF EXISTS numeric_check_pk;
CREATE TABLE numeric_check_pk
(
    numeric_check_pk numeric PRIMARY KEY
);

DROP TABLE IF EXISTS numeric_def_const_unique_check_pk;
CREATE TABLE numeric_def_const_unique_check_pk
(
    numeric_def_const_unique_check_pk numeric PRIMARY KEY UNIQUE DEFAULT 42
);

DROP TABLE IF EXISTS numeric_def_const_unique_check_pk_ref;
CREATE TABLE numeric_def_const_unique_check_pk_ref
(
    numeric_def_const_unique_check_pk_ref numeric PRIMARY KEY UNIQUE DEFAULT 42 REFERENCES numeric_ref (numeric_ref)
);

DROP TABLE IF EXISTS numeric_def_func_unique_check_pk;
CREATE TABLE numeric_def_func_unique_check_pk
(
    numeric_def_func_unique_check_pk numeric PRIMARY KEY UNIQUE DEFAULT 42
);

DROP TABLE IF EXISTS numeric_def_func_unique_check_pk_ref;
CREATE TABLE numeric_def_func_unique_check_pk_ref
(
    numeric_def_func_unique_check_pk_ref numeric PRIMARY KEY UNIQUE DEFAULT 42 REFERENCES numeric_ref (numeric_ref)
);
