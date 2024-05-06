package placeholder

import (
	"platform/authorization/identity"
	"platform/services"
	"strings"
)

func RegisterPlaceholderUserStore() {
	err := services.AddSingleton(func() identity.UserStore {
		return &PlaceholderUserStore{}
	})
	if err != nil {
		panic(err)
	}
}

type PlaceholderUserStore struct{}

var users = map[int]identity.User{
	1: identity.NewBasicUser(1, "Alice", "Administrator"),
	2: identity.NewBasicUser(2, "Bob"),
}

func (store *PlaceholderUserStore) GetUserByID(id int) (identity.User, bool) {
	user, found := users[id]
	return user, found
}

func (store *PlaceholderUserStore) GetUserByName(name string) (identity.User, bool) {
	for _, user := range users {
		if strings.EqualFold(user.GetDisplayName(), name) {
			return user, true
		}
	}
	return nil, false
}
