package server

import "net/http"

func (s *Server) registerHandlers(mux *http.ServeMux) {
	mux.HandleFunc("/assets/", s.staticFilesHandler)

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/" {
			switch r.Method {
			case http.MethodGet:
				b, err := s.IndexPage()
				if err != nil {
					s.error(w, r, err)
					return
				}
				s.success(w, r, b)
			default:
				s.notFoundHandler(w, r)
			}
		} else {
			s.notFoundHandler(w, r)
		}
	})
}
