DROP TABLE IF EXISTS decimal_ref CASCADE;
CREATE TABLE decimal_ref
(
    decimal_ref decimal UNIQUE
);

DROP TABLE IF EXISTS decimal_table;
CREATE TABLE decimal_table
(
    col                            decimal,

    decimal_nn                     decimal NOT NULL,
    decimal_nn_unique              decimal NOT NULL UNIQUE,
    decimal_nn_check               decimal NOT NULL,
    decimal_nn_ref                 decimal NOT NULL REFERENCES decimal_ref (decimal_ref),
    decimal_nn_def_const           decimal NOT NULL DEFAULT 42,
    decimal_nn_def_func            decimal NOT NULL DEFAULT 42,
    decimal_nn_unique_check        decimal NOT NULL UNIQUE,

    decimal_unique                 decimal UNIQUE,
    decimal_unique_check           decimal UNIQUE,
    decimal_unique_ref             decimal UNIQUE REFERENCES decimal_ref (decimal_ref),
    decimal_unique_def_const       decimal UNIQUE   DEFAULT 42,
    decimal_unique_def_func        decimal UNIQUE   DEFAULT 42,

    decimal_check                  decimal,
    decimal_check_ref              decimal REFERENCES decimal_ref (decimal_ref),
    decimal_check_def_const        decimal          DEFAULT 42,
    decimal_check_def_func         decimal          DEFAULT 42,

    decimal_ref                    decimal REFERENCES decimal_ref (decimal_ref),
    decimal_ref_unique_check       decimal UNIQUE REFERENCES decimal_ref (decimal_ref),

    decimal_def_const              decimal          DEFAULT 42,
    decimal_def_const_unique_check decimal UNIQUE   DEFAULT 42,

    decimal_def_func               decimal          DEFAULT 42,
    decimal_def_func_unique_check  decimal UNIQUE   DEFAULT 42
);

DROP TABLE IF EXISTS decimal_pk;
CREATE TABLE decimal_pk
(
    decimal_pk decimal PRIMARY KEY
);

DROP TABLE IF EXISTS decimal_pk_ref;
CREATE TABLE decimal_pk_ref
(
    decimal_pk_ref decimal PRIMARY KEY REFERENCES decimal_ref (decimal_ref)
);

DROP TABLE IF EXISTS decimal_pk_def_const;
CREATE TABLE decimal_pk_def_const
(
    decimal_pk_def_const decimal PRIMARY KEY DEFAULT 42
);

DROP TABLE IF EXISTS decimal_pk_def_func;
CREATE TABLE decimal_pk_def_func
(
    decimal_pk_def_func decimal PRIMARY KEY DEFAULT 42
);

DROP TABLE IF EXISTS decimal_nn_pk;
CREATE TABLE decimal_nn_pk
(
    decimal_nn_pk decimal NOT NULL PRIMARY KEY
);

DROP TABLE IF EXISTS decimal_nn_unique_check_pk;
CREATE TABLE decimal_nn_unique_check_pk
(
    decimal_nn_unique_check_pk decimal PRIMARY KEY NOT NULL UNIQUE
);

DROP TABLE IF EXISTS decimal_nn_unique_check_pk_ref;
CREATE TABLE decimal_nn_unique_check_pk_ref
(
    decimal_nn_unique_check_pk_ref decimal PRIMARY KEY NOT NULL UNIQUE REFERENCES decimal_ref (decimal_ref)
);

DROP TABLE IF EXISTS decimal_unique_pk;
CREATE TABLE decimal_unique_pk
(
    decimal_unique_pk decimal PRIMARY KEY UNIQUE
);

DROP TABLE IF EXISTS decimal_unique_check_pk;
CREATE TABLE decimal_unique_check_pk
(
    decimal_unique_check_pk decimal PRIMARY KEY UNIQUE
);

DROP TABLE IF EXISTS decimal_unique_check_pk_ref;
CREATE TABLE decimal_unique_check_pk_ref
(
    decimal_unique_check_pk_ref decimal PRIMARY KEY UNIQUE REFERENCES decimal_ref (decimal_ref)
);

DROP TABLE IF EXISTS decimal_check_pk;
CREATE TABLE decimal_check_pk
(
    decimal_check_pk decimal PRIMARY KEY
);

DROP TABLE IF EXISTS decimal_def_const_unique_check_pk;
CREATE TABLE decimal_def_const_unique_check_pk
(
    decimal_def_const_unique_check_pk decimal PRIMARY KEY UNIQUE DEFAULT 42
);

DROP TABLE IF EXISTS decimal_def_const_unique_check_pk_ref;
CREATE TABLE decimal_def_const_unique_check_pk_ref
(
    decimal_def_const_unique_check_pk_ref decimal PRIMARY KEY UNIQUE DEFAULT 42 REFERENCES decimal_ref (decimal_ref)
);

DROP TABLE IF EXISTS decimal_def_func_unique_check_pk;
CREATE TABLE decimal_def_func_unique_check_pk
(
    decimal_def_func_unique_check_pk decimal PRIMARY KEY UNIQUE DEFAULT 42
);

DROP TABLE IF EXISTS decimal_def_func_unique_check_pk_ref;
CREATE TABLE decimal_def_func_unique_check_pk_ref
(
    decimal_def_func_unique_check_pk_ref decimal PRIMARY KEY UNIQUE DEFAULT 42 REFERENCES decimal_ref (decimal_ref)
);
