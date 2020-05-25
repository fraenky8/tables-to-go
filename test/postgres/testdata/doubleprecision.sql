DROP TABLE IF EXISTS double_precision_ref CASCADE;
CREATE TABLE double_precision_ref (
    double_precision_ref double precision UNIQUE
);

DROP TABLE IF EXISTS double_precision;
CREATE TABLE double_precision (
    double_precision double precision,
    double_precision_nn double precision NOT NULL,
    double_precision_nn_unique double precision NOT NULL UNIQUE,
    double_precision_nn_check double precision NOT NULL CHECK ( double_precision > 0 ),
    double_precision_nn_ref double precision NOT NULL REFERENCES double_precision_ref(double_precision_ref),
    double_precision_nn_def_const double precision NOT NULL DEFAULT 42,
    double_precision_nn_def_func double precision NOT NULL DEFAULT pi(),
    double_precision_nn_unique_check double precision NOT NULL UNIQUE CHECK ( double_precision > 0 ),

    double_precision_unique double precision UNIQUE,
    double_precision_unique_check double precision UNIQUE CHECK ( double_precision > 0 ),
    double_precision_unique_ref double precision UNIQUE REFERENCES double_precision_ref(double_precision_ref),
    double_precision_unique_def_const double precision UNIQUE DEFAULT 42,
    double_precision_unique_def_func double precision UNIQUE DEFAULT pi(),

    double_precision_check double precision CHECK ( double_precision > 0 ),
    double_precision_check_ref double precision CHECK ( double_precision > 0 ) REFERENCES double_precision_ref(double_precision_ref),
    double_precision_check_def_const double precision CHECK ( double_precision > 0 ) DEFAULT 42,
    double_precision_check_def_func double precision CHECK ( double_precision > 0 ) DEFAULT pi(),

    double_precision_ref double precision REFERENCES double_precision_ref(double_precision_ref),
    double_precision_ref_def_const double precision REFERENCES double_precision_ref(double_precision_ref) DEFAULT 42,
    double_precision_ref_def_func double precision REFERENCES double_precision_ref(double_precision_ref) DEFAULT pi(),
    double_precision_ref_unique_check double precision UNIQUE CHECK ( double_precision > 0 ) REFERENCES double_precision_ref(double_precision_ref),

    double_precision_def_const double precision DEFAULT 42,
    double_precision_def_const_unique_check double precision UNIQUE CHECK ( double_precision > 0 )DEFAULT 42,

    double_precision_def_func double precision DEFAULT pi(),
    double_precision_def_func_unique_check double precision UNIQUE CHECK ( double_precision > 0 ) DEFAULT pi()
);

DROP TABLE IF EXISTS double_precision_pk;
CREATE TABLE double_precision_pk (
    double_precision_pk double precision PRIMARY KEY
);

DROP TABLE IF EXISTS double_precision_pk_ref;
CREATE TABLE double_precision_pk_ref (
    double_precision_pk_ref double precision PRIMARY KEY REFERENCES double_precision_ref(double_precision_ref)
);

DROP TABLE IF EXISTS double_precision_pk_def_const;
CREATE TABLE double_precision_pk_def_const (
    double_precision_pk_def_const double precision PRIMARY KEY DEFAULT 42
);

DROP TABLE IF EXISTS double_precision_pk_def_func;
CREATE TABLE double_precision_pk_def_func (
    double_precision_pk_def_func double precision PRIMARY KEY DEFAULT pi()
);

DROP TABLE IF EXISTS double_precision_nn_pk;
CREATE TABLE double_precision_nn_pk (
    double_precision_nn_pk double precision NOT NULL PRIMARY KEY
);

DROP TABLE IF EXISTS double_precision_nn_unique_check_pk;
CREATE TABLE double_precision_nn_unique_check_pk (
    double_precision_nn_unique_check_pk double precision PRIMARY KEY NOT NULL UNIQUE CHECK ( double_precision_nn_unique_check_pk > 0)
);

DROP TABLE IF EXISTS double_precision_nn_unique_check_pk_ref;
CREATE TABLE double_precision_nn_unique_check_pk_ref (
    double_precision_nn_unique_check_pk_ref double precision PRIMARY KEY NOT NULL UNIQUE CHECK ( double_precision_nn_unique_check_pk_ref > 0) REFERENCES double_precision_ref(double_precision_ref)
);

DROP TABLE IF EXISTS double_precision_unique_pk;
CREATE TABLE double_precision_unique_pk (
    double_precision_unique_pk double precision PRIMARY KEY UNIQUE
);

DROP TABLE IF EXISTS double_precision_unique_check_pk;
CREATE TABLE double_precision_unique_check_pk (
    double_precision_unique_check_pk double precision PRIMARY KEY UNIQUE CHECK ( double_precision_unique_check_pk > 0 )
);

DROP TABLE IF EXISTS double_precision_unique_check_pk_ref;
CREATE TABLE double_precision_unique_check_pk_ref (
    double_precision_unique_check_pk_ref double precision PRIMARY KEY UNIQUE CHECK ( double_precision_unique_check_pk_ref > 0) REFERENCES double_precision_ref(double_precision_ref)
);

DROP TABLE IF EXISTS double_precision_check_pk;
CREATE TABLE double_precision_check_pk (
    double_precision_check_pk double precision PRIMARY KEY CHECK ( double_precision_check_pk > 0 )
);

DROP TABLE IF EXISTS double_precision_def_const_unique_check_pk;
CREATE TABLE double_precision_def_const_unique_check_pk (
    double_precision_def_const_unique_check_pk double precision PRIMARY KEY UNIQUE CHECK ( double_precision_def_const_unique_check_pk > 0 ) DEFAULT 42
);

DROP TABLE IF EXISTS double_precision_def_const_unique_check_pk_ref;
CREATE TABLE double_precision_def_const_unique_check_pk_ref (
    double_precision_def_const_unique_check_pk_ref double precision PRIMARY KEY UNIQUE CHECK ( double_precision_def_const_unique_check_pk_ref > 0 ) DEFAULT 42 REFERENCES double_precision_ref(double_precision_ref)
);

DROP TABLE IF EXISTS double_precision_def_func_unique_check_pk;
CREATE TABLE double_precision_def_func_unique_check_pk (
    double_precision_def_func_unique_check_pk double precision PRIMARY KEY UNIQUE CHECK ( double_precision_def_func_unique_check_pk > 0 ) DEFAULT pi()
);

DROP TABLE IF EXISTS double_precision_def_func_unique_check_pk_ref;
CREATE TABLE double_precision_def_func_unique_check_pk_ref (
    double_precision_def_func_unique_check_pk_ref double precision PRIMARY KEY UNIQUE CHECK ( double_precision_def_func_unique_check_pk_ref > 0 ) DEFAULT pi() REFERENCES double_precision_ref(double_precision_ref)
);
