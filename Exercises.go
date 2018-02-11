package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
)

// Exercises handles GET "/exercises"
func Exercises(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		panic(err)
	}
	postedExerciseNumber := r.PostFormValue("exercise")
	number, err := strconv.Atoi(postedExerciseNumber)
	if err != nil {
		panic(err)
	}
	exerciseTemplateName := fmt.Sprintf("data/exercises/%d/exercise.go.tmpl", number)
	exerciseTestTemplateName := fmt.Sprintf("data/exercises/%d/exercise_test.go.tmpl", number)
	exerciseTemplate, err := Asset(exerciseTemplateName)
	if err != nil {
		panic(err)
	}
	exerciseTestTemplate, err := Asset(exerciseTestTemplateName)
	if err != nil {
		panic(err)
	}

	exerciseFileName := fmt.Sprintf("exercise%d.go", number)
	exerciseTestFileName := fmt.Sprintf("exercise%d_test.go", number)

	ioutil.WriteFile(exerciseFileName, exerciseTemplate, 0644)
	ioutil.WriteFile(exerciseTestFileName, exerciseTestTemplate, 0644)

	v := url.Values{}
	v.Add("exercise", postedExerciseNumber)
	redirectURL := "/instructions?" + v.Encode()
	http.Redirect(w, r, redirectURL, http.StatusSeeOther)
}
