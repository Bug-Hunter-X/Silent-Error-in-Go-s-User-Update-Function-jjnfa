func UpdateUser(user *User) error { 
  if user.ID == 0 { 
    return errors.New("user ID cannot be zero") 
  } 

  // ... other code to update the user ... 
  return nil 
}