package accountHolder

type AccountHolder struct {
	name, taxIdentifier, profession string
	age                             int
}

func CreateAccountHolder(name string, taxIdentifier string, profession string, age int) AccountHolder {
	newHolder := AccountHolder{}
	newHolder.name = name
	newHolder.taxIdentifier = taxIdentifier
	newHolder.profession = profession
	newHolder.age = age
	return newHolder
}
