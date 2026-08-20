package user

import "testing"

func TestNewRepositorySeedsThreeUsers(t *testing.T) {
	repository := NewRepository()
	users := repository.FindAll()

	if len(users) != 3 {
		t.Fatalf("expected 3 users, got %d", len(users))
	}
	if users[0].ID != "usr-001" || users[1].ID != "usr-002" || users[2].ID != "usr-003" {
		t.Fatalf("unexpected seed data: %+v", users)
	}
}

func TestFindAllReturnsCopy(t *testing.T) {
	repository := NewRepository()
	users := repository.FindAll()
	users[0].Name = "changed"

	if repository.users[0].Name == "changed" {
		t.Fatal("expected FindAll to return a copy of the repository slice")
	}
}

func TestFindByID(t *testing.T) {
	repository := NewRepository()

	user, found := repository.FindByID("usr-002")
	if !found {
		t.Fatal("expected usr-002 to exist")
	}
	if user.Name != "Tran Minh Chau" {
		t.Fatalf("unexpected user: %+v", user)
	}

	if _, found := repository.FindByID("usr-missing"); found {
		t.Fatal("expected missing user lookup to return false")
	}
}