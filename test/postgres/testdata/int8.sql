DROP TABLE IF EXISTS int8_ref CASCADE;
CREATE TABLE int8_ref (
    int8_ref int8 UNIQUE
);

DROP TABLE IF EXISTS int8;
CREATE TABLE int8 (
    int8 int8,
    int8_nn int8 NOT NULL,
    int8_nn_unique int8 NOT NULL UNIQUE,
    int8_nn_check int8 NOT NULL CHECK ( int8 > 0 ),
    int8_nn_ref int8 NOT NULL REFERENCES int8_ref(int8_ref),
    int8_nn_def_const int8 NOT NULL DEFAULT 42,
    int8_nn_def_func int8 NOT NULL DEFAULT pi(),
    int8_nn_unique_check int8 NOT NULL UNIQUE CHECK ( int8 > 0 ),

    int8_unique int8 UNIQUE,
    int8_unique_check int8 UNIQUE CHECK ( int8 > 0 ),
    int8_unique_ref int8 UNIQUE REFERENCES int8_ref(int8_ref),
    int8_unique_def_const int8 UNIQUE DEFAULT 42,
    int8_unique_def_func int8 UNIQUE DEFAULT pi(),

    int8_check int8 CHECK ( int8 > 0 ),
    int8_check_ref int8 CHECK ( int8 > 0 ) REFERENCES int8_ref(int8_ref),
    int8_check_def_const int8 CHECK ( int8 > 0 ) DEFAULT 42,
    int8_check_def_func int8 CHECK ( int8 > 0 ) DEFAULT pi(),

    int8_ref int8 REFERENCES int8_ref(int8_ref),
    int8_ref_def_const int8 REFERENCES int8_ref(int8_ref) DEFAULT 42,
    int8_ref_def_func int8 REFERENCES int8_ref(int8_ref) DEFAULT pi(),
    int8_ref_unique_check int8 UNIQUE CHECK ( int8 > 0 ) REFERENCES int8_ref(int8_ref),

    int8_def_const int8 DEFAULT 42,
    int8_def_const_unique_check int8 UNIQUE CHECK ( int8 > 0 )DEFAULT 42,

    int8_def_func int8 DEFAULT pi(),
    int8_def_func_unique_check int8 UNIQUE CHECK ( int8 > 0 ) DEFAULT pi()
);

DROP TABLE IF EXISTS int8_pk;
CREATE TABLE int8_pk (
    int8_pk int8 PRIMARY KEY
);

DROP TABLE IF EXISTS int8_pk_ref;
CREATE TABLE int8_pk_ref (
    int8_pk_ref int8 PRIMARY KEY REFERENCES int8_ref(int8_ref)
);

DROP TABLE IF EXISTS int8_pk_def_const;
CREATE TABLE int8_pk_def_const (
    int8_pk_def_const int8 PRIMARY KEY DEFAULT 42
);

DROP TABLE IF EXISTS int8_pk_def_func;
CREATE TABLE int8_pk_def_func (
    int8_pk_def_func int8 PRIMARY KEY DEFAULT pi()
);

DROP TABLE IF EXISTS int8_nn_pk;
CREATE TABLE int8_nn_pk (
    int8_nn_pk int8 NOT NULL PRIMARY KEY
);

DROP TABLE IF EXISTS int8_nn_unique_check_pk;
CREATE TABLE int8_nn_unique_check_pk (
    int8_nn_unique_check_pk int8 PRIMARY KEY NOT NULL UNIQUE CHECK ( int8_nn_unique_check_pk > 0)
);

DROP TABLE IF EXISTS int8_nn_unique_check_pk_ref;
CREATE TABLE int8_nn_unique_check_pk_ref (
    int8_nn_unique_check_pk_ref int8 PRIMARY KEY NOT NULL UNIQUE CHECK ( int8_nn_unique_check_pk_ref > 0) REFERENCES int8_ref(int8_ref)
);

DROP TABLE IF EXISTS int8_unique_pk;
CREATE TABLE int8_unique_pk (
    int8_unique_pk int8 PRIMARY KEY UNIQUE
);

DROP TABLE IF EXISTS int8_unique_check_pk;
CREATE TABLE int8_unique_check_pk (
    int8_unique_check_pk int8 PRIMARY KEY UNIQUE CHECK ( int8_unique_check_pk > 0 )
);

DROP TABLE IF EXISTS int8_unique_check_pk_ref;
CREATE TABLE int8_unique_check_pk_ref (
    int8_unique_check_pk_ref int8 PRIMARY KEY UNIQUE CHECK ( int8_unique_check_pk_ref > 0) REFERENCES int8_ref(int8_ref)
);

DROP TABLE IF EXISTS int8_check_pk;
CREATE TABLE int8_check_pk (
    int8_check_pk int8 PRIMARY KEY CHECK ( int8_check_pk > 0 )
);

DROP TABLE IF EXISTS int8_def_const_unique_check_pk;
CREATE TABLE int8_def_const_unique_check_pk (
    int8_def_const_unique_check_pk int8 PRIMARY KEY UNIQUE CHECK ( int8_def_const_unique_check_pk > 0 ) DEFAULT 42
);

DROP TABLE IF EXISTS int8_def_const_unique_check_pk_ref;
CREATE TABLE int8_def_const_unique_check_pk_ref (
    int8_def_const_unique_check_pk_ref int8 PRIMARY KEY UNIQUE CHECK ( int8_def_const_unique_check_pk_ref > 0 ) DEFAULT 42 REFERENCES int8_ref(int8_ref)
);

DROP TABLE IF EXISTS int8_def_func_unique_check_pk;
CREATE TABLE int8_def_func_unique_check_pk (
    int8_def_func_unique_check_pk int8 PRIMARY KEY UNIQUE CHECK ( int8_def_func_unique_check_pk > 0 ) DEFAULT pi()
);

DROP TABLE IF EXISTS int8_def_func_unique_check_pk_ref;
CREATE TABLE int8_def_func_unique_check_pk_ref (
    int8_def_func_unique_check_pk_ref int8 PRIMARY KEY UNIQUE CHECK ( int8_def_func_unique_check_pk_ref > 0 ) DEFAULT pi() REFERENCES int8_ref(int8_ref)
);
