package controllers

import (
    "net/http/httptest"
    "testing"
    "time"

    "github.com/gin-gonic/gin"
    "github.com/gorilla/websocket"
)

// SetupRouter creates a new Gin router for testing
func SetupRouter() *gin.Engine {
    r := gin.Default()
    r.GET("/ws", func(c *gin.Context) {
        HandleWebSocket(c.Writer, c.Request)
    })
    return r
}

// TestWebSocketStreaming tests the WebSocket connection and real-time streaming
func TestWebSocketStreaming(t *testing.T) {
    // Set up the test server
    server := httptest.NewServer(SetupRouter())
    defer server.Close()

    // Create WebSocket URL based on the server's address
    wsURL := "ws" + server.URL[4:] + "/ws"
    ws, _, err := websocket.DefaultDialer.Dial(wsURL, nil)
    if err != nil {
        t.Fatalf("Failed to connect to WebSocket: %v", err)
    }
    defer ws.Close()

    // Simulate WAV data stream by sending test data chunks
    wavData := []byte("simulated wav data")
    for i := 0; i < 5; i++ {
        if err := ws.WriteMessage(websocket.BinaryMessage, wavData); err != nil {
            t.Fatalf("Failed to send WAV data: %v", err)
        }
        time.Sleep(100 * time.Millisecond) // Mimic streaming delay
    }

    // Read and validate response as FLAC data
    for i := 0; i < 5; i++ {
        _, message, err := ws.ReadMessage()
        if err != nil {
            t.Fatalf("Failed to read FLAC data: %v", err)
        }

        if len(message) == 0 {
            t.Error("Expected FLAC data, but got empty message")
        }
    }
}

// TestWebSocketErrorHandling tests how WebSocket handles connection and streaming errors
func TestWebSocketErrorHandling(t *testing.T) {
    server := httptest.NewServer(SetupRouter())
    defer server.Close()

    wsURL := "ws" + server.URL[4:] + "/ws"
    ws, _, err := websocket.DefaultDialer.Dial(wsURL, nil)
    if err != nil {
        t.Fatalf("Failed to connect to WebSocket: %v", err)
    }
    defer ws.Close()

    // Simulate invalid data
    invalidData := []byte("invalid data")
    if err := ws.WriteMessage(websocket.BinaryMessage, invalidData); err != nil {
        t.Fatalf("Failed to send invalid data: %v", err)
    }

    // Attempt to read response; expect potential error or empty response
    _, message, err := ws.ReadMessage()
    if err == nil && len(message) != 0 {
        t.Error("Expected error or no data for invalid input, but got a response")
    }
}
