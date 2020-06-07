DROP TABLE IF EXISTS smallint_ref CASCADE;
CREATE TABLE smallint_ref (
    smallint_ref smallint UNIQUE
);

DROP TABLE IF EXISTS smallint_table;
CREATE TABLE smallint_table (
    i smallint,
    smallint_nn smallint NOT NULL,
    smallint_nn_unique smallint NOT NULL UNIQUE,
    smallint_nn_check smallint NOT NULL CHECK ( smallint_nn_check > 0 ),

    smallint_unique smallint UNIQUE,
    smallint_unique_check smallint UNIQUE CHECK ( smallint_unique_check > 0 ),
    smallint_unique_ref smallint UNIQUE REFERENCES smallint_ref(smallint_ref),
    smallint_unique_def_const smallint UNIQUE DEFAULT 42,
    smallint_unique_def_func smallint UNIQUE DEFAULT (pi()),

    smallint_check smallint CHECK ( smallint_check > 0 ),
    smallint_check_ref smallint CHECK ( smallint_check_ref > 0 ) REFERENCES smallint_ref(smallint_ref),
    smallint_check_def_const smallint CHECK ( smallint_check_def_const > 0 ) DEFAULT 42,
    smallint_check_def_func smallint CHECK ( smallint_check_def_func > 0 ) DEFAULT (pi()),

    smallint_ref smallint REFERENCES smallint_ref(smallint_ref),
    smallint_ref_unique_check smallint UNIQUE CHECK ( smallint_ref_unique_check > 0 ) REFERENCES smallint_ref(smallint_ref),

    smallint_def_const smallint DEFAULT 42,
    smallint_def_const_unique_check smallint UNIQUE CHECK ( smallint_def_const_unique_check > 0 ) DEFAULT 42,

    smallint_def_func smallint DEFAULT (pi()),
    smallint_def_func_unique_check smallint UNIQUE CHECK ( smallint_def_func_unique_check > 0 ) DEFAULT (pi())
);

DROP TABLE IF EXISTS smallint_pk;
CREATE TABLE smallint_pk (
    smallint_pk smallint PRIMARY KEY
);

DROP TABLE IF EXISTS smallint_pk_ref;
CREATE TABLE smallint_pk_ref (
    smallint_pk_ref smallint PRIMARY KEY REFERENCES smallint_ref(smallint_ref)
);

DROP TABLE IF EXISTS smallint_pk_def_const;
CREATE TABLE smallint_pk_def_const (
    smallint_pk_def_const smallint PRIMARY KEY DEFAULT 42
);

DROP TABLE IF EXISTS smallint_pk_def_func;
CREATE TABLE smallint_pk_def_func (
    smallint_pk_def_func smallint PRIMARY KEY DEFAULT (pi())
);

DROP TABLE IF EXISTS smallint_nn_pk;
CREATE TABLE smallint_nn_pk (
    smallint_nn_pk smallint NOT NULL PRIMARY KEY
);

DROP TABLE IF EXISTS smallint_nn_unique_check_pk;
CREATE TABLE smallint_nn_unique_check_pk (
    smallint_nn_unique_check_pk smallint PRIMARY KEY NOT NULL UNIQUE CHECK ( smallint_nn_unique_check_pk > 0)
);

DROP TABLE IF EXISTS smallint_nn_unique_check_pk_ref;
CREATE TABLE smallint_nn_unique_check_pk_ref (
    smallint_nn_unique_check_pk_ref smallint PRIMARY KEY NOT NULL UNIQUE CHECK ( smallint_nn_unique_check_pk_ref > 0) REFERENCES smallint_ref(smallint_ref)
);

DROP TABLE IF EXISTS smallint_unique_pk;
CREATE TABLE smallint_unique_pk (
    smallint_unique_pk smallint PRIMARY KEY UNIQUE
);

DROP TABLE IF EXISTS smallint_unique_check_pk;
CREATE TABLE smallint_unique_check_pk (
    smallint_unique_check_pk smallint PRIMARY KEY UNIQUE CHECK ( smallint_unique_check_pk > 0 )
);

DROP TABLE IF EXISTS smallint_unique_check_pk_ref;
CREATE TABLE smallint_unique_check_pk_ref (
    smallint_unique_check_pk_ref smallint PRIMARY KEY UNIQUE CHECK ( smallint_unique_check_pk_ref > 0) REFERENCES smallint_ref(smallint_ref)
);

DROP TABLE IF EXISTS smallint_check_pk;
CREATE TABLE smallint_check_pk (
    smallint_check_pk smallint PRIMARY KEY CHECK ( smallint_check_pk > 0 )
);

DROP TABLE IF EXISTS smallint_def_const_unique_check_pk;
CREATE TABLE smallint_def_const_unique_check_pk (
    smallint_def_const_unique_check_pk smallint PRIMARY KEY UNIQUE CHECK ( smallint_def_const_unique_check_pk > 0 ) DEFAULT 42
);

DROP TABLE IF EXISTS smallint_def_const_unique_check_pk_ref;
CREATE TABLE smallint_def_const_unique_check_pk_ref (
    smallint_def_const_unique_check_pk_ref smallint PRIMARY KEY UNIQUE CHECK ( smallint_def_const_unique_check_pk_ref > 0 ) DEFAULT 42 REFERENCES smallint_ref(smallint_ref)
);

DROP TABLE IF EXISTS smallint_def_func_unique_check_pk;
CREATE TABLE smallint_def_func_unique_check_pk (
    smallint_def_func_unique_check_pk smallint PRIMARY KEY UNIQUE CHECK ( smallint_def_func_unique_check_pk > 0 ) DEFAULT (pi())
);

DROP TABLE IF EXISTS smallint_def_func_unique_check_pk_ref;
CREATE TABLE smallint_def_func_unique_check_pk_ref (
    smallint_def_func_unique_check_pk_ref smallint PRIMARY KEY UNIQUE CHECK ( smallint_def_func_unique_check_pk_ref > 0 ) DEFAULT (pi()) REFERENCES smallint_ref(smallint_ref)
);
