package main

type UpdateEventArguments struct {
	ID   string `json:"id"`
	Data struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	} `json:"data"`
}
