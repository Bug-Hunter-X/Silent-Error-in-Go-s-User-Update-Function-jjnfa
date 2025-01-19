func UpdateUser(user *User) error {
  if user.ID == 0 {
    return errors.New("user ID cannot be zero")
  }

  // Simulate database interaction which can return error
  err := updateUserData(user)
  if err != nil {
    return fmt.Errorf("failed to update user: %w", err)
  }
  return nil
}

func updateUserData(user *User) error {
  //Simulate a database error
  if user.ID % 2 == 0 {
    return fmt.Errorf("database error")
  }
  return nil
}
