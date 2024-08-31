export type Follow ={
  id: number;
  name: string;
  designation: string;
  image: string;
}

export type FollowData = {
  user_id: number,
  user_name: string,
  display_name: string,
  icon_image: string,
  created_at: string
}

export type Music = {
  music_id: number,
  name: string,
  image: string,
  artist_name: string,
  spotify_link: string
};

export type MusicCard = {
  src: string;
  title: string;
  category: string;
  content: string;
};

export type UserInfo = {
  displayName: string;
  iconImage: string;
  xLink: string;
  instagramLink: string;
};

export type ProfileProps = {
  displayName?: string
  iconImage?: string
  xLink?: string
  instagramLink?: string
  myUserId?: number | null
}

export type Follower = {
  userId: number;
  displayName: string;
  iconImage: string;
}

export type FollowerProps = {
  followers: Follower[]
}

export type FollowingsProps = {
  followings: Follower[]
}

export type HeaderProps = {
  iconImage?: string
  myUserId?: number | null
}