package tests

import (
	"testing"

	"github.com/yuedun/ginode/db"
	"github.com/yuedun/ginode/pkg/shortUrl"
)

func TestGeneratorShortUrl(t *testing.T) {
	service := shortUrl.NewService(db.Mysql)
	long := shortUrl.Long2ShortRequest{OriginUrl: "https://www.yuedun.wang"}
	shor, err := service.Long2Short(&long)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(shor)
}
