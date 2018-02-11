package main

import (
	"fmt"
	"net/http"
	"strconv"

	blackfriday "gopkg.in/russross/blackfriday.v2"
)

func isExerciseNumberValid(exerciseNumber string) bool {
	exercisesDir, err := AssetDir("data/exercises")
	if err != nil || exerciseNumber == "" {
		return false
	}
	for _, item := range exercisesDir {
		if item == exerciseNumber {
			return true
		}
	}

	return false
}

// Instructions handles GET "/instructions"
func Instructions(w http.ResponseWriter, r *http.Request) {
	queryExercise := r.URL.Query().Get("exercise")
	exercise, err := strconv.Atoi(queryExercise)
	if err != nil || !isExerciseNumberValid(queryExercise) {
		http.Redirect(w, r, "/error", http.StatusSeeOther)
		return
	}

	instructionsTemplate := fmt.Sprintf("data/exercises/%d/instructions.md.tmpl", exercise)
	instructions, err := Asset(instructionsTemplate)
	if err != nil {
		panic(err)
	}
	output := blackfriday.Run(instructions)
	fmt.Fprint(w, string(output))
}
