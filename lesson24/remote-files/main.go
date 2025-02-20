package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/google/uuid"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

const bucketName = "files"

func newStorage() (*minio.Client, error) {
	ctx := context.Background()

	endpoint := os.Getenv("MINIO_ENDPOINT")
	accessKeyID := os.Getenv("MINIO_ROOT_USER")
	secretAccessKey := os.Getenv("MINIO_ROOT_PASSWORD")

	minioClient, err := minio.New(
		endpoint,
		&minio.Options{
			Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
			Secure: false,
		})
	if err != nil {
		return nil, err
	}

	if err := createBucket(ctx, minioClient); err != nil {
		return nil, err
	}

	return minioClient, nil
}

func createBucket(ctx context.Context, client *minio.Client) error {
	// Make a new bucket called files.
	region := "us-east-1"

	if err := client.MakeBucket(ctx, bucketName, minio.MakeBucketOptions{Region: region}); err != nil {
		// Check to see if we already own this bucket (which happens if you run this twice)
		exists, errBucketExists := client.BucketExists(ctx, bucketName)
		if errBucketExists != nil || !exists {
			return err
		}
	}
	log.Printf("Successfully created bucket: %s\n", bucketName)

	policy := `{
        "Version": "2012-10-17",
        "Statement": [
            {
                "Effect": "Allow",
                "Principal": {"AWS": ["*"]},
                "Action": ["s3:GetObject"],
                "Resource": ["arn:aws:s3:::` + bucketName + `/*"]
            }
        ]
    }`

	return client.SetBucketPolicy(context.Background(), bucketName, policy)
}

type handler struct {
	storage *minio.Client
}

func newHandler(storage *minio.Client) *handler {
	return &handler{
		storage: storage,
	}
}

func (h *handler) ping(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Pong")
}

type uploadResponse struct {
	FileName string `json:"fileName"`
	FileSize int64  `json:"fileSize"`
	FileType string `json:"fileType"`
	FilePath string `json:"filePath"`
}

func newUploadResponse(fileName, fileType, filePath string, fileSize int64) *uploadResponse {
	return &uploadResponse{
		FileName: fileName,
		FileSize: fileSize,
		FileType: fileType,
		FilePath: filePath,
	}
}

func (h *handler) upload(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	// Limit file size to 10MB
	r.Body = http.MaxBytesReader(w, r.Body, 100<<20)

	// Parse multipart form with 32MB max memory
	if err := r.ParseMultipartForm(32 << 20); err != nil {
		http.Error(w, "File too large", http.StatusBadRequest)
		return
	}

	// Get file from form
	file, header, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "Error retrieving file from form", http.StatusBadRequest)
		return
	}
	defer file.Close()

	// Generate unique file path
	fileExt := filepath.Ext(header.Filename)
	uniqueID := uuid.New().String()
	objectName := fmt.Sprintf("%s%s", uniqueID, fileExt)

	// Get content type
	contentType := header.Header.Get("Content-Type")
	if contentType == "" {
		contentType = "application/octet-stream"
	}

	// Upload to MinIO
	info, err := h.storage.PutObject(
		ctx,
		bucketName,
		objectName,
		file,
		header.Size,
		minio.PutObjectOptions{
			ContentType: contentType,
			UserMetadata: map[string]string{
				"x-amz-acl": "public-read",
			},
		},
	)
	if err != nil {
		log.Printf("Error uploading to MinIO: %v", err)
		http.Error(w, "Error uploading file", http.StatusInternalServerError)
		return
	}

	// Send response
	resp := newUploadResponse(header.Filename, contentType, objectName, info.Size)
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		log.Printf("Error encoding response: %v", err)
		http.Error(w, "Error encoding response", http.StatusInternalServerError)
		return
	}
}

type downloadResponse struct {
	FileURL   string    `json:"file_url"`
	ExpiresAt time.Time `json:"expires_at"`
}

func newDownloadResponse(fileUrl string, expiresAt time.Time) *downloadResponse {
	return &downloadResponse{
		FileURL:   fileUrl,
		ExpiresAt: expiresAt,
	}
}

func (h *handler) download(w http.ResponseWriter, r *http.Request) {
	// Get file path from query parameter
	filePath := r.URL.Query().Get("file")
	if filePath == "" {
		http.Error(w, "File path is required", http.StatusBadRequest)
		return
	}

	// Set expiry time (e.g., 1 hour)
	expiry := time.Hour

	// Generate presigned URL
	presignedURL, err := h.storage.PresignedGetObject(
		r.Context(),
		bucketName,
		filePath,
		expiry,
		nil,
	)
	if err != nil {
		log.Printf("Error generating presigned URL: %v", err)
		http.Error(w, "Error generating download URL", http.StatusInternalServerError)
		return
	}

	// Prepare response
	response := newDownloadResponse(presignedURL.String(), time.Now().Add(expiry))

	// Send response
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		log.Printf("Error encoding response: %v", err)
		http.Error(w, "Error encoding response", http.StatusInternalServerError)
		return
	}
}

func main() {
	storage, err := newStorage()
	if err != nil {
		log.Fatal(err)
	}

	h := newHandler(storage)

	http.HandleFunc("GET /ping", h.ping)
	http.HandleFunc("POST /upload", h.upload)
	http.HandleFunc("GET /download", h.download)

	port := os.Getenv("PORT")
	log.Printf("Starting server on :%s\n", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}
}
