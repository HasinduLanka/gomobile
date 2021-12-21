package GoApp

import (
	"log"
	"net/http"
	"os"
	"strconv"

	_ "golang.org/x/mobile/bind"
)

func Run(cwd string) {
	os.Chdir(cwd)
	println("Hello, World!")
	run_server()
}

func Add(a, b int) string {
	return strconv.Itoa(a + b)
}

// simple http server

func run_server() {
	http.HandleFunc("/", func(res http.ResponseWriter, r *http.Request) {
		path, err := os.Getwd()

		if err != nil {
			path = err.Error()
		}

		data := []byte(`
<!DOCTYPE html>
<html>
<head>
<title>Go flutter</title>
</head>

<body>
<br />
<br />
<h1>Hello, from go flutter!</h1>
<br />
<p > cwd : ` + path + ` </p>
</body>

</html>
`)

		os.WriteFile("./file1", data, 0644)
		RBs, RE := os.ReadFile("./file1")

		if RE != nil {
			println(RE.Error())
			RBs = []byte("<br/><br/>" + RE.Error())
		}

		res.Header().Set("Content-Type", "text/html")
		res.WriteHeader(200)
		res.Write(RBs)

	})
	log.Fatal(http.ListenAndServe(":35298", nil))
}
