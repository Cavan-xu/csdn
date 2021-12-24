package entity

type User struct {
	Id      int64  `json:"id"`
	Name    string `json:"name"`
	Age     int64  `json:"age"`
	country string `json:"country"`
}

// User 值接收者方法
func (u User) GetId(id int64) int64 {
	u.Id = id
	return u.Id
}

// *User 指针接收者方法
func (u *User) GetName() string {
	return u.Name
}

func (u *User) GetAge() int64 {
	return u.Age
}

func (u *User) getCountry() string {
	return u.country
}
