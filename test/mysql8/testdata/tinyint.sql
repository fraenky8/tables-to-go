DROP TABLE IF EXISTS tinyint_ref CASCADE;
CREATE TABLE tinyint_ref (
    tinyint_ref tinyint UNIQUE
);

DROP TABLE IF EXISTS tinyint_table;
CREATE TABLE tinyint_table (
    i tinyint,
    tinyint_nn tinyint NOT NULL,
    tinyint_nn_unique tinyint NOT NULL UNIQUE,
    tinyint_nn_check tinyint NOT NULL CHECK ( tinyint_nn_check > 0 ),

    tinyint_unique tinyint UNIQUE,
    tinyint_unique_check tinyint UNIQUE CHECK ( tinyint_unique_check > 0 ),
    tinyint_unique_ref tinyint UNIQUE REFERENCES tinyint_ref(tinyint_ref),
    tinyint_unique_def_const tinyint UNIQUE DEFAULT 42,
    tinyint_unique_def_func tinyint UNIQUE DEFAULT (pi()),

    tinyint_check tinyint CHECK ( tinyint_check > 0 ),
    tinyint_check_ref tinyint CHECK ( tinyint_check_ref > 0 ) REFERENCES tinyint_ref(tinyint_ref),
    tinyint_check_def_const tinyint CHECK ( tinyint_check_def_const > 0 ) DEFAULT 42,
    tinyint_check_def_func tinyint CHECK ( tinyint_check_def_func > 0 ) DEFAULT (pi()),

    tinyint_ref tinyint REFERENCES tinyint_ref(tinyint_ref),
    tinyint_ref_unique_check tinyint UNIQUE CHECK ( tinyint_ref_unique_check > 0 ) REFERENCES tinyint_ref(tinyint_ref),

    tinyint_def_const tinyint DEFAULT 42,
    tinyint_def_const_unique_check tinyint UNIQUE CHECK ( tinyint_def_const_unique_check > 0 ) DEFAULT 42,

    tinyint_def_func tinyint DEFAULT (pi()),
    tinyint_def_func_unique_check tinyint UNIQUE CHECK ( tinyint_def_func_unique_check > 0 ) DEFAULT (pi())
);

DROP TABLE IF EXISTS tinyint_pk;
CREATE TABLE tinyint_pk (
    tinyint_pk tinyint PRIMARY KEY
);

DROP TABLE IF EXISTS tinyint_pk_ref;
CREATE TABLE tinyint_pk_ref (
    tinyint_pk_ref tinyint PRIMARY KEY REFERENCES tinyint_ref(tinyint_ref)
);

DROP TABLE IF EXISTS tinyint_pk_def_const;
CREATE TABLE tinyint_pk_def_const (
    tinyint_pk_def_const tinyint PRIMARY KEY DEFAULT 42
);

DROP TABLE IF EXISTS tinyint_pk_def_func;
CREATE TABLE tinyint_pk_def_func (
    tinyint_pk_def_func tinyint PRIMARY KEY DEFAULT (pi())
);

DROP TABLE IF EXISTS tinyint_nn_pk;
CREATE TABLE tinyint_nn_pk (
    tinyint_nn_pk tinyint NOT NULL PRIMARY KEY
);

DROP TABLE IF EXISTS tinyint_nn_unique_check_pk;
CREATE TABLE tinyint_nn_unique_check_pk (
    tinyint_nn_unique_check_pk tinyint PRIMARY KEY NOT NULL UNIQUE CHECK ( tinyint_nn_unique_check_pk > 0)
);

DROP TABLE IF EXISTS tinyint_nn_unique_check_pk_ref;
CREATE TABLE tinyint_nn_unique_check_pk_ref (
    tinyint_nn_unique_check_pk_ref tinyint PRIMARY KEY NOT NULL UNIQUE CHECK ( tinyint_nn_unique_check_pk_ref > 0) REFERENCES tinyint_ref(tinyint_ref)
);

DROP TABLE IF EXISTS tinyint_unique_pk;
CREATE TABLE tinyint_unique_pk (
    tinyint_unique_pk tinyint PRIMARY KEY UNIQUE
);

DROP TABLE IF EXISTS tinyint_unique_check_pk;
CREATE TABLE tinyint_unique_check_pk (
    tinyint_unique_check_pk tinyint PRIMARY KEY UNIQUE CHECK ( tinyint_unique_check_pk > 0 )
);

DROP TABLE IF EXISTS tinyint_unique_check_pk_ref;
CREATE TABLE tinyint_unique_check_pk_ref (
    tinyint_unique_check_pk_ref tinyint PRIMARY KEY UNIQUE CHECK ( tinyint_unique_check_pk_ref > 0) REFERENCES tinyint_ref(tinyint_ref)
);

DROP TABLE IF EXISTS tinyint_check_pk;
CREATE TABLE tinyint_check_pk (
    tinyint_check_pk tinyint PRIMARY KEY CHECK ( tinyint_check_pk > 0 )
);

DROP TABLE IF EXISTS tinyint_def_const_unique_check_pk;
CREATE TABLE tinyint_def_const_unique_check_pk (
    tinyint_def_const_unique_check_pk tinyint PRIMARY KEY UNIQUE CHECK ( tinyint_def_const_unique_check_pk > 0 ) DEFAULT 42
);

DROP TABLE IF EXISTS tinyint_def_const_unique_check_pk_ref;
CREATE TABLE tinyint_def_const_unique_check_pk_ref (
    tinyint_def_const_unique_check_pk_ref tinyint PRIMARY KEY UNIQUE CHECK ( tinyint_def_const_unique_check_pk_ref > 0 ) DEFAULT 42 REFERENCES tinyint_ref(tinyint_ref)
);

DROP TABLE IF EXISTS tinyint_def_func_unique_check_pk;
CREATE TABLE tinyint_def_func_unique_check_pk (
    tinyint_def_func_unique_check_pk tinyint PRIMARY KEY UNIQUE CHECK ( tinyint_def_func_unique_check_pk > 0 ) DEFAULT (pi())
);

DROP TABLE IF EXISTS tinyint_def_func_unique_check_pk_ref;
CREATE TABLE tinyint_def_func_unique_check_pk_ref (
    tinyint_def_func_unique_check_pk_ref tinyint PRIMARY KEY UNIQUE CHECK ( tinyint_def_func_unique_check_pk_ref > 0 ) DEFAULT (pi()) REFERENCES tinyint_ref(tinyint_ref)
);
