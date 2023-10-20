package main

//imports
import ( //compute html files
	//handle http requests
	"html/template"
	"log"
	"net/http"
)

func main() {
	//this handles any call to static assets will be resolved to the static directory
	http.Handle("/static/",
		http.StripPrefix("/static", http.FileServer(http.Dir("static"))))

	//what happens here is that when the route "/" is called, the HandleFunc runs the anon
	//function that is defined there, it loads a file from disk(index.html), compiles that file
	//then sends that file as the response
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("./templates/index.html"))
		tmpl.Execute(w, nil)
	})

	//This function here handles all requests that come to the /search endpoint
	//The lambda is executed similar to above
	http.HandleFunc("/search", func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(
			template.ParseFiles("./templates/fragments/results.html"))
		//The Execute method executes(compiles and runs) the parsed html template and writes the response to
		//the http response writer. The execute method takes in two arguments, response writer and data to be used
		//when rendering the template so in this case data is generated from the SearchTicker method in stocks.go
		//What's happening below is a map is being created with a key of string(results) and the value of that key is what
		//will be produced from the SearchTicker Method. The string passed into the method is whatever is coming as a parameter
		//e.g. url?key=29;
		data := map[string][]Stock{
			"Results": SearchTicker(r.URL.Query().Get("key"))}
		tmpl.Execute(w, data)
	})

	http.HandleFunc("/stock/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "POST":
			ticker := r.PostFormValue("ticker")
			stk := SearchTicker(ticker)[0]
			val := GetDailyValues(ticker)
			tmpl := template.Must((template.ParseFiles("./templates/index.html")))
			tmpl.ExecuteTemplate(w, "stock-element",
				Stock{Ticker: stk.Ticker, Name: stk.Name, Price: val.Open})
		}
	})

	//shows wherer the app is actually running
	log.Println("App running on 8080...")
	//fatal wraps the actual run of the server and logs a message if the server fails then exits the application
	//ListenAndServe blocks anything more and runs the server indefinitely, it should be the last line in your http server
	log.Fatal(http.ListenAndServe(":8000", nil))
}
