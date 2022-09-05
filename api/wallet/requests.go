package api

import validation "github.com/go-ozzo/ozzo-validation/v4"

type CreateRequest struct {
	DocumentNumber string `json:"document_number"`
}

func (c CreateRequest) Validate() error {
	return validation.ValidateStruct(&c,
		validation.Field(
			&c.DocumentNumber, validation.Required,
			validation.Length(11, 14).Error("must be a cpf or a cnpj"),
			// Other useful documentation checks omitted.
		),
	)
}

type OperationRequest struct {
	Amount      int64  `json:"amount"`
	Description string `json:"description,omitempty"`
}

func (o OperationRequest) Validate() error {
	return validation.ValidateStruct(&o,
		validation.Field(&o.Amount, validation.Required, validation.Min(1).Error("must be positive")),
	)
}
