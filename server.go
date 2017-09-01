package main

import (
	"fmt"
	"net/http"
)

var policy = "script-src http: https: 'self' 'unsafe-inline' 'report-sample' 'strict-dynamic' 'nonce-0yNME4F15OSPevU+UdZXGg=='; report-uri http://localhost:8081/report/;"
var page = `
<html>

<head>
    <title>CSP Selenium Test Page</title>
</head>

<body>
    <h3 id="hello">hello!</h3>
    <input id="title1" type="button" value="change title 1"/>
    <br/>
    <input id="title2" type="button" value="change title 2"/>
</body>

<script nonce="0yNME4F15OSPevU+UdZXGg==">
function changeTitle1() {
    document.title = 'Set by script with nonce';
}
</script>

<!--
// this script doesn't specify a nonce so it shouldn't execute in the browser
// for it to execute the script tag should be as follows
// <script nonce="0yNME4F15OSPevU+UdZXGg==">
-->
<script>
function changeTitle2() {
    document.title = "Set by script without nonce";
}
</script>

<!-- add event listeners for both buttons -->
<script nonce="0yNME4F15OSPevU+UdZXGg==">
// CSP does not allow inline JS handlers, e.g. onclick=changeTitle()
document.addEventListener('DOMContentLoaded', function () {
    document.getElementById('title1').addEventListener('click', changeTitle1);
    document.getElementById('title2').addEventListener('click', changeTitle2);
});
</script>
</html>
`

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Security-Policy-Report-Only", policy)
	//w.Header().Set("Content-Security-Policy", policy)
	fmt.Fprintf(w, page)
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
