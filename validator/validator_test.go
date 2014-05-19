//
// Tests for the validator.
//
// Copyright 2014 TLD dotHIV Registry GmbH.
// @author Markus Tacker <m@dotHIV.org>
//
package validator

import (
	assert "github.com/dothiv/translations-updater/testing"
	"testing"
)

func TestNotEmpty(t *testing.T) {
	v := NewCommitMessageValidator("")
	assert.False(t, v.IsValidCommitMessage(), "IsValidCommitMessage('')")
	assert.Equals(t, v.ValidateCommitMessage().Error(), "Empty commit message")
}

func TestValidFormat(t *testing.T) {
	assert.True(t, NewCommitMessageValidator("test(valid-commit-msg): implement tests").IsValidCommitMessage(), "IsValidCommitMessage")
}

func TestInvalidFormat(t *testing.T) {
	v := NewCommitMessageValidator("bla")
	assert.False(t, v.IsValidCommitMessage(), "IsValidCommitMessage('bla')")
	assert.Equals(t, v.ValidateCommitMessage().Error(), "First commit message line (header) does not follow format: type(scope): message")
}

func TestNoCapitalLetter(t *testing.T) {
	v := NewCommitMessageValidator("test(valid-commit-msg): Implement tests")
	assert.False(t, v.IsValidCommitMessage(), "IsValidCommitMessage('test(valid-commit-msg): Implement tests')")
	assert.Equals(t, v.ValidateCommitMessage().Error(), "Commit subject 'Implement tests' must not have a capital first letter!")
}

func TestNotDotAtEnd(t *testing.T) {
	v := NewCommitMessageValidator("test(valid-commit-msg): implement tests.")
	assert.False(t, v.IsValidCommitMessage(), "IsValidCommitMessage('test(valid-commit-msg): implement tests.')")
	assert.Equals(t, v.ValidateCommitMessage().Error(), "Commit subject 'implement tests.' must not end with a dot (.)!")
}

func TestValidType(t *testing.T) {
	v := NewCommitMessageValidator("foo(valid-commit-msg): implement tests")
	assert.False(t, v.IsValidCommitMessage(), "IsValidCommitMessage('foo(valid-commit-msg): implement tests')")
	assert.Equals(t, v.ValidateCommitMessage().Error(), "Commit type 'foo' not in valid ones: feat, fix, docs, style, refactor, test, chore!")
}

func TestEmptySecondLine(t *testing.T) {
	v := NewCommitMessageValidator("test(valid-commit-msg): implement tests\nLine 2\nLine 3")
	assert.False(t, v.IsValidCommitMessage(), "IsValidCommitMessage('test(valid-commit-msg): implement tests\nLine 2\nLine 3')")
	assert.Equals(t, v.ValidateCommitMessage().Error(), "Second commit message line must be empty!")
}

func TestLengthFirstLine(t *testing.T) {
	v := NewCommitMessageValidator("test(valid-commit-msg): the message of the first line should not be longer than 70 characters")
	assert.False(t, v.IsValidCommitMessage(), "IsValidCommitMessage('test(valid-commit-msg): the message of the first line should not be longer than 70 characters')")
	assert.Equals(t, v.ValidateCommitMessage().Error(), "First line should not be longer than 70 characters! It has 93 characters.")
}

func TestLengthOtherLines(t *testing.T) {
	v := NewCommitMessageValidator("test(valid-commit-msg): short first line\n\nAlso other lines should not be longer than 100 characters as rendering the message without soft-wraps makes them hard to read")
	assert.False(t, v.IsValidCommitMessage(), "IsValidCommitMessage('test(valid-commit-msg): short first line\n\nAlso other lines should not be longer than 100 characters as rendering the message without soft-wraps makes them hard to read')")
	assert.Equals(t, v.ValidateCommitMessage().Error(), "Other lines should not be longer than 100 characters! It has 125 characters.")
}

func TestIgnoreLongComments(t *testing.T) {
	v := NewCommitMessageValidator("test(valid-commit-msg): short first line\n\n# This is a long comment that should be ignored. This is a long comment that should be ignored. This is a long comment that should be ignored.")
	assert.True(t, v.IsValidCommitMessage(), "IsValidCommitMessage('test(valid-commit-msg): short first line\n\n# This is a long comment that should be ignored. This is a long comment that should be ignored.')")
}
