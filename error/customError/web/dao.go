package web

func getResult(id int) (Model, error) {
	return Do(id)
}
