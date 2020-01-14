package fixtures

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/gofrs/uuid"
	"github.com/jinzhu/gorm"

	"github.com/lavaninho/Projet-GO/models"
	"github.com/lavaninho/Projet-GO/utils"
)

var db, err = gorm.Open("postgres", "host=localhost port=5432 user=postgres-dev dbname=dev sslmode=disable password=s3cr3tp4ssw0rd")

// TestInsertUser fixture users
func TestInsertUser(t *testing.T) {

	if err != nil {
		t.Fail()
	}

	for i := 1; i < 5; i++ {
		fmt.Println(strconv.Itoa(i))
		user := models.User{
			FirstName:   "testuser" + strconv.Itoa(i),
			LastName:    "testuser" + strconv.Itoa(i),
			Email:       "testuser" + strconv.Itoa(i) + "@test.com",
			Password:    utils.Hash("testuser" + strconv.Itoa(i)),
			UUID:        uuid.Must(uuid.NewV4()).String(),
			AccessLevel: 2,
			DateOfBirth: "2000-01-01T00:00:00.000Z",
		}

		if err := db.Create(&user).Error; err != nil {
			t.Fail()
		}
	}
}

// TestInsertAdmin fixture admin
func TestInsertAdmin(t *testing.T) {

	if err != nil {
		t.Fail()
	}

	user := models.User{
		FirstName:   "testadmin",
		LastName:    "testadmin",
		Email:       "testadmin" + "@test.com",
		Password:    utils.Hash("testadmin"),
		UUID:        uuid.Must(uuid.NewV4()).String(),
		AccessLevel: 1,
		DateOfBirth: "2000-01-01T00:00:00.000Z",
	}

	if err := db.Create(&user).Error; err != nil {
		t.Fail()
	}
}

func TestInserVotes(t *testing.T) {
	var user []models.User
	var userVote []string

	if err != nil {
		t.Fail()
	}

	if err := db.Find(&user).Error; err != nil {
		t.Fail()
	}

	for i := 1; i < 4; i++ {
		userVote = append(userVote, user[i].UUID)
	}

	vote := models.Vote{
		Title:       "testvote",
		Description: "testvotedescription",
		UUIDVote:    userVote,
		StartDate:   "2000-06-01T00:00:00.000Z",
		EndDate:     "2000-06-10T00:00:00.000Z",
		UUID:        user[0].UUID,
	}

	if err := db.Create(&vote).Error; err != nil {
		t.Fail()
	}
}
