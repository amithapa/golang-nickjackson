package handlers

import (
	"net/http"
	"strconv"

	"../data"
	"github.com/gorilla/mux"
)

// swagger:route DELETE /products/{id} products deleteProduct
//Delete's a product from the database
// responses:
//   201: noContentResponse

// DeleteProduct deletes a  product from the database
func (p *Products) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])

	if err != nil {
		http.Error(w, "Unable to convert id", http.StatusBadRequest)
	}

	p.l.Println("Handle DELETE Products", id)

	err = data.DeleteProduct(id)
	if err == data.ErrProductNotFound {
		p.l.Println("[ERROR] deleting record id does not exist")

		w.WriteHeader(http.StatusNotFound)
		// data.ToJSON(w)
		return
	}

	if err != nil {
		p.l.Println("[ERROR] deleting record", err)

		w.WriteHeader(http.StatusInternalServerError)
		// data.ToJSON(w)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
