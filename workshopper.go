package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"gopkg.in/russross/blackfriday.v2"
)

type example struct {
	Name string
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		index, err := Asset("data/index.tmpl")
		if err != nil {
			panic(err)
		}
		fmt.Fprint(w, string(index))
	})

	http.HandleFunc("/exercises", func(w http.ResponseWriter, r *http.Request) {
		err := r.ParseForm()
		if err != nil {
			panic(err)
		}
		number, err := strconv.Atoi(r.PostForm.Get("number"))
		if err != nil {
			panic(err)
		}
		exerciseTemplateName := fmt.Sprintf("data/exercises/%d/exercise.go.tmpl", number)
		exerciseTestTemplateName := fmt.Sprintf("data/exercises/%d/exercise_test.go.tmpl", number)
		exercise, err := Asset(exerciseTemplateName)
		if err != nil {
			panic(err)
		}
		exerciseTest, err := Asset(exerciseTestTemplateName)
		if err != nil {
			panic(err)
		}

		exerciseFileName := fmt.Sprintf("data/exercise%d.go", number)
		exerciseTestFileName := fmt.Sprintf("data/exercise%d_test.go", number)

		ioutil.WriteFile(exerciseFileName, exercise, 0644)
		ioutil.WriteFile(exerciseTestFileName, exerciseTest, 0644)

		http.Redirect(w, r, "/instructions", http.StatusSeeOther)
	})

	http.HandleFunc("/instructions", func(w http.ResponseWriter, r *http.Request) {
		instructionsTemplate := fmt.Sprintf("data/exercises/%d/instructions.md.tmpl", 1)
		instructions, err := Asset(instructionsTemplate)
		if err != nil {
			panic(err)
		}
		output := blackfriday.Run(instructions)
		fmt.Fprint(w, string(output))
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
