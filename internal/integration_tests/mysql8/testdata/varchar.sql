DROP TABLE IF EXISTS varchar_ref CASCADE;
CREATE TABLE varchar_ref (
    varchar_ref varchar(10) UNIQUE
);

DROP TABLE IF EXISTS varchar_table;
CREATE TABLE varchar_table (
    col varchar(10),
    
    varchar_cap varchar(10),
    varchar_nn varchar(10) NOT NULL,
    varchar_nn_unique varchar(10) NOT NULL UNIQUE,
    varchar_nn_check_cmp varchar(10) NOT NULL CHECK ( varchar_nn_check_cmp = '42' ),
    varchar_nn_check_fn varchar(10) NOT NULL CHECK ( length(varchar_nn_check_fn) > 0 ),
    varchar_nn_ref varchar(10) NOT NULL REFERENCES varchar_ref(varchar_ref),
    varchar_nn_def_const varchar(10) NOT NULL DEFAULT ('42'),
    varchar_nn_def_func varchar(10) NOT NULL DEFAULT (pi()),
    varchar_nn_unique_check varchar(10) NOT NULL UNIQUE CHECK ( length(varchar_nn_unique_check) > 0 ),

    varchar_unique varchar(10) UNIQUE,
    varchar_unique_check varchar(10) UNIQUE CHECK ( length(varchar_unique_check) > 0 ),
    varchar_unique_ref varchar(10) UNIQUE REFERENCES varchar_ref(varchar_ref),
    varchar_unique_def_const varchar(10) UNIQUE DEFAULT ('42'),
    varchar_unique_def_func varchar(10) UNIQUE DEFAULT (pi()),

    varchar_check varchar(10) CHECK ( length(varchar_check) > 0 ),
    varchar_check_ref varchar(10) CHECK ( length(varchar_check_ref) > 0 ) REFERENCES varchar_ref(varchar_ref),
    varchar_check_def_const varchar(10) CHECK ( length(varchar_check_def_const) > 0 ) DEFAULT ('42'),
    varchar_check_def_func varchar(10) CHECK ( length(varchar_check_def_func) > 0 ) DEFAULT (pi()),

    varchar_ref varchar(10) REFERENCES varchar_ref(varchar_ref),
    varchar_ref_unique_check varchar(10) UNIQUE CHECK ( length(varchar_ref_unique_check) > 0 ) REFERENCES varchar_ref(varchar_ref),

    varchar_def_const varchar(10) DEFAULT ('42'),
    varchar_def_const_unique_check varchar(10) UNIQUE CHECK ( length(varchar_def_const_unique_check) > 0 ) DEFAULT ('42'),

    varchar_def_func varchar(10) DEFAULT (pi()),
    varchar_def_func_unique_check varchar(10) UNIQUE CHECK ( length(varchar_def_func_unique_check) > 0 ) DEFAULT (pi())
);

DROP TABLE IF EXISTS varchar_pk;
CREATE TABLE varchar_pk (
    varchar_pk varchar(10) PRIMARY KEY
);

DROP TABLE IF EXISTS varchar_pk_ref;
CREATE TABLE varchar_pk_ref (
    varchar_pk_ref varchar(10) PRIMARY KEY REFERENCES varchar_ref(varchar_ref)
);

DROP TABLE IF EXISTS varchar_pk_def_const;
CREATE TABLE varchar_pk_def_const (
    varchar_pk_def_const varchar(10) PRIMARY KEY DEFAULT ('42')
);

DROP TABLE IF EXISTS varchar_pk_def_func;
CREATE TABLE varchar_pk_def_func (
    varchar_pk_def_func varchar(10) PRIMARY KEY DEFAULT (pi())
);

DROP TABLE IF EXISTS varchar_nn_pk;
CREATE TABLE varchar_nn_pk (
    varchar_nn_pk varchar(10) NOT NULL PRIMARY KEY
);

DROP TABLE IF EXISTS varchar_nn_unique_check_pk;
CREATE TABLE varchar_nn_unique_check_pk (
    varchar_nn_unique_check_pk varchar(10) PRIMARY KEY NOT NULL UNIQUE CHECK ( length(varchar_nn_unique_check_pk) > 0)
);

DROP TABLE IF EXISTS varchar_nn_unique_check_pk_ref;
CREATE TABLE varchar_nn_unique_check_pk_ref (
    varchar_nn_unique_check_pk_ref varchar(10) PRIMARY KEY NOT NULL UNIQUE CHECK ( length(varchar_nn_unique_check_pk_ref) > 0) REFERENCES varchar_ref(varchar_ref)
);

DROP TABLE IF EXISTS varchar_unique_pk;
CREATE TABLE varchar_unique_pk (
    varchar_unique_pk varchar(10) PRIMARY KEY UNIQUE
);

DROP TABLE IF EXISTS varchar_unique_check_pk;
CREATE TABLE varchar_unique_check_pk (
    varchar_unique_check_pk varchar(10) PRIMARY KEY UNIQUE CHECK ( length(varchar_unique_check_pk) > 0 )
);

DROP TABLE IF EXISTS varchar_unique_check_pk_ref;
CREATE TABLE varchar_unique_check_pk_ref (
    varchar_unique_check_pk_ref varchar(10) PRIMARY KEY UNIQUE CHECK ( length(varchar_unique_check_pk_ref) > 0) REFERENCES varchar_ref(varchar_ref)
);

DROP TABLE IF EXISTS varchar_check_pk;
CREATE TABLE varchar_check_pk (
    varchar_check_pk varchar(10) PRIMARY KEY CHECK ( length(varchar_check_pk) > 0 )
);

DROP TABLE IF EXISTS varchar_def_const_unique_check_pk;
CREATE TABLE varchar_def_const_unique_check_pk (
    varchar_def_const_unique_check_pk varchar(10) PRIMARY KEY UNIQUE CHECK ( length(varchar_def_const_unique_check_pk) > 0 ) DEFAULT ('42')
);

DROP TABLE IF EXISTS varchar_def_const_unique_check_pk_ref;
CREATE TABLE varchar_def_const_unique_check_pk_ref (
    varchar_def_const_unique_check_pk_ref varchar(10) PRIMARY KEY UNIQUE CHECK ( length(varchar_def_const_unique_check_pk_ref) > 0 ) DEFAULT ('42') REFERENCES varchar_ref(varchar_ref)
);

DROP TABLE IF EXISTS varchar_def_func_unique_check_pk;
CREATE TABLE varchar_def_func_unique_check_pk (
    varchar_def_func_unique_check_pk varchar(10) PRIMARY KEY UNIQUE CHECK ( length(varchar_def_func_unique_check_pk) > 0 ) DEFAULT (pi())
);

DROP TABLE IF EXISTS varchar_def_func_unique_check_pk_ref;
CREATE TABLE varchar_def_func_unique_check_pk_ref (
    varchar_def_func_unique_check_pk_ref varchar(10) PRIMARY KEY UNIQUE CHECK ( length(varchar_def_func_unique_check_pk_ref) > 0 ) DEFAULT (pi()) REFERENCES varchar_ref(varchar_ref)
);
