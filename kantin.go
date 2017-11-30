package main

import ( 
 "fmt"
 "log"
 "net/http"
 "encoding/json"
 "database/sql"
 _ "github.com/go-sql-driver/mysql")
 


type MyMenu struct {
	Makanan string 
	Harga int
	Kantin string 
}

func main() {
	 
	port := 8181
	http.HandleFunc("/kantinitb/", func(w http.ResponseWriter, r *http.Request){
		switch r.Method {
		case "GET" :
			s := r.URL.Path[len("/kantinitb/"):]
			if s != "" {
				GetFood(w,r,s)
			} else {
				GetAllFood(w,r)
			}
		case "POST" :
		
		case "DELETE" :
			default : http.Error(w, "Invalid method", 405)
		}
	})
	log.Printf("Server starting on port %v\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v",port),nil))
}


func GetAllFood(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/kantin")
	
	if err!= nil {
		log.Fatal(err)
	}
	defer db.Close()

	mymenu := MyMenu{};

	rows, err := db.Query("SELECT Makanan,Harga,Kantin from foodservice")
	if err!= nil {
		log.Fatal(err)
	}
	defer db.Close()

	for rows.Next() {
		err := rows.Scan(&mymenu.Makanan, &mymenu.Harga, &mymenu.Kantin)
		if err!= nil {
			log.Fatal(err)
		}
		
		json.NewEncoder(w).Encode(&mymenu);
	}
	err = rows.Err()
}

func GetFood(w http.ResponseWriter, r *http.Request, Makanan string) {
	//idfood, _ := strconv.Atoi(id)

	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/kantin")
	
	if err!= nil {
		log.Fatal(err)
	}
	defer db.Close()

	mymenu := MyMenu{};

	rows, err := db.Query("select * from foodservice WHERE Makanan = ?", Makanan)
	
	if err!= nil {
		log.Fatal(err)
	}

	defer db.Close()

	for rows.Next() {
		err := rows.Scan(&mymenu.Makanan, &mymenu.Harga, &mymenu.Kantin)
		if err!= nil {
			log.Fatal(err)
		}
		
		json.NewEncoder(w).Encode(&mymenu);
	}
	err = rows.Err()
}
