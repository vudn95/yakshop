package controller

import (
	"encoding/json"
	"io/ioutil"
	"yakshop/model"
)

func CaculateYak(dayT int) (output *model.Output, err error) {
	var data model.Yak

	jsonFile, err := ioutil.ReadFile("data.json")
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal([]byte(jsonFile), &data)
	if err != nil {
		return nil, err
	}

	var milkTotal float32 = 0
	var woolTotal int64 = 0

	for i, herd := range data.Herd {
		var ageDay, woolDay float32
		if herd.Age >= 1 {
			woolDay = herd.Age * 100
			woolTotal++
		}

		for day := 0; day <= dayT-1; day++ {
			ageDay = herd.Age*100 + float32(day)
			if ageDay >= 1000 {
				break
			}

			woolLaps := 8 + ageDay*0.01
			milkTotal += (50 - ageDay*0.03)

			if woolLaps < ageDay-woolDay && ageDay >= 100 {
				woolDay = ageDay
				woolTotal++
			}
		}

		data.Herd[i].Age = (ageDay + 1) / 100
	}

	output = &model.Output{
		MilkTotal: milkTotal,
		WoolTotal: woolTotal,
		Yak:       data,
	}

	return output, nil
}
