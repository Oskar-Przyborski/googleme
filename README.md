# googleme
Googleme is a cli tool, that just googles pased query.
I don't know if it's legal, so if anyone has any objections, please contact me :>>

# Usage
Built the executable file with
```console
$ go build
```
And add it to PATH

Then just type
```console
$ googleme <query>
```
For example
```console
$ googleme how to brush your teeth
```
Note, that you can use quotes too
```console
$ googleme "why the Earth is flat"
```

The resulted output looks like (with clickable links)
```console
$ googleme how to make http request
Searching for: how to make http request

cloud.google.com
Make an HTTP request | Workflows | Google Cloud
Make an HTTP request · On this page · Invoke an HTTP endpoint · Access HTTP response data saved in a variable · Samples. Assign the response from an API call; Make ...


www.tutorialspoint.com
HTTP - Requests
HTTP - Requests, An HTTP client sends an HTTP request to a server in the form of a ... A POST request is used to send data to the server, for example, ...


www.freecodecamp.org
Here are the most popular ways to make an HTTP request in ...
May 8, 2018 ... To make an HTTP call in Ajax, you need to initialize a new XMLHttpRequest() method, specify the URL endpoint and HTTP method (in this case GET).


www.apirequest.io
ApiRequest.io | Make HTTP requests and share
A free debugging tool to make and capture RESTful API Requests. Share your workspace to get help.


www.freecodecamp.org
JavaScript Get Request – How to Make an HTTP Request in JS
5 days ago ... This interaction between your frontend application and the backend server is possible through HTTP requests. There are five popular HTTP methods ...

...
```
