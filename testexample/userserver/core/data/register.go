package data

func Register() {
	err := CreateUser()
	if err != nil {
		panic(err)
	}
}
