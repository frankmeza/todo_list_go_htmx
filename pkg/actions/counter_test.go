package actions

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/frankmeza/todo_list_go_htmx/pkg/utils"
)

// var counter Counter
var testUtils utils.TestUtils

func resetCounter() {
	counter.value = 0
}

func TestIncrement(t *testing.T) {
	resetCounter()

	request, err := http.NewRequest("GET", "/increment_counter", nil)
	if err != nil {
		t.Fatal(err)
	}

	recorder := httptest.NewRecorder()
	handler := http.NewServeMux()

	handler.ServeHTTP(recorder, request)

	testUtils.CheckForExpectedResponseCode(t, recorder.Code, http.StatusOK)

	expected := `<p>1</p>`
	testUtils.CheckBodyForError(t, recorder.Body.String(), expected)
}

func TestDecrement(t *testing.T) {
	counter.value = 1

	request, err := http.NewRequest("GET", "/decrement_counter", nil)
	if err != nil {
		t.Fatal(err)
	}

	recorder := httptest.NewRecorder()
	handler := http.NewServeMux()

	handler.ServeHTTP(recorder, request)

	testUtils.CheckForExpectedResponseCode(t, recorder.Code, http.StatusOK)

	expected := `<p>0</p>`
	testUtils.CheckBodyForError(t, recorder.Body.String(), expected)
}

func TestReset(t *testing.T) {
	counter.value = 39

	request, err := http.NewRequest("GET", "/reset_counter", nil)
	if err != nil {
		t.Fatal(err)
	}

	recorder := httptest.NewRecorder()
	handler := http.NewServeMux()

	handler.ServeHTTP(recorder, request)

	testUtils.CheckForExpectedResponseCode(t, recorder.Code, http.StatusOK)

	expected := `<p>0</p>`
	testUtils.CheckBodyForError(t, recorder.Body.String(), expected)
}
