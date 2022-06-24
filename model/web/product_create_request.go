package web

type ProductCreateRequestWeb struct {
	IdUser     int    `json:"id_user,omitempty"`
	NamaProduk string `validate:"required,min=6,max=10" json:"nama_produk,omitempty"`
	Harga      int    `validate:"required,numeric,gte=0" json:"harga,omitempty"`
	Kategori   string `validate:"required" json:"kategori,omitempty"`
	Deskripsi  string `json:"deskripsi,omitempty"`
	Stok       int    `validate:"required,numeric,gte=0" json:"stok,omitempty"`
}
