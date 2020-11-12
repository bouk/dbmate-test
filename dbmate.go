package dbmate

import (
	"context"
	"dbmate/models"
	_ "github.com/amacneil/dbmate"
	_ "github.com/volatiletech/sqlboiler/v4"
	_ "github.com/volatiletech/sqlboiler/v4/drivers/sqlboiler-psql"
)

func ok() {
	ctx := context.Background()
	var u models.User
	u.Insert(ctx)
}
