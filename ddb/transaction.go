package ddb

type Tx struct {
	dp *Dbmp
}

func (t *Tx) Create(modal interface{})  {
	if m := t.dp.Create(modal); m.Error != nil {
		t.dp = m //事务需要覆盖tx.dp
	}
}

func (t *Tx) Update(modal interface{})  {
	if m := t.dp.Update(modal); m.Error != nil {
		t.dp = m
	}
}

func (t *Tx) Delete(modal interface{})  {
	if m := t.dp.Delete(modal); m.Error != nil {
		t.dp = m
	}
}

func Transaction(fc func(t *Tx)) error {
	tm := &Tx{
		dp: 	NewDbmp(),
	}
	return tm.dp.transaction(func(m *Dbmp) error {
		fc(tm)
		return tm.dp.Error //必须返回fc的error
	})
}