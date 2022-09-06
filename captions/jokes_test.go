package captions

import (
	"testing"
	"github.com/jarcoal/httpmock"
)

func TestGetJokeSimpleJoke(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("GET", "https://v2.jokeapi.dev/joke/Programming", httpmock.NewJsonResponderOrPanic(200, httpmock.File("./fixtures/joke.json")))
	text1, text2 := GetJoke("https://v2.jokeapi.dev/joke/Programming")
	expected2 := "I've got a really good UDP joke to tell you but I don’t know if you'll get it."

	if text1 != "" {
		t.Errorf("First text is not empty: %s", text1)
	}

	if text2 != expected2 {
		t.Fatalf("Second text is wrong: %s", text2)
		t.Errorf("Second text should be: %s", expected2)
	}
}

func TestGetJokeSetupDeliveryJoke(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("GET", "https://v2.jokeapi.dev/joke/Programming", httpmock.NewJsonResponderOrPanic(200, httpmock.File("./fixtures/setup_delivery.json")))
	text1, text2 := GetJoke("https://v2.jokeapi.dev/joke/Programming")
	expected1 := "Why did the Python data scientist get arrested at customs?"
	expected2 := "She was caught trying to import pandas!"

	if text1 != expected1 {
		t.Errorf("First text is wrong: %s", text1)
		t.Errorf("First text should be: %s", expected1)
	}

	if text2 != expected2 {
		t.Fatalf("Second text is wrong: %s", text2)
		t.Errorf("Second text should be: %s", expected2)
	}
}

// TODO Test other scenarios, e.g. timeout

func TestGetJokeAPIError(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("GET", "https://v2.jokeapi.dev/joke/Programming", httpmock.NewStringResponder(500, "error"))
	text1, text2 := GetJoke("https://v2.jokeapi.dev/joke/Programming")
	expected1 := "How do you know God is a shitty programmer?"
	expected2 := "He wrote the OS for an entire universe, but didn't leave a single useful comment."

	if text1 != expected1 {
		t.Errorf("First text is wrong: %s", text1)
		t.Errorf("First text should be: %s", expected1)
	}

	if text2 != expected2 {
		t.Fatalf("Second text is wrong: %s", text2)
		t.Errorf("Second text should be: %s", expected2)
	}
}