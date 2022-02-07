package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"reflect"
	"strings"
	"testing"

	"github.com/go-playground/assert/v2"
)

var (
	input    string          = "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Nulla gravida enim rutrum ligula gravida, id eleifend ex pretium. Etiam vitae enim erat posuere, mollis enim vitae vitae, euismod lacus. Praesent commodo erat ac lectus placerat placerat. Pellentesque mi dolor, volutpat sed ligula id, sollicitudin hendrerit dui. In hac habitasse platea dictumst. Morbi commodo in mauris aliquam euismod. Donec varius eros sed porttitor tincidunt. Vivamus vitae neque sed arcu auctor viverra. Morbi vel lacus eu sem pharetra venenatis. Mauris tristique, velit eget luctus pulvinar, felis massa mattis orci, tempus consectetur turpis enim eu erat. Maecenas quis ante malesuada, pulvinar enim eget, tristique nulla. Nunc commodo, odio eget tristique fermentum, urna enim vehicula neque, a laoreet eros risus non neque"
	words                    = []string{"lorem", "ipsum", "dolor", "sit", "amet", "consectetur", "adipiscing", "elit", "nulla", "gravida", "enim", "rutrum", "ligula", "gravida", "id", "eleifend", "ex", "pretium", "etiam", "vitae", "enim", "erat", "posuere", "mollis", "enim", "vitae", "vitae", "euismod", "lacus", "praesent", "commodo", "erat", "ac", "lectus", "placerat", "placerat", "pellentesque", "mi", "dolor", "volutpat", "sed", "ligula", "id", "sollicitudin", "hendrerit", "dui", "in", "hac", "habitasse", "platea", "dictumst", "morbi", "commodo", "in", "mauris", "aliquam", "euismod", "donec", "varius", "eros", "sed", "porttitor", "tincidunt", "vivamus", "vitae", "neque", "sed", "arcu", "auctor", "viverra", "morbi", "vel", "lacus", "eu", "sem", "pharetra", "venenatis", "mauris", "tristique", "velit", "eget", "luctus", "pulvinar", "felis", "massa", "mattis", "orci", "tempus", "consectetur", "turpis", "enim", "eu", "erat", "maecenas", "quis", "ante", "malesuada", "pulvinar", "enim", "eget", "tristique", "nulla", "nunc", "commodo", "odio", "eget", "tristique", "fermentum", "urna", "enim", "vehicula", "neque", "a", "laoreet", "eros", "risus", "non", "neque"}
	response []WordOccurence = []WordOccurence{
		{
			Word:       "enim",
			Occurrence: 6,
		},
		{
			Word:       "vitae",
			Occurrence: 4,
		},
		{
			Word:       "commodo",
			Occurrence: 3,
		},
		{
			Word:       "tristique",
			Occurrence: 3,
		},
		{
			Word:       "sed",
			Occurrence: 3,
		},
		{
			Word:       "erat",
			Occurrence: 3,
		},
		{
			Word:       "eget",
			Occurrence: 3,
		},
		{
			Word:       "neque",
			Occurrence: 3,
		},
		{
			Word:       "gravida",
			Occurrence: 2,
		},
		{
			Word:       "consectetur",
			Occurrence: 2,
		},
	}
)

func TestSort(t *testing.T) {
	t.Run("Test setup router", func(t *testing.T) {
		r := setUpRouter()
		w := httptest.NewRecorder()
		request := fmt.Sprintf(`{"text": "%s"}`, input)
		req, _ := http.NewRequest("POST", "/occurrence", strings.NewReader(request))
		r.ServeHTTP(w, req)

		assert.Equal(t, 200, w.Code)
	})

	t.Run("Test getWordFrequency", func(t *testing.T) {
		got := getWordFrequency(words)
		want := response
		if !reflect.DeepEqual(got, want) {
			t.Errorf("getWordFrequency: want %v, got %v", want, got)
		}
	})
}
