package dto

import (
	"database/sql"
)

type NonStrictAffinities struct {
	AffinityInt      sql.NullString `db:"affinity_int"`
	AffinityBigint   sql.NullString `db:"affinity_bigint"`
	AffinitySmallint sql.NullString `db:"affinity_smallint"`
	AffinityDecimal  sql.NullString `db:"affinity_decimal"`
	AffinityDouble   sql.NullString `db:"affinity_double"`
	AffinityBool     sql.NullBool   `db:"affinity_bool"`
	AffinityDate     sql.NullString `db:"affinity_date"`
	AffinityDatetime sql.NullString `db:"affinity_datetime"`
	AffinityVarchar  sql.NullString `db:"affinity_varchar"`
	AffinityChar     sql.NullString `db:"affinity_char"`
	AffinityClob     sql.NullString `db:"affinity_clob"`
	AffinityNchar    sql.NullString `db:"affinity_nchar"`
	AffinityNvarchar sql.NullString `db:"affinity_nvarchar"`
	AffinityJSON     sql.NullString `db:"affinity_json"`
	AffinityUuID     sql.NullString `db:"affinity_uuid"`
	AffinityBinary   sql.NullString `db:"affinity_binary"`
	AffinityVarbin   sql.NullString `db:"affinity_varbin"`
}
