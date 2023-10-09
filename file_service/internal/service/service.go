package service

import (
	"github.com/gookit/slog"
	"sync"
)

type IFileRepository interface {
	ProcessAndStoreChunk(fileID, sequenceNumber int, chunk []byte) error
}

type FileService struct {
	repo IFileRepository
}

func NewFileService(repo IFileRepository) *FileService {
	return &FileService{repo: repo}
}

func (fileSvc *FileService) UploadFile(fileID int, data []byte) error {
	chunkSize := calculateChunkSize(len(data)) // Implement a function to determine chunk size.

	var wg sync.WaitGroup
	// Iterate through chunks and create Goroutines.
	for i, chunk := range splitIntoChunks(data, chunkSize) {
		wg.Add(1)
		go func(chunk []byte, sequenceNumber int) {
			defer wg.Done()
			if err := fileSvc.repo.ProcessAndStoreChunk(fileID, sequenceNumber, chunk); err != nil {
				slog.Errorf("Failed to store the chunk of data: %v", err)
			}
		}(chunk, i+1)
	}
	return nil
}

func splitIntoChunks(data []byte, chunkSize int) [][]byte {
	var chunks [][]byte

	for i := 0; i < len(data); i += chunkSize {
		end := i + chunkSize
		if end > len(data) {
			end = len(data)
		}
		chunks = append(chunks, data[i:end])
	}

	return chunks
}

func calculateChunkSize(fileSize int) int {
	// Define a reasonable default chunk size.
	defaultChunkSize := 1024 * 1024 // 1 MB

	// You can adjust the chunk size based on your requirements.
	// For example, you can consider the file size and available memory.
	// Here, we use a 10 MB chunk size as an example.
	if fileSize <= 0 {
		return defaultChunkSize
	}

	// Calculate the chunk size as a fraction of the file size,
	// or use the default chunk size if the file size is too small.
	chunkSize := fileSize / 10 // 10% of file size
	if chunkSize < defaultChunkSize {
		return defaultChunkSize
	}

	return chunkSize
}
