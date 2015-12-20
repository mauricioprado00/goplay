package shopify

type Products struct {
    Products  []Product `json:"products"`
}

type Product struct {
    Id          int         `json:"id"`
    Title       string      `json:"title"`
    Variants    []Variants  `json:"variants"`
    CreatedAt   string      `json:"created_at"`
    UpdatedAt   string      `json:"updated_at"`
    Hanle       string      `json:"handle"`
}

type Variants struct {
    Sku         string  `json:"sku"`
    Price       string  `json:"price"`
    Quantity    int     `json:"inventory_quantity"`
    Weight      float64 `json:"weight"`
    WeightUnit  string  `json:"weight_unit"`
}