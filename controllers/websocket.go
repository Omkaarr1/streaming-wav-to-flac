package controllers

import (
    "io"
    "log"
    "os/exec"
    "github.com/gorilla/websocket"
    "net/http"
)

// WebSocket upgrader to establish WebSocket connections
var upgrader = websocket.Upgrader{
    ReadBufferSize:  1024,
    WriteBufferSize: 1024,
    CheckOrigin: func(r *http.Request) bool {
        return true
    },
}

// HandleWebSocket processes WAV to FLAC conversion in real-time
func HandleWebSocket(w http.ResponseWriter, r *http.Request) {
    conn, err := upgrader.Upgrade(w, r, nil)
    if err != nil {
        log.Println("Failed to set WebSocket upgrade:", err)
        return
    }
    log.Println("WebSocket connection opened.")
    defer func() {
        conn.Close()
        log.Println("WebSocket connection closed.")
    }()

    // Start ffmpeg command to convert WAV to FLAC in real-time
    ffmpegCmd := exec.Command("ffmpeg", "-f", "webm", "-i", "pipe:0", "-c:a", "flac", "-f", "flac", "pipe:1")

    // Get stdin and stdout pipes
    ffmpegStdin, err := ffmpegCmd.StdinPipe()
    if err != nil {
        log.Println("Failed to get stdin pipe:", err)
        return
    }
    ffmpegStdout, err := ffmpegCmd.StdoutPipe()
    if err != nil {
        log.Println("Failed to get stdout pipe:", err)
        return
    }

    // Start the ffmpeg process
    if err := ffmpegCmd.Start(); err != nil {
        log.Println("Failed to start ffmpeg:", err)
        return
    }

    // Goroutine to read WAV data from WebSocket and pipe to ffmpeg
    go func() {
        for {
            // Read WAV data from WebSocket
            _, wavData, err := conn.ReadMessage()
            if err != nil {
                log.Println("Error reading from WebSocket:", err)
                break
            }

            // Write WAV data to ffmpeg stdin
            if _, err := ffmpegStdin.Write(wavData); err != nil {
                log.Println("Error writing to ffmpeg stdin:", err)
                break
            }
        }

        // Close ffmpeg stdin to indicate no more data
        ffmpegStdin.Close()
    }()

    // Stream FLAC data from ffmpeg stdout to WebSocket
    buf := make([]byte, 1024)
    for {
        // Read FLAC data from ffmpeg stdout
        n, err := ffmpegStdout.Read(buf)
        if err != nil {
            if err == io.EOF {
                break
            }
            log.Println("Error reading from ffmpeg stdout:", err)
            break
        }

        // Send FLAC data back to WebSocket client
        if err := conn.WriteMessage(websocket.BinaryMessage, buf[:n]); err != nil {
            log.Println("Error writing to WebSocket:", err)
            break
        }
    }

    // Wait for ffmpeg to finish
    if err := ffmpegCmd.Wait(); err != nil {
        log.Println("ffmpeg process exited with error:", err)
    }
}
