package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type CrearSchemaOikos_20191105_155035 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &CrearSchemaOikos_20191105_155035{}
	m.Created = "20191105_155035"

	migration.Register("CrearSchemaOikos_20191105_155035", m)
}

// Run the migrations
func (m *CrearSchemaOikos_20191105_155035) Up() {
	m.SQL("CREATE SCHEMA IF NOT EXISTS oikos;")
	m.SQL("SET search_path TO pg_catalog,public,oikos;")
}

// Reverse the migrations
func (m *CrearSchemaOikos_20191105_155035) Down() {
	m.SQL("DROP SCHEMA IF EXISTS oikos")
}
