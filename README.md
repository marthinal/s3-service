# s3-service
Service client for AWS S3 using Wire for dependency injection

### Example: Save a file

```go
package main

import (
	"fmt"
	s3 "github.com/marthinal/s3-service"
	"log"
	"net/http"
)

const MaxUploadSize = 1024 * 1024

func main()  {
	fs := s3.GetFS()
	s := &Service{fs: fs}
	err := http.ListenAndServe("0.0.0.0:6090", s)
	if err != nil {
		log.Fatal("Error starting the service")
	}
}

type Service struct {
	fs s3.FS
}

func (s Service) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fs := s.fs
	switch r.URL.Path {
	default:
		http.Error(w, "not found", http.StatusNotFound)
		return
	case "/upload":
		if r.Method != "POST" {
			w.WriteHeader(http.StatusMethodNotAllowed)
		}
		w.Header().Set("Access-Control-Allow-Origin", "*")
		err := r.ParseMultipartForm(100000)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		for _, h := range r.MultipartForm.File["images"] {
			if h.Size > MaxUploadSize {
				http.Error(w, fmt.Sprintf("The uploaded image is too big: %s", h.Filename), http.StatusBadRequest)
				return
			}
			file, err := h.Open()
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			defer file.Close()

			err = fs.Save("EXAMPLE_FILENAME.png", file)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			return
		}
	}
}
```

This is WIP. I just added the function to upload files.

You need these env variables:
- AWS_ACCESS_KEY_ID
- AWS_SECRET_ACCESS_KEY
- AWS_REGION
- AWS_BUCKET_NAME

AWS SDK gets the values by default. 

See https://github.com/google/wire for more info about Wire.