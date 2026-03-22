DROP TABLE IF EXISTS year_ref CASCADE;
CREATE TABLE year_ref (
    year_ref year UNIQUE
);

DROP TABLE IF EXISTS year_table;
CREATE TABLE year_table (
    year year,
    year_nn year NOT NULL,
    year_nn_unique year NOT NULL UNIQUE,
    year_nn_check year NOT NULL CHECK ( year_nn_check > '2020' ),
    year_nn_ref year NOT NULL REFERENCES year_ref(year_ref),
    year_nn_def_const year NOT NULL DEFAULT ('2020'),
    year_nn_def_func year NOT NULL DEFAULT (now()),
    year_nn_unique_check year NOT NULL UNIQUE CHECK ( year_nn_unique_check > '2020' ),

    year_unique year UNIQUE,
    year_unique_check year UNIQUE CHECK ( year_unique_check > '2020' ),
    year_unique_ref year UNIQUE REFERENCES year_ref(year_ref),
    year_unique_def_const year UNIQUE DEFAULT ('2020'),
    year_unique_def_func year UNIQUE DEFAULT (now()),

    year_check year CHECK ( year_check > '2020' ),
    year_check_ref year CHECK ( year_check_ref > '2020' ) REFERENCES year_ref(year_ref),
    year_check_def_const year CHECK ( year_check_def_const > '2020' ) DEFAULT ('2020'),
    year_check_def_func year CHECK ( year_check_def_func > '2020' ) DEFAULT (now()),

    year_ref year REFERENCES year_ref(year_ref),
    year_ref_unique_check year UNIQUE CHECK ( year_ref_unique_check > '2020' ) REFERENCES year_ref(year_ref),

    year_def_const year DEFAULT ('2020'),
    year_def_const_unique_check year UNIQUE CHECK ( year_def_const_unique_check > '2020' ) DEFAULT ('2020'),

    year_def_func year DEFAULT (now()),
    year_def_func_unique_check year UNIQUE CHECK ( year_def_func_unique_check > '2020' ) DEFAULT (now())
);

DROP TABLE IF EXISTS year_pk;
CREATE TABLE year_pk (
    year_pk year PRIMARY KEY
);

DROP TABLE IF EXISTS year_pk_ref;
CREATE TABLE year_pk_ref (
    year_pk_ref year PRIMARY KEY REFERENCES year_ref(year_ref)
);

DROP TABLE IF EXISTS year_pk_def_const;
CREATE TABLE year_pk_def_const (
    year_pk_def_const year PRIMARY KEY DEFAULT ('2020')
);

DROP TABLE IF EXISTS year_pk_def_func;
CREATE TABLE year_pk_def_func (
    year_pk_def_func year PRIMARY KEY DEFAULT (now())
);

DROP TABLE IF EXISTS year_nn_pk;
CREATE TABLE year_nn_pk (
    year_nn_pk year NOT NULL PRIMARY KEY
);

DROP TABLE IF EXISTS year_nn_unique_check_pk;
CREATE TABLE year_nn_unique_check_pk (
    year_nn_unique_check_pk year PRIMARY KEY NOT NULL UNIQUE CHECK ( year_nn_unique_check_pk > '2020' )
);

DROP TABLE IF EXISTS year_nn_unique_check_pk_ref;
CREATE TABLE year_nn_unique_check_pk_ref (
    year_nn_unique_check_pk_ref year PRIMARY KEY NOT NULL UNIQUE CHECK ( year_nn_unique_check_pk_ref > '2020' ) REFERENCES year_ref(year_ref)
);

DROP TABLE IF EXISTS year_unique_pk;
CREATE TABLE year_unique_pk (
    year_unique_pk year PRIMARY KEY UNIQUE
);

DROP TABLE IF EXISTS year_unique_check_pk;
CREATE TABLE year_unique_check_pk (
    year_unique_check_pk year PRIMARY KEY UNIQUE CHECK ( year_unique_check_pk > '2020' )
);

DROP TABLE IF EXISTS year_unique_check_pk_ref;
CREATE TABLE year_unique_check_pk_ref (
    year_unique_check_pk_ref year PRIMARY KEY UNIQUE CHECK ( year_unique_check_pk_ref > '2020' ) REFERENCES year_ref(year_ref)
);

DROP TABLE IF EXISTS year_check_pk;
CREATE TABLE year_check_pk (
    year_check_pk year PRIMARY KEY CHECK ( year_check_pk > '2020' )
);

DROP TABLE IF EXISTS year_def_const_unique_check_pk;
CREATE TABLE year_def_const_unique_check_pk (
    year_def_const_unique_check_pk year PRIMARY KEY UNIQUE CHECK ( year_def_const_unique_check_pk > '2020' ) DEFAULT ('2020')
);

DROP TABLE IF EXISTS year_def_const_unique_check_pk_ref;
CREATE TABLE year_def_const_unique_check_pk_ref (
    year_def_const_unique_check_pk_ref year PRIMARY KEY UNIQUE CHECK ( year_def_const_unique_check_pk_ref > '2020' ) DEFAULT ('2020') REFERENCES year_ref(year_ref)
);

DROP TABLE IF EXISTS year_def_func_unique_check_pk;
CREATE TABLE year_def_func_unique_check_pk (
    year_def_func_unique_check_pk year PRIMARY KEY UNIQUE CHECK ( year_def_func_unique_check_pk > '2020' ) DEFAULT (now())
);

DROP TABLE IF EXISTS year_def_func_unique_check_pk_ref;
CREATE TABLE year_def_func_unique_check_pk_ref (
    year_def_func_unique_check_pk_ref year PRIMARY KEY UNIQUE CHECK ( year_def_func_unique_check_pk_ref > '2020' ) DEFAULT (now()) REFERENCES year_ref(year_ref)
);
