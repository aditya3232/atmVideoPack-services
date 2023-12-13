package jwt

import (
	"errors"
	"time"

	"github.com/aditya3232/atmVideoPack-services.git/config"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

var JWT_SECRET = []byte(config.CONFIG.JWT_KEY)

/*
	- generate jwt token
	- expired is in days
	- if expired = 0, token will not expire
*/

func GenerateToken(userID int, expired int) (string, error) {
	claims := jwt.MapClaims{
		"user": gin.H{
			"id": userID,
		},
	}

	/*
		- jika != 0, maka expired => waktu saat ini ditambah dengan jumlah hari, waktu dalam bentuk unix timestamp
		- dan jika expired bernilai 0, maka waktu kadaluarsa jwt adalah 100 tahun dari waktu saat ini
	*/
	if expired != 0 {
		claims["exp"] = time.Now().Add(time.Hour * 24 * time.Duration(expired)).Unix()
	} else {
		claims["exp"] = time.Now().Add(time.Hour * 24 * 365 * 100).Unix()
	}

	/*
		- jwt.NewWithClaims => membuat token baru dengan claims yang telah dibuat
		- jwt.SigningMethodHS256 => algoritma yang digunakan untuk membuat token
		- token.SignedString => membuat token yang telah ditandatangani dengan secret key
	*/
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString(JWT_SECRET)
	if err != nil {
		return signedToken, err
	}

	return signedToken, nil
}

/*
- validate token => berfungsi untuk memvalidasi token JWT yang sudah dienkripsi

  - func ValidateToken(encodedToken string) (*jwt.Token, error) {
    Ini adalah definisi fungsi ValidateToken yang menerima sebuah string yang berisi token JWT yang dienkripsi.
    Fungsi ini mengembalikan dua nilai: sebuah pointer ke objek jwt.Token dan sebuah nilai error jika terjadi kesalahan.

  - token, err := jwt.Parse(encodedToken, func(token *jwt.Token) (interface{}, error) {
    Fungsi jwt.Parse digunakan untuk mem-parsing token yang dienkripsi.
    Parameter pertama adalah token yang dienkripsi,
    dan parameter kedua adalah sebuah fungsi yang akan dijalankan selama proses parsing untuk mendapatkan kunci rahasia yang digunakan untuk memverifikasi tanda tangan digital pada token.

  - _, ok := token.Method.(*jwt.SigningMethodHMAC)
    Baris ini memeriksa apakah metode penandatanganan dari token adalah HMAC (HMAC-SHA-256).
    Jika tidak, itu berarti token tidak valid.
    Variabel ok akan bernilai true jika metode penandatanganan sesuai dengan yang diharapkan.

  - if !ok { return nil, errors.New("invalid token") }
    Jika metode penandatanganan tidak sesuai, fungsi ini mengembalikan nilai nil untuk objek jwt.Token dan sebuah kesalahan dengan pesan "invalid token".

  - return []byte(JWT_SECRET), nil
    Jika metode penandatanganan sesuai, fungsi ini mengembalikan kunci rahasia yang digunakan untuk memverifikasi tanda tangan digital pada token.
    Kunci rahasia diambil dari konstanta JWT_SECRET yang mungkin telah didefinisikan sebelumnya.

- Setelah selesai, jwt.Parse akan mengembalikan objek jwt.Token yang sudah diisi dengan informasi dari token JWT yang dienkripsi.

  - if err != nil { return token, err }
    Fungsi ini melakukan pemeriksaan terhadap kemungkinan kesalahan selama proses parsing.
    Jika terdapat kesalahan, fungsi akan mengembalikan nilai nil untuk objek jwt.Token dan nilai kesalahan (err).

  - return token, nil
    Jika tidak ada kesalahan, fungsi akan mengembalikan objek jwt.Token yang berisi informasi dari token JWT yang telah di-parse dan nilai kesalahan (nil).

  - Jadi, fungsi ini memvalidasi token JWT, memastikan bahwa metode penandatanganannya sesuai dengan yang diharapkan,
    dan mengembalikan objek jwt.Token jika token valid, bersama dengan nilai error jika terjadi kesalahan.
*/
func ValidateToken(encodedToken string) (*jwt.Token, error) {
	token, err := jwt.Parse(encodedToken, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errors.New("invalid token")
		}
		return []byte(JWT_SECRET), nil
	})
	if err != nil {
		return token, err
	}

	return token, nil

}

/*
- ini digunakan untuk memvalidasi apakah jwt token yang diberikan benar
- dipakai saat ada orang yg memalsukan token
*/
func GetUserIDFromToken(encodedToken string) (int, error) {
	// Parse the JWT token
	token, err := ValidateToken(encodedToken)
	if err != nil {
		return 0, err
	}

	// Extract the user ID from the "sub" claim
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		userID, ok := claims["user"].(map[string]interface{})["id"].(float64)
		if !ok {
			return 0, errors.New("invalid token claims")
		}

		return int(userID), nil
	} else {
		return 0, errors.New("invalid token claims")
	}
}
