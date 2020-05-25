DROP TABLE IF EXISTS integer_ref CASCADE;
CREATE TABLE integer_ref (
    integer_ref integer UNIQUE
);

DROP TABLE IF EXISTS integer;
CREATE TABLE integer (
    integer integer,
    integer_nn integer NOT NULL,
    integer_nn_unique integer NOT NULL UNIQUE,
    integer_nn_check integer NOT NULL CHECK ( integer > 0 ),
    integer_nn_ref integer NOT NULL REFERENCES integer_ref(integer_ref),
    integer_nn_def_const integer NOT NULL DEFAULT 42,
    integer_nn_def_func integer NOT NULL DEFAULT pi(),
    integer_nn_unique_check integer NOT NULL UNIQUE CHECK ( integer > 0 ),

    integer_unique integer UNIQUE,
    integer_unique_check integer UNIQUE CHECK ( integer > 0 ),
    integer_unique_ref integer UNIQUE REFERENCES integer_ref(integer_ref),
    integer_unique_def_const integer UNIQUE DEFAULT 42,
    integer_unique_def_func integer UNIQUE DEFAULT pi(),

    integer_check integer CHECK ( integer > 0 ),
    integer_check_ref integer CHECK ( integer > 0 ) REFERENCES integer_ref(integer_ref),
    integer_check_def_const integer CHECK ( integer > 0 ) DEFAULT 42,
    integer_check_def_func integer CHECK ( integer > 0 ) DEFAULT pi(),

    integer_ref integer REFERENCES integer_ref(integer_ref),
    integer_ref_def_const integer REFERENCES integer_ref(integer_ref) DEFAULT 42,
    integer_ref_def_func integer REFERENCES integer_ref(integer_ref) DEFAULT pi(),
    integer_ref_unique_check integer UNIQUE CHECK ( integer > 0 ) REFERENCES integer_ref(integer_ref),

    integer_def_const integer DEFAULT 42,
    integer_def_const_unique_check integer UNIQUE CHECK ( integer > 0 )DEFAULT 42,

    integer_def_func integer DEFAULT pi(),
    integer_def_func_unique_check integer UNIQUE CHECK ( integer > 0 ) DEFAULT pi()
);

DROP TABLE IF EXISTS integer_pk;
CREATE TABLE integer_pk (
    integer_pk integer PRIMARY KEY
);

DROP TABLE IF EXISTS integer_pk_ref;
CREATE TABLE integer_pk_ref (
    integer_pk_ref integer PRIMARY KEY REFERENCES integer_ref(integer_ref)
);

DROP TABLE IF EXISTS integer_pk_def_const;
CREATE TABLE integer_pk_def_const (
    integer_pk_def_const integer PRIMARY KEY DEFAULT 42
);

DROP TABLE IF EXISTS integer_pk_def_func;
CREATE TABLE integer_pk_def_func (
    integer_pk_def_func integer PRIMARY KEY DEFAULT pi()
);

DROP TABLE IF EXISTS integer_nn_pk;
CREATE TABLE integer_nn_pk (
    integer_nn_pk integer NOT NULL PRIMARY KEY
);

DROP TABLE IF EXISTS integer_nn_unique_check_pk;
CREATE TABLE integer_nn_unique_check_pk (
    integer_nn_unique_check_pk integer PRIMARY KEY NOT NULL UNIQUE CHECK ( integer_nn_unique_check_pk > 0)
);

DROP TABLE IF EXISTS integer_nn_unique_check_pk_ref;
CREATE TABLE integer_nn_unique_check_pk_ref (
    integer_nn_unique_check_pk_ref integer PRIMARY KEY NOT NULL UNIQUE CHECK ( integer_nn_unique_check_pk_ref > 0) REFERENCES integer_ref(integer_ref)
);

DROP TABLE IF EXISTS integer_unique_pk;
CREATE TABLE integer_unique_pk (
    integer_unique_pk integer PRIMARY KEY UNIQUE
);

DROP TABLE IF EXISTS integer_unique_check_pk;
CREATE TABLE integer_unique_check_pk (
    integer_unique_check_pk integer PRIMARY KEY UNIQUE CHECK ( integer_unique_check_pk > 0 )
);

DROP TABLE IF EXISTS integer_unique_check_pk_ref;
CREATE TABLE integer_unique_check_pk_ref (
    integer_unique_check_pk_ref integer PRIMARY KEY UNIQUE CHECK ( integer_unique_check_pk_ref > 0) REFERENCES integer_ref(integer_ref)
);

DROP TABLE IF EXISTS integer_check_pk;
CREATE TABLE integer_check_pk (
    integer_check_pk integer PRIMARY KEY CHECK ( integer_check_pk > 0 )
);

DROP TABLE IF EXISTS integer_def_const_unique_check_pk;
CREATE TABLE integer_def_const_unique_check_pk (
    integer_def_const_unique_check_pk integer PRIMARY KEY UNIQUE CHECK ( integer_def_const_unique_check_pk > 0 ) DEFAULT 42
);

DROP TABLE IF EXISTS integer_def_const_unique_check_pk_ref;
CREATE TABLE integer_def_const_unique_check_pk_ref (
    integer_def_const_unique_check_pk_ref integer PRIMARY KEY UNIQUE CHECK ( integer_def_const_unique_check_pk_ref > 0 ) DEFAULT 42 REFERENCES integer_ref(integer_ref)
);

DROP TABLE IF EXISTS integer_def_func_unique_check_pk;
CREATE TABLE integer_def_func_unique_check_pk (
    integer_def_func_unique_check_pk integer PRIMARY KEY UNIQUE CHECK ( integer_def_func_unique_check_pk > 0 ) DEFAULT pi()
);

DROP TABLE IF EXISTS integer_def_func_unique_check_pk_ref;
CREATE TABLE integer_def_func_unique_check_pk_ref (
    integer_def_func_unique_check_pk_ref integer PRIMARY KEY UNIQUE CHECK ( integer_def_func_unique_check_pk_ref > 0 ) DEFAULT pi() REFERENCES integer_ref(integer_ref)
);
