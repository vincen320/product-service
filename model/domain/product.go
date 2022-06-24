package domain

type Product struct {
	Id           int    `json:"id,omitempty"`
	IdUser       int    `json:"id_user,omitempty"`
	NamaProduk   string `json:"nama_produk,omitempty"`
	Harga        int    `json:"harga,omitempty"`
	Kategori     string `json:"kategori,omitempty"`
	Deskripsi    string `json:"deskripsi,omitempty"`
	Stok         int    `json:"stok,omitempty"`
	LastModified int64  `json:"last_modified,omitempty"`
}
