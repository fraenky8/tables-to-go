DROP TABLE IF EXISTS timestamp_without_time_zone_ref CASCADE;
CREATE TABLE timestamp_without_time_zone_ref (
    timestamp_without_time_zone_ref timestamp without time zone UNIQUE
);

DROP TABLE IF EXISTS timestamp_without_time_zone;
CREATE TABLE timestamp_without_time_zone (
    timestamp_without_time_zone timestamp without time zone,
    timestamp_without_time_zone_nn timestamp without time zone NOT NULL,
    timestamp_without_time_zone_nn_unique timestamp without time zone NOT NULL UNIQUE,
    timestamp_without_time_zone_nn_check timestamp without time zone NOT NULL CHECK ( timestamp_without_time_zone > '2020-03-01 12:34:56' ),
    timestamp_without_time_zone_nn_ref timestamp without time zone NOT NULL REFERENCES timestamp_without_time_zone_ref(timestamp_without_time_zone_ref),
    timestamp_without_time_zone_nn_def_const timestamp without time zone NOT NULL DEFAULT '2020-03-01 12:34:56',
    timestamp_without_time_zone_nn_def_func timestamp without time zone NOT NULL DEFAULT now(),
    timestamp_without_time_zone_nn_unique_check timestamp without time zone NOT NULL UNIQUE CHECK ( timestamp_without_time_zone > '2020-03-01 12:34:56' ),

    timestamp_without_time_zone_unique timestamp without time zone UNIQUE,
    timestamp_without_time_zone_unique_check timestamp without time zone UNIQUE CHECK ( timestamp_without_time_zone > '2020-03-01 12:34:56' ),
    timestamp_without_time_zone_unique_ref timestamp without time zone UNIQUE REFERENCES timestamp_without_time_zone_ref(timestamp_without_time_zone_ref),
    timestamp_without_time_zone_unique_def_const timestamp without time zone UNIQUE DEFAULT '2020-03-01 12:34:56',
    timestamp_without_time_zone_unique_def_func timestamp without time zone UNIQUE DEFAULT now(),

    timestamp_without_time_zone_check timestamp without time zone CHECK ( timestamp_without_time_zone > '2020-03-01 12:34:56' ),
    timestamp_without_time_zone_check_ref timestamp without time zone CHECK ( timestamp_without_time_zone > '2020-03-01 12:34:56' ) REFERENCES timestamp_without_time_zone_ref(timestamp_without_time_zone_ref),
    timestamp_without_time_zone_check_def_const timestamp without time zone CHECK ( timestamp_without_time_zone > '2020-03-01 12:34:56' ) DEFAULT '2020-03-01 12:34:56',
    timestamp_without_time_zone_check_def_func timestamp without time zone CHECK ( timestamp_without_time_zone > '2020-03-01 12:34:56' ) DEFAULT now(),

    timestamp_without_time_zone_ref timestamp without time zone REFERENCES timestamp_without_time_zone_ref(timestamp_without_time_zone_ref),
    timestamp_without_time_zone_ref_def_const timestamp without time zone REFERENCES timestamp_without_time_zone_ref(timestamp_without_time_zone_ref) DEFAULT '2020-03-01 12:34:56',
    timestamp_without_time_zone_ref_def_func timestamp without time zone REFERENCES timestamp_without_time_zone_ref(timestamp_without_time_zone_ref) DEFAULT now(),
    timestamp_without_time_zone_ref_unique_check timestamp without time zone UNIQUE CHECK ( timestamp_without_time_zone > '2020-03-01 12:34:56' ) REFERENCES timestamp_without_time_zone_ref(timestamp_without_time_zone_ref),

    timestamp_without_time_zone_def_const timestamp without time zone DEFAULT '2020-03-01 12:34:56',
    timestamp_without_time_zone_def_const_unique_check timestamp without time zone UNIQUE CHECK ( timestamp_without_time_zone > '2020-03-01 12:34:56' )DEFAULT '2020-03-01 12:34:56',

    timestamp_without_time_zone_def_func timestamp without time zone DEFAULT now(),
    timestamp_without_time_zone_def_func_unique_check timestamp without time zone UNIQUE CHECK ( timestamp_without_time_zone > '2020-03-01 12:34:56' ) DEFAULT now()
);

DROP TABLE IF EXISTS timestamp_without_time_zone_pk;
CREATE TABLE timestamp_without_time_zone_pk (
    timestamp_without_time_zone_pk timestamp without time zone PRIMARY KEY
);

DROP TABLE IF EXISTS timestamp_without_time_zone_pk_ref;
CREATE TABLE timestamp_without_time_zone_pk_ref (
    timestamp_without_time_zone_pk_ref timestamp without time zone PRIMARY KEY REFERENCES timestamp_without_time_zone_ref(timestamp_without_time_zone_ref)
);

DROP TABLE IF EXISTS timestamp_without_time_zone_pk_def_const;
CREATE TABLE timestamp_without_time_zone_pk_def_const (
    timestamp_without_time_zone_pk_def_const timestamp without time zone PRIMARY KEY DEFAULT '2020-03-01 12:34:56'
);

DROP TABLE IF EXISTS timestamp_without_time_zone_pk_def_func;
CREATE TABLE timestamp_without_time_zone_pk_def_func (
    timestamp_without_time_zone_pk_def_func timestamp without time zone PRIMARY KEY DEFAULT now()
);

DROP TABLE IF EXISTS timestamp_without_time_zone_nn_pk;
CREATE TABLE timestamp_without_time_zone_nn_pk (
    timestamp_without_time_zone_nn_pk timestamp without time zone NOT NULL PRIMARY KEY
);

DROP TABLE IF EXISTS timestamp_without_time_zone_nn_unique_check_pk;
CREATE TABLE timestamp_without_time_zone_nn_unique_check_pk (
    timestamp_without_time_zone_nn_unique_check_pk timestamp without time zone PRIMARY KEY NOT NULL UNIQUE CHECK ( timestamp_without_time_zone_nn_unique_check_pk > '2020-03-01 12:34:56' )
);

DROP TABLE IF EXISTS timestamp_without_time_zone_nn_unique_check_pk_ref;
CREATE TABLE timestamp_without_time_zone_nn_unique_check_pk_ref (
    timestamp_without_time_zone_nn_unique_check_pk_ref timestamp without time zone PRIMARY KEY NOT NULL UNIQUE CHECK ( timestamp_without_time_zone_nn_unique_check_pk_ref > '2020-03-01 12:34:56' ) REFERENCES timestamp_without_time_zone_ref(timestamp_without_time_zone_ref)
);

DROP TABLE IF EXISTS timestamp_without_time_zone_unique_pk;
CREATE TABLE timestamp_without_time_zone_unique_pk (
    timestamp_without_time_zone_unique_pk timestamp without time zone PRIMARY KEY UNIQUE
);

DROP TABLE IF EXISTS timestamp_without_time_zone_unique_check_pk;
CREATE TABLE timestamp_without_time_zone_unique_check_pk (
    timestamp_without_time_zone_unique_check_pk timestamp without time zone PRIMARY KEY UNIQUE CHECK ( timestamp_without_time_zone_unique_check_pk > '2020-03-01 12:34:56' )
);

DROP TABLE IF EXISTS timestamp_without_time_zone_unique_check_pk_ref;
CREATE TABLE timestamp_without_time_zone_unique_check_pk_ref (
    timestamp_without_time_zone_unique_check_pk_ref timestamp without time zone PRIMARY KEY UNIQUE CHECK ( timestamp_without_time_zone_unique_check_pk_ref > '2020-03-01 12:34:56' ) REFERENCES timestamp_without_time_zone_ref(timestamp_without_time_zone_ref)
);

DROP TABLE IF EXISTS timestamp_without_time_zone_check_pk;
CREATE TABLE timestamp_without_time_zone_check_pk (
    timestamp_without_time_zone_check_pk timestamp without time zone PRIMARY KEY CHECK ( timestamp_without_time_zone_check_pk > '2020-03-01 12:34:56' )
);

DROP TABLE IF EXISTS timestamp_without_time_zone_def_const_unique_check_pk;
CREATE TABLE timestamp_without_time_zone_def_const_unique_check_pk (
    timestamp_without_time_zone_def_const_unique_check_pk timestamp without time zone PRIMARY KEY UNIQUE CHECK ( timestamp_without_time_zone_def_const_unique_check_pk > '2020-03-01 12:34:56' ) DEFAULT '2020-03-01 12:34:56'
);

DROP TABLE IF EXISTS timestamp_without_time_zone_def_const_unique_check_pk_ref;
CREATE TABLE timestamp_without_time_zone_def_const_unique_check_pk_ref (
    timestamp_without_time_zone_def_const_unique_check_pk_ref timestamp without time zone PRIMARY KEY UNIQUE CHECK ( timestamp_without_time_zone_def_const_unique_check_pk_ref > '2020-03-01 12:34:56' ) DEFAULT '2020-03-01 12:34:56' REFERENCES timestamp_without_time_zone_ref(timestamp_without_time_zone_ref)
);

DROP TABLE IF EXISTS timestamp_without_time_zone_def_func_unique_check_pk;
CREATE TABLE timestamp_without_time_zone_def_func_unique_check_pk (
    timestamp_without_time_zone_def_func_unique_check_pk timestamp without time zone PRIMARY KEY UNIQUE CHECK ( timestamp_without_time_zone_def_func_unique_check_pk > '2020-03-01 12:34:56' ) DEFAULT now()
);

DROP TABLE IF EXISTS timestamp_without_time_zone_def_func_unique_check_pk_ref;
CREATE TABLE timestamp_without_time_zone_def_func_unique_check_pk_ref (
    timestamp_without_time_zone_def_func_unique_check_pk_ref timestamp without time zone PRIMARY KEY UNIQUE CHECK ( timestamp_without_time_zone_def_func_unique_check_pk_ref > '2020-03-01 12:34:56' ) DEFAULT now() REFERENCES timestamp_without_time_zone_ref(timestamp_without_time_zone_ref)
);
