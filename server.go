package main

import (
  "log"
  "strings"
  "net/http"

  "github.com/gorilla/mux"
)

// MyServer struct for mux router
type MyServer struct {
  r *mux.Router
}

func main() {

  htmlRouter := mux.NewRouter().StrictSlash(true)
  /*
  htmlRouter.HandleFunc("/documentation", func(w http.ResponseWriter, r *http.Request) {
    http.ServeFile(w, r, "./html/documentation.html")
  })
  htmlRouter.HandleFunc("/examples", func(w http.ResponseWriter, r *http.Request) {
    http.ServeFile(w, r, "./html/examples.html")
  })
  htmlRouter.HandleFunc("/examples/repository/minimal", func(w http.ResponseWriter, r *http.Request) {
    http.ServeFile(w, r, "./html/static/schema/examples/repository/minimal.jsonld")
  })
  htmlRouter.HandleFunc("/examples/repository/full", func(w http.ResponseWriter, r *http.Request) {
    http.ServeFile(w, r, "./html/static/schema/examples/repository/full.jsonld")
  })
  htmlRouter.HandleFunc("/examples/dataset/minimal", func(w http.ResponseWriter, r *http.Request) {
    http.ServeFile(w, r, "./html/static/schema/examples/resource/dataset-minimal.jsonld")
  })
  htmlRouter.HandleFunc("/examples/dataset/full", func(w http.ResponseWriter, r *http.Request) {
    http.ServeFile(w, r, "./html/static/schema/examples/resource/dataset-full.jsonld")
  })
  */
  // htmlRouter.HandleFunc("/schema", Conneg)
  htmlRouter.HandleFunc("/schema/", Conneg)
  htmlRouter.HandleFunc("/schema/rdf.{ext}", Conneg)
  htmlRouter.PathPrefix("/static").Handler(http.StripPrefix("/static", http.FileServer(http.Dir("./html/static"))))
  htmlRouter.Handle("/", http.RedirectHandler("/schema/", http.StatusMovedPermanently))
  htmlRouter.HandleFunc("/schema/{resource}", func(w http.ResponseWriter, r *http.Request) {
    http.ServeFile(w, r, "./html/index.html")
  })

  http.Handle("/", &MyServer{htmlRouter})

  err := http.ListenAndServe(":9900", nil)
  // http 2.0 http.ListenAndServeTLS(":443", "server.crt", "server.key", nil)
  if err != nil {
    log.Fatal(err)
  }
}

func (s *MyServer) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
  rw.Header().Set("Access-Control-Allow-Origin", "*")
  rw.Header().Set("Access-Control-Allow-Methods", "POST, GET")
  rw.Header().Set("Access-Control-Allow-Headers",
    "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

  // Let the Gorilla work
  s.r.ServeHTTP(rw, req)
}

func addDefaultHeaders(fn http.HandlerFunc) http.HandlerFunc {
  return func(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Access-Control-Allow-Origin", "*")
    fn(w, r)
  }
}

// Conneg handles content negotiation for RDF requests
func Conneg(w http.ResponseWriter, r *http.Request) {
  vars := mux.Vars(r)
  fileExtension := vars["ext"]

  // log.Printf("File Extension: '" + fileExtension + "'")
  switch fileExtension {
    case "owl":
      http.ServeFile(w, r, "./html/static/schema/odo.owl")
      break

    case "rdf":
      http.ServeFile(w, r, "./html/static/schema/odo.owl")
      break

    case "xml":
      http.ServeFile(w, r, "./html/static/schema/odo.owl")
      break

    case "ttl":
      http.ServeFile(w, r, "./html/static/schema/odo.ttl")
      break

    case "nt":
      http.ServeFile(w, r, "./html/static/schema/odo.nt")
      break

    case "json":
      http.ServeFile(w, r, "./html/static/schema/odo.jsonld")
      break

    case "jsonld":
      http.ServeFile(w, r, "./html/static/schema/odo.jsonld")
      break

    default:
      // Check for Accept header
      accept := r.Header.Get("Accept")
      // log.Printf("Accept: '" + accept + "'")
      switch {
        case CaseInsensitiveContains(accept, "application/rdf+xml"):
          http.ServeFile(w, r, "./html/static/schema/odo.owl")
          break

        case CaseInsensitiveContains(accept, "text/html"):
          http.ServeFile(w, r, "./html/index.html")
          break

        case CaseInsensitiveContains(accept, "application/xml"):
          http.ServeFile(w, r, "./html/static/schema/odo.owl")
          break

        case CaseInsensitiveContains(accept, "application/rdf+turtle"):
          http.ServeFile(w, r, "./html/static/schema/odo.ttl")
          break

        case CaseInsensitiveContains(accept, "application/rdf+n3"):
          http.ServeFile(w, r, "./html/static/schema/odo.n3")
          break

        case CaseInsensitiveContains(accept, "application/json"):
          http.ServeFile(w, r, "./html/static/schema/odo.jsonld")
          break

        case CaseInsensitiveContains(accept, "application/ld+json"):
          http.ServeFile(w, r, "./html/static/schema/odo.jsonld")
          break

        default:
          http.ServeFile(w, r, "./html/index.html")
          break
      }
  }
}

func CaseInsensitiveContains(s, substr string) bool {
    s, substr = strings.ToUpper(s), strings.ToUpper(substr)
    return strings.Contains(s, substr)
}

