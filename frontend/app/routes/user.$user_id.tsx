import { json, ActionFunctionArgs, LoaderFunctionArgs } from '@remix-run/node';
import { useLoaderData } from '@remix-run/react';
import { Profile } from "~/components/Profile";
import { MusicList } from "~/components/MusicList";
import { Follower, Music, UserInfo } from '~/types/types';
import { Followings } from '~/components/Followings';
import { Followers } from '~/components/Followers';
import { Header } from '~/components/Header';
import { useEffect, useState } from 'react';

export const loader = async ({ params }: LoaderFunctionArgs) => {
  const userId = params.user_id;
  const userResponse = await fetch(`http://backend:8080/api/v1/user/${userId}`);
  if (!userResponse.ok) {
    throw new Error (`Failed to fetch user data: ${userResponse.statusText}`)
  }
  const userData = await userResponse.json();

  const musicResponse = await fetch(`http://backend:8080/api/v1/music/${userId}`);
  if (!musicResponse.ok) {
    throw new Error (`Failed to fetch music data: ${musicResponse.statusText}`)
  }
  const musicData = await musicResponse.json();

  const followerResponse = await fetch(`http://backend:8080/api/v1/follower/${userId}`);
  if (!followerResponse.ok) {
    throw new Error (`Failed to fetch follower data: ${followerResponse.statusText}`)
  }
  const followerData = await followerResponse.json();

  const followingsResponse = await fetch(`http://backend:8080/api/v1/followee/${userId}`);
  if (!followingsResponse.ok) {
    throw new Error (`Failed to fetch followings data: ${followingsResponse.statusText}`)
  }
  const followingsData = await followingsResponse.json();

  const userInfo: UserInfo = {
    displayName: userData.user.display_name,
    iconImage: userData.user.icon_image,
    xLink: userData.user.x_link,
    instagramLink: userData.user.instagram_link,
  };

  const musicList = musicData.musicList.musics.map((music: Music) => {
    return {
      src: music.image,
      title: music.name,
      category: music.artist_name,
      content: music.spotify_link,
    };
  });

  const followers = followerData.followers.map((follower: Follower) => {
    return {
      user_id: follower.user_id,
      icon_image: follower.icon_image,
    };
  });

  const followings = followingsData.followees.map((following: Follower) => {
    return {
      user_id: following.user_id,
      icon_image: following.icon_image,
    };
  });

  return json({ userInfo, musicList, followers, followings });
};

export const action = async ({
  request,
  params,
}: ActionFunctionArgs) => {
  const userId = params.user_id;
  const response = await fetch(`http://backend:8080/api/v1/music/${userId}`, {
    method: 'POST',
    headers: {
      'Cookie': request.headers.get('Cookie') || '',
    },
  });

  if (!response.ok) {
    throw new Error('Failed to update music list');
  }

  return json({ success: true });
};

export default function User() {
  const { userInfo, musicList, followers, followings } = useLoaderData<typeof loader>();

  const [myUserId, setMyUserId] = useState<number | null>(null);

  useEffect(() => {
    const getMyUserId = async () => {
      try {
        const response = await fetch(`http://localhost:8080/api/v1/myuserid`, { credentials: "include" });
        if (!response.ok) {
          throw new Error(`HTTP error! status: ${response.status}`);
        }
        const result = await response.json();
        setMyUserId(result.user_id);
      } catch (error) {
        console.error("There was a problem with the fetch operation:", error);
      }
    }
    getMyUserId();
  }, []);

  useEffect(() => {
    if (myUserId !== null) {
      console.log('myUserId:', myUserId);
    }

  }, [myUserId]);

  if (myUserId === null) {
    return <div>Loading...</div>;
  }

  return (
    <div>
      <Header
        myUserId={myUserId}
      />
      <div className="font-sans p-4 pt-20 flex flex-col gap-8">
        <Profile
          displayName={userInfo.displayName}
          iconImage={userInfo.iconImage}
          xLink={userInfo.xLink}
          instagramLink={userInfo.instagramLink}
          myUserId={myUserId}
        />
        <MusicList musicList={musicList}/>
        <Followings followings={followings}/>
        <Followers followers={followers}/>
      </div>
    </div>
  );
}

