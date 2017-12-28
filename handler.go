package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"

	"github.com/julienschmidt/httprouter"
)

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	fmt.Fprintf(w, "{\"greeting\":\"Welcome!\"}")
}

func AnnotationIndex(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.WriteHeader(http.StatusOK)

	/* EndTimeを過ぎたアノテーションはキャッシュから消す */
	var _annotations Annotations
	for i := 0; i < len(annotations); i++ {
		t, _ := time.Parse("2006-01-02 15:04:05 -0700 MST", annotations[i].EndTime)
		now := time.Now()
		if t.Sub(now) > 0 {
			_annotations = append(_annotations, annotations[i])
		}
	}
	annotations = _annotations

	if err := json.NewEncoder(w).Encode(annotations); err != nil {
		panic(err)
	}
}

func AnnotationShow(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id, _ := strconv.Atoi(ps.ByName("annotationId"))
	t := RepoFindAnnotation(id)
	if t.ID == 0 && t.Title == "" {
		w.WriteHeader(http.StatusNotFound)
		return
	}
}

func AnnotationCreate(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var annotation Annotation

	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))

	if err != nil {
		panic(err)
	}
	defer r.Body.Close()

	if err := json.Unmarshal(body, &annotation); err != nil {
		w.WriteHeader(500)
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
		return
	}

	t := RepoCreateAnnotation(annotation)
	location := fmt.Sprintf("http://%s/%d", r.Host, t.ID)
	w.Header().Set("Location", location)
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(t); err != nil {
		panic(err)
	}
}

func AnnotationDelete(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id, _ := strconv.Atoi(ps.ByName("annotationId"))
	if err := RepoDestroyAnnotation(id); err != nil {
		w.WriteHeader(http.StatusNotFound)
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
		return
	}

	w.Header().Del("Content-Type")
	w.WriteHeader(204) // 204 No Content
}

func TransAnnotationIndex(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.WriteHeader(http.StatusOK)

	transAnnotations := getTransAnnotations()
	if err := json.NewEncoder(w).Encode(transAnnotations); err != nil {
		panic(err)
	}
}

func TransAnnotationShow(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id, _ := strconv.Atoi(ps.ByName("annotationId"))
	t := getTransAnnotation(id)
	if t.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	if err := json.NewEncoder(w).Encode(t); err != nil {
		panic(err)
	}
}

func TransAnnotationUpdate(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var transAnnotation TransAnnotation

	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	id, _ := strconv.Atoi(ps.ByName("annotationId"))
	if err != nil {
		panic(err)
	}
	defer r.Body.Close()
	transAnnotation.ID = id
	if err := json.Unmarshal(body, &transAnnotation); err != nil {
		w.WriteHeader(500)
		// fmt.Println("CHECK", transAnnotation)
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
		return
	}
	type ResSchema struct {
		TagIds []int `json:"tagIds"`
	}
	tagIds := updateTransAnnotation(id, transAnnotation)
	resBody := ResSchema{tagIds}
	location := fmt.Sprintf("http://%s/%d", r.Host, id)
	w.Header().Set("Location", location)
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(resBody); err != nil {
		panic(err)
	}
}

func TagIndex(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.WriteHeader(http.StatusOK)

	tags := getTags()
	if err := json.NewEncoder(w).Encode(tags); err != nil {
		panic(err)
	}
}

func TagCreate(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var tag Tag

	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))

	if err != nil {
		panic(err)
	}
	defer r.Body.Close()

	if err := json.Unmarshal(body, &tag); err != nil {
		w.WriteHeader(500)
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
		return
	}

	createdId := insertTag(tag)
	tag.ID = createdId
	location := fmt.Sprintf("http://%s/%d", r.Host, createdId)
	w.Header().Set("Location", location)
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(tag); err != nil {
		panic(err)
	}
}

func TagDelete(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id, _ := strconv.Atoi(ps.ByName("tagId"))

	deleteTag(id)

	w.Header().Del("Content-Type")
	w.WriteHeader(204) // 204 No Content
}

func IncrementNiceNum(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id, _ := strconv.Atoi(ps.ByName("annotationId"))

	transAnnotation := incrementNiceNum(id)
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(transAnnotation); err != nil {
		panic(err)
	}
}

func DecrementNiceNum(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id, _ := strconv.Atoi(ps.ByName("annotationId"))

	transAnnotation := decrementNiceNum(id)
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(transAnnotation); err != nil {
		panic(err)
	}
}

func Ranking(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.WriteHeader(http.StatusOK)

	/* EndTimeを過ぎたアノテーションはキャッシュから消す */
	var _fullAnnotations FullAnnotations

	_fullAnnotations = getRanking()

	if err := json.NewEncoder(w).Encode(_fullAnnotations); err != nil {
		panic(err)
	}
}
