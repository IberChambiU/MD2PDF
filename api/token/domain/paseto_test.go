package domain_test

import (
	"log"
	"testing"
	"time"

	d "github.com/IberChambiU/MD2PDF/api/token/domain"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func TestPasetoMaker(t *testing.T) {

	// maker, err := d.NewPasetoMaker("12345678901234567890123456789012")
	maker, err := d.NewPasetoMaker("ComprobemosSiAlguienPuedeLeerlos")

	require.NoError(t, err)

	username := "admin"
	duration := time.Minute

	issuedAt := time.Now()
	expiredAt := issuedAt.Add(duration)

	token, err := maker.CreateToken(username, uuid.New(), duration)
	require.NoError(t, err)
	require.NotEmpty(t, token)

	payload, err := maker.ValidateToken(token)
	require.NoError(t, err)
	require.NotEmpty(t, payload)

	require.NotZero(t, payload.Id)
	require.Equal(t, username, payload.Username)
	require.WithinDuration(t, issuedAt, payload.IssuedAt, time.Second)
	require.WithinDuration(t, expiredAt, payload.ExpiredAt, time.Second)
}

func TestValidToken(t *testing.T) {

	// maker, err := d.NewPasetoMaker("12345678901234567890123456789012")
	maker, err := d.NewPasetoMaker("ComprobemosSiAlguienPuedeLeerlos")

	require.NoError(t, err)

	username := "admin"

	token := "v2.local.14mUOYbfidri8qPB7lgP10k6g4IRgoZIr1sQHfXeChvLlRWhw_zkIntZNUXWkr_b_wPgEMEHTFbXEEvZKsCgmbF7PxHqkxiHBGk8GtxzfrMKpAuT_y-ggJJuLdiDTmxEadj7cYvgXJ_EneJnwJ3iDhurbDyRcg49xPkoLUYc4Z-_zglTUG9XBBBLl8wi0HkOOaDqYoT_C9nBYKXyoFQ9GZ4ZIHHWooGV1ei6W2O1ptcGghgmtUunPM2ct0azO1BmGEpLRzmf.bnVsbA"
	require.NotEmpty(t, token)

	payload, err := maker.ValidateToken(token)
	require.NoError(t, err)
	require.NotEmpty(t, payload)

	require.NotZero(t, payload.Id)
	require.Equal(t, username, payload.Username)
}

func TestPasetoToken(t *testing.T) {

	// maker, err := d.NewPasetoMaker("12345678901234567890123456789012")
	// maker, err := d.NewPasetoMaker("ComprobemosSiAlguienPuedeLeerlos")
	maker, err := d.NewPasetoMaker("00829440655557043386732036027194")

	require.NoError(t, err)

	username := "admin"
	duration := time.Now().AddDate(0, 1, 0).Sub(time.Now()) // 1 month
	// duration := time.Now().AddDate(1, 0, 0).Sub(time.Now()) // 1 year

	log.Printf("duration: %s", duration.String())

	token, err := maker.CreateToken(username, uuid.MustParse("a8fba77e-c2f0-484b-b976-a81da82c7393"), duration)
	require.NoError(t, err)
	require.NotEmpty(t, token)

	log.Printf("token: %s", token)
}

func TestExpiredPasetoToken(t *testing.T) {

	// maker, err := d.NewPasetoMaker("12345678901234567890123456789012")
	maker, err := d.NewPasetoMaker("ComprobemosSiAlguienPuedeLeerlos")

	require.NoError(t, err)

	username := "admin"
	duration := -time.Minute

	token, err := maker.CreateToken(username, uuid.MustParse("a8fba77e-c2f0-484b-b976-a81da82c7393"), duration)
	require.NoError(t, err)
	require.NotEmpty(t, token)

	payload, err := maker.ValidateToken(token)
	require.Error(t, err)
	require.EqualError(t, err, "token expirado o invalido")
	require.Nil(t, payload)
}
