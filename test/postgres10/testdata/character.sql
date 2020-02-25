DROP TABLE IF EXISTS character_ref CASCADE;
CREATE TABLE character_ref (
    character_ref character UNIQUE
);

DROP TABLE IF EXISTS character;
CREATE TABLE character (
    character character,
    character_cap character(255),
    character_nn character NOT NULL,
    character_nn_unique character NOT NULL UNIQUE,
    character_nn_check_cmp character NOT NULL CHECK ( character = '42' ),
    character_nn_check_fn character NOT NULL CHECK ( length(character) > 0 ),
    character_nn_ref character NOT NULL REFERENCES character_ref(character_ref),
    character_nn_def_const character NOT NULL DEFAULT '42',
    character_nn_def_func character NOT NULL DEFAULT pi(),
    character_nn_unique_check character NOT NULL UNIQUE CHECK ( length(character) > 0 ),

    character_unique character UNIQUE,
    character_unique_check character UNIQUE CHECK ( length(character) > 0 ),
    character_unique_ref character UNIQUE REFERENCES character_ref(character_ref),
    character_unique_def_const character UNIQUE DEFAULT '42',
    character_unique_def_func character UNIQUE DEFAULT pi(),

    character_check character CHECK ( length(character) > 0 ),
    character_check_ref character CHECK ( length(character) > 0 ) REFERENCES character_ref(character_ref),
    character_check_def_const character CHECK ( length(character) > 0 ) DEFAULT '42',
    character_check_def_func character CHECK ( length(character) > 0 ) DEFAULT pi(),

    character_ref character REFERENCES character_ref(character_ref),
    character_ref_def_const character REFERENCES character_ref(character_ref) DEFAULT '42',
    character_ref_def_func character REFERENCES character_ref(character_ref) DEFAULT pi(),
    character_ref_unique_check character UNIQUE CHECK ( length(character) > 0 ) REFERENCES character_ref(character_ref),

    character_def_const character DEFAULT '42',
    character_def_const_unique_check character UNIQUE CHECK ( length(character) > 0 ) DEFAULT '42',

    character_def_func character DEFAULT pi(),
    character_def_func_unique_check character UNIQUE CHECK ( length(character) > 0 ) DEFAULT pi()
);

DROP TABLE IF EXISTS character_pk;
CREATE TABLE character_pk (
    character_pk character PRIMARY KEY
);

DROP TABLE IF EXISTS character_pk_ref;
CREATE TABLE character_pk_ref (
    character_pk_ref character PRIMARY KEY REFERENCES character_ref(character_ref)
);

DROP TABLE IF EXISTS character_pk_def_const;
CREATE TABLE character_pk_def_const (
    character_pk_def_const character PRIMARY KEY DEFAULT '42'
);

DROP TABLE IF EXISTS character_pk_def_func;
CREATE TABLE character_pk_def_func (
    character_pk_def_func character PRIMARY KEY DEFAULT pi()
);

DROP TABLE IF EXISTS character_nn_pk;
CREATE TABLE character_nn_pk (
    character_nn_pk character NOT NULL PRIMARY KEY
);

DROP TABLE IF EXISTS character_nn_unique_check_pk;
CREATE TABLE character_nn_unique_check_pk (
    character_nn_unique_check_pk character PRIMARY KEY NOT NULL UNIQUE CHECK ( length(character_nn_unique_check_pk) > 0)
);

DROP TABLE IF EXISTS character_nn_unique_check_pk_ref;
CREATE TABLE character_nn_unique_check_pk_ref (
    character_nn_unique_check_pk_ref character PRIMARY KEY NOT NULL UNIQUE CHECK ( length(character_nn_unique_check_pk_ref) > 0) REFERENCES character_ref(character_ref)
);

DROP TABLE IF EXISTS character_unique_pk;
CREATE TABLE character_unique_pk (
    character_unique_pk character PRIMARY KEY UNIQUE
);

DROP TABLE IF EXISTS character_unique_check_pk;
CREATE TABLE character_unique_check_pk (
    character_unique_check_pk character PRIMARY KEY UNIQUE CHECK ( length(character_unique_check_pk) > 0 )
);

DROP TABLE IF EXISTS character_unique_check_pk_ref;
CREATE TABLE character_unique_check_pk_ref (
    character_unique_check_pk_ref character PRIMARY KEY UNIQUE CHECK ( length(character_unique_check_pk_ref) > 0) REFERENCES character_ref(character_ref)
);

DROP TABLE IF EXISTS character_check_pk;
CREATE TABLE character_check_pk (
    character_check_pk character PRIMARY KEY CHECK ( length(character_check_pk) > 0 )
);

DROP TABLE IF EXISTS character_def_const_unique_check_pk;
CREATE TABLE character_def_const_unique_check_pk (
    character_def_const_unique_check_pk character PRIMARY KEY UNIQUE CHECK ( length(character_def_const_unique_check_pk) > 0 ) DEFAULT '42'
);

DROP TABLE IF EXISTS character_def_const_unique_check_pk_ref;
CREATE TABLE character_def_const_unique_check_pk_ref (
    character_def_const_unique_check_pk_ref character PRIMARY KEY UNIQUE CHECK ( length(character_def_const_unique_check_pk_ref) > 0 ) DEFAULT '42' REFERENCES character_ref(character_ref)
);

DROP TABLE IF EXISTS character_def_func_unique_check_pk;
CREATE TABLE character_def_func_unique_check_pk (
    character_def_func_unique_check_pk character PRIMARY KEY UNIQUE CHECK ( length(character_def_func_unique_check_pk) > 0 ) DEFAULT pi()
);

DROP TABLE IF EXISTS character_def_func_unique_check_pk_ref;
CREATE TABLE character_def_func_unique_check_pk_ref (
    character_def_func_unique_check_pk_ref character PRIMARY KEY UNIQUE CHECK ( length(character_def_func_unique_check_pk_ref) > 0 ) DEFAULT pi() REFERENCES character_ref(character_ref)
);
