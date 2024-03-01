package parser

import (
    "testing"
)

func TestCountParts_Positive1(t *testing.T) {
    emailText := `From: sender@example.com
To: recipient@example.com
Subject: Test email

This is a test email with a single part.`
    expected := 1

    actual := CountParts(emailText)

    if actual != expected {
        t.Errorf("CountParts() = %d; want %d", actual, expected)
    }
}

func TestCountParts_Positive2(t *testing.T) {
    emailText := `From: sender@example.com
To: recipient@example.com
Subject: Test email

This is a test email with multiple parts.

--Boundary_123456789
Content-Type: text/plain; charset="UTF-8"

This is the plain text part of the email.

--Boundary_123456789
Content-Type: text/html; charset="UTF-8"

<html>
<head>
<title>Test Email</title>
</head>
<body>
<p>This is the HTML part of the email.</p>
</body>
</html>

--Boundary_123456789--`
    expected := 3

    actual := CountParts(emailText)

    if actual != expected {
        t.Errorf("CountParts() = %d; want %d", actual, expected)
    }
}

func TestCountParts_Negative1(t *testing.T) {
    emailText := `From: sender@example.com
To: recipient@example.com
Subject: Test email`
    expected := 0

    actual := CountParts(emailText)

    if actual != expected {
        t.Errorf("CountParts() = %d; want %d", actual, expected)
    }
}

func TestCountParts_Negative2(t *testing.T) {
    emailText := ""
    expected := 0

    actual := CountParts(emailText)

    if actual != expected {
        t.Errorf("CountParts() = %d; want %d", actual, expected)
    }
}

func TestParseEmail_Positive1(t *testing.T) {
    emailText := `From: sender@example.com
To: recipient@example.com
Subject: Test email

This is a test email with a single part.`
    expectedCount := 1
    expectedContents := []string{"This is a test email with a single part."}

    actualCount, actualContents := ParseEmail(emailText)

    if actualCount != expectedCount {
        t.Errorf("ParseEmail() count = %d; want %d", actualCount, expectedCount)
    }

    for i := range expectedContents {
        if actualContents[i] != expectedContents[i] {
            t.Errorf("ParseEmail() content = %s; want %s", actualContents[i], expectedContents[i])
        }
    }
}

func TestParseEmail_Positive2(t *testing.T) {
    emailText := `From: sender@example.com
To: recipient@example.com
Subject: Test email

This is a test email with multiple parts.

--Boundary_123456789
Content-Type: text/plain; charset="UTF-8"

This is the plain text part of the email.

--Boundary_123456789
Content-Type: text/html; charset="UTF-8"

<html>
<head>
<title>Test Email</title>
</head>
<body>
<p>This is the HTML part of the email.</p>
</body>
</html>

--Boundary_123456789--`
    expectedCount := 3
    expectedContents := []string{
        "This is a test email with multiple parts.",
        "This is the plain text part of the email.",
        "<html>\n<head>\n<title>Test Email</title>\n</head>\n<body>\n<p>This is the HTML part of the email.</p>\n</body>\n</html>\n",
    }

    actualCount, actualContents := ParseEmail(emailText)

    if actualCount != expectedCount {
        t.Errorf("ParseEmail() count = %d; want %d", actualCount, expectedCount)
    }

    for i := range expectedContents {
        if actualContents[i] != expectedContents[i] {
            t.Errorf("ParseEmail() content = %s; want %s", actualContents[i], expectedContents[i])
        }
    }
}

func TestParseEmail_Negative1(t *testing.T) {
    emailText := `From: sender@example.com
To: recipient@example.com
Subject: Test email`
    expectedCount := 0
    var expectedContents []string

    actualCount, actualContents := ParseEmail(emailText)

    if actualCount != expectedCount {
        t.Errorf("ParseEmail() count = %d; want %d", actualCount, expectedCount)
    }

    if len(actualContents) != len(expectedContents) {
        t.Errorf("ParseEmail() content length = %d; want %d", len(actualContents), len(expectedContents))
    }
}

func TestParseEmail_Negative2(t *testing.T) {
    emailText := ""
    expectedCount := 0
    var expectedContents []string

    actualCount, actualContents := ParseEmail(emailText)

    if actualCount != expectedCount {
        t.Errorf("ParseEmail() count = %d; want %d", actualCount, expectedCount)
    }

    if len(actualContents) != len(expectedContents) {
        t.Errorf("ParseEmail() content length = %d; want %d", len(actualContents), len(expectedContents))
    }
}

