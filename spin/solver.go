package main

import (
	"fmt"
	"io"
	"net/http"

	spinhttp "github.com/fermyon/spin/sdk/go/http"
	nytimes "github.com/timgluz/nydigits/pkg"
)

func init() {
	spinhttp.Handle(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain")

		body, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		switch r.Method {
		case http.MethodGet:
			w.WriteHeader(http.StatusOK)
			fmt.Fprintln(w, "Solver of the NY Times Digits")
			fmt.Fprintln(w, "Send a POST request with target and digits as JSON body")
			return
		case http.MethodPost:
			solution, err := solve(body)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			w.WriteHeader(http.StatusOK)
			w.Header().Set("Content-Type", "application/json")
			w.Write(solution)
			return
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}
	})
}

func main() {}

// dont forget to generate easyjson code whenever you change the structs
func parseProblem(body []byte) (nytimes.Problem, error) {
	var problem nytimes.Problem
	if err := problem.UnmarshalJSON(body); err != nil {
		return nytimes.Problem{}, err
	}

	return problem, nil
}

func solve(body []byte) ([]byte, error) {
	problem, err := parseProblem(body)
	if err != nil {
		return nil, err
	}

	if err := problem.Validate(); err != nil {
		return nil, err
	}

	solution, err := nytimes.SolveProblem(problem)
	if err != nil {
		return nil, err
	}

	return solution.MarshalJSON()
}
