// This is a sumsup webhook client example

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/sirupsen/logrus"
)

type WebhookPayload struct {
	ApplicantId               string          `json:"applicantId"`
	InspectionId              string          `json:"inspectionId"`
	CorrelationId             string          `json:"correlationId"`
	LevelName                 string          `json:"levelName,omitempty"`
	ExternalUserId            string          `json:"externalUserId,omitempty"`
	Type                      string          `json:"type"`
	SandboxMode               bool            `json:"sandboxMode"`
	ReviewStatus              string          `json:"reviewStatus"`
	CreatedAtMs               string          `json:"createdAtMs"`
	ApplicantType             string          `json:"applicantType,omitempty"`
	ReviewResult              json.RawMessage `json:"reviewResult,omitempty"`
	ApplicantMemberOf         []string        `json:"applicantMemberOf,omitempty"`
	VideoIdentReviewStatus    string          `json:"videoIdentReviewStatus,omitempty"`
	ApplicantActionId         string          `json:"applicantActionId,omitempty"`
	ExternalApplicantActionId string          `json:"externalApplicantActionId,omitempty"`
	ClientId                  string          `json:"clientId,omitempty"`
}

func HandleWebhook(w http.ResponseWriter, r *http.Request) {
	// Parse the webhook payload from the request body
	var payload WebhookPayload
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		log.Println("failed to decode payload:", err)
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}

	fmt.Println(payload.Type)
	logrus.Printf("%+v", payload)

	// Handle the webhook event based on its type
	switch payload.Type {
	case "applicantReviewed":
		// Handle verification result
	case "applicantPending":
		// Handle pending applicant status
	case "applicantCreated":
		// Handle applicant creation
	// Add cases for other webhook types as needed
	default:
		log.Println("unknown webhook type:", payload.Type)
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}

	// Send a 200 OK response to the third-party service
	w.WriteHeader(http.StatusOK)
}
