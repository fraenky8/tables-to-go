DROP TABLE IF EXISTS date_ref CASCADE;
CREATE TABLE date_ref (
    date_ref date UNIQUE
);

DROP TABLE IF EXISTS date;
CREATE TABLE date (
    date date,
    date_nn date NOT NULL,
    date_nn_unique date NOT NULL UNIQUE,
    date_nn_check date NOT NULL CHECK ( date > '2020-03-01' ),
    date_nn_ref date NOT NULL REFERENCES date_ref(date_ref),
    date_nn_def_const date NOT NULL DEFAULT '2020-03-01',
    date_nn_def_func date NOT NULL DEFAULT now(),
    date_nn_unique_check date NOT NULL UNIQUE CHECK ( date > '2020-03-01' ),

    date_unique date UNIQUE,
    date_unique_check date UNIQUE CHECK ( date > '2020-03-01' ),
    date_unique_ref date UNIQUE REFERENCES date_ref(date_ref),
    date_unique_def_const date UNIQUE DEFAULT '2020-03-01',
    date_unique_def_func date UNIQUE DEFAULT now(),

    date_check date CHECK ( date > '2020-03-01' ),
    date_check_ref date CHECK ( date > '2020-03-01' ) REFERENCES date_ref(date_ref),
    date_check_def_const date CHECK ( date > '2020-03-01' ) DEFAULT '2020-03-01',
    date_check_def_func date CHECK ( date > '2020-03-01' ) DEFAULT now(),

    date_ref date REFERENCES date_ref(date_ref),
    date_ref_def_const date REFERENCES date_ref(date_ref) DEFAULT '2020-03-01',
    date_ref_def_func date REFERENCES date_ref(date_ref) DEFAULT now(),
    date_ref_unique_check date UNIQUE CHECK ( date > '2020-03-01' ) REFERENCES date_ref(date_ref),

    date_def_const date DEFAULT '2020-03-01',
    date_def_const_unique_check date UNIQUE CHECK ( date > '2020-03-01' )DEFAULT '2020-03-01',

    date_def_func date DEFAULT now(),
    date_def_func_unique_check date UNIQUE CHECK ( date > '2020-03-01' ) DEFAULT now()
);

DROP TABLE IF EXISTS date_pk;
CREATE TABLE date_pk (
    date_pk date PRIMARY KEY
);

DROP TABLE IF EXISTS date_pk_ref;
CREATE TABLE date_pk_ref (
    date_pk_ref date PRIMARY KEY REFERENCES date_ref(date_ref)
);

DROP TABLE IF EXISTS date_pk_def_const;
CREATE TABLE date_pk_def_const (
    date_pk_def_const date PRIMARY KEY DEFAULT '2020-03-01'
);

DROP TABLE IF EXISTS date_pk_def_func;
CREATE TABLE date_pk_def_func (
    date_pk_def_func date PRIMARY KEY DEFAULT now()
);

DROP TABLE IF EXISTS date_nn_pk;
CREATE TABLE date_nn_pk (
    date_nn_pk date NOT NULL PRIMARY KEY
);

DROP TABLE IF EXISTS date_nn_unique_check_pk;
CREATE TABLE date_nn_unique_check_pk (
    date_nn_unique_check_pk date PRIMARY KEY NOT NULL UNIQUE CHECK ( date_nn_unique_check_pk > '2020-03-01' )
);

DROP TABLE IF EXISTS date_nn_unique_check_pk_ref;
CREATE TABLE date_nn_unique_check_pk_ref (
    date_nn_unique_check_pk_ref date PRIMARY KEY NOT NULL UNIQUE CHECK ( date_nn_unique_check_pk_ref > '2020-03-01' ) REFERENCES date_ref(date_ref)
);

DROP TABLE IF EXISTS date_unique_pk;
CREATE TABLE date_unique_pk (
    date_unique_pk date PRIMARY KEY UNIQUE
);

DROP TABLE IF EXISTS date_unique_check_pk;
CREATE TABLE date_unique_check_pk (
    date_unique_check_pk date PRIMARY KEY UNIQUE CHECK ( date_unique_check_pk > '2020-03-01' )
);

DROP TABLE IF EXISTS date_unique_check_pk_ref;
CREATE TABLE date_unique_check_pk_ref (
    date_unique_check_pk_ref date PRIMARY KEY UNIQUE CHECK ( date_unique_check_pk_ref > '2020-03-01' ) REFERENCES date_ref(date_ref)
);

DROP TABLE IF EXISTS date_check_pk;
CREATE TABLE date_check_pk (
    date_check_pk date PRIMARY KEY CHECK ( date_check_pk > '2020-03-01' )
);

DROP TABLE IF EXISTS date_def_const_unique_check_pk;
CREATE TABLE date_def_const_unique_check_pk (
    date_def_const_unique_check_pk date PRIMARY KEY UNIQUE CHECK ( date_def_const_unique_check_pk > '2020-03-01' ) DEFAULT '2020-03-01'
);

DROP TABLE IF EXISTS date_def_const_unique_check_pk_ref;
CREATE TABLE date_def_const_unique_check_pk_ref (
    date_def_const_unique_check_pk_ref date PRIMARY KEY UNIQUE CHECK ( date_def_const_unique_check_pk_ref > '2020-03-01' ) DEFAULT '2020-03-01' REFERENCES date_ref(date_ref)
);

DROP TABLE IF EXISTS date_def_func_unique_check_pk;
CREATE TABLE date_def_func_unique_check_pk (
    date_def_func_unique_check_pk date PRIMARY KEY UNIQUE CHECK ( date_def_func_unique_check_pk > '2020-03-01' ) DEFAULT now()
);

DROP TABLE IF EXISTS date_def_func_unique_check_pk_ref;
CREATE TABLE date_def_func_unique_check_pk_ref (
    date_def_func_unique_check_pk_ref date PRIMARY KEY UNIQUE CHECK ( date_def_func_unique_check_pk_ref > '2020-03-01' ) DEFAULT now() REFERENCES date_ref(date_ref)
);
