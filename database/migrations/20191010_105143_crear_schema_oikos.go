package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type CrearSchemaOikos_20191010_105143 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &CrearSchemaOikos_20191010_105143{}
	m.Created = "20191010_105143"

	migration.Register("CrearSchemaOikos_20191010_105143", m)
}

// Run the migrations
func (m *CrearSchemaOikos_20191010_105143) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("CREATE SCHEMA IF NOT EXISTS oikos;")
	m.SQL("ALTER SCHEMA oikos OWNER TO desarrollooas;")
	m.SQL("SET search_path TO pg_catalog,public,oikos;")
}

// Reverse the migrations
func (m *CrearSchemaOikos_20191010_105143) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("DROP SCHEMA IF EXISTS oikos");
}
