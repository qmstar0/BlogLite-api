package user

import (
	"common/domainevent"
)

type User struct {
	Uid      uint32
	Name     UserName
	Email    Email
	Password Password
	Rights   UserRights

	events []domainevent.DomainEvent
}

func NewUser(name UserName, email Email, pwd Password) *User {
	uid := email.ToID()
	return &User{
		Uid:      uid,
		Name:     name,
		Email:    email,
		Password: pwd,
		Rights:   Normally,
		events: []domainevent.DomainEvent{domainevent.NewDomainEvent(
			uid,
			domainevent.Created,
			UserCreated{
				Uid:      uid,
				Name:     name.String(),
				Email:    email.String(),
				Password: pwd.String(),
				Rights:   uint16(Normally),
			}),
		},
	}
}

func (u *User) ChangeUsername(name UserName) {
	u.events = append(u.events, domainevent.NewDomainEvent(u.Uid, domainevent.Updated, UsernameUpdated{
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

func (u *User) AddPermissions(rights ...UserRights) {
	newRights := u.Rights
	for _, right := range rights {
		newRights |= right
	}
	u.changePermissions(newRights)
}

func (u *User) CancelPermissions(rights ...UserRights) {
	newRights := u.Rights
	for _, right := range rights {
		newRights &= ^right
	}
	u.changePermissions(newRights)
}

func (u *User) changePermissions(newRights UserRights) {
	u.events = append(u.events, domainevent.NewDomainEvent(u.Uid, RightsChanged, UserRightsChanged{
		Uid:       u.Uid,
		OldRights: uint16(u.Rights),
		NewRights: uint16(newRights),
	}))
	u.Rights = newRights
}

func EventFromAggregate(agg *User) []domainevent.DomainEvent {
	return agg.events
}
