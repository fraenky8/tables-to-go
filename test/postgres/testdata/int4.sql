DROP TABLE IF EXISTS int4_ref CASCADE;
CREATE TABLE int4_ref (
    int4_ref int4 UNIQUE
);

DROP TABLE IF EXISTS int4;
CREATE TABLE int4 (
    int4 int4,
    int4_nn int4 NOT NULL,
    int4_nn_unique int4 NOT NULL UNIQUE,
    int4_nn_check int4 NOT NULL CHECK ( int4 > 0 ),
    int4_nn_ref int4 NOT NULL REFERENCES int4_ref(int4_ref),
    int4_nn_def_const int4 NOT NULL DEFAULT 42,
    int4_nn_def_func int4 NOT NULL DEFAULT pi(),
    int4_nn_unique_check int4 NOT NULL UNIQUE CHECK ( int4 > 0 ),

    int4_unique int4 UNIQUE,
    int4_unique_check int4 UNIQUE CHECK ( int4 > 0 ),
    int4_unique_ref int4 UNIQUE REFERENCES int4_ref(int4_ref),
    int4_unique_def_const int4 UNIQUE DEFAULT 42,
    int4_unique_def_func int4 UNIQUE DEFAULT pi(),

    int4_check int4 CHECK ( int4 > 0 ),
    int4_check_ref int4 CHECK ( int4 > 0 ) REFERENCES int4_ref(int4_ref),
    int4_check_def_const int4 CHECK ( int4 > 0 ) DEFAULT 42,
    int4_check_def_func int4 CHECK ( int4 > 0 ) DEFAULT pi(),

    int4_ref int4 REFERENCES int4_ref(int4_ref),
    int4_ref_def_const int4 REFERENCES int4_ref(int4_ref) DEFAULT 42,
    int4_ref_def_func int4 REFERENCES int4_ref(int4_ref) DEFAULT pi(),
    int4_ref_unique_check int4 UNIQUE CHECK ( int4 > 0 ) REFERENCES int4_ref(int4_ref),

    int4_def_const int4 DEFAULT 42,
    int4_def_const_unique_check int4 UNIQUE CHECK ( int4 > 0 )DEFAULT 42,

    int4_def_func int4 DEFAULT pi(),
    int4_def_func_unique_check int4 UNIQUE CHECK ( int4 > 0 ) DEFAULT pi()
);

DROP TABLE IF EXISTS int4_pk;
CREATE TABLE int4_pk (
    int4_pk int4 PRIMARY KEY
);

DROP TABLE IF EXISTS int4_pk_ref;
CREATE TABLE int4_pk_ref (
    int4_pk_ref int4 PRIMARY KEY REFERENCES int4_ref(int4_ref)
);

DROP TABLE IF EXISTS int4_pk_def_const;
CREATE TABLE int4_pk_def_const (
    int4_pk_def_const int4 PRIMARY KEY DEFAULT 42
);

DROP TABLE IF EXISTS int4_pk_def_func;
CREATE TABLE int4_pk_def_func (
    int4_pk_def_func int4 PRIMARY KEY DEFAULT pi()
);

DROP TABLE IF EXISTS int4_nn_pk;
CREATE TABLE int4_nn_pk (
    int4_nn_pk int4 NOT NULL PRIMARY KEY
);

DROP TABLE IF EXISTS int4_nn_unique_check_pk;
CREATE TABLE int4_nn_unique_check_pk (
    int4_nn_unique_check_pk int4 PRIMARY KEY NOT NULL UNIQUE CHECK ( int4_nn_unique_check_pk > 0)
);

DROP TABLE IF EXISTS int4_nn_unique_check_pk_ref;
CREATE TABLE int4_nn_unique_check_pk_ref (
    int4_nn_unique_check_pk_ref int4 PRIMARY KEY NOT NULL UNIQUE CHECK ( int4_nn_unique_check_pk_ref > 0) REFERENCES int4_ref(int4_ref)
);

DROP TABLE IF EXISTS int4_unique_pk;
CREATE TABLE int4_unique_pk (
    int4_unique_pk int4 PRIMARY KEY UNIQUE
);

DROP TABLE IF EXISTS int4_unique_check_pk;
CREATE TABLE int4_unique_check_pk (
    int4_unique_check_pk int4 PRIMARY KEY UNIQUE CHECK ( int4_unique_check_pk > 0 )
);

DROP TABLE IF EXISTS int4_unique_check_pk_ref;
CREATE TABLE int4_unique_check_pk_ref (
    int4_unique_check_pk_ref int4 PRIMARY KEY UNIQUE CHECK ( int4_unique_check_pk_ref > 0) REFERENCES int4_ref(int4_ref)
);

DROP TABLE IF EXISTS int4_check_pk;
CREATE TABLE int4_check_pk (
    int4_check_pk int4 PRIMARY KEY CHECK ( int4_check_pk > 0 )
);

DROP TABLE IF EXISTS int4_def_const_unique_check_pk;
CREATE TABLE int4_def_const_unique_check_pk (
    int4_def_const_unique_check_pk int4 PRIMARY KEY UNIQUE CHECK ( int4_def_const_unique_check_pk > 0 ) DEFAULT 42
);

DROP TABLE IF EXISTS int4_def_const_unique_check_pk_ref;
CREATE TABLE int4_def_const_unique_check_pk_ref (
    int4_def_const_unique_check_pk_ref int4 PRIMARY KEY UNIQUE CHECK ( int4_def_const_unique_check_pk_ref > 0 ) DEFAULT 42 REFERENCES int4_ref(int4_ref)
);

DROP TABLE IF EXISTS int4_def_func_unique_check_pk;
CREATE TABLE int4_def_func_unique_check_pk (
    int4_def_func_unique_check_pk int4 PRIMARY KEY UNIQUE CHECK ( int4_def_func_unique_check_pk > 0 ) DEFAULT pi()
);

DROP TABLE IF EXISTS int4_def_func_unique_check_pk_ref;
CREATE TABLE int4_def_func_unique_check_pk_ref (
    int4_def_func_unique_check_pk_ref int4 PRIMARY KEY UNIQUE CHECK ( int4_def_func_unique_check_pk_ref > 0 ) DEFAULT pi() REFERENCES int4_ref(int4_ref)
);
