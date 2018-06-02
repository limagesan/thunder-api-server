package main

import (
	"encoding/json"
	"errors"
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

func returnMonth(month int) (time.Month, error) {
	switch month {
	case 1:
		return time.January, nil
	case 2:
		return time.February, nil
	case 3:
		return time.March, nil
	case 4:
		return time.April, nil
	case 5:
		return time.May, nil
	case 6:
		return time.June, nil
	case 7:
		return time.July, nil
	case 8:
		return time.August, nil
	case 9:
		return time.September, nil
	case 10:
		return time.October, nil
	case 11:
		return time.November, nil
	case 12:
		return time.December, nil
	default:
		return time.January, errors.New("Unexpected Number.")
	}
}

func AnnotationIndex(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	w.WriteHeader(http.StatusOK)

	headTime := time.Now()
	tailTime := headTime.Add(30 * 24 * time.Hour)

	annotations := getAnnotations(headTime, tailTime)

	if err := json.NewEncoder(w).Encode(annotations); err != nil {
		panic(err)
	}
}

func AnnotationShow(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id, _ := strconv.Atoi(ps.ByName("annotationId"))
	t := getAnnotation(id)
	if t.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	if err := json.NewEncoder(w).Encode(t); err != nil {
		panic(err)
	}
}

func SelectAnnotationIndex(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.WriteHeader(http.StatusOK)
	year, _ := strconv.Atoi(ps.ByName("year"))
	monthInt, _ := strconv.Atoi(ps.ByName("month"))
	month, _ := returnMonth(monthInt)
	day, _ := strconv.Atoi(ps.ByName("day"))
	hour, _ := strconv.Atoi(ps.ByName("hour"))
	min, _ := strconv.Atoi(ps.ByName("min"))

	headTime := time.Date(year, month, day, hour, min, 0, 0, time.UTC)
	tailTime := headTime.Add(12 * time.Hour)

	annotations := getAnnotations(headTime, tailTime)

	if err := json.NewEncoder(w).Encode(annotations); err != nil {
		panic(err)
	}
}

func SelectFeaturedAnnotationIndex(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.WriteHeader(http.StatusOK)
	areaId, _ := strconv.Atoi(ps.ByName("areaId"))
	year, _ := strconv.Atoi(ps.ByName("year"))
	monthInt, _ := strconv.Atoi(ps.ByName("month"))
	month, _ := returnMonth(monthInt)
	day, _ := strconv.Atoi(ps.ByName("day"))

	headTime := time.Date(year, month, day, 0, 0, 0, 0, time.UTC)
	tailTime := headTime.Add(24 * time.Hour)

	annotations := getFeaturedAnnotations(areaId, headTime, tailTime)

	if err := json.NewEncoder(w).Encode(annotations); err != nil {
		panic(err)
	}
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

func UpdateFeatured(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id, err := strconv.Atoi(ps.ByName("annotationId"))
	featured, err2 := strconv.ParseBool(ps.ByName("featured"))
	if err != nil {
		w.WriteHeader(400)
		return
	}
	if err2 != nil {
		w.WriteHeader(400)
		return
	}
	fmt.Println("Host", r.Host, "remoteAddr", r.RemoteAddr)
	updateFeatured(id, featured)
	w.WriteHeader(http.StatusAccepted)
}

func IncrementNiceNum(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id, _ := strconv.Atoi(ps.ByName("annotationId"))
	fmt.Println("Host", r.Host, "remoteAddr", r.RemoteAddr)
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
