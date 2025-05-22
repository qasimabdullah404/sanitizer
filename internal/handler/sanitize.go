package handler

import (
	"bufio"
	"io"
	"net/http"

	"gihtub.com/qasimabdullah404/sanitizer/internal/obfuscator"
	"go.uber.org/zap"
)

func UploadHandler(log *zap.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		file, _, err := r.FormFile("logfile")
		if err != nil {
			log.Warn("Invalid file upload", zap.Error(err))
			http.Error(w, "Invalid file upload", http.StatusBadRequest)
			return
		}
		defer file.Close()

		w.Header().Set("Content-Type", "text/plain")
		w.Header().Set("Content-Disposition", "attachment; filename=\"sanitized.log\"")

		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			sanitized := obfuscator.ObfuscateLine(scanner.Text())
			io.WriteString(w, sanitized+"\n")
		}

		if err := scanner.Err(); err != nil {
			log.Error("Error reading uploaded file", zap.Error(err))
			http.Error(w, "File reading error", http.StatusInternalServerError)
		}
	}
}
