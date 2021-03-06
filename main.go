package main

import (
	"github.com/zakuro9715/aojgo"
	"github.com/zenazn/goji"
	"html/template"
	"net/http"
	"sort"
)

var t = template.Must(template.ParseFiles("./templates/dashboard.html"))

type Problem struct {
	Id, Name string
	Level    int
	Solved   bool
}

type ProblemList []Problem

func (p ProblemList) Len() int {
	return len(p)
}

func (p ProblemList) Less(i, j int) bool {
	return p[i].Level < p[j].Level
}

func (p ProblemList) Swap(i, j int) {
	tmp := p[i]
	p[i] = p[j]
	p[j] = tmp
}

func newProblemFromAoj(p aoj.Problem) *Problem {
	np := &Problem{}
	np.Id = p.Id
	np.Name = p.Name
	np.Level = getLevel(p.Id)
	return np
}

func checkIsSolved(userId, problemId string) bool {
	r, _ := aoj.SolvedRecordSearchApi(userId, problemId, "", -1, -1)
	return len(r) > 0
}

func showDashboard(w http.ResponseWriter, r *http.Request) {
	rawProblems5, _ := aoj.ProblemListSearchApi(5)
	rawProblems6, _ := aoj.ProblemListSearchApi(6)

	problems := make([]Problem, len(rawProblems5)+len(rawProblems6))
	for i, p := range rawProblems5 {
		problems[i] = *newProblemFromAoj(p)
	}
	for i, p := range rawProblems6 {
		problems[len(rawProblems5)+i] = *newProblemFromAoj(p)
	}
	sort.Sort(ProblemList(problems))

	userId := r.URL.Query().Get("userId")
	if len(userId) > 0 {
		for i, p := range problems {
			problems[i].Solved = checkIsSolved(userId, p.Id)
		}
	}

	data := struct {
		Problems []Problem
		UserId   string
	}{
		problems,
		userId,
	}
	t.Execute(w, data)
}

func main() {
	goji.Get("/static/*", http.FileServer(http.Dir(".")))
	goji.Get("/dashboard", showDashboard)
	goji.Get("/", http.RedirectHandler("/dashboard", 302))
	goji.Serve()
}
