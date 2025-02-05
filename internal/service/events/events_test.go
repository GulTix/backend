package events_test

import (
	repository "backend/internal/repository/events"
	"backend/internal/service/events"
	"backend/pkg/database"
	oauth_pkg "backend/pkg/oauth"
	"backend/pkg/testutil"
	"context"
	"log"
	"net/http"
	"os"
	"reflect"
	"testing"

	"golang.org/x/oauth2"
)

var (
	ctx     context.Context
	db      database.DB
	repo    repository.Repository
	oauth   oauth_pkg.OAuth
	s       events.Service
	eventId string
)

func TestCreateEvent(t *testing.T) {
	s := events.NewService(repo, oauth)

	res, err := s.Create(ctx, events.CreateBody{
		Name:           "Test Event",
		GoogleFormLink: "https://forms.gle/123",
	})

	if err != nil {
		t.Error(err)
	}

	if res.StatusCode != http.StatusCreated {
		t.Errorf("Expected status code 201, got %d", res.StatusCode)
	}

	if res.Message != "Event berhasil dibuat" {
		t.Errorf("Expected message 'Event  berhasil dibuat', got %s", res.Message)
	}

	if res.Data.Name != "Test Event" {
		t.Errorf("Expected event name 'Test Event', got %s", res.Data.Name)
	}

	if res.Data.GoogleFormLink != "https://forms.gle/123" {
		t.Errorf("Expected google form link 'https://forms.gle/123', got %s", res.Data.GoogleFormLink)
	}

	if res.Data.Deleted != false {
		t.Errorf("Expected deleted false, got %t", res.Data.Deleted)
	}

	if res.Data.Id == "" {
		t.Error("Expected event id not empty")
	}

	eventId = res.Data.Id
}

func TestFindAll(t *testing.T) {
	s := events.NewService(repo, oauth)

	res, err := s.FindAll(ctx)

	if err != nil {
		t.Error(err)
	}

	if len(res.Data) == 0 && res.StatusCode != http.StatusNotFound {
		t.Errorf("Status Code does'nt match got %d", res.StatusCode)
	}

	if len(res.Data) != 0 && res.StatusCode != http.StatusOK {
		t.Errorf("Status Code does'nt match got %d", res.StatusCode)
	}
}

func TestMain(t *testing.M) {
	err := testutil.LoadEnv()

	if err != nil {
		log.Println(err)
	}

	ctx = context.Background()
	db = database.NewDB(ctx, os.Getenv("DATABASE_URL"))
	repo = repository.NewRepository(db)
	oauth = oauth_pkg.NewClient([]string{
		"https://www.googleapis.com/auth/userinfo.email",
	})

	s = events.NewService(repo, oauth)

	defer repo.Delete(ctx, eventId)
}

func TestGetGoogleToken(t *testing.T) {
	eventsRes, _ := s.FindAll(ctx)

	eventId := eventsRes.Data[1].Id

	res, err := s.GetBlasterToken(ctx, eventId)

	if err != nil {
		t.Error(err)
	}

	if reflect.TypeOf(res) != reflect.TypeOf(&oauth2.Token{}) {
		t.Error("it doesnt return token")
	}

	if res.RefreshToken == "" {
		t.Error("Refresh Token is null. Token not usable")
	}
}
