package webutils

import (
	"net/http"
	"strconv"
)

func HanleParams(w http.ResponseWriter, r *http.Request, one, two string) (param1, param2 int) {
	r.ParseForm()
	if len(r.Form[one]) > 0 && len(r.Form[two]) > 0 {
		param1, _ = strconv.Atoi(r.Form[one][0])
		param2, _ = strconv.Atoi(r.Form[two][0])
	}

	return

}
