package minio

type ObjectInfo struct {
	Bucket string `json:"bucket"`
	Object string `json:"object"`
	URL    string `json:"url"`
}


