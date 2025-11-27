package domain

func BaseCSS() string {
	return `
:root {
	--font: system-ui, -apple-system, Segoe UI, Roboto, Ubuntu, Cantarell, "Helvetica Neue", Arial, "Noto Sans", "Liberation Sans", sans-serif;
	max-width: 100dvh;
}
body {
	font-family: var(--font);
	color: #222;
	line-height: 1.4;
	background: #fafbfc;
}
.main, .container {
	margin: 1rem auto;
	padding: 0 1rem;
	max-width: 900px;
}
h1, h2, h3 {
	line-height: 1.25;
	margin-top: 1.2rem;
	margin-bottom: .2rem;
}
pre {
	padding: 1rem;
	overflow: auto;
	background: #f6f8fa;
	border-radius: 8px;
	border: 2px solid #e1e4e8;
	box-shadow: 0 2px 8px rgba(0,0,0,0.04);
	margin: 1.2em 0;
}
pre code {
	background: none;
	padding: 0;
	border-radius: 0;
	border: none;
	box-shadow: none;
}
code {
	background: #f6f8fa;
	padding: .2rem .4rem;
	border-radius: 4px;
	border: 1px solid #e1e4e8;
}
table {
	border-collapse: collapse;
	width: 100%;
	margin: 1.5rem 0;
	background: #fff;
	box-shadow: 0 2px 8px rgba(0,0,0,0.03);
}
th, td {
	border: 1px solid #bbb;
	padding: .7rem;
	text-align: left;
}
th {
	background: #f3f4f6;
	font-weight: 600;
}
tr:nth-child(even) {
	background: #f8f9fa;
}
tr:hover {
	background: #eef2fb;
}
blockquote {
	border-left: 4px solid #6c63ff;
	padding-left: 1rem;
	color: #555;
	background: #f6f8fa;
	margin: 1rem 0;
}
img {
	max-width: 100%;
	height: auto;
	display: block;
	margin: 1rem auto;
	box-shadow: 0 2px 8px rgba(0,0,0,0.05);
}
a {
	color: #0b5fff;
	text-decoration: none;
	border-bottom: 1px dotted #0b5fff;
}
a:hover {
	text-decoration: underline;
	background: #eaf4ff;
}
ul, ol {
	padding-left: 1.5rem;
	margin-bottom: 1rem;
}
li {
	margin-bottom: .4rem;
}
hr {
	border: none;
	border-top: 1px solid #e1e4e8;
	margin: 2rem 0;
}
/* Iconos check y cruz para tablas */
.check {
	color: #6c63ff;
	font-size: 1.2em;
}
.cross {
	color: #e74c3c;
	font-size: 1.2em;
}
`
}
