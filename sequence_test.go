package main

import (
	"context"
	"github.com/jackc/pgx/v4"
	"testing"
	"time"
)

func TestSequence(t *testing.T) {
	conn := Connect()
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	for _, level := range IsoLevels {
		t.Run("TestUncommittedSequence-"+string(level), func(t *testing.T) {
			IDs := make([]uint, 0, 3)
			for i := 0; i < 3; i++ {
				tx, err := conn.BeginTx(ctx, pgx.TxOptions{
					IsoLevel: level,
				})

				if err != nil {
					t.Fatalf("starting transaction error: %v", err)
				}

				row := tx.QueryRow(ctx, "INSERT INTO books(author, name) VALUES ('А. С. Пушкин', 'Руслан и Людмила') RETURNING id")
				var id uint

				scanErr := row.Scan(&id)
				err = tx.Rollback(ctx)

				if err != nil {
					t.Fatalf("commiting transaction error: %v", err)
				}

				if scanErr != nil {
					t.Fatalf("getting row id error: %v", scanErr)
				}

				IDs = append(IDs, id)
			}

			t.Logf("sequence: %v", IDs)
			t.Logf("sequence is %s", idChangeType(IDs))
		})
	}
}

func idChangeType(IDs []uint) string {
	if len(IDs) < 2 {
		return "unknown"
	}

	growing := true
	constant := true
	for i := 1; i < len(IDs); i++ {
		// если меняются значения
		if IDs[i-1] != IDs[i] {
			constant = false
		}
		// если предыдущее значение больше или такое же как текущее
		if IDs[i-1] >= IDs[i] {
			growing = false
		}
	}

	switch {
	case growing:
		return "growing"
	case constant:
		return "constant"
	default:
		return "unknown"
	}
}
