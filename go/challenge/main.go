package main

import (
	"context"
	"errors"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// ==============================================================================
// LEVEL 1: STRUCTS, TAGS, & METHODS
// ==============================================================================

// TODO 1: Define the SensorReading struct.
// Requirements:
// - Field "ID" (string) -> mapped to JSON "sensor_id"
// - Field "Value" (float64) -> mapped to JSON "value"
// - Field "Timestamp" (time.Time) -> (no json tag needed)
type SensorReading struct {
	ID        string  `json:"sensor_id"`
	Value     float64 `json:"value"`
	Timestamp time.Time
}

// TODO 2: Implement the fmt.Stringer interface.
// The output should be strictly in this format: "SENSOR[<ID>]: <Value>"
// Hint: Use Sprintf. Receiver should be a pointer to match the rest of the app.
func (r *SensorReading) String() string {
	return fmt.Sprintf("SENSOR[%s]: %.5f", r.ID, r.Value)
}

// ==============================================================================
// LEVEL 2: THREAD-SAFE STORAGE (MUTEX, MAPS, SLICES)
// ==============================================================================

type SafeDataStore struct {
	// TODO 3: Add a RWMutex field here (embedded or named).
	rwmu sync.RWMutex
	// TODO 4: Add a map field named 'data' where Key=string (SensorID) and Value=[]float64 (Slice of readings)
	data map[string][]float64
}

// NewDataStore initializes the map.
// (Go maps must be initialized with 'make' before use!)
func NewDataStore() *SafeDataStore {
	return &SafeDataStore{
		// TODO 5: Initialize the map here using make(map[...])
		data: make(map[string][]float64),
	}
}

// Add safely appends a reading to the store.
func (ds *SafeDataStore) Add(id string, val float64) {
	// TODO 6: Lock the mutex for writing.
	ds.rwmu.Lock()
	// TODO 7: Defer the unlock. (Critical for interviews!)
	defer ds.rwmu.Unlock()
	// TODO 8: Append the value to the slice in the map.
	// Hint: You do not need to check if the key exists. Append handles nil slices automatically.
	ds.data[id] = append(ds.data[id], val)
}

// GetAverage safely calculates the average for a sensor.
func (ds *SafeDataStore) GetAverage(id string) float64 {
	var output float64
	// TODO 9: Lock the mutex for READING (RLock).
	ds.rwmu.RLock()
	// TODO 10: Defer the RUnlock.
	defer ds.rwmu.RUnlock()
	// TODO 11: Retrieve the slice. If it doesn't exist or len is 0, return 0.
	values, ok := ds.data[id]
	if !ok || len(values) == 0 {
		return output
	}
	// TODO 12: Calculate and return average.
	for _, v := range values {
		output += v
	}
	return output / float64(len(values))
}

// ==============================================================================
// LEVEL 3: INTERFACES, TYPE SWITCHES, & CUSTOM ERRORS
// ==============================================================================

// TODO 13: Define a custom error struct named 'ErrInvalidReading'.
// It should have a field 'Msg' (string).
type ErrInvalidReading struct {
	Msg string
}

// TODO 14: Implement the 'error' interface for ErrInvalidReading.
func (e ErrInvalidReading) Error() string {
	return e.Msg
}

// Validate checks the data type and value.
// Input is interface{} (Empty Interface) -> can be anything.
func Validate(data interface{}) (*SensorReading, error) {
	// TODO 15: Use a "Type Switch" (switch v := data.(type))
	// Case *SensorReading:
	//    If v.Value is < -100 or > 100, return nil and ErrInvalidReading{Msg: "Unrealistic Temp"}.
	//    Otherwise return v, nil.
	// Case default:
	//    Return nil, and a standard error (errors.New or fmt.Errorf) saying "unknown type".
	//
	// Note: The return type is *SensorReading (a specific pointer), so returning 'nil' for the struct
	// is safe and idiomatic when an error occurs.
	switch v := data.(type) {
	case *SensorReading:
		if v.Value < -100 || v.Value > 100 {
			return nil, ErrInvalidReading{Msg: "Unrealistic Temp"}
		} else {
			return v, nil
		}
	default:
		return nil, errors.New("unknown type")
	}
}

// ==============================================================================
// LEVEL 4: CONCURRENCY (GOROUTINES, CHANNELS, CONTEXT)
// ==============================================================================

func main() {
	// Setup (Go 1.20+ seeds global random automatically, no need for rand.Seed)
	store := NewDataStore()
	var wg sync.WaitGroup

	// TODO 16: Create a Context with a Timeout of 2 seconds.
	// Use context.WithTimeout.
	// defer the cancel function.
	// ctx, cancel := ...
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	// TODO 17: Create a buffered channel of *SensorReading with capacity 10.
	// readingsChan := ...
	readingsChan := make(chan *SensorReading, 10)
	// --- GENERATOR ---
	// This goroutine simulates sensors pushing data.
	wg.Add(1)
	go func() {
		defer wg.Done()

		sensors := []string{"S1", "S2", "S3"}

		for {
			// Create a mock reading
			id := sensors[rand.Intn(len(sensors))]
			val := rand.Float64()*200 - 50 // Range -50 to 150
			// reading := &SensorReading{...}
			reading := &SensorReading{
				ID:    id,
				Value: val,
			}
			// TODO 18: Use a Select Statement to handle sending safely.
			// Critical: If you simply send to the channel without checking ctx.Done(),
			// this goroutine might block forever if workers exit early (deadlock).
			//
			// Pattern:
			// select {
			// case <-ctx.Done():
			//     return // Stop generating
			// case readingsChan <- reading:
			//     time.Sleep(50 * time.Millisecond) // Throttle
			// }
			select {
			case <-ctx.Done():
				return
			case readingsChan <- reading:
				time.Sleep(50 * time.Millisecond)
			}
		}
	}()

	// --- WORKER POOL ---
	// We will launch 3 workers to process the data concurrently.
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func(workerID int) {
			defer wg.Done()
			// TODO 19: Use a for-select loop to process data until context is done.
			for {
				select {
				// TODO 20: Handle Context Cancellation (return)
				case <-ctx.Done():
					return
					// TODO 21: Handle reading from channel inside the select.
					// Syntax: case reading := <-readingsChan:
					//    1. Call Validate(reading).
					//    2. If error, fmt.Printf("Worker %d: dropped bad data: %v\n", workerID, err).
					//       (Continue to next loop iteration).
					//    3. If valid, store.Add(reading.ID, reading.Value).
				case reading := <-readingsChan:
					sr, err := Validate(reading)
					if err != nil {
						fmt.Printf("Worker %d: dropped bad data: %v\n", workerID, err)
					} else {
						store.Add(sr.ID, sr.Value)
					}
				}
			}
		}(i)
	}

	// Wait for the context to timeout and workers to finish
	fmt.Println("System running... waiting for timeout...")
	wg.Wait()

	// Report
	fmt.Println("\n--- FINAL REPORT ---")
	for _, id := range []string{"S1", "S2", "S3"} {
		avg := store.GetAverage(id)
		fmt.Printf("Sensor %s Average: %.2f\n", id, avg)
	}
}
