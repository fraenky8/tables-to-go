DROP TABLE IF EXISTS float8_ref CASCADE;
CREATE TABLE float8_ref (
    float8_ref float8 UNIQUE
);

DROP TABLE IF EXISTS float8;
CREATE TABLE float8 (
    float8 float8,
    float8_nn float8 NOT NULL,
    float8_nn_unique float8 NOT NULL UNIQUE,
    float8_nn_check float8 NOT NULL CHECK ( float8 > 0 ),
    float8_nn_ref float8 NOT NULL REFERENCES float8_ref(float8_ref),
    float8_nn_def_const float8 NOT NULL DEFAULT 42,
    float8_nn_def_func float8 NOT NULL DEFAULT pi(),
    float8_nn_unique_check float8 NOT NULL UNIQUE CHECK ( float8 > 0 ),

    float8_unique float8 UNIQUE,
    float8_unique_check float8 UNIQUE CHECK ( float8 > 0 ),
    float8_unique_ref float8 UNIQUE REFERENCES float8_ref(float8_ref),
    float8_unique_def_const float8 UNIQUE DEFAULT 42,
    float8_unique_def_func float8 UNIQUE DEFAULT pi(),

    float8_check float8 CHECK ( float8 > 0 ),
    float8_check_ref float8 CHECK ( float8 > 0 ) REFERENCES float8_ref(float8_ref),
    float8_check_def_const float8 CHECK ( float8 > 0 ) DEFAULT 42,
    float8_check_def_func float8 CHECK ( float8 > 0 ) DEFAULT pi(),

    float8_ref float8 REFERENCES float8_ref(float8_ref),
    float8_ref_def_const float8 REFERENCES float8_ref(float8_ref) DEFAULT 42,
    float8_ref_def_func float8 REFERENCES float8_ref(float8_ref) DEFAULT pi(),
    float8_ref_unique_check float8 UNIQUE CHECK ( float8 > 0 ) REFERENCES float8_ref(float8_ref),

    float8_def_const float8 DEFAULT 42,
    float8_def_const_unique_check float8 UNIQUE CHECK ( float8 > 0 )DEFAULT 42,

    float8_def_func float8 DEFAULT pi(),
    float8_def_func_unique_check float8 UNIQUE CHECK ( float8 > 0 ) DEFAULT pi()
);

DROP TABLE IF EXISTS float8_pk;
CREATE TABLE float8_pk (
    float8_pk float8 PRIMARY KEY
);

DROP TABLE IF EXISTS float8_pk_ref;
CREATE TABLE float8_pk_ref (
    float8_pk_ref float8 PRIMARY KEY REFERENCES float8_ref(float8_ref)
);

DROP TABLE IF EXISTS float8_pk_def_const;
CREATE TABLE float8_pk_def_const (
    float8_pk_def_const float8 PRIMARY KEY DEFAULT 42
);

DROP TABLE IF EXISTS float8_pk_def_func;
CREATE TABLE float8_pk_def_func (
    float8_pk_def_func float8 PRIMARY KEY DEFAULT pi()
);

DROP TABLE IF EXISTS float8_nn_pk;
CREATE TABLE float8_nn_pk (
    float8_nn_pk float8 NOT NULL PRIMARY KEY
);

DROP TABLE IF EXISTS float8_nn_unique_check_pk;
CREATE TABLE float8_nn_unique_check_pk (
    float8_nn_unique_check_pk float8 PRIMARY KEY NOT NULL UNIQUE CHECK ( float8_nn_unique_check_pk > 0)
);

DROP TABLE IF EXISTS float8_nn_unique_check_pk_ref;
CREATE TABLE float8_nn_unique_check_pk_ref (
    float8_nn_unique_check_pk_ref float8 PRIMARY KEY NOT NULL UNIQUE CHECK ( float8_nn_unique_check_pk_ref > 0) REFERENCES float8_ref(float8_ref)
);

DROP TABLE IF EXISTS float8_unique_pk;
CREATE TABLE float8_unique_pk (
    float8_unique_pk float8 PRIMARY KEY UNIQUE
);

DROP TABLE IF EXISTS float8_unique_check_pk;
CREATE TABLE float8_unique_check_pk (
    float8_unique_check_pk float8 PRIMARY KEY UNIQUE CHECK ( float8_unique_check_pk > 0 )
);

DROP TABLE IF EXISTS float8_unique_check_pk_ref;
CREATE TABLE float8_unique_check_pk_ref (
    float8_unique_check_pk_ref float8 PRIMARY KEY UNIQUE CHECK ( float8_unique_check_pk_ref > 0) REFERENCES float8_ref(float8_ref)
);

DROP TABLE IF EXISTS float8_check_pk;
CREATE TABLE float8_check_pk (
    float8_check_pk float8 PRIMARY KEY CHECK ( float8_check_pk > 0 )
);

DROP TABLE IF EXISTS float8_def_const_unique_check_pk;
CREATE TABLE float8_def_const_unique_check_pk (
    float8_def_const_unique_check_pk float8 PRIMARY KEY UNIQUE CHECK ( float8_def_const_unique_check_pk > 0 ) DEFAULT 42
);

DROP TABLE IF EXISTS float8_def_const_unique_check_pk_ref;
CREATE TABLE float8_def_const_unique_check_pk_ref (
    float8_def_const_unique_check_pk_ref float8 PRIMARY KEY UNIQUE CHECK ( float8_def_const_unique_check_pk_ref > 0 ) DEFAULT 42 REFERENCES float8_ref(float8_ref)
);

DROP TABLE IF EXISTS float8_def_func_unique_check_pk;
CREATE TABLE float8_def_func_unique_check_pk (
    float8_def_func_unique_check_pk float8 PRIMARY KEY UNIQUE CHECK ( float8_def_func_unique_check_pk > 0 ) DEFAULT pi()
);

DROP TABLE IF EXISTS float8_def_func_unique_check_pk_ref;
CREATE TABLE float8_def_func_unique_check_pk_ref (
    float8_def_func_unique_check_pk_ref float8 PRIMARY KEY UNIQUE CHECK ( float8_def_func_unique_check_pk_ref > 0 ) DEFAULT pi() REFERENCES float8_ref(float8_ref)
);
