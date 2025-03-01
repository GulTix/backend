package volunteers

import (
	"backend/internal/entity"
	"context"
	"log"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

func (r *RepositoryImpl) CreateDevVolunteer(ctx context.Context) error {
	checkQuery := `SELECT * FROM volunteers WHERE username = $1`

	volunteers := make([]entity.Volunteer, 0)

	volunteers = append(
		volunteers,
		entity.Volunteer{
			Id:       uuid.NewString(),
			Username: "avei",
			Email:    "muhammadabdulazizalghofari@gmail.com",
			Role:     "ADMIN",
			Deleted:  false,
		},
		entity.Volunteer{
			Id:       uuid.NewString(),
			Username: "jason",
			Email:    "jasonjeremyw@gmail.com",
			Role:     "ADMIN",
			Deleted:  false,
		},
		entity.Volunteer{
			Id:       uuid.NewString(),
			Username: "razan",
			Email:    "razanfawwaz1905@gmail.com",
			Role:     "ADMIN",
			Deleted:  false,
		},
		entity.Volunteer{
			Id: uuid.NewString(),
			Username: "jeeehaan",
			Email: "jessyhanifiah@gmail.com",
			Role: "ADMIN",
			Deleted: false,
		}
	)

	query := `INSERT INTO volunteers (id, username, email, role, deleted) VALUES
		($1, $2, $3, $4, $5);`

	for _, volunteer := range volunteers {

		rows, err := r.db.Query(ctx, checkQuery, volunteer.Username)

		if err != nil {
			return err
		}

		defer rows.Close()

		user, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[entity.Volunteer])

		if err != nil {
			if err == pgx.ErrNoRows {
				// Create User
				_, err := r.db.Exec(ctx, query, volunteer.Id, volunteer.Username, volunteer.Email, volunteer.Role, volunteer.Deleted)
				if err != nil {
					return err
				}
				continue
			}
			return nil
		}

		if user.Deleted == false {
			log.Println("User already exists")
			continue
		}
	}

	return nil
}
