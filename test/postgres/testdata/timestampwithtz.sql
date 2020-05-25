DROP TABLE IF EXISTS timestamp_with_time_zone_ref CASCADE;
CREATE TABLE timestamp_with_time_zone_ref (
    timestamp_with_time_zone_ref timestamp with time zone UNIQUE
);

DROP TABLE IF EXISTS timestamp_with_time_zone;
CREATE TABLE timestamp_with_time_zone (
    timestamp_with_time_zone timestamp with time zone,
    timestamp_with_time_zone_nn timestamp with time zone NOT NULL,
    timestamp_with_time_zone_nn_unique timestamp with time zone NOT NULL UNIQUE,
    timestamp_with_time_zone_nn_check timestamp with time zone NOT NULL CHECK ( timestamp_with_time_zone > '2020-03-01 12:34:56+8' ),
    timestamp_with_time_zone_nn_ref timestamp with time zone NOT NULL REFERENCES timestamp_with_time_zone_ref(timestamp_with_time_zone_ref),
    timestamp_with_time_zone_nn_def_const timestamp with time zone NOT NULL DEFAULT '2020-03-01 12:34:56+8',
    timestamp_with_time_zone_nn_def_func timestamp with time zone NOT NULL DEFAULT now(),
    timestamp_with_time_zone_nn_unique_check timestamp with time zone NOT NULL UNIQUE CHECK ( timestamp_with_time_zone > '2020-03-01 12:34:56+8' ),

    timestamp_with_time_zone_unique timestamp with time zone UNIQUE,
    timestamp_with_time_zone_unique_check timestamp with time zone UNIQUE CHECK ( timestamp_with_time_zone > '2020-03-01 12:34:56+8' ),
    timestamp_with_time_zone_unique_ref timestamp with time zone UNIQUE REFERENCES timestamp_with_time_zone_ref(timestamp_with_time_zone_ref),
    timestamp_with_time_zone_unique_def_const timestamp with time zone UNIQUE DEFAULT '2020-03-01 12:34:56+8',
    timestamp_with_time_zone_unique_def_func timestamp with time zone UNIQUE DEFAULT now(),

    timestamp_with_time_zone_check timestamp with time zone CHECK ( timestamp_with_time_zone > '2020-03-01 12:34:56+8' ),
    timestamp_with_time_zone_check_ref timestamp with time zone CHECK ( timestamp_with_time_zone > '2020-03-01 12:34:56+8' ) REFERENCES timestamp_with_time_zone_ref(timestamp_with_time_zone_ref),
    timestamp_with_time_zone_check_def_const timestamp with time zone CHECK ( timestamp_with_time_zone > '2020-03-01 12:34:56+8' ) DEFAULT '2020-03-01 12:34:56+8',
    timestamp_with_time_zone_check_def_func timestamp with time zone CHECK ( timestamp_with_time_zone > '2020-03-01 12:34:56+8' ) DEFAULT now(),

    timestamp_with_time_zone_ref timestamp with time zone REFERENCES timestamp_with_time_zone_ref(timestamp_with_time_zone_ref),
    timestamp_with_time_zone_ref_def_const timestamp with time zone REFERENCES timestamp_with_time_zone_ref(timestamp_with_time_zone_ref) DEFAULT '2020-03-01 12:34:56+8',
    timestamp_with_time_zone_ref_def_func timestamp with time zone REFERENCES timestamp_with_time_zone_ref(timestamp_with_time_zone_ref) DEFAULT now(),
    timestamp_with_time_zone_ref_unique_check timestamp with time zone UNIQUE CHECK ( timestamp_with_time_zone > '2020-03-01 12:34:56+8' ) REFERENCES timestamp_with_time_zone_ref(timestamp_with_time_zone_ref),

    timestamp_with_time_zone_def_const timestamp with time zone DEFAULT '2020-03-01 12:34:56+8',
    timestamp_with_time_zone_def_const_unique_check timestamp with time zone UNIQUE CHECK ( timestamp_with_time_zone > '2020-03-01 12:34:56+8' )DEFAULT '2020-03-01 12:34:56+8',

    timestamp_with_time_zone_def_func timestamp with time zone DEFAULT now(),
    timestamp_with_time_zone_def_func_unique_check timestamp with time zone UNIQUE CHECK ( timestamp_with_time_zone > '2020-03-01 12:34:56+8' ) DEFAULT now()
);

DROP TABLE IF EXISTS timestamp_with_time_zone_pk;
CREATE TABLE timestamp_with_time_zone_pk (
    timestamp_with_time_zone_pk timestamp with time zone PRIMARY KEY
);

DROP TABLE IF EXISTS timestamp_with_time_zone_pk_ref;
CREATE TABLE timestamp_with_time_zone_pk_ref (
    timestamp_with_time_zone_pk_ref timestamp with time zone PRIMARY KEY REFERENCES timestamp_with_time_zone_ref(timestamp_with_time_zone_ref)
);

DROP TABLE IF EXISTS timestamp_with_time_zone_pk_def_const;
CREATE TABLE timestamp_with_time_zone_pk_def_const (
    timestamp_with_time_zone_pk_def_const timestamp with time zone PRIMARY KEY DEFAULT '2020-03-01 12:34:56+8'
);

DROP TABLE IF EXISTS timestamp_with_time_zone_pk_def_func;
CREATE TABLE timestamp_with_time_zone_pk_def_func (
    timestamp_with_time_zone_pk_def_func timestamp with time zone PRIMARY KEY DEFAULT now()
);

DROP TABLE IF EXISTS timestamp_with_time_zone_nn_pk;
CREATE TABLE timestamp_with_time_zone_nn_pk (
    timestamp_with_time_zone_nn_pk timestamp with time zone NOT NULL PRIMARY KEY
);

DROP TABLE IF EXISTS timestamp_with_time_zone_nn_unique_check_pk;
CREATE TABLE timestamp_with_time_zone_nn_unique_check_pk (
    timestamp_with_time_zone_nn_unique_check_pk timestamp with time zone PRIMARY KEY NOT NULL UNIQUE CHECK ( timestamp_with_time_zone_nn_unique_check_pk > '2020-03-01 12:34:56+8' )
);

DROP TABLE IF EXISTS timestamp_with_time_zone_nn_unique_check_pk_ref;
CREATE TABLE timestamp_with_time_zone_nn_unique_check_pk_ref (
    timestamp_with_time_zone_nn_unique_check_pk_ref timestamp with time zone PRIMARY KEY NOT NULL UNIQUE CHECK ( timestamp_with_time_zone_nn_unique_check_pk_ref > '2020-03-01 12:34:56+8' ) REFERENCES timestamp_with_time_zone_ref(timestamp_with_time_zone_ref)
);

DROP TABLE IF EXISTS timestamp_with_time_zone_unique_pk;
CREATE TABLE timestamp_with_time_zone_unique_pk (
    timestamp_with_time_zone_unique_pk timestamp with time zone PRIMARY KEY UNIQUE
);

DROP TABLE IF EXISTS timestamp_with_time_zone_unique_check_pk;
CREATE TABLE timestamp_with_time_zone_unique_check_pk (
    timestamp_with_time_zone_unique_check_pk timestamp with time zone PRIMARY KEY UNIQUE CHECK ( timestamp_with_time_zone_unique_check_pk > '2020-03-01 12:34:56+8' )
);

DROP TABLE IF EXISTS timestamp_with_time_zone_unique_check_pk_ref;
CREATE TABLE timestamp_with_time_zone_unique_check_pk_ref (
    timestamp_with_time_zone_unique_check_pk_ref timestamp with time zone PRIMARY KEY UNIQUE CHECK ( timestamp_with_time_zone_unique_check_pk_ref > '2020-03-01 12:34:56+8' ) REFERENCES timestamp_with_time_zone_ref(timestamp_with_time_zone_ref)
);

DROP TABLE IF EXISTS timestamp_with_time_zone_check_pk;
CREATE TABLE timestamp_with_time_zone_check_pk (
    timestamp_with_time_zone_check_pk timestamp with time zone PRIMARY KEY CHECK ( timestamp_with_time_zone_check_pk > '2020-03-01 12:34:56+8' )
);

DROP TABLE IF EXISTS timestamp_with_time_zone_def_const_unique_check_pk;
CREATE TABLE timestamp_with_time_zone_def_const_unique_check_pk (
    timestamp_with_time_zone_def_const_unique_check_pk timestamp with time zone PRIMARY KEY UNIQUE CHECK ( timestamp_with_time_zone_def_const_unique_check_pk > '2020-03-01 12:34:56+8' ) DEFAULT '2020-03-01 12:34:56+8'
);

DROP TABLE IF EXISTS timestamp_with_time_zone_def_const_unique_check_pk_ref;
CREATE TABLE timestamp_with_time_zone_def_const_unique_check_pk_ref (
    timestamp_with_time_zone_def_const_unique_check_pk_ref timestamp with time zone PRIMARY KEY UNIQUE CHECK ( timestamp_with_time_zone_def_const_unique_check_pk_ref > '2020-03-01 12:34:56+8' ) DEFAULT '2020-03-01 12:34:56+8' REFERENCES timestamp_with_time_zone_ref(timestamp_with_time_zone_ref)
);

DROP TABLE IF EXISTS timestamp_with_time_zone_def_func_unique_check_pk;
CREATE TABLE timestamp_with_time_zone_def_func_unique_check_pk (
    timestamp_with_time_zone_def_func_unique_check_pk timestamp with time zone PRIMARY KEY UNIQUE CHECK ( timestamp_with_time_zone_def_func_unique_check_pk > '2020-03-01 12:34:56+8' ) DEFAULT now()
);

DROP TABLE IF EXISTS timestamp_with_time_zone_def_func_unique_check_pk_ref;
CREATE TABLE timestamp_with_time_zone_def_func_unique_check_pk_ref (
    timestamp_with_time_zone_def_func_unique_check_pk_ref timestamp with time zone PRIMARY KEY UNIQUE CHECK ( timestamp_with_time_zone_def_func_unique_check_pk_ref > '2020-03-01 12:34:56+8' ) DEFAULT now() REFERENCES timestamp_with_time_zone_ref(timestamp_with_time_zone_ref)
);
