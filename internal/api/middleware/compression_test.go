package middleware

import (
	"compress/gzip"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestCompressHandler(t *testing.T) {
	// Crée un handler de test qui renvoie une grande chaîne
	testHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain")
		largeString := strings.Repeat("test data", 100)
		w.Write([]byte(largeString))
	})

	// Crée une requête de test avec support gzip
	req := httptest.NewRequest("GET", "/test", nil)
	req.Header.Set("Accept-Encoding", "gzip")
	rr := httptest.NewRecorder()

	// Applique le middleware de compression
	handler := CompressHandler(testHandler)
	handler.ServeHTTP(rr, req)

	// Vérifie les headers
	if rr.Header().Get("Content-Encoding") != "gzip" {
		t.Error("Content-Encoding header non défini à gzip")
	}

	// Décompresse la réponse
	reader, err := gzip.NewReader(rr.Body)
	if err != nil {
		t.Fatal("Erreur lors de la création du reader gzip:", err)
	}
	defer reader.Close()

	decompressed, err := io.ReadAll(reader)
	if err != nil {
		t.Fatal("Erreur lors de la décompression:", err)
	}

	// Vérifie que les données décompressées sont correctes
	expectedData := strings.Repeat("test data", 100)
	if string(decompressed) != expectedData {
		t.Error("Les données décompressées ne correspondent pas aux données attendues")
	}
}
