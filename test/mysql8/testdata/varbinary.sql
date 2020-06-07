DROP TABLE IF EXISTS varbinary_ref CASCADE;
CREATE TABLE varbinary_ref (
    varbinary_ref varbinary(10) UNIQUE
);

DROP TABLE IF EXISTS varbinary_table;
CREATE TABLE varbinary_table (
    col varbinary(10),
    
    varbinary_cap varbinary(10),
    varbinary_nn varbinary(10) NOT NULL,
    varbinary_nn_unique varbinary(10) NOT NULL UNIQUE,
    varbinary_nn_check_cmp varbinary(10) NOT NULL CHECK ( varbinary_nn_check_cmp = '42' ),
    varbinary_nn_check_fn varbinary(10) NOT NULL CHECK ( length(varbinary_nn_check_fn) > 0 ),
    varbinary_nn_ref varbinary(10) NOT NULL REFERENCES varbinary_ref(varbinary_ref),
    varbinary_nn_def_const varbinary(10) NOT NULL DEFAULT ('42'),
    varbinary_nn_def_func varbinary(10) NOT NULL DEFAULT (pi()),
    varbinary_nn_unique_check varbinary(10) NOT NULL UNIQUE CHECK ( length(varbinary_nn_unique_check) > 0 ),

    varbinary_unique varbinary(10) UNIQUE,
    varbinary_unique_check varbinary(10) UNIQUE CHECK ( length(varbinary_unique_check) > 0 ),
    varbinary_unique_ref varbinary(10) UNIQUE REFERENCES varbinary_ref(varbinary_ref),
    varbinary_unique_def_const varbinary(10) UNIQUE DEFAULT ('42'),
    varbinary_unique_def_func varbinary(10) UNIQUE DEFAULT (pi()),

    varbinary_check varbinary(10) CHECK ( length(varbinary_check) > 0 ),
    varbinary_check_ref varbinary(10) CHECK ( length(varbinary_check_ref) > 0 ) REFERENCES varbinary_ref(varbinary_ref),
    varbinary_check_def_const varbinary(10) CHECK ( length(varbinary_check_def_const) > 0 ) DEFAULT ('42'),
    varbinary_check_def_func varbinary(10) CHECK ( length(varbinary_check_def_func) > 0 ) DEFAULT (pi()),

    varbinary_ref varbinary(10) REFERENCES varbinary_ref(varbinary_ref),
    varbinary_ref_unique_check varbinary(10) UNIQUE CHECK ( length(varbinary_ref_unique_check) > 0 ) REFERENCES varbinary_ref(varbinary_ref),

    varbinary_def_const varbinary(10) DEFAULT ('42'),
    varbinary_def_const_unique_check varbinary(10) UNIQUE CHECK ( length(varbinary_def_const_unique_check) > 0 ) DEFAULT ('42'),

    varbinary_def_func varbinary(10) DEFAULT (pi()),
    varbinary_def_func_unique_check varbinary(10) UNIQUE CHECK ( length(varbinary_def_func_unique_check) > 0 ) DEFAULT (pi())
);

DROP TABLE IF EXISTS varbinary_pk;
CREATE TABLE varbinary_pk (
    varbinary_pk varbinary(10) PRIMARY KEY
);

DROP TABLE IF EXISTS varbinary_pk_ref;
CREATE TABLE varbinary_pk_ref (
    varbinary_pk_ref varbinary(10) PRIMARY KEY REFERENCES varbinary_ref(varbinary_ref)
);

DROP TABLE IF EXISTS varbinary_pk_def_const;
CREATE TABLE varbinary_pk_def_const (
    varbinary_pk_def_const varbinary(10) PRIMARY KEY DEFAULT ('42')
);

DROP TABLE IF EXISTS varbinary_pk_def_func;
CREATE TABLE varbinary_pk_def_func (
    varbinary_pk_def_func varbinary(10) PRIMARY KEY DEFAULT (pi())
);

DROP TABLE IF EXISTS varbinary_nn_pk;
CREATE TABLE varbinary_nn_pk (
    varbinary_nn_pk varbinary(10) NOT NULL PRIMARY KEY
);

DROP TABLE IF EXISTS varbinary_nn_unique_check_pk;
CREATE TABLE varbinary_nn_unique_check_pk (
    varbinary_nn_unique_check_pk varbinary(10) PRIMARY KEY NOT NULL UNIQUE CHECK ( length(varbinary_nn_unique_check_pk) > 0)
);

DROP TABLE IF EXISTS varbinary_nn_unique_check_pk_ref;
CREATE TABLE varbinary_nn_unique_check_pk_ref (
    varbinary_nn_unique_check_pk_ref varbinary(10) PRIMARY KEY NOT NULL UNIQUE CHECK ( length(varbinary_nn_unique_check_pk_ref) > 0) REFERENCES varbinary_ref(varbinary_ref)
);

DROP TABLE IF EXISTS varbinary_unique_pk;
CREATE TABLE varbinary_unique_pk (
    varbinary_unique_pk varbinary(10) PRIMARY KEY UNIQUE
);

DROP TABLE IF EXISTS varbinary_unique_check_pk;
CREATE TABLE varbinary_unique_check_pk (
    varbinary_unique_check_pk varbinary(10) PRIMARY KEY UNIQUE CHECK ( length(varbinary_unique_check_pk) > 0 )
);

DROP TABLE IF EXISTS varbinary_unique_check_pk_ref;
CREATE TABLE varbinary_unique_check_pk_ref (
    varbinary_unique_check_pk_ref varbinary(10) PRIMARY KEY UNIQUE CHECK ( length(varbinary_unique_check_pk_ref) > 0) REFERENCES varbinary_ref(varbinary_ref)
);

DROP TABLE IF EXISTS varbinary_check_pk;
CREATE TABLE varbinary_check_pk (
    varbinary_check_pk varbinary(10) PRIMARY KEY CHECK ( length(varbinary_check_pk) > 0 )
);

DROP TABLE IF EXISTS varbinary_def_const_unique_check_pk;
CREATE TABLE varbinary_def_const_unique_check_pk (
    varbinary_def_const_unique_check_pk varbinary(10) PRIMARY KEY UNIQUE CHECK ( length(varbinary_def_const_unique_check_pk) > 0 ) DEFAULT ('42')
);

DROP TABLE IF EXISTS varbinary_def_const_unique_check_pk_ref;
CREATE TABLE varbinary_def_const_unique_check_pk_ref (
    varbinary_def_const_unique_check_pk_ref varbinary(10) PRIMARY KEY UNIQUE CHECK ( length(varbinary_def_const_unique_check_pk_ref) > 0 ) DEFAULT ('42') REFERENCES varbinary_ref(varbinary_ref)
);

DROP TABLE IF EXISTS varbinary_def_func_unique_check_pk;
CREATE TABLE varbinary_def_func_unique_check_pk (
    varbinary_def_func_unique_check_pk varbinary(10) PRIMARY KEY UNIQUE CHECK ( length(varbinary_def_func_unique_check_pk) > 0 ) DEFAULT (pi())
);

DROP TABLE IF EXISTS varbinary_def_func_unique_check_pk_ref;
CREATE TABLE varbinary_def_func_unique_check_pk_ref (
    varbinary_def_func_unique_check_pk_ref varbinary(10) PRIMARY KEY UNIQUE CHECK ( length(varbinary_def_func_unique_check_pk_ref) > 0 ) DEFAULT (pi()) REFERENCES varbinary_ref(varbinary_ref)
);
