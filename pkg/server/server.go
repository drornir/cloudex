package server

import (
	"bytes"
	"fmt"
	"html/template"
	"io"
	"io/fs"
	"log/slog"
	"net/http"

	"github.com/drornir/cloudex/pkg/app"
	"github.com/drornir/cloudex/pkg/render"
)

type Server struct {
	log     *slog.Logger
	fs      fs.FS
	tpls    *template.Template
	handler http.Handler
}

func New(log *slog.Logger, fs fs.FS) (*Server, error) {
	tpls, err := template.ParseFS(fs, "html/*.*", "css/*.*")
	if err != nil {
		return nil, fmt.Errorf("parsing templates: %w", err)
	}

	return &Server{
		log:  log,
		fs:   fs,
		tpls: tpls,
	}, nil
}

func (s *Server) HTTPHandler() http.Handler {
	mux := http.NewServeMux()
	s.registerHandlers(mux)

	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			go func() {
				if r.Body != nil {
					_, _ = io.Copy(io.Discard, r.Body)
					_ = r.Body.Close()
				}
			}()
		}()

		if _, p := mux.Handler(r); p == "" {
			s.notFoundHandler(w, r)
			return
		}

		// TODO user
		r = r.WithContext(app.ContextWithUser(r.Context(), app.User{ID: "dror-id"}))

		mux.ServeHTTP(w, r)
	})

	return handler
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
	return s.IndexPageData().Render(s.tpls)
}

func (s *Server) IndexPageData() render.Document {
	return render.Document{
		Title: "Home",
	}
}

func (s *Server) notFoundHandler(w http.ResponseWriter, r *http.Request) {
	doc := render.Document{
		Title:        "Page Not Found",
		PageNotFound: true,
	}

	b, err := doc.Render(s.tpls)
	if err != nil {
		s.log.ErrorContext(r.Context(), "rendering 404 document",
			"error", err.Error())

		http.NotFoundHandler().ServeHTTP(w, r)
		return
	}

	w.Header().Set("content-type", "text/html")

	w.WriteHeader(http.StatusNotFound)

	if _, err := bytes.NewBuffer(b).WriteTo(w); err != nil {
		s.log.ErrorContext(r.Context(), "writing 404 response to client",
			"error", err.Error())
		_, _ = w.Write([]byte("<code>Page Not Found</code>"))
	}
}
