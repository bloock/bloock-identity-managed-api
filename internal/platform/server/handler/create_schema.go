package handler

import (
	"bloock-identity-managed-api/internal/domain"
	"bloock-identity-managed-api/internal/services/create"
	"bloock-identity-managed-api/internal/services/create/request"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

type CreateSchemaRequest struct {
	Attributes  map[string]AttributeSchema `json:"attributes" binding:"required"`
	DisplayName string                     `json:"display_name" binding:"required"`
	SchemaType  string                     `json:"schema_type" binding:"required"`
	Version     string                     `json:"version" binding:"required"`
	Description string                     `json:"description" binding:"required"`
}

type AttributeSchema struct {
	Name        string `json:"name" binding:"required"`
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
		var req CreateSchemaRequest
		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, NewBadRequestAPIError(err.Error()))
			return
		}

		attrSchema := make([]request.AttributeSchema, 0)
		for id, attr := range req.Attributes {
			attrSchema = append(attrSchema, request.AttributeSchema{
				Id:          id,
				Name:        attr.Name,
				DataType:    attr.DataType,
				Description: attr.Description,
				Required:    attr.Required,
			})
		}
		sr := request.CreateSchemaRequest{
			DisplayName: req.DisplayName,
			SchemaType:  req.SchemaType,
			Version:     req.Version,
			Description: req.Description,
			Attributes:  attrSchema,
		}

		res, err := schema.Create(ctx, sr)
		if err != nil {
			if errors.Is(domain.ErrInvalidDataType, err) {
				ctx.JSON(http.StatusBadRequest, NewBadRequestAPIError(err.Error()))
				return
			}
			ctx.JSON(http.StatusInternalServerError, NewInternalServerAPIError(err.Error()))
			return
		}

		ctx.JSON(http.StatusCreated, mapToCreateSchemaResponse(res.(string)))
	}
}
