package lib

import "fmt"

type Any interface{}

type SCar struct {
	Model string
	Manufacturer string
	BuildYear    int
}

type Cars []*SCar

func (c Cars) Process(f func(car *SCar))  {
	for _,item := range c {
		f(item)
	}
}

func (c Cars) FindAll(f func(car *SCar) bool) Cars {
	cars := make([]*SCar, 0)
	c.Process(func(car *SCar) {
		if f(car) {
			cars = append(cars, car)
		}
	})
	return cars
}

func MakeSortedAppender(manufacturers []string)(func(car *SCar), map[string]Cars) {
	sortedCars := make(map[string]Cars)
	for _,item := range manufacturers {
		sortedCars[item] = make([]*SCar, 0)
	}
	sortedCars["Default"] = make([]*SCar, 0)
	appender := func(c *SCar) {
		if _,ok := sortedCars[c.Manufacturer]; ok {
			sortedCars[c.Manufacturer] = append(sortedCars[c.Manufacturer], c)
		}else {
			sortedCars["Default"] = append(sortedCars["Default"], c)
		}
	}
	return appender, sortedCars
}

func (c SCar) String() string {
	return fmt.Sprintf("Manufacturer:%s, BuildYear:%d ", c.Manufacturer, c.BuildYear)
}

//example
/*
	ford := &lib.SCar{"Fiesta", "Ford", 2008}
	bmw := &lib.SCar{"XL 450", "BMW", 2011}
	merc := &lib.SCar{"D600", "Mercedes", 2009}
	bmw2 := &lib.SCar{"X 800", "BMW", 2008}

	allCars := lib.Cars{ford, bmw, merc, bmw2}
	fmt.Printf("allCars: %+v \n", allCars)	//output:allCars: [Manufacturer:Ford, BuildYear:2008  Manufacturer:BMW, BuildYear:2011
											//			Manufacturer:Mercedes, BuildYear:2009  Manufacturer:BMW, BuildYear:2008 ]

	allNewBMW := allCars.FindAll(func(car *lib.SCar) bool {
		return car.Manufacturer == "BMW" && car.BuildYear > 2008
	})
	fmt.Printf("allNewBMW: %+v \n", allNewBMW)	//output:allNewBMW: [Manufacturer:BMW, BuildYear:2011 ]

	manufacturers := []string{"Ford", "BMW", "Jaguar"}
	appender, list := lib.MakeSortedAppender(manufacturers)
	allCars.Process(appender)
	fmt.Printf("list: %+v \n", list)	//output:list: map[BMW:[Manufacturer:BMW, BuildYear:2011  Manufacturer:BMW, BuildYear:2008 ]
								//Default:[Manufacturer:Mercedes, BuildYear:2009 ] Ford:[Manufacturer:Ford, BuildYear:2008 ] Jaguar:[]]
*/