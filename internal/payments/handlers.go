package payments

import (
	"net/http"

	"github.com/Just-Mike4/go-ecom/internal/json"
)

type handler struct{}

func NewHandler() *handler {
	return &handler{}
}

// Payment handlers
func (h *handler) ListPayments(w http.ResponseWriter, r *http.Request) {
	json.WriteJSON(w, http.StatusOK, map[string]string{"message": "List payments endpoint"})
}

func (h *handler) InitializePayment(w http.ResponseWriter, r *http.Request) {
	json.WriteJSON(w, http.StatusOK, map[string]string{"message": "Initialize payment endpoint"})
}

func (h *handler) VerifyPayment(w http.ResponseWriter, r *http.Request) {
	json.WriteJSON(w, http.StatusOK, map[string]string{"message": "Verify payment endpoint"})
}

func (h *handler) PaymentWebhook(w http.ResponseWriter, r *http.Request) {
	json.WriteJSON(w, http.StatusOK, map[string]string{"message": "Payment webhook endpoint"})
}

func (h *handler) GetPayment(w http.ResponseWriter, r *http.Request) {
	json.WriteJSON(w, http.StatusOK, map[string]string{"message": "Get payment endpoint"})
}

func (h *handler) RefundPayment(w http.ResponseWriter, r *http.Request) {
	json.WriteJSON(w, http.StatusOK, map[string]string{"message": "Refund payment endpoint"})
}
