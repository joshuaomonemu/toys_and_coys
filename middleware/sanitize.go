//package middleware
//
//import "app/models"
//
//func (user *models.Users) HashPassword(password string) error {
//	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
//	if err != nil {
//		return err
//	}
//	user.Password = string(bytes)
//	return nil
//}
//func (user *User) CheckPassword(providedPassword string) error {
//	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(providedPassword))
//	if err != nil {
//		return err
//	}
//	return nil
//}
