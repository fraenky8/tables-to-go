DROP TABLE IF EXISTS tinyint_ref CASCADE;
CREATE TABLE tinyint_ref
(
    tinyint_ref tinyint UNIQUE
);

DROP TABLE IF EXISTS tinyint_table;
CREATE TABLE tinyint_table
(
    i                              tinyint,
    tinyint_nn                     tinyint NOT NULL,
    tinyint_nn_unique              tinyint NOT NULL UNIQUE,
    tinyint_nn_check               tinyint NOT NULL,

    tinyint_unique                 tinyint UNIQUE,
    tinyint_unique_check           tinyint UNIQUE,
    tinyint_unique_ref             tinyint UNIQUE REFERENCES tinyint_ref (tinyint_ref),
    tinyint_unique_def_const       tinyint UNIQUE DEFAULT 42,
    tinyint_unique_def_func        tinyint UNIQUE DEFAULT 42,

    tinyint_check                  tinyint,
    tinyint_check_ref              tinyint REFERENCES tinyint_ref (tinyint_ref),
    tinyint_check_def_const        tinyint        DEFAULT 42,
    tinyint_check_def_func         tinyint        DEFAULT 42,

    tinyint_ref                    tinyint REFERENCES tinyint_ref (tinyint_ref),
    tinyint_ref_unique_check       tinyint UNIQUE REFERENCES tinyint_ref (tinyint_ref),

    tinyint_def_const              tinyint        DEFAULT 42,
    tinyint_def_const_unique_check tinyint UNIQUE DEFAULT 42,

    tinyint_def_func               tinyint        DEFAULT 42,
    tinyint_def_func_unique_check  tinyint UNIQUE DEFAULT 42
);

DROP TABLE IF EXISTS tinyint_pk;
CREATE TABLE tinyint_pk
(
    tinyint_pk tinyint PRIMARY KEY
);

DROP TABLE IF EXISTS tinyint_pk_ref;
CREATE TABLE tinyint_pk_ref
(
    tinyint_pk_ref tinyint PRIMARY KEY REFERENCES tinyint_ref (tinyint_ref)
);

DROP TABLE IF EXISTS tinyint_pk_def_const;
CREATE TABLE tinyint_pk_def_const
(
    tinyint_pk_def_const tinyint PRIMARY KEY DEFAULT 42
);

DROP TABLE IF EXISTS tinyint_pk_def_func;
CREATE TABLE tinyint_pk_def_func
(
    tinyint_pk_def_func tinyint PRIMARY KEY DEFAULT 42
);

DROP TABLE IF EXISTS tinyint_nn_pk;
CREATE TABLE tinyint_nn_pk
(
    tinyint_nn_pk tinyint NOT NULL PRIMARY KEY
);

DROP TABLE IF EXISTS tinyint_nn_unique_check_pk;
CREATE TABLE tinyint_nn_unique_check_pk
(
    tinyint_nn_unique_check_pk tinyint PRIMARY KEY NOT NULL UNIQUE
);

DROP TABLE IF EXISTS tinyint_nn_unique_check_pk_ref;
CREATE TABLE tinyint_nn_unique_check_pk_ref
(
    tinyint_nn_unique_check_pk_ref tinyint PRIMARY KEY NOT NULL UNIQUE REFERENCES tinyint_ref (tinyint_ref)
);

DROP TABLE IF EXISTS tinyint_unique_pk;
CREATE TABLE tinyint_unique_pk
(
    tinyint_unique_pk tinyint PRIMARY KEY UNIQUE
);

DROP TABLE IF EXISTS tinyint_unique_check_pk;
CREATE TABLE tinyint_unique_check_pk
(
    tinyint_unique_check_pk tinyint PRIMARY KEY UNIQUE
);

DROP TABLE IF EXISTS tinyint_unique_check_pk_ref;
CREATE TABLE tinyint_unique_check_pk_ref
(
    tinyint_unique_check_pk_ref tinyint PRIMARY KEY UNIQUE REFERENCES tinyint_ref (tinyint_ref)
);

DROP TABLE IF EXISTS tinyint_check_pk;
CREATE TABLE tinyint_check_pk
(
    tinyint_check_pk tinyint PRIMARY KEY
);

DROP TABLE IF EXISTS tinyint_def_const_unique_check_pk;
CREATE TABLE tinyint_def_const_unique_check_pk
(
    tinyint_def_const_unique_check_pk tinyint PRIMARY KEY UNIQUE DEFAULT 42
);

DROP TABLE IF EXISTS tinyint_def_const_unique_check_pk_ref;
CREATE TABLE tinyint_def_const_unique_check_pk_ref
(
    tinyint_def_const_unique_check_pk_ref tinyint PRIMARY KEY UNIQUE DEFAULT 42 REFERENCES tinyint_ref (tinyint_ref)
);

DROP TABLE IF EXISTS tinyint_def_func_unique_check_pk;
CREATE TABLE tinyint_def_func_unique_check_pk
(
    tinyint_def_func_unique_check_pk tinyint PRIMARY KEY UNIQUE DEFAULT 42
);

DROP TABLE IF EXISTS tinyint_def_func_unique_check_pk_ref;
CREATE TABLE tinyint_def_func_unique_check_pk_ref
(
    tinyint_def_func_unique_check_pk_ref tinyint PRIMARY KEY UNIQUE DEFAULT 42 REFERENCES tinyint_ref (tinyint_ref)
);
