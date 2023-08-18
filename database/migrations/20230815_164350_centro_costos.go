package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type CentroCostos_20230815_164350 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &CentroCostos_20230815_164350{}
	m.Created = "20230815_164350"

	migration.Register("CentroCostos_20230815_164350", m)
}

const script = "../scripts/20230815_164350_centro_costos_"

// Run the migrations
func (m *CentroCostos_20230815_164350) Up() {
	file, err := os.ReadFile(script + "up.sql")

	if err != nil {
		// handle error
		fmt.Println(err)
	}

	requests := strings.Split(string(file), ";")

	for _, request := range requests {
		fmt.Println(request)
		m.SQL(request)
	}
}

// Reverse the migrations
func (m *CentroCostos_20230815_164350) Down() {
	file, err := os.ReadFile(script + "down.sql")

	if err != nil {
		// handle error
		fmt.Println(err)
	}

	requests := strings.Split(string(file), ";")

	for _, request := range requests {
		fmt.Println(request)
		m.SQL(request)
	}
}
