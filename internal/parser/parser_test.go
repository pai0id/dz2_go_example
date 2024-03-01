package parser

import (
    "testing"
    "reflect"
)

func TestCountParts_Positive1(t *testing.T) {
    emailText := `From: sender@example.com
To: recipient@example.com
Subject: Test email

This is a test email with a single part.`

    expected := 1
    result := CountParts(emailText)
    if result != expected {
        t.Errorf("CountParts returned incorrect result, got: %d, want: %d.", result, expected)
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

    expected := 3 // Одна часть заголовка и две части контента
    result := CountParts(emailText)
    if result != expected {
        t.Errorf("CountParts returned incorrect result, got: %d, want: %d.", result, expected)
    }
}

func TestCountParts_Negative1(t *testing.T) {
    emailText := "" // Пустой текст письма

    expected := 0 // Ожидаем, что не будет ни одной части
    result := CountParts(emailText)
    if result != expected {
        t.Errorf("CountParts returned incorrect result, got: %d, want: %d.", result, expected)
    }
}

func TestCountParts_Negative2(t *testing.T) {
    emailText := `This is not an SMTP formatted email.`

    expected := 0 // Ожидаем, что не будет ни одной части
    result := CountParts(emailText)
    if result != expected {
        t.Errorf("CountParts returned incorrect result, got: %d, want: %d.", result, expected)
    }
}

func TestParseEmail_Positive1(t *testing.T) {
    emailText := `From: sender@example.com
To: recipient@example.com
Subject: Test email

This is a test email with a single part.`

    expectedParts := 1
    expectedContents := []string{"From: sender@example.com\nTo: recipient@example.com\nSubject: Test email\n\nThis is a test email with a single part."}

    partsCount, partContents := ParseEmail(emailText)
    if partsCount != expectedParts {
        t.Errorf("ParseEmail returned incorrect number of parts, got: %d, want: %d.", partsCount, expectedParts)
    }
    if !reflect.DeepEqual(partContents, expectedContents) {
        t.Errorf("ParseEmail returned incorrect part content, got: %v, want: %v.", partContents, expectedContents)
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

    expectedParts := 3
    expectedContents := []string{
        "From: sender@example.com\nTo: recipient@example.com\nSubject: Test email\n\nThis is a test email with multiple parts.",
        "Content-Type: text/plain; charset=\"UTF-8\"\n\nThis is the plain text part of the email.",
        "Content-Type: text/html; charset=\"UTF-8\"\n\n<html>\n<head>\n<title>Test Email</title>\n</head>\n<body>\n<p>This is the HTML part of the email.</p>\n</body>\n</html>",
    }

    partsCount, partContents := ParseEmail(emailText)
    if partsCount != expectedParts {
        t.Errorf("ParseEmail returned incorrect number of parts, got: %d, want: %d.", partsCount, expectedParts)
    }
    if !reflect.DeepEqual(partContents, expectedContents) {
        t.Errorf("ParseEmail returned incorrect part content, got: %v, want: %v.", partContents, expectedContents)
    }
}

func TestParseEmail_Negative1(t *testing.T) {
    emailText := "" // Пустой текст письма

    expectedParts := 0
    expectedContents := []string{}

    partsCount, partContents := ParseEmail(emailText)
    if partsCount != expectedParts {
        t.Errorf("ParseEmail returned incorrect number of parts, got: %d, want: %d.", partsCount, expectedParts)
    }
    if !reflect.DeepEqual(partContents, expectedContents) {
        t.Errorf("ParseEmail returned incorrect part content, got: %v, want: %v.", partContents, expectedContents)
    }
}

func TestParseEmail_Negative2(t *testing.T) {
    emailText := "This is not an SMTP formatted email."

    expectedParts := 0
    expectedContents := []string{}

    partsCount, partContents := ParseEmail(emailText)
    if partsCount != expectedParts {
        t.Errorf("ParseEmail returned incorrect number of parts, got: %d, want: %d.", partsCount, expectedParts)
    }
    if !reflect.DeepEqual(partContents, expectedContents) {
        t.Errorf("ParseEmail returned incorrect part content, got: %v, want: %v.", partContents, expectedContents)
    }
}

