DROP TABLE IF EXISTS datetime_ref CASCADE;
CREATE TABLE datetime_ref (
    datetime_ref datetime UNIQUE
);

DROP TABLE IF EXISTS datetime_table;
CREATE TABLE datetime_table (
    datetime datetime,
    datetime_nn datetime NOT NULL,
    datetime_nn_unique datetime NOT NULL UNIQUE,
    datetime_nn_check datetime NOT NULL CHECK ( datetime_nn_check > '2020-03-01 12:34:56' ),
    datetime_nn_ref datetime NOT NULL REFERENCES datetime_ref(datetime_ref),
    datetime_nn_def_const datetime NOT NULL DEFAULT ('2020-03-01 12:34:56'),
    datetime_nn_def_func datetime NOT NULL DEFAULT (now()),
    datetime_nn_unique_check datetime NOT NULL UNIQUE CHECK ( datetime_nn_unique_check > '2020-03-01 12:34:56' ),

    datetime_unique datetime UNIQUE,
    datetime_unique_check datetime UNIQUE CHECK ( datetime_unique_check > '2020-03-01 12:34:56' ),
    datetime_unique_ref datetime UNIQUE REFERENCES datetime_ref(datetime_ref),
    datetime_unique_def_const datetime UNIQUE DEFAULT ('2020-03-01 12:34:56'),
    datetime_unique_def_func datetime UNIQUE DEFAULT (now()),

    datetime_check datetime CHECK ( datetime_check > '2020-03-01 12:34:56' ),
    datetime_check_ref datetime CHECK ( datetime_check_ref > '2020-03-01 12:34:56' ) REFERENCES datetime_ref(datetime_ref),
    datetime_check_def_const datetime CHECK ( datetime_check_def_const > '2020-03-01 12:34:56' ) DEFAULT ('2020-03-01 12:34:56'),
    datetime_check_def_func datetime CHECK ( datetime_check_def_func > '2020-03-01 12:34:56' ) DEFAULT (now()),

    datetime_ref datetime REFERENCES datetime_ref(datetime_ref),
    datetime_ref_unique_check datetime UNIQUE CHECK ( datetime_ref_unique_check > '2020-03-01 12:34:56' ) REFERENCES datetime_ref(datetime_ref),

    datetime_def_const datetime DEFAULT ('2020-03-01 12:34:56'),
    datetime_def_const_unique_check datetime UNIQUE CHECK ( datetime_def_const_unique_check > '2020-03-01 12:34:56' ) DEFAULT ('2020-03-01 12:34:56'),

    datetime_def_func datetime DEFAULT (now()),
    datetime_def_func_unique_check datetime UNIQUE CHECK ( datetime_def_func_unique_check > '2020-03-01 12:34:56' ) DEFAULT (now())
);

DROP TABLE IF EXISTS datetime_pk;
CREATE TABLE datetime_pk (
    datetime_pk datetime PRIMARY KEY
);

DROP TABLE IF EXISTS datetime_pk_ref;
CREATE TABLE datetime_pk_ref (
    datetime_pk_ref datetime PRIMARY KEY REFERENCES datetime_ref(datetime_ref)
);

DROP TABLE IF EXISTS datetime_pk_def_const;
CREATE TABLE datetime_pk_def_const (
    datetime_pk_def_const datetime PRIMARY KEY DEFAULT ('2020-03-01 12:34:56')
);

DROP TABLE IF EXISTS datetime_pk_def_func;
CREATE TABLE datetime_pk_def_func (
    datetime_pk_def_func datetime PRIMARY KEY DEFAULT (now())
);

DROP TABLE IF EXISTS datetime_nn_pk;
CREATE TABLE datetime_nn_pk (
    datetime_nn_pk datetime NOT NULL PRIMARY KEY
);

DROP TABLE IF EXISTS datetime_nn_unique_check_pk;
CREATE TABLE datetime_nn_unique_check_pk (
    datetime_nn_unique_check_pk datetime PRIMARY KEY NOT NULL UNIQUE CHECK ( datetime_nn_unique_check_pk > '2020-03-01 12:34:56' )
);

DROP TABLE IF EXISTS datetime_nn_unique_check_pk_ref;
CREATE TABLE datetime_nn_unique_check_pk_ref (
    datetime_nn_unique_check_pk_ref datetime PRIMARY KEY NOT NULL UNIQUE CHECK ( datetime_nn_unique_check_pk_ref > '2020-03-01 12:34:56' ) REFERENCES datetime_ref(datetime_ref)
);

DROP TABLE IF EXISTS datetime_unique_pk;
CREATE TABLE datetime_unique_pk (
    datetime_unique_pk datetime PRIMARY KEY UNIQUE
);

DROP TABLE IF EXISTS datetime_unique_check_pk;
CREATE TABLE datetime_unique_check_pk (
    datetime_unique_check_pk datetime PRIMARY KEY UNIQUE CHECK ( datetime_unique_check_pk > '2020-03-01 12:34:56' )
);

DROP TABLE IF EXISTS datetime_unique_check_pk_ref;
CREATE TABLE datetime_unique_check_pk_ref (
    datetime_unique_check_pk_ref datetime PRIMARY KEY UNIQUE CHECK ( datetime_unique_check_pk_ref > '2020-03-01 12:34:56' ) REFERENCES datetime_ref(datetime_ref)
);

DROP TABLE IF EXISTS datetime_check_pk;
CREATE TABLE datetime_check_pk (
    datetime_check_pk datetime PRIMARY KEY CHECK ( datetime_check_pk > '2020-03-01 12:34:56' )
);

DROP TABLE IF EXISTS datetime_def_const_unique_check_pk;
CREATE TABLE datetime_def_const_unique_check_pk (
    datetime_def_const_unique_check_pk datetime PRIMARY KEY UNIQUE CHECK ( datetime_def_const_unique_check_pk > '2020-03-01 12:34:56' ) DEFAULT ('2020-03-01 12:34:56')
);

DROP TABLE IF EXISTS datetime_def_const_unique_check_pk_ref;
CREATE TABLE datetime_def_const_unique_check_pk_ref (
    datetime_def_const_unique_check_pk_ref datetime PRIMARY KEY UNIQUE CHECK ( datetime_def_const_unique_check_pk_ref > '2020-03-01 12:34:56' ) DEFAULT ('2020-03-01 12:34:56') REFERENCES datetime_ref(datetime_ref)
);

DROP TABLE IF EXISTS datetime_def_func_unique_check_pk;
CREATE TABLE datetime_def_func_unique_check_pk (
    datetime_def_func_unique_check_pk datetime PRIMARY KEY UNIQUE CHECK ( datetime_def_func_unique_check_pk > '2020-03-01 12:34:56' ) DEFAULT (now())
);

DROP TABLE IF EXISTS datetime_def_func_unique_check_pk_ref;
CREATE TABLE datetime_def_func_unique_check_pk_ref (
    datetime_def_func_unique_check_pk_ref datetime PRIMARY KEY UNIQUE CHECK ( datetime_def_func_unique_check_pk_ref > '2020-03-01 12:34:56' ) DEFAULT (now()) REFERENCES datetime_ref(datetime_ref)
);
