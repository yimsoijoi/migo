package svc

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func ExecDB(db *sqlx.DB) ([]DBModel, error) {
	models := []DBModel{}
	err := db.Select(&models, queryStr)
	if err != nil {
		return nil, wrapError("failed to query db", err)
	}
	return models, nil
}

func ConnectDB(host, user, password, dbName string) (*sqlx.DB, error) {
	db, err := sqlx.Connect("postgres", getDataSourceName(host, user, password, dbName))
	if err != nil {
		return nil, wrapError("failed to connect db", err)
	}
	return db, nil
}

func getDataSourceName(host, user, password, dbName string) string {
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
		host,
		user,
		password,
		dbName,
		"5432",
		"disable",
		"UTC",
	)
}

const queryStr = `
WITH EnumValues AS (
    SELECT
        n.nspname AS enum_schema,
        t.typname AS enum_name,
        string_agg(e.enumlabel, ', ') AS enum_value
    FROM
        pg_type t
            JOIN pg_enum e ON t.oid = e.enumtypid
            JOIN pg_catalog.pg_namespace n ON n.oid = t.typnamespace
    WHERE
            n.nspname = 'public'
    GROUP BY
        enum_schema,
        enum_name
)

SELECT
    c.table_name,
    c.column_name,
    CASE
        WHEN EXISTS (
                SELECT 1
                FROM information_schema.table_constraints tc
                         JOIN information_schema.key_column_usage kcu ON tc.constraint_name = kcu.constraint_name
                WHERE tc.table_name = c.table_name AND kcu.column_name = c.column_name AND tc.constraint_type = 'PRIMARY KEY'
            ) THEN 'PRIMARY KEY'
        WHEN EXISTS (
                SELECT 1
                FROM information_schema.table_constraints tc
                         JOIN information_schema.key_column_usage kcu ON tc.constraint_name = kcu.constraint_name
                WHERE tc.table_name = c.table_name AND kcu.column_name = c.column_name AND tc.constraint_type = 'FOREIGN KEY'
            ) THEN 'FOREIGN KEY'

        WHEN EXISTS (
                SELECT 1
                FROM information_schema.table_constraints tc
                         JOIN information_schema.key_column_usage kcu ON tc.constraint_name = kcu.constraint_name
                WHERE tc.table_name = c.table_name AND kcu.column_name = c.column_name AND tc.constraint_type = 'UNIQUE'
            ) THEN 'UNIQUE'
        ELSE '-'
        END AS constraint,
    CASE
        WHEN c.data_type = 'USER-DEFINED' THEN ev.enum_name
        ELSE c.data_type
        END AS data_type,
    CASE
        WHEN c.character_maximum_length IS NULL AND c.numeric_precision IS NOT NULL
            THEN CONCAT('(', COALESCE(c.numeric_precision::varchar, ''), ', ', COALESCE(c.numeric_scale::varchar, ''), ')')
        WHEN c.data_type = 'uuid'
            THEN '36'
        ELSE COALESCE(CAST(c.character_maximum_length AS varchar), '-')
        END AS size,
    CASE
        WHEN c.is_nullable = 'NO' THEN 'Y'
        WHEN c.is_nullable = 'YES' THEN 'N'
        END AS required,
    CASE
        WHEN c.data_type = 'uuid' THEN '25869d9c-65e2-4a46-aa2c-f67e1e3447bd'
        WHEN c.data_type = 'timestamp without time zone' THEN '2024-01-01 00:00:00.000000'
        WHEN c.data_type = 'timestamp with time zone' THEN '2023-01-01 00:00:00.000000 +00:00'
        WHEN c.data_type = 'integer' THEN '178932'
        WHEN c.data_type = 'bigint' THEN '178932456'
        WHEN c.data_type = 'numeric' AND c.numeric_scale = 2 THEN '1789.32'
        WHEN c.data_type = 'numeric' AND c.numeric_scale = 3 THEN '1789.324'
        WHEN c.column_name = 'name' THEN 'John'
        WHEN c.column_name = 'last_name' THEN 'Doe'
        WHEN ev.enum_value IS NOT NULL THEN ev.enum_value
        ELSE ''
        END AS example_value,
    COALESCE(c.column_default, '-') AS default_value,
    CASE
        WHEN EXISTS (
                SELECT  am.amname
                FROM pg_index i
                         JOIN pg_attribute a ON a.attnum = ANY(i.indkey)
                         JOIN pg_class cl ON cl.oid = i.indexrelid
                         JOIN pg_am am ON am.oid = cl.relam
                WHERE i.indrelid = c.table_name::regclass::oid
                  AND a.attname = c.column_name
            ) THEN (
            SELECT STRING_AGG(DISTINCT am.amname, ', ')
            FROM pg_index i
                     JOIN pg_attribute a ON a.attnum = ANY(i.indkey)
                     JOIN pg_class cl ON cl.oid = i.indexrelid
                     JOIN pg_am am ON am.oid = cl.relam
            WHERE i.indrelid = c.table_name::regclass::oid
              AND a.attname = c.column_name
        )
        ELSE '-'
        END AS index
FROM
    information_schema.columns c
        LEFT JOIN
    EnumValues ev ON c.udt_name = ev.enum_name
WHERE
        c.table_schema = 'public'
ORDER BY
    c.table_name,
    CASE
        WHEN c.column_name = 'id' THEN 0
        WHEN c.column_name LIKE '%id' THEN 1
        WHEN c.column_name IN ('created_at', 'created_by','created_by_fullname','created_by_full_name', 'updated_at', 'updated_by','updated_by_fullname','updated_by_full_name', 'deleted_at', 'deleted_by','deleted_by_fullname','deleted_by_full_name') THEN 3
        ELSE 2
        END;
`
