package models

import "time"

// User is the generic "user" model implementation.
// User should be passed around anywhere that involves authentication.
type User struct {
	Model

	ID        string     `json:"id"`
	Name      string     `json:"name"`
	Username  string     `json:"username"`
	Email     string     `json:"email"`
	Password  string     `json:"password"`
	Level     UserLevel  `json:"level"`
	UpdatedAt *time.Time `json:"updated_at"`
	CreatedAt *time.Time `json:"created_at"`
}

// Anonymize will return an AnonymousUser from a user object for API output.
func (u User) Anonymize() AnonymousUser {
	return AnonymousUser{
		Name:     u.Name,
		Username: u.Username,
	}
}

// AnonymousUser is a type returned from public endpoints or to non-admins.
// There is a function on User to return this type.
type AnonymousUser struct {
	Name     string `json:"name"`
	Username string `json:"username"`
}

// Level is an enum of user levels
type Level int

// Level enum values.
const (
	LevelBanned      = -1
	LevelDefault     = 0
	LevelJunior      = 10
	LevelLolibrarian = 50
	LevelSenior      = 100
	LevelAdmin       = 500
	LevelDeveloper   = 1000
)

// ItemStatusExternal is an external GraphQL representation of ItemStatus.
type UserLevel string

// UserLevel enum values.
const (
	UserLevelBanned      = "BANNED"
	UserLevelDefault     = "DEFAULT"
	UserLevelJunior      = "JUNIOR_LOLIBRARIAN"
	UserLevelLolibrarian = "LOLIBRARIAN"
	UserLevelSenior      = "SENIOR_LOLIBRARIAN"
	UserLevelAdmin       = "ADMIN"
	UserLevelDeveloper   = "DEVELOPER"
)

// UserLevels is a map of Level (int) to UserLevel (graphql string).
var UserLevels = map[Level]UserLevel{
	LevelBanned:      UserLevelBanned,
	LevelDefault:     UserLevelDefault,
	LevelJunior:      UserLevelJunior,
	LevelLolibrarian: UserLevelLolibrarian,
	LevelSenior:      UserLevelSenior,
	LevelAdmin:       UserLevelAdmin,
	LevelDeveloper:   UserLevelDeveloper,
}

// Levels is a map of UserLevel (graphql string) to Level (int).
var Levels = map[UserLevel]Level{
	UserLevelBanned:      LevelBanned,
	UserLevelDefault:     LevelDefault,
	UserLevelJunior:      LevelJunior,
	UserLevelLolibrarian: LevelLolibrarian,
	UserLevelSenior:      LevelSenior,
	UserLevelAdmin:       LevelAdmin,
	UserLevelDeveloper:   LevelDeveloper,
}
