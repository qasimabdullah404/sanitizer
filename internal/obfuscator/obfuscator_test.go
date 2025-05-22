package obfuscator_test

import (
	"testing"

	"gihtub.com/qasimabdullah404/sanitizer/internal/obfuscator"
)

func TestObfuscateLine_Password(t *testing.T) {
	input := "User=admin Password=supersecret123"
	expected := "User=admin password=***REDACTED***"
	output := obfuscator.ObfuscateLine(input)

	if output != expected {
		t.Errorf("Expected '%s', got '%s'", expected, output)
	}
}

func TestObfuscateLine_Token(t *testing.T) {
	input := "Bearer abc123TOKENVALUE"
	expected := "Bearer ***REDACTED***"
	output := obfuscator.ObfuscateLine(input)

	if output != expected {
		t.Errorf("Expected '%s', got '%s'", expected, output)
	}
}

func TestObfuscateLine_NoSecrets(t *testing.T) {
	input := "INFO: service started on port 8080"
	expected := input
	output := obfuscator.ObfuscateLine(input)

	if output != expected {
		t.Errorf("Expected '%s', got '%s'", expected, output)
	}
}

func TestObfuscateLine_MultipleSecrets(t *testing.T) {
	input := "API_KEY=abcd1234 SECRET=topsecret"
	expected := "API_KEY=***REDACTED*** SECRET=***REDACTED***"
	output := obfuscator.ObfuscateLine(input)

	if output != expected {
		t.Errorf("Expected '%s', got '%s'", expected, output)
	}
}
