package database

import "go.uber.org/dig"

type MyDataBase struct {
	User    UserModel
	Fllower FllowerModel
}

type UserModel interface {
	getUser()
}

type UserImpl struct {
}

func (Impl *UserImpl) getUser() {

}

type FllowerModel interface {
	setValue()
}

type FllowerImpl struct {
}

func (Impl *FllowerImpl) setValue() {

}

func initialDI() *MyDataBase {
	db := &MyDataBase{
		User:    &UserImpl{},
		Fllower: &FllowerImpl{},
	}
	return db

}

func BuildContainer() *dig.Container {
	container := dig.New()

	container.Provide(func() *MyDataBase {
		return initialDI()
	})

	return container
}

func Get() {
	container := BuildContainer()
	container.Invoke(func(MyDB *MyDataBase) {
		MyDB.User.getUser()

		// sql.DB is ready to use here
	})
}
