package product

import (
	"fmt"
	"net/http"

	"github.com/Siddhant6674/ECOM/types"
	"github.com/Siddhant6674/ECOM/utils"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
)

type Handler struct {
	store types.ProductStore
}

// factor function which create instance to Handler struct
func NewHandler(store types.ProductStore) *Handler {
	return &Handler{
		store: store,
	}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/Product", h.handelGetProduct).Methods("GET")
	router.HandleFunc("/Product", h.handelCreateProduct).Methods("POST")
}

func (h *Handler) handelGetProduct(w http.ResponseWriter, r *http.Request) {

	ps, err := h.store.GetProducts()
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}
	utils.WriteJSON(w, http.StatusOK, ps)

}

func (h *Handler) handelCreateProduct(w http.ResponseWriter, r *http.Request) {
	var payload types.ProductPayload

	if err := utils.ParseJSON(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
	}
	//validating payload
	if err := utils.Validate.Struct(payload); err != nil {
		errors := err.(validator.ValidationErrors)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid payload %v", errors))
	}

	err := h.store.CreateProduct(types.Product{
		Name:        payload.Name,
		Description: payload.Description,
		Image:       payload.Image,
		Price:       payload.Price,
		Quantity:    payload.Quantity,
	})
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}
	utils.WriteJSON(w, http.StatusCreated, nil)

}
