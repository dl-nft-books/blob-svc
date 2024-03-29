package handlers

import (
	"net/http"

	"github.com/google/uuid"
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"github.com/dl-nft-books/blob-svc/internal/service/helpers"
	"github.com/dl-nft-books/blob-svc/internal/service/requests"
	"github.com/dl-nft-books/blob-svc/internal/service/responses"
)

func CreateDocument(w http.ResponseWriter, r *http.Request) {
	key, document, header, err := requests.NewCreateDocumentRequest(r)
	if err != nil {
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	ext, err := helpers.CheckDocumentMimeType(header.Header.Get("Content-Type"), r)
	if err != nil {
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	awsConfig := helpers.AwsConfig(r)

	if key == "" {
		key = uuid.New().String()
	} else {
		// checking if key exists (only in case of custom keys, not uuid-generated)
		// to not overwrite the existing document

		exists, err := helpers.IsKeyExists(key+"."+ext, awsConfig)
		if err != nil || exists {
			helpers.Log(r).WithError(err).Debug("failed to check key existence or key was found")
			ape.RenderErr(w, problems.BadRequest(
				errors.New("Document with such key already exists or it cannot be checked"))...)
			return
		}
	}
	key += "." + ext

	err = helpers.UploadFile(document, key, awsConfig)
	if err != nil {
		helpers.Log(r).WithError(err).Debug("failed to upload file")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	ape.Render(w, responses.NewKeyResponse(key))
}
