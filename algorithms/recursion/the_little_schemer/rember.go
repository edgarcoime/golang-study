package main

// Prepends to list
func cons(a string, lat []string) []string {
	res := []string{a}
	copy(res[1:], lat)
	return res
}

// "Remove member"
// Removes first occurence of string (atom) in list
func rember(a string, lat []string) []string {
	// (null? lat (quote ()))
	if len(lat) <= 0 {
		return lat // should be empty
	}

	// (eq? a (car lat)) (cdr lat)
	if a == lat[0] {
		return lat[1:]
	}

	// else cons (car lat) (rember a (cdr lat))
	return append([]string{lat[0]}, (rember(a, lat[1:]))...)
}

func main() {
	a := "and"
	lat := []string{"bacon", "lettuce", "and", "tomato"}
	ans := rember(a, lat)
	println(ans)

	a1 := "yes"
	lat2 := []string{"yes"}
	lat1 := []string{"sir", "why", "not"}
	ans2 := cons(a1, lat1)
	ans3 := append(lat1, lat2...)
	println(ans2)
	println(ans3)
}
