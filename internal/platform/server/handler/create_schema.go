package handler

import (
	"bloock-identity-managed-api/internal/services/create"
	"bloock-identity-managed-api/internal/services/create/request"
	"github.com/gin-gonic/gin"
	"net/http"
)

type CreateSchemaRequest struct {
	Attributes  map[string]AttributeSchema `json:"attributes" binding:"required"`
	Title       string                     `json:"title" binding:"required"`
	SchemaType  string                     `json:"schema_type" binding:"required"`
	Version     string                     `json:"version" binding:"required"`
	Description string                     `json:"description" binding:"required"`
}

type AttributeSchema struct {
	Title       string `json:"title" binding:"required"`
	DataType    string `json:"data_type" binding:"required"`
	Description string `json:"description" binding:"required"`
	Required    bool   `json:"required" binding:"required"`
}

type CreateSchemaResponse struct {
	SchemaID string `json:"schema_id"`
}

func mapToCreateSchemaResponse(schemaID string) CreateSchemaResponse {
	return CreateSchemaResponse{
		SchemaID: schemaID,
	}
}

func CreateSchema(schema create.Schema) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		issuerDid := ctx.Param("issuer_did")
		if issuerDid == "" {
			ctx.JSON(http.StatusBadRequest, "empty issuer did")
			return
		}

		var req CreateSchemaRequest
		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, NewBadRequestAPIError(err.Error()))
			return
		}

		attrSchema := make([]request.AttributeSchema, 0)
		for key, attr := range req.Attributes {
			attrSchema = append(attrSchema, request.AttributeSchema{
				Key:         key,
				Title:       attr.Title,
				DataType:    attr.DataType,
				Description: attr.Description,
				Required:    attr.Required,
			})
		}
		sr := request.CreateSchemaRequest{
			IssuerDID:   issuerDid,
			Title:       req.Title,
			SchemaType:  req.SchemaType,
			Version:     req.Version,
			Description: req.Description,
			Attributes:  attrSchema,
		}

		res, err := schema.Create(ctx, sr)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, NewInternalServerAPIError(err.Error()))
			return
		}

		ctx.JSON(http.StatusCreated, mapToCreateIssuerResponse(res.(string)))
	}
}
