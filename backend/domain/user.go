package domain

type User struct {
	Aggregate
	FirstName      string
	LastName       string
	Email          string
	HashedPassword string
	Address        string
	Bio            string
	PhotoURL       string
}

func (user *User) SetPassword(newHashedPassword string) {
	user.HashedPassword = newHashedPassword
}

func (user *User) SetAddress(newAddress string) {
	user.Address = newAddress
}

func (user *User) SetBio(newBio string) {
	user.Bio = newBio
}

func (user *User) SetPhotoURL(newPhotoURL string) {
	user.PhotoURL = newPhotoURL
}

func (user *User) SetFirstName(newFirstName string) {
	user.FirstName = newFirstName
}

func (user *User) SetLastName(newLastName string) {
	user.FirstName = newLastName
}
