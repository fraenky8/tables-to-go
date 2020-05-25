DROP TABLE IF EXISTS float4_ref CASCADE;
CREATE TABLE float4_ref (
    float4_ref float4 UNIQUE
);

DROP TABLE IF EXISTS float4;
CREATE TABLE float4 (
    float4 float4,
    float4_nn float4 NOT NULL,
    float4_nn_unique float4 NOT NULL UNIQUE,
    float4_nn_check float4 NOT NULL CHECK ( float4 > 0 ),
    float4_nn_ref float4 NOT NULL REFERENCES float4_ref(float4_ref),
    float4_nn_def_const float4 NOT NULL DEFAULT 42,
    float4_nn_def_func float4 NOT NULL DEFAULT pi(),
    float4_nn_unique_check float4 NOT NULL UNIQUE CHECK ( float4 > 0 ),

    float4_unique float4 UNIQUE,
    float4_unique_check float4 UNIQUE CHECK ( float4 > 0 ),
    float4_unique_ref float4 UNIQUE REFERENCES float4_ref(float4_ref),
    float4_unique_def_const float4 UNIQUE DEFAULT 42,
    float4_unique_def_func float4 UNIQUE DEFAULT pi(),

    float4_check float4 CHECK ( float4 > 0 ),
    float4_check_ref float4 CHECK ( float4 > 0 ) REFERENCES float4_ref(float4_ref),
    float4_check_def_const float4 CHECK ( float4 > 0 ) DEFAULT 42,
    float4_check_def_func float4 CHECK ( float4 > 0 ) DEFAULT pi(),

    float4_ref float4 REFERENCES float4_ref(float4_ref),
    float4_ref_def_const float4 REFERENCES float4_ref(float4_ref) DEFAULT 42,
    float4_ref_def_func float4 REFERENCES float4_ref(float4_ref) DEFAULT pi(),
    float4_ref_unique_check float4 UNIQUE CHECK ( float4 > 0 ) REFERENCES float4_ref(float4_ref),

    float4_def_const float4 DEFAULT 42,
    float4_def_const_unique_check float4 UNIQUE CHECK ( float4 > 0 )DEFAULT 42,

    float4_def_func float4 DEFAULT pi(),
    float4_def_func_unique_check float4 UNIQUE CHECK ( float4 > 0 ) DEFAULT pi()
);

DROP TABLE IF EXISTS float4_pk;
CREATE TABLE float4_pk (
    float4_pk float4 PRIMARY KEY
);

DROP TABLE IF EXISTS float4_pk_ref;
CREATE TABLE float4_pk_ref (
    float4_pk_ref float4 PRIMARY KEY REFERENCES float4_ref(float4_ref)
);

DROP TABLE IF EXISTS float4_pk_def_const;
CREATE TABLE float4_pk_def_const (
    float4_pk_def_const float4 PRIMARY KEY DEFAULT 42
);

DROP TABLE IF EXISTS float4_pk_def_func;
CREATE TABLE float4_pk_def_func (
    float4_pk_def_func float4 PRIMARY KEY DEFAULT pi()
);

DROP TABLE IF EXISTS float4_nn_pk;
CREATE TABLE float4_nn_pk (
    float4_nn_pk float4 NOT NULL PRIMARY KEY
);

DROP TABLE IF EXISTS float4_nn_unique_check_pk;
CREATE TABLE float4_nn_unique_check_pk (
    float4_nn_unique_check_pk float4 PRIMARY KEY NOT NULL UNIQUE CHECK ( float4_nn_unique_check_pk > 0)
);

DROP TABLE IF EXISTS float4_nn_unique_check_pk_ref;
CREATE TABLE float4_nn_unique_check_pk_ref (
    float4_nn_unique_check_pk_ref float4 PRIMARY KEY NOT NULL UNIQUE CHECK ( float4_nn_unique_check_pk_ref > 0) REFERENCES float4_ref(float4_ref)
);

DROP TABLE IF EXISTS float4_unique_pk;
CREATE TABLE float4_unique_pk (
    float4_unique_pk float4 PRIMARY KEY UNIQUE
);

DROP TABLE IF EXISTS float4_unique_check_pk;
CREATE TABLE float4_unique_check_pk (
    float4_unique_check_pk float4 PRIMARY KEY UNIQUE CHECK ( float4_unique_check_pk > 0 )
);

DROP TABLE IF EXISTS float4_unique_check_pk_ref;
CREATE TABLE float4_unique_check_pk_ref (
    float4_unique_check_pk_ref float4 PRIMARY KEY UNIQUE CHECK ( float4_unique_check_pk_ref > 0) REFERENCES float4_ref(float4_ref)
);

DROP TABLE IF EXISTS float4_check_pk;
CREATE TABLE float4_check_pk (
    float4_check_pk float4 PRIMARY KEY CHECK ( float4_check_pk > 0 )
);

DROP TABLE IF EXISTS float4_def_const_unique_check_pk;
CREATE TABLE float4_def_const_unique_check_pk (
    float4_def_const_unique_check_pk float4 PRIMARY KEY UNIQUE CHECK ( float4_def_const_unique_check_pk > 0 ) DEFAULT 42
);

DROP TABLE IF EXISTS float4_def_const_unique_check_pk_ref;
CREATE TABLE float4_def_const_unique_check_pk_ref (
    float4_def_const_unique_check_pk_ref float4 PRIMARY KEY UNIQUE CHECK ( float4_def_const_unique_check_pk_ref > 0 ) DEFAULT 42 REFERENCES float4_ref(float4_ref)
);

DROP TABLE IF EXISTS float4_def_func_unique_check_pk;
CREATE TABLE float4_def_func_unique_check_pk (
    float4_def_func_unique_check_pk float4 PRIMARY KEY UNIQUE CHECK ( float4_def_func_unique_check_pk > 0 ) DEFAULT pi()
);

DROP TABLE IF EXISTS float4_def_func_unique_check_pk_ref;
CREATE TABLE float4_def_func_unique_check_pk_ref (
    float4_def_func_unique_check_pk_ref float4 PRIMARY KEY UNIQUE CHECK ( float4_def_func_unique_check_pk_ref > 0 ) DEFAULT pi() REFERENCES float4_ref(float4_ref)
);
