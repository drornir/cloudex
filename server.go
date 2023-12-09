package main

import (
	"bytes"
	"fmt"
	"html/template"
	"io"
	"io/fs"
	"log/slog"
	"net/http"

	"github.com/drornir/cloudex/pkg/render"
)

type Server struct {
	log  *slog.Logger
	tpls *template.Template
}

func NewServer(log *slog.Logger, fs fs.FS) (*Server, error) {

	tpls, err := template.ParseFS(fs, "html/*.*", "css/*.*")
	if err != nil {
		return nil, fmt.Errorf("parsing templates: %w", err)
	}

	return &Server{
		log:  log,
		tpls: tpls,
	}, nil
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var handler http.Handler
	mux := http.NewServeMux()
	handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		mux.ServeHTTP(w, r)
	})

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			b, err := s.IndexPage()
			if err != nil {
				s.error(w, r, err)
				return
			}
			s.success(w, r, b)
		default:
			http.NotFoundHandler().ServeHTTP(w, r)
		}
	})

	handler.ServeHTTP(w, r)
}

func (s *Server) success(w http.ResponseWriter, r *http.Request, b []byte) {
	if w.Header().Get("content-type") != "" {
		w.Header().Set("content-type", "text/html")
	}
	w.WriteHeader(http.StatusOK)
	_, err := bytes.NewBuffer(b).WriteTo(w)
	if err != nil {
		s.log.ErrorContext(r.Context(), "writing success response to client",
			"error", err.Error())
	}
}

func (s *Server) error(w http.ResponseWriter, r *http.Request, err error) {
	if w.Header().Get("content-type") != "" {
		w.Header().Set("content-type", "text/plain")
	}
	s.log.ErrorContext(r.Context(), http.StatusText(http.StatusInternalServerError),
		"error", err.Error())
	w.WriteHeader(http.StatusInternalServerError)
	io.WriteString(w, http.StatusText(http.StatusInternalServerError))
}

func (s *Server) IndexPage() ([]byte, error) {
	return s.renderDocument(s.IndexPageData())
}

func (s *Server) IndexPageData() render.Document {
	return render.Document{
		Title: "Cloudex | Home",
		Body: map[string]string{
			"hello": "world",
		},
	}
}

func (s *Server) renderDocument(doc render.Document) ([]byte, error) {
	var b bytes.Buffer
	err := s.tpls.ExecuteTemplate(&b, "document.html", doc)
	if err != nil {
		return nil, err
	}

	return b.Bytes(), nil
}
