package easypostdevtools

import "math/rand"

type Random struct {
}

func (r *Random) GetRandomBoolean() bool {
	return rand.Intn(2) == 1
}

func (r *Random) GetRandomIntInRange(min int, max int) int {
	return rand.Intn(max-min) + min
}

func (r *Random) GetRandomInt() int {
	return r.GetRandomIntInRange(0, 100)
}

func (r *Random) GetRandomDoubleInRange(min float32, max float32) float32 {
	return min + rand.Float32()*(max-min)
}

func (r *Random) GetRandomDouble() float32 {
	return r.GetRandomDoubleInRange(0.0, 100.0)
}

func (r *Random) GetRandomFloatInRange(min float32, max float32) float32 {
	return min + rand.Float32()*(max-min)
}

func (r *Random) GetRandomFloat() float32 {
	return r.GetRandomFloatInRange(0.0, 100.0)
}

func (r *Random) GetRandomChar() string {
	return r.GetRandomStringOfLength(1)
}

func (r *Random) GetRandomStringOfLength(length int) string {
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	b := make([]rune, length)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func (r *Random) GetRandomString() string {
	return r.GetRandomStringOfLength(r.GetRandomIntInRange(3, 10))
}

func (r *Random) RemoveItemFromList(s []interface{}, i int) []interface{} {
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}

func (r *Random) GetRandomItemsFromList(list []interface{}, amount int, allow_duplicates bool) []interface{} {
	if !allow_duplicates && amount > len(list) {
		panic("Amount must be less than or equal to list size when unique is true")
	}
	var randomItems []interface{}
	for i := 0; i < amount; i++ {
		loc := r.GetRandomIntInRange(0, len(list)-1)
		choice := list[loc]
		randomItems = append(randomItems, choice)
		if !allow_duplicates {
			list = r.RemoveItemFromList(list, loc)
		}
	}
	return randomItems
}

func (r *Random) GetRandomItemFromList(list []interface{}) interface{} {
	return r.GetRandomItemsFromList(list, 1, false)[0]
}
