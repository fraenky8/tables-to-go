DROP TABLE IF EXISTS non_strict_affinities;
CREATE TABLE non_strict_affinities
(
    affinity_int      int,
    affinity_bigint   bigint,
    affinity_smallint smallint,
    affinity_decimal  decimal(10,2),
    affinity_double   double,
    affinity_bool     boolean,
    affinity_date     date,
    affinity_datetime datetime,
    affinity_varchar  varchar(255),
    affinity_char     char(32),
    affinity_clob     clob,
    affinity_nchar    nchar(50),
    affinity_nvarchar nvarchar(50),
    affinity_json     json,
    affinity_uuid     uuid,
    affinity_binary   binary,
    affinity_varbin   varbinary(32)
);
