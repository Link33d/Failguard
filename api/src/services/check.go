package services

import (
	"server/src/config"
	"server/src/models"
	"time"

	"github.com/lib/pq"
)

func SaveCheck(check *models.Check) error {

	now := time.Now()
	var id string

	err := config.DB.QueryRow(`
		INSERT INTO checks (
			user_parent, 
			service_name, 
			address_type, 
			addresses, 
			check_type, 
			port, 
			interval_seconds, 
			created_at, 
			updated_at
		) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
		RETURNING id`,
		check.UserParent,
		check.ServiceName,
		check.AddressType,
		pq.Array(check.Addresses),
		check.CheckType,
		check.Port,
		check.IntervalSeconds,
		now,
		now,
	).Scan(&id)

	if err != nil {
		return err
	}

	check.ID = id
	check.CreatedAt = now.Format(time.RFC3339)
	check.UpdatedAt = now.Format(time.RFC3339)

	return nil

}
