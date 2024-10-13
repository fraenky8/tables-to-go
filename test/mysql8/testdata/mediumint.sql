DROP TABLE IF EXISTS mediumint_ref CASCADE;
CREATE TABLE mediumint_ref (
    mediumint_ref mediumint UNIQUE
);

DROP TABLE IF EXISTS mediumint_table;
CREATE TABLE mediumint_table (
    i mediumint,
    mediumint_nn mediumint NOT NULL,
    mediumint_nn_unique mediumint NOT NULL UNIQUE,
    mediumint_nn_check mediumint NOT NULL CHECK ( mediumint_nn_check > 0 ),

    mediumint_unique mediumint UNIQUE,
    mediumint_unique_check mediumint UNIQUE CHECK ( mediumint_unique_check > 0 ),
    mediumint_unique_ref mediumint UNIQUE REFERENCES mediumint_ref(mediumint_ref),
    mediumint_unique_def_const mediumint UNIQUE DEFAULT 42,
    mediumint_unique_def_func mediumint UNIQUE DEFAULT (pi()),

    mediumint_check mediumint CHECK ( mediumint_check > 0 ),
    mediumint_check_ref mediumint CHECK ( mediumint_check_ref > 0 ) REFERENCES mediumint_ref(mediumint_ref),
    mediumint_check_def_const mediumint CHECK ( mediumint_check_def_const > 0 ) DEFAULT 42,
    mediumint_check_def_func mediumint CHECK ( mediumint_check_def_func > 0 ) DEFAULT (pi()),

    mediumint_ref mediumint REFERENCES mediumint_ref(mediumint_ref),
    mediumint_ref_unique_check mediumint UNIQUE CHECK ( mediumint_ref_unique_check > 0 ) REFERENCES mediumint_ref(mediumint_ref),

    mediumint_def_const mediumint DEFAULT 42,
    mediumint_def_const_unique_check mediumint UNIQUE CHECK ( mediumint_def_const_unique_check > 0 ) DEFAULT 42,

    mediumint_def_func mediumint DEFAULT (pi()),
    mediumint_def_func_unique_check mediumint UNIQUE CHECK ( mediumint_def_func_unique_check > 0 ) DEFAULT (pi())
);

DROP TABLE IF EXISTS mediumint_pk;
CREATE TABLE mediumint_pk (
    mediumint_pk mediumint PRIMARY KEY
);

DROP TABLE IF EXISTS mediumint_pk_ref;
CREATE TABLE mediumint_pk_ref (
    mediumint_pk_ref mediumint PRIMARY KEY REFERENCES mediumint_ref(mediumint_ref)
);

DROP TABLE IF EXISTS mediumint_pk_def_const;
CREATE TABLE mediumint_pk_def_const (
    mediumint_pk_def_const mediumint PRIMARY KEY DEFAULT 42
);

DROP TABLE IF EXISTS mediumint_pk_def_func;
CREATE TABLE mediumint_pk_def_func (
    mediumint_pk_def_func mediumint PRIMARY KEY DEFAULT (pi())
);

DROP TABLE IF EXISTS mediumint_nn_pk;
CREATE TABLE mediumint_nn_pk (
    mediumint_nn_pk mediumint NOT NULL PRIMARY KEY
);

DROP TABLE IF EXISTS mediumint_nn_unique_check_pk;
CREATE TABLE mediumint_nn_unique_check_pk (
    mediumint_nn_unique_check_pk mediumint PRIMARY KEY NOT NULL UNIQUE CHECK ( mediumint_nn_unique_check_pk > 0)
);

DROP TABLE IF EXISTS mediumint_nn_unique_check_pk_ref;
CREATE TABLE mediumint_nn_unique_check_pk_ref (
    mediumint_nn_unique_check_pk_ref mediumint PRIMARY KEY NOT NULL UNIQUE CHECK ( mediumint_nn_unique_check_pk_ref > 0) REFERENCES mediumint_ref(mediumint_ref)
);

DROP TABLE IF EXISTS mediumint_unique_pk;
CREATE TABLE mediumint_unique_pk (
    mediumint_unique_pk mediumint PRIMARY KEY UNIQUE
);

DROP TABLE IF EXISTS mediumint_unique_check_pk;
CREATE TABLE mediumint_unique_check_pk (
    mediumint_unique_check_pk mediumint PRIMARY KEY UNIQUE CHECK ( mediumint_unique_check_pk > 0 )
);

DROP TABLE IF EXISTS mediumint_unique_check_pk_ref;
CREATE TABLE mediumint_unique_check_pk_ref (
    mediumint_unique_check_pk_ref mediumint PRIMARY KEY UNIQUE CHECK ( mediumint_unique_check_pk_ref > 0) REFERENCES mediumint_ref(mediumint_ref)
);

DROP TABLE IF EXISTS mediumint_check_pk;
CREATE TABLE mediumint_check_pk (
    mediumint_check_pk mediumint PRIMARY KEY CHECK ( mediumint_check_pk > 0 )
);

DROP TABLE IF EXISTS mediumint_def_const_unique_check_pk;
CREATE TABLE mediumint_def_const_unique_check_pk (
    mediumint_def_const_unique_check_pk mediumint PRIMARY KEY UNIQUE CHECK ( mediumint_def_const_unique_check_pk > 0 ) DEFAULT 42
);

DROP TABLE IF EXISTS mediumint_def_const_unique_check_pk_ref;
CREATE TABLE mediumint_def_const_unique_check_pk_ref (
    mediumint_def_const_unique_check_pk_ref mediumint PRIMARY KEY UNIQUE CHECK ( mediumint_def_const_unique_check_pk_ref > 0 ) DEFAULT 42 REFERENCES mediumint_ref(mediumint_ref)
);

DROP TABLE IF EXISTS mediumint_def_func_unique_check_pk;
CREATE TABLE mediumint_def_func_unique_check_pk (
    mediumint_def_func_unique_check_pk mediumint PRIMARY KEY UNIQUE CHECK ( mediumint_def_func_unique_check_pk > 0 ) DEFAULT (pi())
);

DROP TABLE IF EXISTS mediumint_def_func_unique_check_pk_ref;
CREATE TABLE mediumint_def_func_unique_check_pk_ref (
    mediumint_def_func_unique_check_pk_ref mediumint PRIMARY KEY UNIQUE CHECK ( mediumint_def_func_unique_check_pk_ref > 0 ) DEFAULT (pi()) REFERENCES mediumint_ref(mediumint_ref)
);
