package balancers

import (
	"database/sql"
	"fmt"
	"github.com/KolmaginDanil/Lab3/server/tools"
)

type Balancer struct {
	Id              int   `json:"id"`
	UsedMachines    []int `json:"usedMachines"`
	NotUsedMachines []int `json:"notUsedMachines"`
}

type MachineStatus struct {
	MachineId int
	IsWork    bool
}

type Store struct {
	Db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{Db: db}
}

func (s *Store) ListBalancers() ([]*Balancer, error) {
	rows, err := s.Db.Query(`SELECT l.all, l.used, l.notUsed FROM (
	SELECT CASE WHEN a.id is NOT NULL THEN a.id ELSE b.id END as all, a.used as used, b.notUsed as notUsed FROM (
		SELECT balancer_id AS "id",
       		array_agg(machine_id) AS used
		FROM ConnectToBalancers, Machines
		WHERE ConnectToBalancers.machine_id = Machines.id AND Machines.isUsed = true
		GROUP BY balancer_id
		) as a
		full join (
		SELECT balancer_id AS "id", array_agg(machine_id) AS notUsed
        		FROM ConnectToBalancers, Machines
        		WHERE ConnectToBalancers.machine_id = Machines.id AND Machines.isUsed = false
        		GROUP BY balancer_id) as b on a.id = b.id) as l ORDER BY l.all`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []*Balancer
	for rows.Next() {
		var blc Balancer
		var UsedMachinesInASCII, NotUsedMachinesInASCII []uint8

		rowErr := rows.Scan(&blc.Id, &UsedMachinesInASCII, &NotUsedMachinesInASCII)
		if rowErr != nil {
			return nil, rowErr
		}

		var convertErr error
		blc.UsedMachines, convertErr = tools.ASCIItoIntArr(UsedMachinesInASCII)
		if convertErr != nil {
			return nil, convertErr
		}
		blc.NotUsedMachines, convertErr = tools.ASCIItoIntArr(NotUsedMachinesInASCII)
		if convertErr != nil {
			return nil, convertErr
		}

		result = append(result, &blc)
	}
	if result == nil {
		result = make([]*Balancer, 0)
	}
	return result, nil
}

func (s *Store) ChangeStatus(ms MachineStatus) error {
	if ms.MachineId <= 0 {
		return fmt.Errorf("this balancers id cannot exist")
	}
	_, err := s.Db.Exec("UPDATE Machines SET isUsed = $1 WHERE id = $2", ms.IsWork, ms.MachineId)
	return err
}
