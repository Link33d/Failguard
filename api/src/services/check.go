package services

import (
	"server/src/config"
	"server/src/models"
	"strings"
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

func GetChecks(checks *[]models.Check) error {

	rows, err := config.DB.Query(`
		SELECT id, user_parent, service_name, address_type, addresses, check_type, port, interval_seconds, created_at, updated_at 
		FROM checks
	`)
	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {
		var check models.Check
		var addressesRaw []byte

		if err := rows.Scan(
			&check.ID,
			&check.UserParent,
			&check.ServiceName,
			&check.AddressType,
			&addressesRaw,
			&check.CheckType,
			&check.Port,
			&check.IntervalSeconds,
			&check.CreatedAt,
			&check.UpdatedAt,
		); err != nil {
			return err
		}

		check.Addresses = strings.Split(string(addressesRaw), ",")

		*checks = append(*checks, check)
	}

	if err := rows.Err(); err != nil {
		return err
	}

	return nil
}
