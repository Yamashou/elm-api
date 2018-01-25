package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/context"
)

const MyContextKey = 1

func (q *Query) CreatExperiment(w http.ResponseWriter, r *http.Request) {
	contextVal := context.Get(r, MyContextKey)
	data := contextVal.(Learn)

	q.SaveExperiment(data)

	m := Message{
		"200",
		"OK",
		data,
	}
	if err := json.NewEncoder(w).Encode(m); err != nil {
		panic(err)
	}
}

func (q *Query) AddFeatureData(w http.ResponseWriter, r *http.Request) {
	contextVal := context.Get(r, MyContextKey)
	data := contextVal.(Learn)

	q.SaveData(data)

	m := Message{
		"200",
		"OK",
		data,
	}
	if err := json.NewEncoder(w).Encode(m); err != nil {
		panic(err)
	}
}

func (q *Query) Learning(w http.ResponseWriter, r *http.Request) {
	contextVal := context.Get(r, MyContextKey)
	data := contextVal.(Learn)

	d := q.GetDatasFromExperimentModelName(data.ModelName)
	ex := q.GetExperimentFromModelName(data.ModelName)
	_, err := ex.Training(d).MarshalBinaryTo(data.ModelName)
	if err != nil {
		panic(err)
	}

	m := Message{
		"200",
		"OK",
		data,
	}

	if err := json.NewEncoder(w).Encode(m); err != nil {
		panic(err)
	}
}

func (q *Query) Fetch(w http.ResponseWriter, r *http.Request) {
	contextVal := context.Get(r, MyContextKey)
	data := contextVal.(Learn)
	ex := q.GetExperimentFromModelName(data.ModelName)
	ans, err := ex.Result(data.Feature[0])
	m := Message{
		"200",
		"OK",
		Ans{ans},
	}
	if err != nil {
		m.Status = "500"
		m.Status = "Model Not Found"
	}

	if err := json.NewEncoder(w).Encode(m); err != nil {
		panic(err)
	}
}

func UseContextPost(handler http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		w.Header().Set("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, Authorization")
		w.Header().Set("Access-Control-Allow-Methods", "POST")

		b, err := ioutil.ReadAll(r.Body)
		if err != nil {
			panic(err)
		}

		data, err := UnmarshalLearn(b)
		if err != nil {
			panic(err)
		}

		context.Set(r, MyContextKey, data)
		handler.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}

func UseContextGet(handler http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		w.Header().Set("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, Authorization")
		w.Header().Set("Access-Control-Allow-Methods", "GET")

		handler.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}

func (q *Query) Index(w http.ResponseWriter, r *http.Request) {
	es := q.GetAllExperiment()
	ejs := make([]ExperimentJson, len(es))
	for i, v := range es {
		ejs[i] = v.NewExperimentJson()
	}
	m := Message{
		"200",
		"OK",
		ejs,
	}
	if err := json.NewEncoder(w).Encode(m); err != nil {
		panic(err)
	}
}

func NotFoundHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Not Found\n"))
}
