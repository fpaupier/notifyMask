package main

import (
	"strconv"
	"testing"
)

func TestGenerateEmailBody(t *testing.T) {
	ts := "2020-10-12T14:14:07.999999999Z"
	recipientName := "Tester"
	alertId := 12
	expectedBody := "<!DOCTYPE html>\n<html lang=\"en\">\n<body>\n<p>\n    Hello " + recipientName + ",<br>\n    Someone has been seen not wearing their mask on " + "Monday, 12-Oct-20 14:14:07 UTC" + ".<br>\n    Check alert ID #" + strconv.Itoa(alertId) + " for more information.<br>\n</p>\n</body>\n</html>"
	body := generateBody(templateFPath, recipientName, ts, alertId)
	if expectedBody != body {
		t.Errorf("expected body:\n %s\ngot:\n%s\n", expectedBody, body)
	}
}
