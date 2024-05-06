package authorization

import (
	"context"
	"platform/authorization/identity"
	"platform/services"
	"platform/sessions"
)

const USER_SESSION_KEY string = "USER"

func RegisterDefaultSignInService() {
	err := services.AddScoped(func(c context.Context) identity.SignInManager {
		return &SessionSignInMgr{Context: c}
	})
	if err != nil {
		panic(err)
	}
}

type SessionSignInMgr struct {
	context.Context
}

// SignIn implements identity.SignInManager.
func (s *SessionSignInMgr) SignIn(user identity.User) (err error) {
	session, err := s.getSession()
	if err == nil {
		session.SetValue(USER_SESSION_KEY, user.GetID())
	}
	return err
}

// SignOut implements identity.SignInManager.
func (s *SessionSignInMgr) SignOut(user identity.User) (err error) {
	session, err := s.getSession()
	if err == nil {
		session.SetValue(USER_SESSION_KEY, nil)
	}
	return
}

func (mgr *SessionSignInMgr) getSession() (s sessions.Session, err error) {
	err = services.GetServiceForContext(mgr.Context, &s)
	return
}
