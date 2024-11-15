package handler

import (
	"encoding/json"
	"net/http"
	"strconv"
	"sync"

	"github.com/kevintamakuwala/retail-pulse/internal/models"
	"github.com/kevintamakuwala/retail-pulse/internal/processor"
	"github.com/kevintamakuwala/retail-pulse/internal/store"
)

// JobHandler manages job-related operations
type JobHandler struct {
	jobs     map[int]*models.Job
	jobID    int
	jobMutex sync.Mutex
}

// NewJobHandler creates a new JobHandler instance
func NewJobHandler() *JobHandler {
	return &JobHandler{
		jobs: make(map[int]*models.Job),
	}
}

// SubmitHandler handles the job submission endpoint
func (h *JobHandler) SubmitHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req models.SubmitRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Invalid request format"})
		return
	}

	if req.Count != len(req.Visits) {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Count mismatch with visits"})
		return
	}

	h.jobMutex.Lock()
	h.jobID++
	currentJobID := h.jobID
	job := &models.Job{
		JobID:  strconv.Itoa(currentJobID),
		Status: "ongoing",
	}
	h.jobs[currentJobID] = job
	h.jobMutex.Unlock()

	// Process job asynchronously
	go h.processJob(currentJobID, req.Visits)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]int{"job_id": currentJobID})
}

// StatusHandler handles the job status endpoint
func (h *JobHandler) StatusHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	jobIDStr := r.URL.Query().Get("jobid")
	if jobIDStr == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Missing job ID"})
		return
	}

	requestedJobID, err := strconv.Atoi(jobIDStr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Invalid job ID format"})
		return
	}

	h.jobMutex.Lock()
	job, exists := h.jobs[requestedJobID]
	h.jobMutex.Unlock()

	if !exists {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Job not found"})
		return
	}

	response := models.Job{
		Status: job.Status,
		JobID:  job.JobID,
	}

	if job.Status == "failed" {
		response.Errors = job.Errors
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

// processJob handles the actual job processing
func (h *JobHandler) processJob(jobID int, visits []models.Visit) {
	h.jobMutex.Lock()
	job := h.jobs[jobID]
	h.jobMutex.Unlock()

	// First validate all store IDs
	for _, visit := range visits {
		if !store.Exists(visit.StoreID) {
			h.jobMutex.Lock()
			job.Status = "failed"
			job.Errors = []models.Error{{
				StoreID: visit.StoreID,
				Message: "Store not found",
			}}
			h.jobMutex.Unlock()
			return
		}
	}

	// Then process all images
	for _, visit := range visits {
		for _, imageURL := range visit.ImageURLs {
			if _, err := processor.ProcessImage(imageURL); err != nil {
				h.jobMutex.Lock()
				job.Status = "failed"
				job.Errors = []models.Error{{
					StoreID: visit.StoreID,
					Message: "Failed to process image",
				}}
				h.jobMutex.Unlock()
				return
			}
		}
	}

	h.jobMutex.Lock()
	job.Status = "completed"
	h.jobMutex.Unlock()
}
