DROP TABLE IF EXISTS float_ref CASCADE;
CREATE TABLE float_ref (
    float_ref float UNIQUE
);

DROP TABLE IF EXISTS float;
CREATE TABLE float (
    float float,
    float_nn float NOT NULL,
    float_nn_unique float NOT NULL UNIQUE,
    float_nn_check float NOT NULL CHECK ( float > 0 ),
    float_nn_ref float NOT NULL REFERENCES float_ref(float_ref),
    float_nn_def_const float NOT NULL DEFAULT 42,
    float_nn_def_func float NOT NULL DEFAULT pi(),
    float_nn_unique_check float NOT NULL UNIQUE CHECK ( float > 0 ),

    float_unique float UNIQUE,
    float_unique_check float UNIQUE CHECK ( float > 0 ),
    float_unique_ref float UNIQUE REFERENCES float_ref(float_ref),
    float_unique_def_const float UNIQUE DEFAULT 42,
    float_unique_def_func float UNIQUE DEFAULT pi(),

    float_check float CHECK ( float > 0 ),
    float_check_ref float CHECK ( float > 0 ) REFERENCES float_ref(float_ref),
    float_check_def_const float CHECK ( float > 0 ) DEFAULT 42,
    float_check_def_func float CHECK ( float > 0 ) DEFAULT pi(),

    float_ref float REFERENCES float_ref(float_ref),
    float_ref_def_const float REFERENCES float_ref(float_ref) DEFAULT 42,
    float_ref_def_func float REFERENCES float_ref(float_ref) DEFAULT pi(),
    float_ref_unique_check float UNIQUE CHECK ( float > 0 ) REFERENCES float_ref(float_ref),

    float_def_const float DEFAULT 42,
    float_def_const_unique_check float UNIQUE CHECK ( float > 0 )DEFAULT 42,

    float_def_func float DEFAULT pi(),
    float_def_func_unique_check float UNIQUE CHECK ( float > 0 ) DEFAULT pi()
);

DROP TABLE IF EXISTS float_pk;
CREATE TABLE float_pk (
    float_pk float PRIMARY KEY
);

DROP TABLE IF EXISTS float_pk_ref;
CREATE TABLE float_pk_ref (
    float_pk_ref float PRIMARY KEY REFERENCES float_ref(float_ref)
);

DROP TABLE IF EXISTS float_pk_def_const;
CREATE TABLE float_pk_def_const (
    float_pk_def_const float PRIMARY KEY DEFAULT 42
);

DROP TABLE IF EXISTS float_pk_def_func;
CREATE TABLE float_pk_def_func (
    float_pk_def_func float PRIMARY KEY DEFAULT pi()
);

DROP TABLE IF EXISTS float_nn_pk;
CREATE TABLE float_nn_pk (
    float_nn_pk float NOT NULL PRIMARY KEY
);

DROP TABLE IF EXISTS float_nn_unique_check_pk;
CREATE TABLE float_nn_unique_check_pk (
    float_nn_unique_check_pk float PRIMARY KEY NOT NULL UNIQUE CHECK ( float_nn_unique_check_pk > 0)
);

DROP TABLE IF EXISTS float_nn_unique_check_pk_ref;
CREATE TABLE float_nn_unique_check_pk_ref (
    float_nn_unique_check_pk_ref float PRIMARY KEY NOT NULL UNIQUE CHECK ( float_nn_unique_check_pk_ref > 0) REFERENCES float_ref(float_ref)
);

DROP TABLE IF EXISTS float_unique_pk;
CREATE TABLE float_unique_pk (
    float_unique_pk float PRIMARY KEY UNIQUE
);

DROP TABLE IF EXISTS float_unique_check_pk;
CREATE TABLE float_unique_check_pk (
    float_unique_check_pk float PRIMARY KEY UNIQUE CHECK ( float_unique_check_pk > 0 )
);

DROP TABLE IF EXISTS float_unique_check_pk_ref;
CREATE TABLE float_unique_check_pk_ref (
    float_unique_check_pk_ref float PRIMARY KEY UNIQUE CHECK ( float_unique_check_pk_ref > 0) REFERENCES float_ref(float_ref)
);

DROP TABLE IF EXISTS float_check_pk;
CREATE TABLE float_check_pk (
    float_check_pk float PRIMARY KEY CHECK ( float_check_pk > 0 )
);

DROP TABLE IF EXISTS float_def_const_unique_check_pk;
CREATE TABLE float_def_const_unique_check_pk (
    float_def_const_unique_check_pk float PRIMARY KEY UNIQUE CHECK ( float_def_const_unique_check_pk > 0 ) DEFAULT 42
);

DROP TABLE IF EXISTS float_def_const_unique_check_pk_ref;
CREATE TABLE float_def_const_unique_check_pk_ref (
    float_def_const_unique_check_pk_ref float PRIMARY KEY UNIQUE CHECK ( float_def_const_unique_check_pk_ref > 0 ) DEFAULT 42 REFERENCES float_ref(float_ref)
);

DROP TABLE IF EXISTS float_def_func_unique_check_pk;
CREATE TABLE float_def_func_unique_check_pk (
    float_def_func_unique_check_pk float PRIMARY KEY UNIQUE CHECK ( float_def_func_unique_check_pk > 0 ) DEFAULT pi()
);

DROP TABLE IF EXISTS float_def_func_unique_check_pk_ref;
CREATE TABLE float_def_func_unique_check_pk_ref (
    float_def_func_unique_check_pk_ref float PRIMARY KEY UNIQUE CHECK ( float_def_func_unique_check_pk_ref > 0 ) DEFAULT pi() REFERENCES float_ref(float_ref)
);
