package handlers

import (
	"errors"
	"net/http"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/hs-zavet/comtools/httpkit"
	"github.com/hs-zavet/comtools/httpkit/problems"
	"github.com/hs-zavet/news-radar/internal/api/rest/requests"
	"github.com/hs-zavet/news-radar/internal/api/rest/responses"
	"github.com/hs-zavet/news-radar/internal/app"
	"github.com/hs-zavet/news-radar/internal/app/ape"
	"github.com/hs-zavet/news-radar/internal/enums"
	"github.com/hs-zavet/tokens"
)

func (h *Handler) CreateTag(w http.ResponseWriter, r *http.Request) {
	user, err := tokens.GetAccountTokenData(r.Context())
	if err != nil {
		h.log.WithError(err).Error("Failed to retrieve account data")
		httpkit.RenderErr(w, problems.Unauthorized(err.Error()))
		return
	}

	req, err := requests.CreateTag(r)
	if err != nil {
		h.log.WithError(err).Warn("Error parsing request")
		httpkit.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	tagStatus, ok := enums.ParseTagStatus(req.Data.Attributes.Status)
	if !ok {
		h.log.WithError(err).Warn("Error parsing tag status")
		httpkit.RenderErr(w, problems.BadRequest(validation.Errors{
			"status": validation.NewError("invalid", "invalid tag status"),
		})...)
		return
	}

	tagType, ok := enums.ParseTagType(req.Data.Attributes.Type)
	if !ok {
		h.log.WithError(err).Warn("Error parsing tag type")
		httpkit.RenderErr(w, problems.BadRequest(validation.Errors{
			"type": validation.NewError("invalid", "invalid tag type"),
		})...)
		return
	}

	tag, err := h.app.CreateTag(r.Context(), app.CreateTagRequest{
		Name:   req.Data.Attributes.Name,
		Type:   tagType,
		Status: tagStatus,
		Color:  req.Data.Attributes.Color,
		Icon:   req.Data.Attributes.Icon,
	})

	if err != nil {
		switch {
		case errors.Is(err, ape.ErrorTagNameAlreadyTaken):
			httpkit.RenderErr(w, problems.Conflict())
		default:
			httpkit.RenderErr(w, problems.InternalError())
		}
		h.log.WithError(err).Error("Failed to create tag")
		return
	}

	h.log.Infof("Created tag %s, by: %s", req.Data.Attributes.Name, user.AccountID)

	httpkit.Render(w, responses.Tag(tag))
}
