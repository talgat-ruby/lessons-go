package db

import (
	"database/sql"
	"errors"
)

type EmployeeModel struct {
	ID       int
	Name     string
	Position string
	IsRobot  bool
}

func (m *Model) SelectEmployee(id int) (*EmployeeModel, error) {
	statement := "SELECT id, name, position, is_robot FROM personal WHERE id=$1"
	row := m.db.QueryRow(statement, id)
	empl := &EmployeeModel{}

	if err := row.Scan(&empl.ID, &empl.Name, &empl.Position, &empl.IsRobot); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return empl, nil
		}
		return nil, err
	}

	return empl, nil
}

func (m *Model) SelectEmployees() ([]*EmployeeModel, error) {
	statement := "SELECT id, name, position, is_robot FROM personal"
	rows, err := m.db.Query(statement)
	if err != nil {
		return nil, err
	}

	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			panic(err)
		}
	}(rows)

	empls := make([]*EmployeeModel, 0)

	for rows.Next() {
		empl := &EmployeeModel{}
		if err := rows.Scan(&empl.ID, &empl.Name, &empl.Position, &empl.IsRobot); err != nil {
			return nil, err
		}
		empls = append(empls, empl)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return empls, nil
}

func (m *Model) InsertEmployee(empl *EmployeeModel) error {
	statement := "INSERT INTO personal (name, position, is_robot)  VALUES ($1, $2, $3)"
	_, err := m.db.Exec(statement, empl.Name, empl.Position, empl.IsRobot)
	if err != nil {
		return err
	}
	return nil
}
