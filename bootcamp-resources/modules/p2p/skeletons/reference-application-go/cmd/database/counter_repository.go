package database

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/coreeng/core-reference-application-go/cmd/structs"
	log "github.com/sirupsen/logrus"
)

var nameColumn = "name"
var counterColumn = "counter"

func IncrementCounterValue(counterName string) structs.Counter {
	database := GetDatabase()

	counter := findOrCreateCounter(counterName)
	counter.IncrementCounter()

	insertQuery := buildInsertToCounterTableQuery(nameColumn, counterColumn, counterName, counter.Counter)

	_, err := database.Exec(insertQuery)
	if err != nil {
		log.Errorf("Error occurred while inserting into the counter table, error: %v", err)
	}
	return counter
}

func buildInsertToCounterTableQuery(nameColumn string, counterColumn string, counterName string, counterValue uint64) string {
	return fmt.Sprintf("INSERT INTO counter(%s, %s)"+
		"VALUES ('%s', %d)"+
		"ON CONFLICT(name) "+
		"DO UPDATE SET counter=%d;",
		nameColumn, counterColumn,
		counterName, counterValue,
		counterValue)
}

func GetCounter(counterName string) structs.Counter {
	return findOrCreateCounter(counterName)
}
func findOrCreateCounter(counterName string) structs.Counter {
	database := GetDatabase()

	var counter structs.Counter
	query := fmt.Sprintf(SelectFromCounterTable, counterName)
	row := database.QueryRow(query)
	err := row.Scan(&counter.Name, &counter.Counter)
	if err != nil && errors.Is(err, sql.ErrNoRows) {
		return structs.Counter{Name: counterName, Counter: 0}
	} else if err != nil {
		log.Errorf("Error occurred while selecting counter, error: %v", err)
	}

	return counter
}
