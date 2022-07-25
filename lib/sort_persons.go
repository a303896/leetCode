package lib

type Person2 struct {
	FirstName string
	LastName string
}

type Persons []Person2

func (list Persons) Len() int {
	return len(list)
}

func (list Persons) Less(i,j int) bool {
	if list[i].LastName > list[j].LastName {
		return false
	}else {
		if list[i].FirstName > list[j].FirstName {
			return false
		}
	}
	return true
}

func (list Persons) Swap(i,j int)  {
	list[i],list[j] = list[j],list[i]
}

//example
/*
list := lib.Persons{{"Yichao", "Peng"},{"Chong", "Cao"},
		{"Kai", "Ding"}, {"Cao", "Cao"},{"Dan", "Cao"}}
fmt.Println(list)			//output:[{Yichao Peng} {Chong Cao} {Kai Ding} {Cao Cao} {Dan Cao}]
sort.Sort(list)
fmt.Println(list)			//output:[{Cao Cao} {Chong Cao} {Dan Cao} {Kai Ding} {Yichao Peng}]
 */