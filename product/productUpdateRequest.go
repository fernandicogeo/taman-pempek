package product

import "mime/multipart"

type ProductUpdateRequest struct {
	CategoryID  int                  `form:"category_id"`
	Name        string               `form:"name"`
	Image       multipart.FileHeader `form:"image"`
	Description string               `form:"description"`
	Price       int                  `form:"price"`
	Stock       int                  `form:"stock"`
}
