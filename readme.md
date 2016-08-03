# Tables-to-Go ##

* Table with name ```foo_table_name``` will become file ```FooTableName.go``` with struct name ```FooTableName```
* all columns get CamelCased, (hopefully correct) typed (eg ```int``` or ```sql.NullInt64```) and annotated (eg. ```db:"country_id"```)