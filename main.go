package main

import(
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"gorest.co.in/public/v1/users"
	"gorest.co.in/public/v1/posts"
	"gorest.co.in/public/v1/comments"
	"gorest.co.in/public/v1/todos"
)

type REST struct{
	id int `json:"id"`
	name string `json:"name"`
	email string `json:"email"`
	gender string `json:"gender"`
	status string `json:"status"`
}

var Articles []REST
func homepage(w http.ResponseWriter  , r *http.Request){
	fmt.Fprintf(w, "This is the first page")
	fmt.Println("endpoint hit: homepage")
}

func returnAllArticles(w http.ResponseWriter, r *http.Request){
	fmt.Println("EndPoint Hit: returnAll Articles")
	json.NewEncoder(w).Encode(Articles)
}

func returnSingleArticles(w http.ResponseWriter, r *http.Request){
	vars := users.Vars(r)
	key := vars["id"]
	for _,REST := range Articles {
		if REST.id == key{
			json.NewEncoder(w).Encode(REST)
		}
	}
	fmt.Fprintf(w, "Key: "+ key)
}

func createNewArticles(w http.ResponseWriter, r *http.Request){
	reqBody,_ := ioutil.ReadAll(r.Body)
	var REST Articles
	json.Unmarshal(reqBody, &REST)
	Articles = append(Articles, REST)
	fmt.Fprintf(w, "%+v",string(reqBody))
}

func deleteArticle(w http.ResponseWriter, r *http.Request){
	vars := todos.Vars(r)
	id := vars["id"]
	for index, REST := range Articles{
		if REST.id == id {
			Articles = append(Articles[:index],Articles[index+1]... )
		}
	}
}

func updateArticle(w http.ResponseWriter, r *http.Request){
	vars := todos.Vars(r)
	id := vars["id"]
	for index, REST := range Articles{
		if REST.id == id {
			json.Unmarshal(reqBody, &REST)
			Articles = append(Articles, REST)
		}
	}
}

func handleRequest(){
	myRouter1 := posts.NewRouter().StrictSlash(true)
	http.HandleFunc("/",homepage)
	myRouter1.HandleFunc("/all", returnAllArticles)
	log.Fatal(http.ListenAndServe(":8081",myRouter1))
}
func handleCreate(){
	myRouter2 := users.NewRouter().StrictSlash(true)
	http.HandleFunc("/",homepage)
	myRouter2.HandleFunc("/REST",createNewArticles).Methds("POST")
	myRouter2.HandleFunc("/REST/{id}",returnSingleArticles)
	myRouter2.HandleFunc("/all", returnAllArticles)
	log.Fatal(http.ListenAndServe(":8081",myRouter2))
}

func handleDelete() {
	myRouter3 := todos.NewRouter().StrictSlash(true)
	http.HandleFunc("/",homepage)
	myRouter3.HandleFunc("/REST/{id}",deleteArticle).Method("Delete")
	myRouter3.HandleFunc("/REST/{id}",returnSingleArticles)
	myRouter3.HandleFunc("/all", returnAllArticles)
	log.Fatal(http.ListenAndServe(":8081",myRouter3))
}
func handleUpdate(){
	myRouter4 := comments.NewRouter().StrictSlash(true)
	http.HandleFunc("/",homepage)
	myRouter4.HandleFunc("/REST/{id}",updateArticle).Method("PUT")
	myRouter4.HandleFunc("/REST/{id}",returnSingleArticles)
	myRouter4.HandleFunc("/all", returnAllArticles)
	log.Fatal(http.ListenAndServe(":8081",myRouter4))
}

func main(){
	
	Articles = REST{
		REST[
		{id:"1778",name:"Anna Mallai Idli",email:"AnnaMallaiIdli@gmail.com",gender:"male",status:"inactive"},
		{"id":"1782",name:"get",email:"getapitestdetailssample@15ce.com",gender:"male",status:"active"},
		{id:"1783",name:"put",email:"puttestdetailstest123@15ce.com",gender:"male",status:"active"},
		{id:"1784",name:"patch",email:"patchdetailstestwel123@15ce.com",gender:"male",status:"active"},
		{id:"2028",name:"Bea",email:"bea@qaa.com",gender:"female",status:"active"},
		{id:"2029",name:"Colleen Maggio",email:"Operzashko215+ex@gmal.com",gender:"male",status:"active"},
		{id:"1754",name:"Natasha nagachithnya samantha",email:"nat50@marvel.com",gender:"female",status:"active"},
		{id:"2031",name:"Rosemary Medhurst",email:"Wilton_Schmidt2@example.com",gender:"male",status:"active"},
		{id:"2032",name:"Charles Towne",email:"Jillian30@example.net",gender:"male",status:"active"},
		{id:"2035",name:"Ben Renner",email:"Winona.Senger@example.net","gender":"female","status":"active"},
		{id:"2036",name:"Aubrey Botsford PhD",email:"Favian_Watsica0@gmail.com",gender:"male",status:"inactive"},
		{id:"2041",name:"Grady Stoltenberg",email:"elwood.vandervort@gmail.com",gender:"female",status:"active"},
		{id:"2284",name:"Gomzyakov_E","email":"gamzik+634531876471638476@mail.ru","gender":"male","status":"active"},
		{id:"2291",name:"PANUGANTI HARSHITHA",email:"harshithapanuganti2@gmail.com",gender:"female",status:"inactive"},
		{id:"2293",name:"PANUGANTI HARSHITHA",email:"harshithapanuganti3@gmail.com",gender:"female",status:"inactive"},
		{id:"2053",name:"Anthony Abbott MD",email:"colene.gorczany@yahoo.com",gender:"female",status:"active"},
		{id:"2054",name:"mlz",email:"mariusz2323lazor@gmail.com",gender:"male",status:"active"},
		{id:"2055",name:"Vernell Beahan",email:"reuben.waters@yahoo.com",gender:"female",status:"active"},
		{id:"2057",name:"Ms. Leroy Greenfelder",email:"chi.stoltenberg@gmail.com",gender:"female",status:"active"},
		{id:"2058",name:"Madison Leannon",email:"dean.keebler@hotmail.com",gender:"female",status:"active"}
		];
	}
	handleRequest()
	fmt.Println("Choose your action(Create, Update ,Delete):")
	var option string
	fmt.Scanln(&option)

	if &option == Create || CREATE || create{

		handleCreate()
	}
	else if &option == Delete || DELETE || delete{
		handleDelete()
	}
	else if &option == Update || UPDATE || update {
		handleUpdate()
		
	}
	else{
		os.Exit()
	}
	
}
	
	
