package repository

type FollowUserDTO struct {
	UserID		int		`json:"user_id"`
	UserName	string	`json:"user_name"`
	DisplayName	string	`json:"display_name"`
	IconImage	string	`json:"icon_image"`
	UpdatedAt	string	`json:"created_at"`
}

type FollowUpdateDTO struct {
	FollowerID	int		`json:"follower_id"`
	FolloweeID	int		`json:"followee_id"`
}

type FollowersRepository interface {
	GetFollowers(userID int) (*[]FollowUserDTO, error)
}

type FolloweesRepository interface {
	GetFollowees(userID int) (*[]FollowUserDTO, error)
}

type FollowRepository interface {
	FollowUser(followerID int, followeeID int) error
}
