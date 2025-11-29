package main

import (
	"fmt"
	"sort"
)

type CompanyInterface interface {
	AddWorkerInfo(name, position string, salary, experience uint) error
	SortWorkers() ([]string, error)
}

type Worker struct{
	Name string
	Position string
	Salary uint
	Experience uint
}


type Company struct{
	Workers []Worker
}

func (c *Company) AddWorkerInfo(name, position string, salary, experience uint) error{
	w := Worker{
		Name: name,
		Position: position,
		Salary: salary,
		Experience: experience,
	}

	c.Workers = append(c.Workers, w)
	return nil
}

func (c *Company) SortWorkers() ([]string, error){
	var positionRank = map[string]int{
		"директор" : 5,
		"зам. директора" : 4,
		"начальник цеха" : 3,
		"мастер" : 2,
		"рабочий" : 1,
	}

	sort.Slice(c.Workers, func(i, j int) bool {
		moneyI := c.Workers[i].Salary * c.Workers[i].Experience
		moneyJ := c.Workers[j].Salary * c.Workers[j].Experience

		rankI := positionRank[c.Workers[i].Position]
		rankJ := positionRank[c.Workers[j].Position]

		if moneyI != moneyJ{
			return  moneyI > moneyJ
		}
		return rankI > rankJ
	})

	result := make([]string, len(c.Workers))
	for i, w := range c.Workers{
		result[i] = w.Name
	}

	for i, w := range c.Workers{
		income := w.Salary * w.Experience
		result[i] = fmt.Sprintf("%s — %d — %s", w.Name, income, w.Position)
	}
	return result, nil
}