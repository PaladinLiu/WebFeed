package main

//generate params
var counter = [6]int{0,0,0,0,0,0}
var typeMultiplier = 1000000

func uuidGenerator(form int) int{
	uuid := 0
	switch form {
	case 1: //user 1
		uuid = typeMultiplier * form + counter[form]
		counter[form] += 1
	case 2: //tweet 2
		uuid = typeMultiplier * form + counter[form]
		counter[form] += 1
	case 3: //comment 3
		uuid = typeMultiplier * form + counter[form]
		counter[form] += 1
	case 4: //like 4
		uuid = typeMultiplier * form + counter[form]
		counter[form] += 1
	case 5: //retweet 5
		uuid = typeMultiplier * form + counter[form]
		counter[form] += 1
	default:
		break
	}
	return uuid
}
