DROP TABLE IF EXISTS varchar_ref CASCADE;
CREATE TABLE varchar_ref (
    varchar_ref varchar UNIQUE
);

DROP TABLE IF EXISTS varchar;
CREATE TABLE varchar (
    varchar varchar,
    varchar_cap varchar(255),
    varchar_nn varchar NOT NULL,
    varchar_nn_unique varchar NOT NULL UNIQUE,
    varchar_nn_check_cmp varchar NOT NULL CHECK ( varchar = '42' ),
    varchar_nn_check_fn varchar NOT NULL CHECK ( length(varchar) > 0 ),
    varchar_nn_ref varchar NOT NULL REFERENCES varchar_ref(varchar_ref),
    varchar_nn_def_const varchar NOT NULL DEFAULT '42',
    varchar_nn_def_func varchar NOT NULL DEFAULT pi(),
    varchar_nn_unique_check varchar NOT NULL UNIQUE CHECK ( length(varchar) > 0 ),

    varchar_unique varchar UNIQUE,
    varchar_unique_check varchar UNIQUE CHECK ( length(varchar) > 0 ),
    varchar_unique_ref varchar UNIQUE REFERENCES varchar_ref(varchar_ref),
    varchar_unique_def_const varchar UNIQUE DEFAULT '42',
    varchar_unique_def_func varchar UNIQUE DEFAULT pi(),

    varchar_check varchar CHECK ( length(varchar) > 0 ),
    varchar_check_ref varchar CHECK ( length(varchar) > 0 ) REFERENCES varchar_ref(varchar_ref),
    varchar_check_def_const varchar CHECK ( length(varchar) > 0 ) DEFAULT '42',
    varchar_check_def_func varchar CHECK ( length(varchar) > 0 ) DEFAULT pi(),

    varchar_ref varchar REFERENCES varchar_ref(varchar_ref),
    varchar_ref_def_const varchar REFERENCES varchar_ref(varchar_ref) DEFAULT '42',
    varchar_ref_def_func varchar REFERENCES varchar_ref(varchar_ref) DEFAULT pi(),
    varchar_ref_unique_check varchar UNIQUE CHECK ( length(varchar) > 0 ) REFERENCES varchar_ref(varchar_ref),

    varchar_def_const varchar DEFAULT '42',
    varchar_def_const_unique_check varchar UNIQUE CHECK ( length(varchar) > 0 ) DEFAULT '42',

    varchar_def_func varchar DEFAULT pi(),
    varchar_def_func_unique_check varchar UNIQUE CHECK ( length(varchar) > 0 ) DEFAULT pi()
);

DROP TABLE IF EXISTS varchar_pk;
CREATE TABLE varchar_pk (
    varchar_pk varchar PRIMARY KEY
);

DROP TABLE IF EXISTS varchar_pk_ref;
CREATE TABLE varchar_pk_ref (
    varchar_pk_ref varchar PRIMARY KEY REFERENCES varchar_ref(varchar_ref)
);

DROP TABLE IF EXISTS varchar_pk_def_const;
CREATE TABLE varchar_pk_def_const (
    varchar_pk_def_const varchar PRIMARY KEY DEFAULT '42'
);

DROP TABLE IF EXISTS varchar_pk_def_func;
CREATE TABLE varchar_pk_def_func (
    varchar_pk_def_func varchar PRIMARY KEY DEFAULT pi()
);

DROP TABLE IF EXISTS varchar_nn_pk;
CREATE TABLE varchar_nn_pk (
    varchar_nn_pk varchar NOT NULL PRIMARY KEY
);

DROP TABLE IF EXISTS varchar_nn_unique_check_pk;
CREATE TABLE varchar_nn_unique_check_pk (
    varchar_nn_unique_check_pk varchar PRIMARY KEY NOT NULL UNIQUE CHECK ( length(varchar_nn_unique_check_pk) > 0)
);

DROP TABLE IF EXISTS varchar_nn_unique_check_pk_ref;
CREATE TABLE varchar_nn_unique_check_pk_ref (
    varchar_nn_unique_check_pk_ref varchar PRIMARY KEY NOT NULL UNIQUE CHECK ( length(varchar_nn_unique_check_pk_ref) > 0) REFERENCES varchar_ref(varchar_ref)
);

DROP TABLE IF EXISTS varchar_unique_pk;
CREATE TABLE varchar_unique_pk (
    varchar_unique_pk varchar PRIMARY KEY UNIQUE
);

DROP TABLE IF EXISTS varchar_unique_check_pk;
CREATE TABLE varchar_unique_check_pk (
    varchar_unique_check_pk varchar PRIMARY KEY UNIQUE CHECK ( length(varchar_unique_check_pk) > 0 )
);

DROP TABLE IF EXISTS varchar_unique_check_pk_ref;
CREATE TABLE varchar_unique_check_pk_ref (
    varchar_unique_check_pk_ref varchar PRIMARY KEY UNIQUE CHECK ( length(varchar_unique_check_pk_ref) > 0) REFERENCES varchar_ref(varchar_ref)
);

DROP TABLE IF EXISTS varchar_check_pk;
CREATE TABLE varchar_check_pk (
    varchar_check_pk varchar PRIMARY KEY CHECK ( length(varchar_check_pk) > 0 )
);

DROP TABLE IF EXISTS varchar_def_const_unique_check_pk;
CREATE TABLE varchar_def_const_unique_check_pk (
    varchar_def_const_unique_check_pk varchar PRIMARY KEY UNIQUE CHECK ( length(varchar_def_const_unique_check_pk) > 0 ) DEFAULT '42'
);

DROP TABLE IF EXISTS varchar_def_const_unique_check_pk_ref;
CREATE TABLE varchar_def_const_unique_check_pk_ref (
    varchar_def_const_unique_check_pk_ref varchar PRIMARY KEY UNIQUE CHECK ( length(varchar_def_const_unique_check_pk_ref) > 0 ) DEFAULT '42' REFERENCES varchar_ref(varchar_ref)
);

DROP TABLE IF EXISTS varchar_def_func_unique_check_pk;
CREATE TABLE varchar_def_func_unique_check_pk (
    varchar_def_func_unique_check_pk varchar PRIMARY KEY UNIQUE CHECK ( length(varchar_def_func_unique_check_pk) > 0 ) DEFAULT pi()
);

DROP TABLE IF EXISTS varchar_def_func_unique_check_pk_ref;
CREATE TABLE varchar_def_func_unique_check_pk_ref (
    varchar_def_func_unique_check_pk_ref varchar PRIMARY KEY UNIQUE CHECK ( length(varchar_def_func_unique_check_pk_ref) > 0 ) DEFAULT pi() REFERENCES varchar_ref(varchar_ref)
);
