type user struct {
	Name string `json:"name"`
	Age  string `json:"age"`
	City string `json:"city"`
}

func main() {

	port := ":3000"

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// fmt.Fprintf(w, "Main Page")
		w.Write([]byte("Main Page"))
	})

	http.HandleFunc("/teachers", func(w http.ResponseWriter, r *http.Request) {
		// fmt.Fprintf(w, "Main Page")

		switch r.Method {
		case http.MethodGet:
			w.Write([]byte("GET Method on Teachers Route Page"))
		case http.MethodPost:

			// Parse form data (necessary for x-www-form-urlencoded)

			// err := r.ParseForm()
			// if err != nil {
			// 	http.Error(w, "Error parsing form", http.StatusBadRequest)
			// 	return
			// }
			// fmt.Println("Form:", r.Form)

			// // Prepare response data
			// response := make(map[string]any)
			// for key, value := range r.Form {
			// 	response[key] = value[0]
			// }
			// fmt.Println("Processed Response Map:", response)

			// name := r.FormValue("name")
			// fmt.Println("Name: ", name)

			// RAW Body

			body, err := io.ReadAll(r.Body)
			if err != nil {
				return
			}
			defer r.Body.Close()

			fmt.Println("RAW Body:", body)
			fmt.Println("RAW Body:", string(body))

			// If we expect json data, then unmarshal it.
			var user user
			err = json.Unmarshal(body, &user)
			if err != nil {
				return
			}
			fmt.Println("Unmarshaled JSON into an instance of user struct: ", user)
			fmt.Println("Recieved user name as:", user.Name)

			// // Prepare response data
			response := make(map[string]any)
			for key, value := range r.Form {
				response[key] = value[0]
			}

			err = json.Unmarshal(body, &response)
			if err != nil {
				return
			}
			fmt.Println("Unmarshaled JSON into a map: ", response)

			w.Write([]byte("POST Method on Teachers Route Page"))
		case http.MethodPut:
			w.Write([]byte("PUT Method on Teachers Route Page"))
		case http.MethodDelete:
			w.Write([]byte("DELETE Method on Teachers Route Page"))
		default:
			w.Write([]byte("Teachers Route Page"))
		}

	})

	// http.HandleFunc("/students", func(w http.ResponseWriter, r *http.Request) {
	// 	// fmt.Fprintf(w, "Main Page")
	// 	w.Write([]byte("Students Page"))
	// })

	// http.HandleFunc("/execs", func(w http.ResponseWriter, r *http.Request) {
	// 	// fmt.Fprintf(w, "Main Page")
	// 	w.Write([]byte("Execs Page"))
	// })

	fmt.Println("Server is running on port:", port)
	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatalln("Error starting the server", err)
	}

	*** Path Params ***
	teachers/{id} -> teachers/9
	fmt.Println(r.URL.Path)
	path := strings.TrimPrefix(r.URL.Path, "/teachers/")
	fmt.Println(path)
	userID := strings.TrimPrefix(path, "/")
	fmt.Println("The User ID:", userID)

	*** Query Params ***
	teachers/?key=value -> teachers/?name=Jhon&age=25
	fmt.Println(r.URL.Query())
	queryParams := r.URL.Query()
	name := queryParams.Get("name")
	fmt.Println(name)
}



	*** apply update ***

	for field, value := range updatedTeacher {

	 	strValue, ok := value.(string)
	 	if !ok {
	 		http.Error(w, "Invalid data type: expected text", http.StatusBadRequest)
	 		return
	 	}

	 	switch field {
	 	case "first_name":
	 		existingTeacher.FirstName = strValue
	 	case "last_name":
	 		existingTeacher.LastName = strValue
	 	case "email":
	 		existingTeacher.Email = strValue
	 	case "class":
	 		existingTeacher.Class = strValue
	 	case "subject":
	 		existingTeacher.Subject = strValue
	 	}
	}




	*** apply updates using reflect ***

	teacherValue := reflect.ValueOf(&existingTeacher).Elem()  --> {100 Alice Brown alice@example.com 6A World History}
	teacherType := teacherValue.Type()                        --> model.Teacher
	teacherValue.Field(0)  --> 100
	teacherType.Field(0))  --> {ID  int json:"id,omitempty" 0 [0] false}
	teacherType.NumField()  --> 6
	field.Tag.Get("json") --> teacher's struct  example: "first_name,omitempty"
	teacherField := teacherValue.Field(i) --> alice@example.com (old)
	teacherField := teacherValue.Field(i).Type() --> string
	reflect.ValueOf(value) --> aliceSmith@example.com (new)