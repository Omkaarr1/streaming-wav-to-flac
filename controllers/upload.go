package controllers

import (
    "fmt"
    "log"
    "net/http"
    "os"
    "os/exec"
    "path/filepath"
    "time"

    "github.com/gin-gonic/gin"
)

// HandleWavUpload receives a WAV file, converts it to FLAC, and provides a download link
func HandleWavUpload(c *gin.Context) {
    // Extract the file from the request
    file, err := c.FormFile("audio")
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to get audio file"})
        return
    }

    // Save the uploaded WAV file temporarily
    tempDir := "./temp"
    os.MkdirAll(tempDir, os.ModePerm) // Ensure temp directory exists
    wavFilePath := filepath.Join(tempDir, fmt.Sprintf("input_%d.wav", time.Now().UnixNano()))
    if err := c.SaveUploadedFile(file, wavFilePath); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save WAV file"})
        return
    }

    // Convert WAV to FLAC
    flacFilePath := filepath.Join(tempDir, fmt.Sprintf("output_%d.flac", time.Now().UnixNano()))
    ffmpegCmd := exec.Command("ffmpeg", "-i", wavFilePath, "-c:a", "flac", flacFilePath)

    if err := ffmpegCmd.Run(); err != nil {
        log.Println("Error converting to FLAC:", err)
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to convert to FLAC"})
        return
    }

    // Clean up the original WAV file
    os.Remove(wavFilePath)

    // Provide the download link for the FLAC file
    c.JSON(http.StatusOK, gin.H{"message": "Conversion successful", "download_url": "/download/" + filepath.Base(flacFilePath)})
}
