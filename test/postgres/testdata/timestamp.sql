DROP TABLE IF EXISTS timestamp_ref CASCADE;
CREATE TABLE timestamp_ref (
    timestamp_ref timestamp UNIQUE
);

DROP TABLE IF EXISTS timestamp;
CREATE TABLE timestamp (
    timestamp timestamp,
    timestamp_nn timestamp NOT NULL,
    timestamp_nn_unique timestamp NOT NULL UNIQUE,
    timestamp_nn_check timestamp NOT NULL CHECK ( timestamp > '2020-03-01 12:34:56' ),
    timestamp_nn_ref timestamp NOT NULL REFERENCES timestamp_ref(timestamp_ref),
    timestamp_nn_def_const timestamp NOT NULL DEFAULT '2020-03-01 12:34:56',
    timestamp_nn_def_func timestamp NOT NULL DEFAULT now(),
    timestamp_nn_unique_check timestamp NOT NULL UNIQUE CHECK ( timestamp > '2020-03-01 12:34:56' ),

    timestamp_unique timestamp UNIQUE,
    timestamp_unique_check timestamp UNIQUE CHECK ( timestamp > '2020-03-01 12:34:56' ),
    timestamp_unique_ref timestamp UNIQUE REFERENCES timestamp_ref(timestamp_ref),
    timestamp_unique_def_const timestamp UNIQUE DEFAULT '2020-03-01 12:34:56',
    timestamp_unique_def_func timestamp UNIQUE DEFAULT now(),

    timestamp_check timestamp CHECK ( timestamp > '2020-03-01 12:34:56' ),
    timestamp_check_ref timestamp CHECK ( timestamp > '2020-03-01 12:34:56' ) REFERENCES timestamp_ref(timestamp_ref),
    timestamp_check_def_const timestamp CHECK ( timestamp > '2020-03-01 12:34:56' ) DEFAULT '2020-03-01 12:34:56',
    timestamp_check_def_func timestamp CHECK ( timestamp > '2020-03-01 12:34:56' ) DEFAULT now(),

    timestamp_ref timestamp REFERENCES timestamp_ref(timestamp_ref),
    timestamp_ref_def_const timestamp REFERENCES timestamp_ref(timestamp_ref) DEFAULT '2020-03-01 12:34:56',
    timestamp_ref_def_func timestamp REFERENCES timestamp_ref(timestamp_ref) DEFAULT now(),
    timestamp_ref_unique_check timestamp UNIQUE CHECK ( timestamp > '2020-03-01 12:34:56' ) REFERENCES timestamp_ref(timestamp_ref),

    timestamp_def_const timestamp DEFAULT '2020-03-01 12:34:56',
    timestamp_def_const_unique_check timestamp UNIQUE CHECK ( timestamp > '2020-03-01 12:34:56' )DEFAULT '2020-03-01 12:34:56',

    timestamp_def_func timestamp DEFAULT now(),
    timestamp_def_func_unique_check timestamp UNIQUE CHECK ( timestamp > '2020-03-01 12:34:56' ) DEFAULT now()
);

DROP TABLE IF EXISTS timestamp_pk;
CREATE TABLE timestamp_pk (
    timestamp_pk timestamp PRIMARY KEY
);

DROP TABLE IF EXISTS timestamp_pk_ref;
CREATE TABLE timestamp_pk_ref (
    timestamp_pk_ref timestamp PRIMARY KEY REFERENCES timestamp_ref(timestamp_ref)
);

DROP TABLE IF EXISTS timestamp_pk_def_const;
CREATE TABLE timestamp_pk_def_const (
    timestamp_pk_def_const timestamp PRIMARY KEY DEFAULT '2020-03-01 12:34:56'
);

DROP TABLE IF EXISTS timestamp_pk_def_func;
CREATE TABLE timestamp_pk_def_func (
    timestamp_pk_def_func timestamp PRIMARY KEY DEFAULT now()
);

DROP TABLE IF EXISTS timestamp_nn_pk;
CREATE TABLE timestamp_nn_pk (
    timestamp_nn_pk timestamp NOT NULL PRIMARY KEY
);

DROP TABLE IF EXISTS timestamp_nn_unique_check_pk;
CREATE TABLE timestamp_nn_unique_check_pk (
    timestamp_nn_unique_check_pk timestamp PRIMARY KEY NOT NULL UNIQUE CHECK ( timestamp_nn_unique_check_pk > '2020-03-01 12:34:56' )
);

DROP TABLE IF EXISTS timestamp_nn_unique_check_pk_ref;
CREATE TABLE timestamp_nn_unique_check_pk_ref (
    timestamp_nn_unique_check_pk_ref timestamp PRIMARY KEY NOT NULL UNIQUE CHECK ( timestamp_nn_unique_check_pk_ref > '2020-03-01 12:34:56' ) REFERENCES timestamp_ref(timestamp_ref)
);

DROP TABLE IF EXISTS timestamp_unique_pk;
CREATE TABLE timestamp_unique_pk (
    timestamp_unique_pk timestamp PRIMARY KEY UNIQUE
);

DROP TABLE IF EXISTS timestamp_unique_check_pk;
CREATE TABLE timestamp_unique_check_pk (
    timestamp_unique_check_pk timestamp PRIMARY KEY UNIQUE CHECK ( timestamp_unique_check_pk > '2020-03-01 12:34:56' )
);

DROP TABLE IF EXISTS timestamp_unique_check_pk_ref;
CREATE TABLE timestamp_unique_check_pk_ref (
    timestamp_unique_check_pk_ref timestamp PRIMARY KEY UNIQUE CHECK ( timestamp_unique_check_pk_ref > '2020-03-01 12:34:56' ) REFERENCES timestamp_ref(timestamp_ref)
);

DROP TABLE IF EXISTS timestamp_check_pk;
CREATE TABLE timestamp_check_pk (
    timestamp_check_pk timestamp PRIMARY KEY CHECK ( timestamp_check_pk > '2020-03-01 12:34:56' )
);

DROP TABLE IF EXISTS timestamp_def_const_unique_check_pk;
CREATE TABLE timestamp_def_const_unique_check_pk (
    timestamp_def_const_unique_check_pk timestamp PRIMARY KEY UNIQUE CHECK ( timestamp_def_const_unique_check_pk > '2020-03-01 12:34:56' ) DEFAULT '2020-03-01 12:34:56'
);

DROP TABLE IF EXISTS timestamp_def_const_unique_check_pk_ref;
CREATE TABLE timestamp_def_const_unique_check_pk_ref (
    timestamp_def_const_unique_check_pk_ref timestamp PRIMARY KEY UNIQUE CHECK ( timestamp_def_const_unique_check_pk_ref > '2020-03-01 12:34:56' ) DEFAULT '2020-03-01 12:34:56' REFERENCES timestamp_ref(timestamp_ref)
);

DROP TABLE IF EXISTS timestamp_def_func_unique_check_pk;
CREATE TABLE timestamp_def_func_unique_check_pk (
    timestamp_def_func_unique_check_pk timestamp PRIMARY KEY UNIQUE CHECK ( timestamp_def_func_unique_check_pk > '2020-03-01 12:34:56' ) DEFAULT now()
);

DROP TABLE IF EXISTS timestamp_def_func_unique_check_pk_ref;
CREATE TABLE timestamp_def_func_unique_check_pk_ref (
    timestamp_def_func_unique_check_pk_ref timestamp PRIMARY KEY UNIQUE CHECK ( timestamp_def_func_unique_check_pk_ref > '2020-03-01 12:34:56' ) DEFAULT now() REFERENCES timestamp_ref(timestamp_ref)
);
