package cart

import (
	"fmt"
	"net/http"

	"github.com/Siddhant6674/ECOM/service/auth"
	"github.com/Siddhant6674/ECOM/types"
	"github.com/Siddhant6674/ECOM/utils"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
)

type Handler struct {
	store        types.OrderStore
	ProductStore types.ProductStore
	userStore    types.UserStore
}

func NewHandler(store types.OrderStore, ProductStore types.ProductStore, userStore types.UserStore) *Handler {
	return &Handler{
		store:        store,
		ProductStore: ProductStore,
		userStore:    userStore,
	}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/cart/checkout", auth.WithJWTAuth(h.handelCheckout, h.userStore)).Methods("POST")
}

func (h *Handler) handelCheckout(w http.ResponseWriter, r *http.Request) {
	userID := auth.GetUserIDFromContext(r.Context())
	var cart types.CartCheckoutPayload
	if err := utils.ParseJSON(r, &cart); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	if err := utils.Validate.Struct(cart); err != nil {
		errors := err.(validator.ValidationErrors)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid payload %v", errors))
		return
	}

	ProductIDs, err := getCartItemsIDs(cart.Items)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}
	//get products
	ps, err := h.ProductStore.GetProductsByIDs(ProductIDs)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}
	orderID, totalPrice, err := h.createOrder(ps, cart.Items, userID)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}
	utils.WriteJSON(w, http.StatusOK, map[string]any{
		"total_price": totalPrice,
		"order_id":    orderID,
	})
}
