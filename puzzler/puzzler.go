package puzzler

import (
	"net/http"
	"slapchop/actions"
	"github.com/bndr/gopencils"
)


func CreatePuzzle(actions.CreateResponse) {
	puzzler := gopencils.Api("http://localhost:8000")
}
