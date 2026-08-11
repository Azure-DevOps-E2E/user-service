package user

type Repository struct {
	users []User
}

func NewRepository() *Repository {
	return &Repository{users: []User{
		{ID: "usr-001", Name: "Nguyen Van An", Email: "an@example.com"},
		{ID: "usr-002", Name: "Tran Minh Chau", Email: "chau@example.com"},
		{ID: "usr-003", Name: "Le Hoang Nam", Email: "nam@example.com"},
	}}
}

func (r *Repository) FindAll() []User {
	result := make([]User, len(r.users))
	copy(result, r.users)
	return result
}

func (r *Repository) FindByID(id string) (User, bool) {
	for _, item := range r.users {
		if item.ID == id {
			return item, true
		}
	}
	return User{}, false
}
