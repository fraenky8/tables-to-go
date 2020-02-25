DROP TABLE IF EXISTS char_ref CASCADE;
CREATE TABLE char_ref (
    char_ref char UNIQUE
);

DROP TABLE IF EXISTS char;
CREATE TABLE char (
    char char,
    char_cap char(255),
    char_nn char NOT NULL,
    char_nn_unique char NOT NULL UNIQUE,
    char_nn_check_cmp char NOT NULL CHECK ( char = '42' ),
    char_nn_check_fn char NOT NULL CHECK ( length(char) > 0 ),
    char_nn_ref char NOT NULL REFERENCES char_ref(char_ref),
    char_nn_def_const char NOT NULL DEFAULT '42',
    char_nn_def_func char NOT NULL DEFAULT pi(),
    char_nn_unique_check char NOT NULL UNIQUE CHECK ( length(char) > 0 ),

    char_unique char UNIQUE,
    char_unique_check char UNIQUE CHECK ( length(char) > 0 ),
    char_unique_ref char UNIQUE REFERENCES char_ref(char_ref),
    char_unique_def_const char UNIQUE DEFAULT '42',
    char_unique_def_func char UNIQUE DEFAULT pi(),

    char_check char CHECK ( length(char) > 0 ),
    char_check_ref char CHECK ( length(char) > 0 ) REFERENCES char_ref(char_ref),
    char_check_def_const char CHECK ( length(char) > 0 ) DEFAULT '42',
    char_check_def_func char CHECK ( length(char) > 0 ) DEFAULT pi(),

    char_ref char REFERENCES char_ref(char_ref),
    char_ref_def_const char REFERENCES char_ref(char_ref) DEFAULT '42',
    char_ref_def_func char REFERENCES char_ref(char_ref) DEFAULT pi(),
    char_ref_unique_check char UNIQUE CHECK ( length(char) > 0 ) REFERENCES char_ref(char_ref),

    char_def_const char DEFAULT '42',
    char_def_const_unique_check char UNIQUE CHECK ( length(char) > 0 ) DEFAULT '42',

    char_def_func char DEFAULT pi(),
    char_def_func_unique_check char UNIQUE CHECK ( length(char) > 0 ) DEFAULT pi()
);

DROP TABLE IF EXISTS char_pk;
CREATE TABLE char_pk (
    char_pk char PRIMARY KEY
);

DROP TABLE IF EXISTS char_pk_ref;
CREATE TABLE char_pk_ref (
    char_pk_ref char PRIMARY KEY REFERENCES char_ref(char_ref)
);

DROP TABLE IF EXISTS char_pk_def_const;
CREATE TABLE char_pk_def_const (
    char_pk_def_const char PRIMARY KEY DEFAULT '42'
);

DROP TABLE IF EXISTS char_pk_def_func;
CREATE TABLE char_pk_def_func (
    char_pk_def_func char PRIMARY KEY DEFAULT pi()
);

DROP TABLE IF EXISTS char_nn_pk;
CREATE TABLE char_nn_pk (
    char_nn_pk char NOT NULL PRIMARY KEY
);

DROP TABLE IF EXISTS char_nn_unique_check_pk;
CREATE TABLE char_nn_unique_check_pk (
    char_nn_unique_check_pk char PRIMARY KEY NOT NULL UNIQUE CHECK ( length(char_nn_unique_check_pk) > 0)
);

DROP TABLE IF EXISTS char_nn_unique_check_pk_ref;
CREATE TABLE char_nn_unique_check_pk_ref (
    char_nn_unique_check_pk_ref char PRIMARY KEY NOT NULL UNIQUE CHECK ( length(char_nn_unique_check_pk_ref) > 0) REFERENCES char_ref(char_ref)
);

DROP TABLE IF EXISTS char_unique_pk;
CREATE TABLE char_unique_pk (
    char_unique_pk char PRIMARY KEY UNIQUE
);

DROP TABLE IF EXISTS char_unique_check_pk;
CREATE TABLE char_unique_check_pk (
    char_unique_check_pk char PRIMARY KEY UNIQUE CHECK ( length(char_unique_check_pk) > 0 )
);

DROP TABLE IF EXISTS char_unique_check_pk_ref;
CREATE TABLE char_unique_check_pk_ref (
    char_unique_check_pk_ref char PRIMARY KEY UNIQUE CHECK ( length(char_unique_check_pk_ref) > 0) REFERENCES char_ref(char_ref)
);

DROP TABLE IF EXISTS char_check_pk;
CREATE TABLE char_check_pk (
    char_check_pk char PRIMARY KEY CHECK ( length(char_check_pk) > 0 )
);

DROP TABLE IF EXISTS char_def_const_unique_check_pk;
CREATE TABLE char_def_const_unique_check_pk (
    char_def_const_unique_check_pk char PRIMARY KEY UNIQUE CHECK ( length(char_def_const_unique_check_pk) > 0 ) DEFAULT '42'
);

DROP TABLE IF EXISTS char_def_const_unique_check_pk_ref;
CREATE TABLE char_def_const_unique_check_pk_ref (
    char_def_const_unique_check_pk_ref char PRIMARY KEY UNIQUE CHECK ( length(char_def_const_unique_check_pk_ref) > 0 ) DEFAULT '42' REFERENCES char_ref(char_ref)
);

DROP TABLE IF EXISTS char_def_func_unique_check_pk;
CREATE TABLE char_def_func_unique_check_pk (
    char_def_func_unique_check_pk char PRIMARY KEY UNIQUE CHECK ( length(char_def_func_unique_check_pk) > 0 ) DEFAULT pi()
);

DROP TABLE IF EXISTS char_def_func_unique_check_pk_ref;
CREATE TABLE char_def_func_unique_check_pk_ref (
    char_def_func_unique_check_pk_ref char PRIMARY KEY UNIQUE CHECK ( length(char_def_func_unique_check_pk_ref) > 0 ) DEFAULT pi() REFERENCES char_ref(char_ref)
);
