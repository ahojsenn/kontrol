package handler

import (
	"fmt"
	"net/http"
	"sort"

		"github.com/ahojsenn/kontrol/booking"
	"github.com/ahojsenn/kontrol/util"
	"github.com/ahojsenn/kontrol/accountSystem"
)

func MakeGetCollectiveAccountHandler(repository accountSystem.AccountSystem) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		a := repository.GetCollectiveAccount()
		a.UpdateSaldo()
		w.Header().Set("Content-Type", "application/json")
		sort.Sort(booking.ByMonth(a.Bookings))
		json := util.Json(a)
		fmt.Fprintf(w, json)
	}
}