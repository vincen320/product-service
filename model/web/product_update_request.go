package web

type ProductUpdateWebRequest struct {
	Id         int    `json:"id,omitempty"`
	IdUser     int    `json:"id_user,omitempty"`
	NamaProduk string `validate:"required,min=6,max=10" json:"nama_produk"`
	Harga      int    `validate:"required,numeric,gte=0" json:"harga"`
	Kategori   string `validate:"required" json:"kategori"`
	Deskripsi  string `json:"deskripsi"`
	Stok       int    `validate:"required,numeric,gte=0" json:"stok"`
}
