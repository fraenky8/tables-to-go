DROP TABLE IF EXISTS real_ref CASCADE;
CREATE TABLE real_ref
(
    real_ref real UNIQUE
);

DROP TABLE IF EXISTS real_table;
CREATE TABLE real_table
(
    col                         real,

    real_nn                     real NOT NULL,
    real_nn_unique              real NOT NULL UNIQUE,
    real_nn_check               real NOT NULL,
    real_nn_ref                 real NOT NULL REFERENCES real_ref (real_ref),
    real_nn_def_const           real NOT NULL DEFAULT 42,
    real_nn_def_func            real NOT NULL DEFAULT 42,
    real_nn_unique_check        real NOT NULL UNIQUE,

    real_unique                 real UNIQUE,
    real_unique_check           real UNIQUE,
    real_unique_ref             real UNIQUE REFERENCES real_ref (real_ref),
    real_unique_def_const       real UNIQUE   DEFAULT 42,
    real_unique_def_func        real UNIQUE   DEFAULT 42,

    real_check                  real,
    real_check_ref              real REFERENCES real_ref (real_ref),
    real_check_def_const        real          DEFAULT 42,
    real_check_def_func         real          DEFAULT 42,

    real_ref                    real REFERENCES real_ref (real_ref),
    real_ref_unique_check       real UNIQUE REFERENCES real_ref (real_ref),

    real_def_const              real          DEFAULT 42,
    real_def_const_unique_check real UNIQUE   DEFAULT 42,

    real_def_func               real          DEFAULT 42,
    real_def_func_unique_check  real UNIQUE   DEFAULT 42
);

DROP TABLE IF EXISTS real_pk;
CREATE TABLE real_pk
(
    real_pk real PRIMARY KEY
);

DROP TABLE IF EXISTS real_pk_ref;
CREATE TABLE real_pk_ref
(
    real_pk_ref real PRIMARY KEY REFERENCES real_ref (real_ref)
);

DROP TABLE IF EXISTS real_pk_def_const;
CREATE TABLE real_pk_def_const
(
    real_pk_def_const real PRIMARY KEY DEFAULT 42
);

DROP TABLE IF EXISTS real_pk_def_func;
CREATE TABLE real_pk_def_func
(
    real_pk_def_func real PRIMARY KEY DEFAULT 42
);

DROP TABLE IF EXISTS real_nn_pk;
CREATE TABLE real_nn_pk
(
    real_nn_pk real NOT NULL PRIMARY KEY
);

DROP TABLE IF EXISTS real_nn_unique_check_pk;
CREATE TABLE real_nn_unique_check_pk
(
    real_nn_unique_check_pk real PRIMARY KEY NOT NULL UNIQUE
);

DROP TABLE IF EXISTS real_nn_unique_check_pk_ref;
CREATE TABLE real_nn_unique_check_pk_ref
(
    real_nn_unique_check_pk_ref real PRIMARY KEY NOT NULL UNIQUE REFERENCES real_ref (real_ref)
);

DROP TABLE IF EXISTS real_unique_pk;
CREATE TABLE real_unique_pk
(
    real_unique_pk real PRIMARY KEY UNIQUE
);

DROP TABLE IF EXISTS real_unique_check_pk;
CREATE TABLE real_unique_check_pk
(
    real_unique_check_pk real PRIMARY KEY UNIQUE
);

DROP TABLE IF EXISTS real_unique_check_pk_ref;
CREATE TABLE real_unique_check_pk_ref
(
    real_unique_check_pk_ref real PRIMARY KEY UNIQUE REFERENCES real_ref (real_ref)
);

DROP TABLE IF EXISTS real_check_pk;
CREATE TABLE real_check_pk
(
    real_check_pk real PRIMARY KEY
);

DROP TABLE IF EXISTS real_def_const_unique_check_pk;
CREATE TABLE real_def_const_unique_check_pk
(
    real_def_const_unique_check_pk real PRIMARY KEY UNIQUE DEFAULT 42
);

DROP TABLE IF EXISTS real_def_const_unique_check_pk_ref;
CREATE TABLE real_def_const_unique_check_pk_ref
(
    real_def_const_unique_check_pk_ref real PRIMARY KEY UNIQUE DEFAULT 42 REFERENCES real_ref (real_ref)
);

DROP TABLE IF EXISTS real_def_func_unique_check_pk;
CREATE TABLE real_def_func_unique_check_pk
(
    real_def_func_unique_check_pk real PRIMARY KEY UNIQUE DEFAULT 42
);

DROP TABLE IF EXISTS real_def_func_unique_check_pk_ref;
CREATE TABLE real_def_func_unique_check_pk_ref
(
    real_def_func_unique_check_pk_ref real PRIMARY KEY UNIQUE DEFAULT 42 REFERENCES real_ref (real_ref)
);
