package server

import "net/http"

func (s *Server) staticFilesHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Cache-Control", "public, max-age=604800, immutable")

	http.FileServer(http.FS(s.fs)).ServeHTTP(w, r)
}
