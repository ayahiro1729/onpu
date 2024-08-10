import React, { useEffect, useState } from 'react';
import { Profile } from "~/components/Profile";
import { Followings } from "~/components/Followings";
import { Followers } from "~/components/Followers";
import { MusicList } from "~/components/MusicList";
import { useParams } from '@remix-run/react';
import axios from 'axios';

type UserData = {
  displayName: string;
  iconImage: string;
  xLink: string;
  instagramLink: string;
};

export default function User() {
  const { user_id } = useParams();
  const [userData, setUserData] = useState<UserData | null>(null);
  const [error, setError] = useState<string | null>(null);
  const [isLoading, setIsLoading] = useState(true);

  useEffect(() => {
    const fetchUserData = async () => {
      try {
        const response = await axios.get(`http://localhost:8080/api/v1/user/${user_id}`);
        const user = response.data.user;
        console.log(user);
        setUserData({
          displayName: user.DisplayName,
          iconImage: user.IconImage,
          xLink: user.XLink,
          instagramLink: user.InstagramLink,
        });
      } catch (error) {
        console.error("Error fetching user data:", error);
        if (axios.isAxiosError(error)) {
          setError(`Error fetching data: ${error.message}`);
        } else {
          setError("An unknown error occurred");
        }
      } finally {
        setIsLoading(false);
      }
    };

    fetchUserData();
  }, [user_id]);
  
  return (
    <div className="font-sans p-4 pt-20 flex flex-col gap-8">
      <Profile 
        displayName={userData?.displayName}
        iconImage={userData?.iconImage}
        xLink={userData?.xLink}
        instagramLink={userData?.instagramLink}
      />
      {/* <MusicList />
      <Followings />
      <Followers /> */}
    </div>
  );
}