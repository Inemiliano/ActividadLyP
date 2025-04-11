package domain


type Iperson interface{
	Save(persona *Person)error
	GetAll()([]Person,error)
	GetCountH()(int,error)
	GetCountM()(int,error)
}