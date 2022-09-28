package cmd

import "testing"

func TestCommand_RegisterDriver(t *testing.T) {
	ucase := businessCreator(t)
	command := NewCommand(ucase)

	output, err := command.RegisterDriver("driver-003")
	if err != nil {
		t.Fatalf("expected to be success register driver, got %v", err)
	}

	expOutput := "Driver driver-003 registered"

	if output != expOutput {
		t.Fatalf("expected output `%s`, got `%s`", expOutput, output)
	}
}

func TestCommand_RegisterDriverInvalid(t *testing.T) {
	ucase := businessCreator(t)
	command := NewCommand(ucase)

	if _, err := command.RegisterDriver("driver-001"); err == nil {
		t.Fatal("expected to throw error, got nil")
	}
}
