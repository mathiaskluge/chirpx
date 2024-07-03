package user

import (
	"os"
	"testing"

	"github.com/mathiaskluge/chirpx/db"
	"github.com/mathiaskluge/chirpx/types"
)

func TestUserStore(t *testing.T) {
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

	t.Run("fails if user is not created of empty DB", func(t *testing.T) {
		userID, err := testStore.GenerateUserID()
		if err != nil {
			t.Fatal(err)
		}

		user := types.User{
			ID:     userID,
			Email:  "New@new.com",
			PwHash: "newhash",
		}

		err = testStore.CreateUser(user)
		if err != nil {
			t.Fatal(err)
		}

		dat, err := testDB.LoadDB()
		if err != nil {
			t.Fatal(err)
		}

		if dat.Users[1] != user {
			t.Errorf("Expected: %v, Got: %v", user, dat.Users[1])
		}

	})

	t.Run("fails if user is not created off existing DB correctly", func(t *testing.T) {
		err = testDB.WriteDB(testData)
		if err != nil {
			t.Fatal(err)
		}

		userID, err := testStore.GenerateUserID()
		if err != nil {
			t.Fatal(err)
		}

		user := types.User{
			ID:     userID,
			Email:  "New@new.com",
			PwHash: "newhash",
		}

		err = testStore.CreateUser(user)
		if err != nil {
			t.Fatal(err)
		}

		dat, err := testDB.LoadDB()
		if err != nil {
			t.Fatal(err)
		}

		if dat.Users[userID] != user {
			t.Errorf("Expected: %v, Got: %v", user, dat.Users[userID])
		}

	})
}
