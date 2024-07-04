package chirp

import (
	"os"
	"testing"

	"github.com/mathiaskluge/chirpx/db"
	"github.com/mathiaskluge/chirpx/types"
)

func TestChirpStore(t *testing.T) {
	testDB, err := db.NewDB("db_test.json")
	if err != nil {
		t.Fatalf("error initializing test database: %v", err)
	}
	defer os.Remove("db_test.json")

	testStore := NewStore(testDB)

	testData := db.DBStructure{
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

	t.Run("Create chirp correctly in empty DB", func(t *testing.T) {
		chirpID, err := testStore.GenerateChirpID()
		if err != nil {
			t.Fatal(err)
		}

		chirp := types.Chirp{
			ID:   chirpID,
			Body: "First chirp",
		}

		err = testStore.CreateChirp(chirp)
		if err != nil {
			t.Fatal(err)
		}

		dat, err := testDB.LoadDB()
		if err != nil {
			t.Fatal(err)
		}

		if dat.Chirps[1] != chirp {
			t.Errorf("Expected: %v, Got: %v", chirp, dat.Chirps[1])
		}

	})

	t.Run("Create chirp correctly in populated DB", func(t *testing.T) {
		err = testDB.WriteDB(testData)
		if err != nil {
			t.Fatal(err)
		}

		chirpID, err := testStore.GenerateChirpID()
		if err != nil {
			t.Fatal(err)
		}

		chirp := types.Chirp{
			ID:   chirpID,
			Body: "First chirp",
		}

		err = testStore.CreateChirp(chirp)
		if err != nil {
			t.Fatal(err)
		}

		dat, err := testDB.LoadDB()
		if err != nil {
			t.Fatal(err)
		}

		if dat.Chirps[chirpID] != chirp {
			t.Errorf("Expected: %v, Got: %v", chirp, dat.Users[chirpID])
		}

	})
}
