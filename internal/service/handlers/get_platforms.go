package handlers

import (
	"net/http"

	"gitlab.com/distributed_lab/ape"
)

func GetPlatforms(w http.ResponseWriter, r *http.Request) {
	platforms := Platforms(r)

	ape.Render(w, platforms.Response)
}
