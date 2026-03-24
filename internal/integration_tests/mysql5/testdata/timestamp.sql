DROP TABLE IF EXISTS timestamp_ref CASCADE;
CREATE TABLE timestamp_ref
(
    timestamp_ref timestamp UNIQUE DEFAULT CURRENT_TIMESTAMP
);

DROP TABLE IF EXISTS timestamp_table;
CREATE TABLE timestamp_table
(
    timestamp                        timestamp                 DEFAULT CURRENT_TIMESTAMP,
    timestamp_nn                     timestamp NOT NULL        DEFAULT CURRENT_TIMESTAMP,
    timestamp_nn_unique              timestamp NOT NULL UNIQUE DEFAULT CURRENT_TIMESTAMP,
    timestamp_nn_check               timestamp NOT NULL        DEFAULT CURRENT_TIMESTAMP,
    timestamp_nn_ref                 timestamp NOT NULL        DEFAULT CURRENT_TIMESTAMP REFERENCES timestamp_ref (timestamp_ref),
    timestamp_nn_def_const           timestamp NOT NULL        DEFAULT CURRENT_TIMESTAMP,
    timestamp_nn_def_func            timestamp NOT NULL        DEFAULT CURRENT_TIMESTAMP,
    timestamp_nn_unique_check        timestamp NOT NULL UNIQUE DEFAULT CURRENT_TIMESTAMP,

    timestamp_unique                 timestamp UNIQUE          DEFAULT CURRENT_TIMESTAMP,
    timestamp_unique_check           timestamp UNIQUE          DEFAULT CURRENT_TIMESTAMP,
    timestamp_unique_ref             timestamp UNIQUE          DEFAULT CURRENT_TIMESTAMP REFERENCES timestamp_ref (timestamp_ref),
    timestamp_unique_def_const       timestamp UNIQUE          DEFAULT CURRENT_TIMESTAMP,
    timestamp_unique_def_func        timestamp UNIQUE          DEFAULT CURRENT_TIMESTAMP,

    timestamp_check                  timestamp                 DEFAULT CURRENT_TIMESTAMP,
    timestamp_check_ref              timestamp                 DEFAULT CURRENT_TIMESTAMP REFERENCES timestamp_ref (timestamp_ref),
    timestamp_check_def_const        timestamp                 DEFAULT CURRENT_TIMESTAMP,
    timestamp_check_def_func         timestamp                 DEFAULT CURRENT_TIMESTAMP,

    timestamp_ref                    timestamp                 DEFAULT CURRENT_TIMESTAMP REFERENCES timestamp_ref (timestamp_ref),
    timestamp_ref_unique_check       timestamp UNIQUE          DEFAULT CURRENT_TIMESTAMP REFERENCES timestamp_ref (timestamp_ref),

    timestamp_def_const              timestamp                 DEFAULT CURRENT_TIMESTAMP,
    timestamp_def_const_unique_check timestamp UNIQUE          DEFAULT CURRENT_TIMESTAMP,

    timestamp_def_func               timestamp                 DEFAULT CURRENT_TIMESTAMP,
    timestamp_def_func_unique_check  timestamp UNIQUE          DEFAULT CURRENT_TIMESTAMP
);

DROP TABLE IF EXISTS timestamp_pk;
CREATE TABLE timestamp_pk
(
    timestamp_pk timestamp PRIMARY KEY DEFAULT CURRENT_TIMESTAMP
);

DROP TABLE IF EXISTS timestamp_pk_ref;
CREATE TABLE timestamp_pk_ref
(
    timestamp_pk_ref timestamp PRIMARY KEY DEFAULT CURRENT_TIMESTAMP REFERENCES timestamp_ref (timestamp_ref)
);

DROP TABLE IF EXISTS timestamp_pk_def_const;
CREATE TABLE timestamp_pk_def_const
(
    timestamp_pk_def_const timestamp PRIMARY KEY DEFAULT CURRENT_TIMESTAMP
);

DROP TABLE IF EXISTS timestamp_pk_def_func;
CREATE TABLE timestamp_pk_def_func
(
    timestamp_pk_def_func timestamp PRIMARY KEY DEFAULT CURRENT_TIMESTAMP
);

DROP TABLE IF EXISTS timestamp_nn_pk;
CREATE TABLE timestamp_nn_pk
(
    timestamp_nn_pk timestamp NOT NULL PRIMARY KEY DEFAULT CURRENT_TIMESTAMP
);

DROP TABLE IF EXISTS timestamp_nn_unique_check_pk;
CREATE TABLE timestamp_nn_unique_check_pk
(
    timestamp_nn_unique_check_pk timestamp PRIMARY KEY NOT NULL UNIQUE DEFAULT CURRENT_TIMESTAMP
);

DROP TABLE IF EXISTS timestamp_nn_unique_check_pk_ref;
CREATE TABLE timestamp_nn_unique_check_pk_ref
(
    timestamp_nn_unique_check_pk_ref timestamp PRIMARY KEY NOT NULL UNIQUE DEFAULT CURRENT_TIMESTAMP REFERENCES timestamp_ref (timestamp_ref)
);

DROP TABLE IF EXISTS timestamp_unique_pk;
CREATE TABLE timestamp_unique_pk
(
    timestamp_unique_pk timestamp PRIMARY KEY UNIQUE DEFAULT CURRENT_TIMESTAMP
);

DROP TABLE IF EXISTS timestamp_unique_check_pk;
CREATE TABLE timestamp_unique_check_pk
(
    timestamp_unique_check_pk timestamp PRIMARY KEY UNIQUE DEFAULT CURRENT_TIMESTAMP
);

DROP TABLE IF EXISTS timestamp_unique_check_pk_ref;
CREATE TABLE timestamp_unique_check_pk_ref
(
    timestamp_unique_check_pk_ref timestamp PRIMARY KEY UNIQUE DEFAULT CURRENT_TIMESTAMP REFERENCES timestamp_ref (timestamp_ref)
);

DROP TABLE IF EXISTS timestamp_check_pk;
CREATE TABLE timestamp_check_pk
(
    timestamp_check_pk timestamp PRIMARY KEY DEFAULT CURRENT_TIMESTAMP
);

DROP TABLE IF EXISTS timestamp_def_const_unique_check_pk;
CREATE TABLE timestamp_def_const_unique_check_pk
(
    timestamp_def_const_unique_check_pk timestamp PRIMARY KEY UNIQUE DEFAULT CURRENT_TIMESTAMP
);

DROP TABLE IF EXISTS timestamp_def_const_unique_check_pk_ref;
CREATE TABLE timestamp_def_const_unique_check_pk_ref
(
    timestamp_def_const_unique_check_pk_ref timestamp PRIMARY KEY UNIQUE DEFAULT CURRENT_TIMESTAMP REFERENCES timestamp_ref (timestamp_ref)
);

DROP TABLE IF EXISTS timestamp_def_func_unique_check_pk;
CREATE TABLE timestamp_def_func_unique_check_pk
(
    timestamp_def_func_unique_check_pk timestamp PRIMARY KEY UNIQUE DEFAULT CURRENT_TIMESTAMP
);

DROP TABLE IF EXISTS timestamp_def_func_unique_check_pk_ref;
CREATE TABLE timestamp_def_func_unique_check_pk_ref
(
    timestamp_def_func_unique_check_pk_ref timestamp PRIMARY KEY UNIQUE DEFAULT CURRENT_TIMESTAMP REFERENCES timestamp_ref (timestamp_ref)
);
