package carts

import (
	"net/http"

	"github.com/Just-Mike4/ecom-go/internal/json"
)

type handler struct{}

func NewHandler() *handler {
	return &handler{}
}

// Cart handlers
func (h *handler) GetCart(w http.ResponseWriter, r *http.Request) {
	json.WriteJSON(w, http.StatusOK, map[string]string{"message": "Get cart endpoint"})
}

func (h *handler) AddCartItem(w http.ResponseWriter, r *http.Request) {
	json.WriteJSON(w, http.StatusOK, map[string]string{"message": "Add cart item endpoint"})
}

func (h *handler) UpdateCartItem(w http.ResponseWriter, r *http.Request) {
	json.WriteJSON(w, http.StatusOK, map[string]string{"message": "Update cart item endpoint"})
}

func (h *handler) DeleteCartItem(w http.ResponseWriter, r *http.Request) {
	json.WriteJSON(w, http.StatusOK, map[string]string{"message": "Delete cart item endpoint"})
}

func (h *handler) ClearCart(w http.ResponseWriter, r *http.Request) {
	json.WriteJSON(w, http.StatusOK, map[string]string{"message": "Clear cart endpoint"})
}

func (h *handler) SyncCart(w http.ResponseWriter, r *http.Request) {
	json.WriteJSON(w, http.StatusOK, map[string]string{"message": "Sync cart endpoint"})
}

// Wishlist handlers
func (h *handler) GetWishlist(w http.ResponseWriter, r *http.Request) {
	json.WriteJSON(w, http.StatusOK, map[string]string{"message": "Get wishlist endpoint"})
}

func (h *handler) AddWishlistItem(w http.ResponseWriter, r *http.Request) {
	json.WriteJSON(w, http.StatusOK, map[string]string{"message": "Add wishlist item endpoint"})
}

func (h *handler) DeleteWishlistItem(w http.ResponseWriter, r *http.Request) {
	json.WriteJSON(w, http.StatusOK, map[string]string{"message": "Delete wishlist item endpoint"})
}
