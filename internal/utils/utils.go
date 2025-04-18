package utils

func NewPage(title string, content string) string {
	return `<!DOCTYPE html>
<html>
<head>
    <title>` + title + `</title>
</head>
<body>
    <h1>` + title + `</h1>
    <p>` + content + `</p>
</body>
</html>`
}
