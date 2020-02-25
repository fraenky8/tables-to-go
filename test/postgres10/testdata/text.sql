DROP TABLE IF EXISTS text_ref CASCADE;
CREATE TABLE text_ref (
    text_ref text UNIQUE
);

DROP TABLE IF EXISTS text;
CREATE TABLE text (
    text text,
    text_nn text NOT NULL,
    text_nn_unique text NOT NULL UNIQUE,
    text_nn_check_cmp text NOT NULL CHECK ( text = '42' ),
    text_nn_check_fn text NOT NULL CHECK ( length(text) > 0 ),
    text_nn_ref text NOT NULL REFERENCES text_ref(text_ref),
    text_nn_def_const text NOT NULL DEFAULT '42',
    text_nn_def_func text NOT NULL DEFAULT pi(),
    text_nn_unique_check text NOT NULL UNIQUE CHECK ( length(text) > 0 ),

    text_unique text UNIQUE,
    text_unique_check text UNIQUE CHECK ( length(text) > 0 ),
    text_unique_ref text UNIQUE REFERENCES text_ref(text_ref),
    text_unique_def_const text UNIQUE DEFAULT '42',
    text_unique_def_func text UNIQUE DEFAULT pi(),

    text_check text CHECK ( length(text) > 0 ),
    text_check_ref text CHECK ( length(text) > 0 ) REFERENCES text_ref(text_ref),
    text_check_def_const text CHECK ( length(text) > 0 ) DEFAULT '42',
    text_check_def_func text CHECK ( length(text) > 0 ) DEFAULT pi(),

    text_ref text REFERENCES text_ref(text_ref),
    text_ref_def_const text REFERENCES text_ref(text_ref) DEFAULT '42',
    text_ref_def_func text REFERENCES text_ref(text_ref) DEFAULT pi(),
    text_ref_unique_check text UNIQUE CHECK ( length(text) > 0 ) REFERENCES text_ref(text_ref),

    text_def_const text DEFAULT '42',
    text_def_const_unique_check text UNIQUE CHECK ( length(text) > 0 ) DEFAULT '42',

    text_def_func text DEFAULT pi(),
    text_def_func_unique_check text UNIQUE CHECK ( length(text) > 0 ) DEFAULT pi()
);

DROP TABLE IF EXISTS text_pk;
CREATE TABLE text_pk (
    text_pk text PRIMARY KEY
);

DROP TABLE IF EXISTS text_pk_ref;
CREATE TABLE text_pk_ref (
    text_pk_ref text PRIMARY KEY REFERENCES text_ref(text_ref)
);

DROP TABLE IF EXISTS text_pk_def_const;
CREATE TABLE text_pk_def_const (
    text_pk_def_const text PRIMARY KEY DEFAULT '42'
);

DROP TABLE IF EXISTS text_pk_def_func;
CREATE TABLE text_pk_def_func (
    text_pk_def_func text PRIMARY KEY DEFAULT pi()
);

DROP TABLE IF EXISTS text_nn_pk;
CREATE TABLE text_nn_pk (
    text_nn_pk text NOT NULL PRIMARY KEY
);

DROP TABLE IF EXISTS text_nn_unique_check_pk;
CREATE TABLE text_nn_unique_check_pk (
    text_nn_unique_check_pk text PRIMARY KEY NOT NULL UNIQUE CHECK ( length(text_nn_unique_check_pk) > 0)
);

DROP TABLE IF EXISTS text_nn_unique_check_pk_ref;
CREATE TABLE text_nn_unique_check_pk_ref (
    text_nn_unique_check_pk_ref text PRIMARY KEY NOT NULL UNIQUE CHECK ( length(text_nn_unique_check_pk_ref) > 0) REFERENCES text_ref(text_ref)
);

DROP TABLE IF EXISTS text_unique_pk;
CREATE TABLE text_unique_pk (
    text_unique_pk text PRIMARY KEY UNIQUE
);

DROP TABLE IF EXISTS text_unique_check_pk;
CREATE TABLE text_unique_check_pk (
    text_unique_check_pk text PRIMARY KEY UNIQUE CHECK ( length(text_unique_check_pk) > 0 )
);

DROP TABLE IF EXISTS text_unique_check_pk_ref;
CREATE TABLE text_unique_check_pk_ref (
    text_unique_check_pk_ref text PRIMARY KEY UNIQUE CHECK ( length(text_unique_check_pk_ref) > 0) REFERENCES text_ref(text_ref)
);

DROP TABLE IF EXISTS text_check_pk;
CREATE TABLE text_check_pk (
    text_check_pk text PRIMARY KEY CHECK ( length(text_check_pk) > 0 )
);

DROP TABLE IF EXISTS text_def_const_unique_check_pk;
CREATE TABLE text_def_const_unique_check_pk (
    text_def_const_unique_check_pk text PRIMARY KEY UNIQUE CHECK ( length(text_def_const_unique_check_pk) > 0 ) DEFAULT '42'
);

DROP TABLE IF EXISTS text_def_const_unique_check_pk_ref;
CREATE TABLE text_def_const_unique_check_pk_ref (
    text_def_const_unique_check_pk_ref text PRIMARY KEY UNIQUE CHECK ( length(text_def_const_unique_check_pk_ref) > 0 ) DEFAULT '42' REFERENCES text_ref(text_ref)
);

DROP TABLE IF EXISTS text_def_func_unique_check_pk;
CREATE TABLE text_def_func_unique_check_pk (
    text_def_func_unique_check_pk text PRIMARY KEY UNIQUE CHECK ( length(text_def_func_unique_check_pk) > 0 ) DEFAULT pi()
);

DROP TABLE IF EXISTS text_def_func_unique_check_pk_ref;
CREATE TABLE text_def_func_unique_check_pk_ref (
    text_def_func_unique_check_pk_ref text PRIMARY KEY UNIQUE CHECK ( length(text_def_func_unique_check_pk_ref) > 0 ) DEFAULT pi() REFERENCES text_ref(text_ref)
);
