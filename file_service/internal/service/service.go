package service

import (
	"errors"
	"github.com/gookit/slog"
	"sync"
)

type IFileRepository interface {
	ProcessAndStoreChunk(fileID uint, sequenceNumber int, chunk []byte) error
	StoreFileData(uid, fileType string) (uint, error)
}

type FileService struct {
	repo IFileRepository
}

func NewFileService(repo IFileRepository) *FileService {
	return &FileService{repo: repo}
}

func (fileSvc *FileService) UploadFileData(data []byte, uid, fileType string) error {
	chunkSize, err := calculateChunkSize(len(data))
	if err != nil {
		return err
	}
	fileID, err := fileSvc.repo.StoreFileData(uid, fileType)
	if err != nil {
		return err
	}

	maxRoutines := 4
	tasks := make(chan struct{}, maxRoutines)
	var wg sync.WaitGroup
	for i, chunk := range splitIntoChunks(data, chunkSize) {
		wg.Add(1)
		tasks <- struct{}{}
		go func(chunk []byte, sequenceNumber int) {
			defer func() {
				<-tasks
				wg.Done()
			}()
			if err := fileSvc.repo.ProcessAndStoreChunk(fileID, sequenceNumber, chunk); err != nil {
				slog.Errorf("Failed to store the chunk of data: %v", err)
			}
		}(chunk, i+1)
	}
	wg.Wait()
	return nil
}

func splitIntoChunks(data []byte, chunkSize int) [][]byte {
	slog.Info("file is being split into chunks")
	var chunks [][]byte

	for i := 0; i < len(data); i += chunkSize {
		end := i + chunkSize
		if end > len(data) {
			end = len(data)
		}
		chunks = append(chunks, data[i:end])
	}
	slog.Info("file successfully split")
	return chunks
}

func calculateChunkSize(fileSize int) (int, error) {
	// Define a reasonable default chunk size.
	defaultChunkSize := 1024 * 1024 // 1 MB

	// You can adjust the chunk size based on your requirements.
	// For example, you can consider the file size and available memory.
	// Here, we use a 10 MB chunk size as an example.
	if fileSize <= 0 {
		return 0, errors.New("file size is equal to 0")
	}

	// Calculate the chunk size as a fraction of the file size,
	// or use the default chunk size if the file size is too small.
	chunkSize := fileSize / 10 // 10% of file size
	if chunkSize < defaultChunkSize {
		return defaultChunkSize, nil
	}

	return chunkSize, nil
}
