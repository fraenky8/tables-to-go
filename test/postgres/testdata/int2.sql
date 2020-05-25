DROP TABLE IF EXISTS int2_ref CASCADE;
CREATE TABLE int2_ref (
    int2_ref int2 UNIQUE
);

DROP TABLE IF EXISTS int2;
CREATE TABLE int2 (
    int2 int2,
    int2_nn int2 NOT NULL,
    int2_nn_unique int2 NOT NULL UNIQUE,
    int2_nn_check int2 NOT NULL CHECK ( int2 > 0 ),
    int2_nn_ref int2 NOT NULL REFERENCES int2_ref(int2_ref),
    int2_nn_def_const int2 NOT NULL DEFAULT 42,
    int2_nn_def_func int2 NOT NULL DEFAULT pi(),
    int2_nn_unique_check int2 NOT NULL UNIQUE CHECK ( int2 > 0 ),

    int2_unique int2 UNIQUE,
    int2_unique_check int2 UNIQUE CHECK ( int2 > 0 ),
    int2_unique_ref int2 UNIQUE REFERENCES int2_ref(int2_ref),
    int2_unique_def_const int2 UNIQUE DEFAULT 42,
    int2_unique_def_func int2 UNIQUE DEFAULT pi(),

    int2_check int2 CHECK ( int2 > 0 ),
    int2_check_ref int2 CHECK ( int2 > 0 ) REFERENCES int2_ref(int2_ref),
    int2_check_def_const int2 CHECK ( int2 > 0 ) DEFAULT 42,
    int2_check_def_func int2 CHECK ( int2 > 0 ) DEFAULT pi(),

    int2_ref int2 REFERENCES int2_ref(int2_ref),
    int2_ref_def_const int2 REFERENCES int2_ref(int2_ref) DEFAULT 42,
    int2_ref_def_func int2 REFERENCES int2_ref(int2_ref) DEFAULT pi(),
    int2_ref_unique_check int2 UNIQUE CHECK ( int2 > 0 ) REFERENCES int2_ref(int2_ref),

    int2_def_const int2 DEFAULT 42,
    int2_def_const_unique_check int2 UNIQUE CHECK ( int2 > 0 )DEFAULT 42,

    int2_def_func int2 DEFAULT pi(),
    int2_def_func_unique_check int2 UNIQUE CHECK ( int2 > 0 ) DEFAULT pi()
);

DROP TABLE IF EXISTS int2_pk;
CREATE TABLE int2_pk (
    int2_pk int2 PRIMARY KEY
);

DROP TABLE IF EXISTS int2_pk_ref;
CREATE TABLE int2_pk_ref (
    int2_pk_ref int2 PRIMARY KEY REFERENCES int2_ref(int2_ref)
);

DROP TABLE IF EXISTS int2_pk_def_const;
CREATE TABLE int2_pk_def_const (
    int2_pk_def_const int2 PRIMARY KEY DEFAULT 42
);

DROP TABLE IF EXISTS int2_pk_def_func;
CREATE TABLE int2_pk_def_func (
    int2_pk_def_func int2 PRIMARY KEY DEFAULT pi()
);

DROP TABLE IF EXISTS int2_nn_pk;
CREATE TABLE int2_nn_pk (
    int2_nn_pk int2 NOT NULL PRIMARY KEY
);

DROP TABLE IF EXISTS int2_nn_unique_check_pk;
CREATE TABLE int2_nn_unique_check_pk (
    int2_nn_unique_check_pk int2 PRIMARY KEY NOT NULL UNIQUE CHECK ( int2_nn_unique_check_pk > 0)
);

DROP TABLE IF EXISTS int2_nn_unique_check_pk_ref;
CREATE TABLE int2_nn_unique_check_pk_ref (
    int2_nn_unique_check_pk_ref int2 PRIMARY KEY NOT NULL UNIQUE CHECK ( int2_nn_unique_check_pk_ref > 0) REFERENCES int2_ref(int2_ref)
);

DROP TABLE IF EXISTS int2_unique_pk;
CREATE TABLE int2_unique_pk (
    int2_unique_pk int2 PRIMARY KEY UNIQUE
);

DROP TABLE IF EXISTS int2_unique_check_pk;
CREATE TABLE int2_unique_check_pk (
    int2_unique_check_pk int2 PRIMARY KEY UNIQUE CHECK ( int2_unique_check_pk > 0 )
);

DROP TABLE IF EXISTS int2_unique_check_pk_ref;
CREATE TABLE int2_unique_check_pk_ref (
    int2_unique_check_pk_ref int2 PRIMARY KEY UNIQUE CHECK ( int2_unique_check_pk_ref > 0) REFERENCES int2_ref(int2_ref)
);

DROP TABLE IF EXISTS int2_check_pk;
CREATE TABLE int2_check_pk (
    int2_check_pk int2 PRIMARY KEY CHECK ( int2_check_pk > 0 )
);

DROP TABLE IF EXISTS int2_def_const_unique_check_pk;
CREATE TABLE int2_def_const_unique_check_pk (
    int2_def_const_unique_check_pk int2 PRIMARY KEY UNIQUE CHECK ( int2_def_const_unique_check_pk > 0 ) DEFAULT 42
);

DROP TABLE IF EXISTS int2_def_const_unique_check_pk_ref;
CREATE TABLE int2_def_const_unique_check_pk_ref (
    int2_def_const_unique_check_pk_ref int2 PRIMARY KEY UNIQUE CHECK ( int2_def_const_unique_check_pk_ref > 0 ) DEFAULT 42 REFERENCES int2_ref(int2_ref)
);

DROP TABLE IF EXISTS int2_def_func_unique_check_pk;
CREATE TABLE int2_def_func_unique_check_pk (
    int2_def_func_unique_check_pk int2 PRIMARY KEY UNIQUE CHECK ( int2_def_func_unique_check_pk > 0 ) DEFAULT pi()
);

DROP TABLE IF EXISTS int2_def_func_unique_check_pk_ref;
CREATE TABLE int2_def_func_unique_check_pk_ref (
    int2_def_func_unique_check_pk_ref int2 PRIMARY KEY UNIQUE CHECK ( int2_def_func_unique_check_pk_ref > 0 ) DEFAULT pi() REFERENCES int2_ref(int2_ref)
);
