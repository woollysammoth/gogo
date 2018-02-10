package main

func main() {
	a := App{}
	a.Initialize("root", "nomysql", "rest_api_example")

	a.Run(":8080")
}