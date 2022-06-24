package web

type ProductUpdatePatchWebRequest struct {
	Id         int    `json:"id,omitempty"`
	IdUser     int    `json:"id_user,omitempty"`
	NamaProduk string `validate:"min=6,max=10" json:"nama_produk,omitempty"`
	Harga      int    `validate:"numeric,gte=0" json:"harga,omitempty"`
	Kategori   string `json:"kategori,omitempty"`
	Deskripsi  string `json:"deskripsi,omitempty"`
	Stok       int    `validate:"numeric,gte=0" json:"stok,omitempty"`
}
