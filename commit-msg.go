//
// Git commit hook:
//  .git/hooks/commit-msg
//
// Check commit message according to angularjs guidelines:
// https://docs.google.com/document/d/1QrDFcIiPjSLDn3EL15IJygNPiHORgU1_OOAqWjiDU5Y/edit#
//
// Copyright 2014 TLD dotHIV Registry GmbH.
// @author Markus Tacker <m@dotHIV.org>
//
package main

import (
	"github.com/dothiv/validate-commit-msg/validator"
	"io/ioutil"
	"os"
)

func main() {
	content, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		os.Stderr.WriteString(err.Error())
		os.Exit(1)
	}

	v := validator.NewCommitMessageValidator(string(content))
	validError := v.ValidateCommitMessage()
	if validError == nil {
		os.Exit(0)
	} else {
		os.Stderr.WriteString(validError.Error() + "\n")
		os.Stderr.WriteString("Refer to commit guide: https://docs.google.com/document/d/1QrDFcIiPjSLDn3EL15IJygNPiHORgU1_OOAqWjiDU5Y/edit#\n")
		os.Stderr.WriteString("\nYour message:\n\n")
		os.Stderr.Write(content)
		os.Stderr.WriteString("\n")
		os.Exit(1)
	}
}
