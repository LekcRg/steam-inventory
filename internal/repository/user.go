package repository

func (r *Repo) GetUsersIds() ([]int, error) {
	var users []int
	err := r.db.Select(&users, "SELECT id FROM users")
	if err != nil {
		return nil, err
	}
	return users, nil
}
