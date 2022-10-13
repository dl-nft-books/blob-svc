package responses

import "gitlab.com/tokend/nft-books/blob-svc/resources"

func NewKeyResponse(resourceKey string) resources.KeyResponseResponse {
	return resources.KeyResponseResponse{
		Data: resources.KeyResponse{
			Key: resources.NewKeyInt64(1, resources.S3KEY),
			Attributes: resources.KeyResponseAttributes{
				S3Key: resourceKey,
			},
		},
	}
}
