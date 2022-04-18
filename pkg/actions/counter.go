package actions

import (
	"fmt"
	"net/http"

	"github.com/frankmeza/todo_list_go_htmx/pkg/utils"
)

type Counter struct {
	value int
}

var counter Counter

// purely presentation
func getCounterComponent() string {
	return "<p>%d</p>"
}

// interpolates data into component,
// i.e. passes in component props
func passDataIntoComponent(componentHtml string, value int) string {
	return fmt.Sprintf(componentHtml, value)
}

func updateAppState(kind string) {
	if kind == "ADD" {
		counter.value++
	} else if kind == "MINUS" {
		counter.value--
	} else if kind == "RESET" {
		counter.value = 0
	}
}

// acts as a reducer for causing side effects,
// generates an instance of the updated component,
// sends component as html string as http response
func renderUpdatedComponent(w http.ResponseWriter, kind string) {
	updateAppState(kind)

	newValue := passDataIntoComponent(
		getCounterComponent(),
		counter.value,
	)

	utils.RenderHtmlString(w, http.StatusOK, newValue)
}

func Increment(w http.ResponseWriter, r *http.Request) {
	renderUpdatedComponent(w, "ADD")
}

func Decrement(w http.ResponseWriter, r *http.Request) {
	renderUpdatedComponent(w, "MINUS")
}

func ResetCounter(w http.ResponseWriter, r *http.Request) {
	renderUpdatedComponent(w, "RESET")
}
