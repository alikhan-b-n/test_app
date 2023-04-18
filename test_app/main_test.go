package main

import (
	"bufio"
	"bytes"
	"io"
	"os"
	"strings"
	"testing"
)

func Test_isPrime(t *testing.T) {
	primeTests := []struct {
		name     string
		testNum  int
		expected bool
		msg      string
	}{
		{"prime", 7, true, "7 is a prime number!"},
		{"not prime", 8, false, "8 is not a prime number because it is divisible by 2!"},
		{"zero", 0, false, "0 is not prime, by definition!"},
		{"one", 1, false, "1 is not prime, by definition!"},
		{"negative number", -11, false, "Negative numbers are not prime, by definition!"},
	}

	for _, e := range primeTests {
		result, msg := isPrime(e.testNum)
		if e.expected && !result {
			t.Errorf("%s: expected true but got false", e.name)
		}

		if !e.expected && result {
			t.Errorf("%s: expected false but got true", e.name)
		}

		if e.msg != msg {
			t.Errorf("%s: expected %s but got %s", e.name, e.msg, msg)
		}
	}
}

func TestCheckNumbers(t *testing.T) {
	numbersCheckTests := []struct {
		name     string
		value    string
		expected bool
		msg      string
	}{
		{"integer", "7", false, "7 is a prime number!"},
		{"string", "qwe", false, "Please enter a whole number!"},
		{"float", "7.5", false, "Please enter a whole number!"},
		{"quit", "q", true, ""},
	}
	for _, e := range numbersCheckTests {
		values := strings.NewReader(e.value)
		scanner := bufio.NewScanner(values)
		msg, actualValue := checkNumbers(scanner)
		if e.expected && !actualValue {
			t.Errorf("%s: expected true but got false", e.name)
		}

		if !e.expected && actualValue {
			t.Errorf("%s: expected false but got true", e.name)
		}

		if e.msg != msg {
			t.Errorf("%s: expected %s but got %s", e.name, msg, e.msg)
		}

	}
}

func TestPrompt(t *testing.T) {
	expectedRes := "-> "

	oldStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	prompt()

	_ = w.Close()
	os.Stdout = oldStdout

	result, _ := io.ReadAll(r)
	if string(result) != "-> " {
		t.Errorf("expected %s but got %s", expectedRes, string(result))
	}

}

func TestIntro(t *testing.T) {
	oldOut := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	intro()

	_ = w.Close()

	os.Stdout = oldOut

	out, _ := io.ReadAll(r)

	if strings.Contains(string(out), "Enter whole number") {
		t.Errorf("Incorrect intro text, got %s", string(out))
	}
}

func TestReadUserInput(t *testing.T) {
	doneChan := make(chan bool)

	var stdin bytes.Buffer

	stdin.Write([]byte("1\nq\n"))

	go readUserInput(&stdin, doneChan)
	<-doneChan
	close(doneChan)
}
