package db

import (
	"os"
	"reflect"
	"testing"

	"github.com/mathiaskluge/chirpx/types"
)

func TestDB(t *testing.T) {
	testDB, _ := NewDB("db_test.json")
	defer os.Remove("db_test.json")

	testData := DBStructure{
		Chirps: map[int]types.Chirp{
			1: {ID: 1, Body: "Hello world!"},
			2: {ID: 2, Body: "This is a chirp."},
			3: {ID: 3, Body: "Testing chirp data."},
		},
		Users: map[int]types.User{
			1: {ID: 1, Email: "abc@def.com", PwHash: "xyz"},
			2: {ID: 2, Email: "xyz@test.com", PwHash: "anotherhast"},
			3: {ID: 3, Email: "my@email.com", PwHash: "superhashihash"},
		},
	}

	t.Run("fails if DB does not exist", func(t *testing.T) {
		if _, err := os.Open(testDB.path); err != nil {
			t.Errorf("Expected database file at: %s", "db_test.json")
		}
	})

	t.Run("fails if write input does not equal load output", func(t *testing.T) {
		err := testDB.WriteDB(testData)
		if err != nil {
			t.Fatalf("writeDB returned error: %v", err)
		}

		loadedData, err := testDB.LoadDB()
		if err != nil {
			t.Fatalf("loadDB returned error: %v", err)
		}

		if !reflect.DeepEqual(loadedData, testData) {
			t.Errorf("loadedData does not match testData")
		}
	})

}
