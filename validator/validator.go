//
// Check commit message according to angularjs guidelines:
// https://docs.google.com/document/d/1QrDFcIiPjSLDn3EL15IJygNPiHORgU1_OOAqWjiDU5Y/edit#
//
// Copyright 2014 TLD dotHIV Registry GmbH.
// @author Markus Tacker <m@dotHIV.org>
//
package validator

import (
	"fmt"
	"regexp"
	"strings"
)

type CommitMessageValidator struct {
	message                 string
	commit_types            []string
	first_line_length_limit int
	other_line_length_limit int
}

func NewCommitMessageValidator(message string) (v *CommitMessageValidator) {
	v = new(CommitMessageValidator)
	v.message = message
	v.commit_types = append(v.commit_types, "feat", "fix", "docs", "style", "refactor", "test", "chore")
	v.first_line_length_limit = 70
	v.other_line_length_limit = 100
	return v
}

func (v *CommitMessageValidator) IsValidCommitMessage() (valid bool) {
	err := v.ValidateCommitMessage()
	if err == nil {
		valid = true
	}
	return
}

func (v *CommitMessageValidator) ValidateCommitMessage() (err error) {
	if len(v.message) == 0 {
		err = fmt.Errorf("Empty commit message")
		return
	}
	lines := strings.Split(v.message, "\n")

	firstline := lines[0]
	validFirstLine := regexp.MustCompile(`^(.*)\((.*)\): (.*)$`)

	groups := validFirstLine.FindStringSubmatch(firstline)
	if len(groups) != 4 {
		err = fmt.Errorf("First commit message line (header) does not follow format: type(scope): message")
		return
	}

	commit_message := groups[3]

	capitalMatch, _ := regexp.MatchString("^[A-Z]", commit_message)
	if capitalMatch {
		err = fmt.Errorf("Commit subject '%s' must not have a capital first letter!", commit_message)
		return
	}

	dotMatch, _ := regexp.MatchString("\\.$", commit_message)
	if dotMatch {
		err = fmt.Errorf("Commit subject '%s' must not end with a dot (.)!", commit_message)
		return
	}

	commit_type := groups[1]
	if !v.isValidType(commit_type) {
		err = fmt.Errorf("Commit type '%s' not in valid ones: %s!", commit_type, strings.Join(v.commit_types, ", "))
		return
	}

	if len(lines) > 1 {
		if lines[1] != "" {
			err = fmt.Errorf("Second commit message line must be empty!")
			return
		}
		for _, l := range lines[1:] {
			if len(l) > v.other_line_length_limit {
				if string(l[0]) == "#" { // Ignore comments.
					continue
				}
				err = fmt.Errorf("Other lines should not be longer than %d characters! It has %d characters.", v.other_line_length_limit, len(l))
				return
			}
		}
	}

	if len(firstline) > v.first_line_length_limit {
		err = fmt.Errorf("First line should not be longer than %d characters! It has %d characters.", v.first_line_length_limit, len(firstline))
		return
	}
	return
}

func (v *CommitMessageValidator) isValidType(t string) (b bool) {
	for _, v := range v.commit_types {
		if v == t {
			return true
		}
	}
	return
}
