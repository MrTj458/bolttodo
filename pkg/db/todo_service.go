package db

import (
	"encoding/json"

	"github.com/mrtj458/bolttodo/pkg/data"
	"go.etcd.io/bbolt"
)

type TodoService struct {
	DB *bbolt.DB
}

func (ts *TodoService) GetAll() ([]data.Todo, error) {
	todos := make([]data.Todo, 0)

	err := ts.DB.View(func(tx *bbolt.Tx) error {
		b := tx.Bucket(todosBucket)

		err := b.ForEach(func(k, v []byte) error {
			var t data.Todo
			err := json.Unmarshal(v, &t)
			if err != nil {
				return err
			}
			todos = append(todos, t)

			return nil
		})

		return err
	})
	if err != nil {
		return nil, err
	}

	return todos, nil
}

func (ts *TodoService) GetByID(id int) (*data.Todo, error) {
	var t *data.Todo

	err := ts.DB.View(func(tx *bbolt.Tx) error {
		b := tx.Bucket(todosBucket)

		jsn := b.Get(itob(id))
		if len(jsn) == 0 {
			return data.ErrNotFound
		}

		return json.Unmarshal(jsn, &t)
	})
	if err != nil {
		return nil, err
	}

	return t, err
}

func (ts *TodoService) Insert(t *data.Todo) error {
	return ts.DB.Update(func(tx *bbolt.Tx) error {
		b := tx.Bucket(todosBucket)

		id, err := b.NextSequence()
		if err != nil {
			return err
		}
		t.ID = int(id)

		jsn, err := json.Marshal(t)
		if err != nil {
			return err
		}

		return b.Put(itob(int(id)), jsn)
	})
}

func (ts *TodoService) Update(id int, t *data.Todo) error {
	t.ID = id // make sure todo's ID matches the key

	return ts.DB.Update(func(tx *bbolt.Tx) error {
		b := tx.Bucket(todosBucket)

		jsn := b.Get(itob(id))
		if len(jsn) == 0 {
			return data.ErrNotFound
		}

		jsn, err := json.Marshal(t)
		if err != nil {
			return err
		}

		return b.Put(itob(id), jsn)
	})
}

func (ts *TodoService) Delete(id int) error {
	return ts.DB.Update(func(tx *bbolt.Tx) error {
		b := tx.Bucket(todosBucket)

		jsn := b.Get(itob(id))
		if len(jsn) == 0 {
			return data.ErrNotFound
		}

		return b.Delete(itob(id))
	})
}
