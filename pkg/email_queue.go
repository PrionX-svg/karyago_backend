package pkg

import (
	"fmt"
	"log"
	"time"
)

// Struktur untuk tiap job email
type EmailJob struct {
	To      string
	Subject string
	Body    string
}

// Channel antrian FIFO
var emailQueue = make(chan EmailJob, 100)

// Fungsi enqueue job
func EnqueueEmail(job EmailJob) {
	select {
	case emailQueue <- job:
		fmt.Printf("[QUEUE] ✅ Enqueued email to: %s (subject: %s)\n", job.To, job.Subject) // 🧠 LOG
	default:
		fmt.Printf("[QUEUE] ⚠️ Queue full, dropping email to: %s\n", job.To) // 🧠 LOG
	}
}

// Jalankan worker background (biasanya dipanggil dari main.go)
func StartEmailWorker() {
	go func() {
		fmt.Println("[WORKER] 🏁 Email worker started and waiting for jobs...") // 🧠 LOG

		for job := range emailQueue {
			fmt.Printf("[WORKER] 📦 Processing email job for: %s\n", job.To) // 🧠 LOG

			start := time.Now()
			err := SendEmail(job.To, job.Subject, job.Body)

			if err != nil {
				log.Printf("[WORKER] ❌ Failed to send email to %s: %v\n", job.To, err)
			} else {
				fmt.Printf("[WORKER] ✅ Successfully sent email to: %s (took %v)\n", job.To, time.Since(start))
			}

			// Delay kecil supaya tidak hit API SMTP terlalu cepat
			time.Sleep(500 * time.Millisecond)
		}
	}()
}
