package main

//generate params
var counter = [4]int{0,0,0,0}
var typeMultiplier = 1000000

func uuidGenerator(form int) int{
	uuid := 0
	switch form {
	case 1: //user 1
		uuid = typeMultiplier * form + counter[form]
		counter[form] += 1
	default:
		uuid = typeMultiplier * 1 + counter[form]
		counter[form] += 1
	}
	return uuid
}
