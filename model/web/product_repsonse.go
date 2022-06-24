package web

type ProductResponse struct {
	Id           int    `json:"id"`
	NamaProduk   string `json:"nama_produk"`
	Harga        int    `json:"harga"`
	Kategori     string `json:"kategori"`
	Deskripsi    string `json:"deskripsi"`
	Stok         int    `json:"stok"`
	LastModified int64  `json:"last_modified"`
}
