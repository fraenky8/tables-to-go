DROP TABLE IF EXISTS character_varying_ref CASCADE;
CREATE TABLE character_varying_ref (
    character_varying_ref character varying UNIQUE
);

DROP TABLE IF EXISTS character_varying;
CREATE TABLE character_varying (
    character_varying character varying,
    character_varying_cap character varying(255),
    character_varying_nn character varying NOT NULL,
    character_varying_nn_unique character varying NOT NULL UNIQUE,
    character_varying_nn_check_cmp character varying NOT NULL CHECK ( character_varying = '42' ),
    character_varying_nn_check_fn character varying NOT NULL CHECK ( length(character_varying) > 0 ),
    character_varying_nn_ref character varying NOT NULL REFERENCES character_varying_ref(character_varying_ref),
    character_varying_nn_def_const character varying NOT NULL DEFAULT '42',
    character_varying_nn_def_func character varying NOT NULL DEFAULT pi(),
    character_varying_nn_unique_check character varying NOT NULL UNIQUE CHECK ( length(character_varying) > 0 ),

    character_varying_unique character varying UNIQUE,
    character_varying_unique_check character varying UNIQUE CHECK ( length(character_varying) > 0 ),
    character_varying_unique_ref character varying UNIQUE REFERENCES character_varying_ref(character_varying_ref),
    character_varying_unique_def_const character varying UNIQUE DEFAULT '42',
    character_varying_unique_def_func character varying UNIQUE DEFAULT pi(),

    character_varying_check character varying CHECK ( length(character_varying) > 0 ),
    character_varying_check_ref character varying CHECK ( length(character_varying) > 0 ) REFERENCES character_varying_ref(character_varying_ref),
    character_varying_check_def_const character varying CHECK ( length(character_varying) > 0 ) DEFAULT '42',
    character_varying_check_def_func character varying CHECK ( length(character_varying) > 0 ) DEFAULT pi(),

    character_varying_ref character varying REFERENCES character_varying_ref(character_varying_ref),
    character_varying_ref_def_const character varying REFERENCES character_varying_ref(character_varying_ref) DEFAULT '42',
    character_varying_ref_def_func character varying REFERENCES character_varying_ref(character_varying_ref) DEFAULT pi(),
    character_varying_ref_unique_check character varying UNIQUE CHECK ( length(character_varying) > 0 ) REFERENCES character_varying_ref(character_varying_ref),

    character_varying_def_const character varying DEFAULT '42',
    character_varying_def_const_unique_check character varying UNIQUE CHECK ( length(character_varying) > 0 ) DEFAULT '42',

    character_varying_def_func character varying DEFAULT pi(),
    character_varying_def_func_unique_check character varying UNIQUE CHECK ( length(character_varying) > 0 ) DEFAULT pi()
);

DROP TABLE IF EXISTS character_varying_pk;
CREATE TABLE character_varying_pk (
    character_varying_pk character varying PRIMARY KEY
);

DROP TABLE IF EXISTS character_varying_pk_ref;
CREATE TABLE character_varying_pk_ref (
    character_varying_pk_ref character varying PRIMARY KEY REFERENCES character_varying_ref(character_varying_ref)
);

DROP TABLE IF EXISTS character_varying_pk_def_const;
CREATE TABLE character_varying_pk_def_const (
    character_varying_pk_def_const character varying PRIMARY KEY DEFAULT '42'
);

DROP TABLE IF EXISTS character_varying_pk_def_func;
CREATE TABLE character_varying_pk_def_func (
    character_varying_pk_def_func character varying PRIMARY KEY DEFAULT pi()
);

DROP TABLE IF EXISTS character_varying_nn_pk;
CREATE TABLE character_varying_nn_pk (
    character_varying_nn_pk character varying NOT NULL PRIMARY KEY
);

DROP TABLE IF EXISTS character_varying_nn_unique_check_pk;
CREATE TABLE character_varying_nn_unique_check_pk (
    character_varying_nn_unique_check_pk character varying PRIMARY KEY NOT NULL UNIQUE CHECK ( length(character_varying_nn_unique_check_pk) > 0)
);

DROP TABLE IF EXISTS character_varying_nn_unique_check_pk_ref;
CREATE TABLE character_varying_nn_unique_check_pk_ref (
    character_varying_nn_unique_check_pk_ref character varying PRIMARY KEY NOT NULL UNIQUE CHECK ( length(character_varying_nn_unique_check_pk_ref) > 0) REFERENCES character_varying_ref(character_varying_ref)
);

DROP TABLE IF EXISTS character_varying_unique_pk;
CREATE TABLE character_varying_unique_pk (
    character_varying_unique_pk character varying PRIMARY KEY UNIQUE
);

DROP TABLE IF EXISTS character_varying_unique_check_pk;
CREATE TABLE character_varying_unique_check_pk (
    character_varying_unique_check_pk character varying PRIMARY KEY UNIQUE CHECK ( length(character_varying_unique_check_pk) > 0 )
);

DROP TABLE IF EXISTS character_varying_unique_check_pk_ref;
CREATE TABLE character_varying_unique_check_pk_ref (
    character_varying_unique_check_pk_ref character varying PRIMARY KEY UNIQUE CHECK ( length(character_varying_unique_check_pk_ref) > 0) REFERENCES character_varying_ref(character_varying_ref)
);

DROP TABLE IF EXISTS character_varying_check_pk;
CREATE TABLE character_varying_check_pk (
    character_varying_check_pk character varying PRIMARY KEY CHECK ( length(character_varying_check_pk) > 0 )
);

DROP TABLE IF EXISTS character_varying_def_const_unique_check_pk;
CREATE TABLE character_varying_def_const_unique_check_pk (
    character_varying_def_const_unique_check_pk character varying PRIMARY KEY UNIQUE CHECK ( length(character_varying_def_const_unique_check_pk) > 0 ) DEFAULT '42'
);

DROP TABLE IF EXISTS character_varying_def_const_unique_check_pk_ref;
CREATE TABLE character_varying_def_const_unique_check_pk_ref (
    character_varying_def_const_unique_check_pk_ref character varying PRIMARY KEY UNIQUE CHECK ( length(character_varying_def_const_unique_check_pk_ref) > 0 ) DEFAULT '42' REFERENCES character_varying_ref(character_varying_ref)
);

DROP TABLE IF EXISTS character_varying_def_func_unique_check_pk;
CREATE TABLE character_varying_def_func_unique_check_pk (
    character_varying_def_func_unique_check_pk character varying PRIMARY KEY UNIQUE CHECK ( length(character_varying_def_func_unique_check_pk) > 0 ) DEFAULT pi()
);

DROP TABLE IF EXISTS character_varying_def_func_unique_check_pk_ref;
CREATE TABLE character_varying_def_func_unique_check_pk_ref (
    character_varying_def_func_unique_check_pk_ref character varying PRIMARY KEY UNIQUE CHECK ( length(character_varying_def_func_unique_check_pk_ref) > 0 ) DEFAULT pi() REFERENCES character_varying_ref(character_varying_ref)
);
