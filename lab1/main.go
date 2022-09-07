package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

type Post struct {
	Id    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"e-mail"`
	Title string `json:"title"`
	Text  string `json:"text"`
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("source/index.html")
	if err != nil {
		log.Fatal("indexHandler ", err)
	}

	tmpl.ExecuteTemplate(w, "index", nil)
}

func historyHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("source/history.html", "source/header.html", "source/footer.html")
	if err != nil {
		log.Fatal("historyHandler ", err)
	}

	tmpl.ExecuteTemplate(w, "history", nil)
}

func infrastructureHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("source/infrastructure.html", "source/header.html", "source/footer.html")
	if err != nil {
		log.Fatal("infrastructureHandler ", err)
	}

	tmpl.ExecuteTemplate(w, "infrastructure", nil)
}

func cultureHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("source/culture.html", "source/header.html", "source/footer.html")
	if err != nil {
		log.Fatal("cultureHandler ", err)
	}

	tmpl.ExecuteTemplate(w, "culture", nil)
}

func contacsHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("source/contacs.html", "source/header.html", "source/footer.html")
	if err != nil {
		log.Fatal("contacsHandler ", err)
	}

	tmpl.ExecuteTemplate(w, "contacs", nil)
}

func sendMessageHandler(w http.ResponseWriter, r *http.Request) { // not work
	// newId := getId()
	// post := Post{
	// 	Id:    newId,
	// 	Name:  r.FormValue("name"),
	// 	Email: r.FormValue("email"),
	// 	Title: r.FormValue("title"),
	// 	Text:  r.FormValue("message"),
	// }

	// file, err := os.OpenFile("list.json", os.O_RDWR, 0644)
	// if err != nil {
	// 	log.Fatal("OpenFile sendMessageHandler ", err)
	// }
	// defer file.Close()

	// b, _ := json.Marshal(post)
	// ioutil.WriteFile("list.json", b, 0666)

	// // if r.Method == "POST" {
	// // 	b, err := json.Marshal(post)

	// // 	fmt.Println(string(b))
	// // 	if err != nil {
	// // 		log.Fatal("Marshal", err)
	// // 	}

	// // 	ioutil.WriteFile("list.json", b, 0666)

	// // }
}

func handlerResponse() {

	port := ":9090"
	fmt.Println("Listen on port", port)
	//-------------------------------------------------

	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/history", historyHandler)
	http.HandleFunc("/infrastructure", infrastructureHandler)
	http.HandleFunc("/culture", cultureHandler)
	http.HandleFunc("/contacs", contacsHandler)
	http.HandleFunc("/contacs/addnewspost/", sendMessageHandler)

	http.Handle("/source/", http.StripPrefix("/source", http.FileServer(http.Dir("./source/"))))

	//-------------------------------------------------
	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatal("ListenAndServe", err)
	}
}

// func getId() string {
// 	id := rand.Int()
// 	return string(id)
// }

func main() {
	handlerResponse()

}
