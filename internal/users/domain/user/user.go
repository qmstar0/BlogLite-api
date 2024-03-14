package user

import (
	"common/domainevent"
)

type User struct {
	Uid      uint32
	Name     UserName
	Email    Email
	Password Password
	Roles    UserRole

	events []domainevent.DomainEvent
}

func NewUser(name UserName, email Email, pwd Password) *User {
	uid := email.ToID()
	return &User{
		Uid:      uid,
		Name:     name,
		Email:    email,
		Password: pwd,
		Roles:    Normally,
		events: []domainevent.DomainEvent{domainevent.NewDomainEvent(
			uid,
			RegistrationSuccess,
			UserRegistrationSuccess{
				Uid:      uid,
				Name:     name.String(),
				Email:    email.String(),
				Password: pwd.String(),
				Rights:   uint16(Normally),
			}),
		},
	}
}

func (u *User) ModifyUserName(name UserName) {
	u.events = append(u.events, domainevent.NewDomainEvent(u.Uid, NameChanged, UsernameChanged{
		Uid:     u.Uid,
		OldName: u.Name.String(),
		NewName: name.String(),
	}))
	u.Name = name
}

func (u *User) ResetPassowrd(pwd Password) {
	u.events = append(u.events, domainevent.NewDomainEvent(u.Uid, PasswordReset, UserPasswordReset{
		Uid:         u.Uid,
		OldPassowrd: u.Password.String(),
		NewPassword: pwd.String(),
	}))
	u.Password = pwd
}

func (u *User) ModifyRoles(roles UserRole) {
	u.events = append(u.events, domainevent.NewDomainEvent(u.Uid, RolesChanged, UserRolesChanged{
		Uid:       u.Uid,
		OldRights: uint16(u.Roles),
		NewRights: uint16(roles),
	}))

	u.Roles = roles
}

func (u *User) Login() {
	u.events = append(u.events, domainevent.NewDomainEvent(u.Uid, Login, UserLogin{
		Uid:   u.Uid,
		Email: u.Email.String(),
	}))
}

func (u *User) Logout() {
	u.events = append(u.events, domainevent.NewDomainEvent(u.Uid, Logout, UserLogout{
		Uid:   u.Uid,
		Email: u.Email.String(),
	}))
}

func EventFromAggregate(agg *User) []domainevent.DomainEvent {
	return agg.events
}
