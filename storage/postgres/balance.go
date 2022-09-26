package postgres

import (
	"fmt"
	"github.com/fokurly/avito-tech-test-task/models"
)

func (d *Db) IncreaseBalance(client models.Client) (*float64, error) {
	const (
		getClientQuery     = `SELECT * FROM Client WHERE id=$1;`
		updateBalanceQuery = `UPDATE Client SET balance=$1 WHERE id=$2`
	)

	balance, err := d.GetBalance(client.Id)
	if err != nil && err.Error() == "[GetBalance] - no rows with such id" {
		err := d.InsertClientBalance(client)
		if err != nil {
			return nil, err
		}
		return &client.Money, nil
	}

	if err != nil {
		return nil, err
	}

	*balance += client.Money
	_, err = d.db.Exec(updateBalanceQuery, *balance, client.Id)
	if err != nil {
		return nil, fmt.Errorf("[IncreaseBalance] - could not update balance. error: %v", err)
	}
	return balance, nil
}

func (d *Db) DecreaseBalance(client models.Client) (*float64, error) {
	const (
		update = `UPDATE Client SET balance=$1 WHERE id=$2`
	)
	currentBalance, err := d.GetBalance(client.Id)
	if err != nil {
		return nil, err
	}

	if *currentBalance < client.Money {
		return nil, fmt.Errorf("[DecreaseBalance] - current balance is lower than money you want to withdraw from balance")
	}

	*currentBalance -= client.Money
	_, err = d.db.Exec(update, *currentBalance, client.Id)
	if err != nil {
		return nil, fmt.Errorf("[DecreaseBalance] - could not update balance. error: %v", err)
	}
	return currentBalance, nil
}

func (d *Db) GetBalance(clientId int64) (*float64, error) {
	const (
		getBalanceQuery = `SELECT balance FROM Client WHERE id=$1`
	)
	rows, err := d.db.Query(getBalanceQuery, clientId)
	if err != nil {
		return nil, fmt.Errorf("[GetBalance] - could not exec query. error: %v", err)
	}

	defer rows.Close()
	var balance float64

	if rows.Next() {
		err := rows.Scan(&balance)
		if err != nil {
			return nil, fmt.Errorf("[GetBalance] - could not scan rows. error: %v", err)
		}
	} else {
		return nil, fmt.Errorf("[GetBalance] - no rows with such id")
	}
	return &balance, nil
}

func (d *Db) InsertClientBalance(client models.Client) error {
	const (
		insertClientQuery = `INSERT INTO Client(id, balance) VALUES($1, $2)`
	)
	_, err := d.db.Exec(insertClientQuery, client.Id, client.Money)
	if err != nil {
		return fmt.Errorf("[InsertClientBalance] - could not exec query. error: %v", err)
	}

	return nil
}

func (d *Db) GetAllClients() ([]models.Client, error) {
	const (
		getAllClientsQuery = `SELECT * 
							  FROM Client`
	)

	rows, err := d.db.Query(getAllClientsQuery)
	if err != nil {
		return nil, fmt.Errorf("[GetAllClients] - could not exec query. error: %v", err)
	}

	defer func() {
		_ = rows.Close()
	}()

	clients := make([]models.Client, 0)

	for rows.Next() {
		client := models.Client{}

		err = rows.Scan(&client.Id, &client.Money)
		if err != nil {
			return nil, fmt.Errorf("[GetAllClients] - could not scan row. error: %v", err)
		}
		clients = append(clients, client)
	}
	return clients, nil
}
