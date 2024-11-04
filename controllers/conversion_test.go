package controllers

import (
    "bytes"
    "os/exec"
    "testing"
)

// TestConvertWavToFlac checks if the ffmpeg command for WAV to FLAC conversion runs successfully.
func TestConvertWavToFlac(t *testing.T) {
    // Simulate WAV data as input
    input := bytes.NewBuffer([]byte("sample wav data"))

    // Run the ffmpeg command with simulated input and output buffers
    cmd := exec.Command("ffmpeg", "-f", "wav", "-i", "pipe:0", "-c:a", "flac", "-f", "flac", "pipe:1")
    cmd.Stdin = input
    output := &bytes.Buffer{}
    cmd.Stdout = output

    // Check if the command runs successfully
    if err := cmd.Run(); err != nil {
        t.Fatalf("Conversion failed: %v", err)
    }

    if output.Len() == 0 {
        t.Error("Expected output data, but got none")
    }
}

// TestInvalidInputErrorHandling tests error handling for invalid WAV input.
func TestInvalidInputErrorHandling(t *testing.T) {
    // Simulate invalid input data that ffmpeg will fail to convert
    invalidInput := bytes.NewBuffer([]byte("invalid wav data"))

    // Run the ffmpeg command
    cmd := exec.Command("ffmpeg", "-f", "wav", "-i", "pipe:0", "-c:a", "flac", "-f", "flac", "pipe:1")
    cmd.Stdin = invalidInput
    var output bytes.Buffer
    cmd.Stdout = &output

    // Run the command and check for an error
    err := cmd.Run()
    if err == nil {
        t.Fatal("Expected an error for invalid input, but got none")
    }

    // Verify the output buffer is empty as expected
    if output.Len() != 0 {
        t.Error("Expected no output data for invalid input, but some data was produced")
    }
}
