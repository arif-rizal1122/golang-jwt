package user

import (
	"database/sql"
	"fmt"
	"github.com/arif-rizal1122/golang-jwt/types"
)

// Store adalah struktur yang bertanggung jawab untuk menyimpan koneksi database
type Store struct {
	db *sql.DB
}

// NewStore adalah fungsi konstruktor untuk Store yang membuat instansiasi dari Store
// dengan menyediakan koneksi database.
func NewStore(db *sql.DB) *Store {
	return &Store{db: db}
}

// GetUserEmail adalah fungsi yang digunakan untuk mengambil informasi pengguna berdasarkan alamat email.
// Fungsi ini menjalankan query SQL untuk mencari pengguna dengan email yang diberikan.
// Jika pengguna ditemukan, informasi pengguna tersebut akan di-scan dari baris hasil query ke struktur types.User.
func (s *Store) GetUserEmail(email string) (*types.User, error) {
	rows, err := s.db.Query("select * from users where email = ?", email)
	// jika ada error maka kembalikan error dan data nil/kosong nya
	if err != nil {
		return nil, err
	}

	u := new(types.User)
	for rows.Next() {
		u, err = scanRowIntoUser(rows)
		if err != nil {
			return nil, err
		}
	}

	// Jika pengguna dengan email yang diberikan tidak ditemukan, kembalikan error "user not found"
	if u.ID == 0 {
		return nil, fmt.Errorf("user not found")
	}
	return u, nil
}

// scanRowIntoUser adalah fungsi bantu yang digunakan untuk memindahkan data dari baris hasil query ke struktur types.User.
func scanRowIntoUser(rows *sql.Rows) (*types.User, error) {
	user := new(types.User)
	err := rows.Scan(
		&user.ID,
		&user.FirstName,
		&user.LastName,
		&user.Email,
		&user.Password,
		&user.CreatedAt,
	)

	if err != nil {
		return nil, err
	}

	return user, nil
}


type mockUserStore struct{}

func GetUserByEmail(email string) (*types.User, error)  {
	return nil, nil
}