import { json, LoaderFunctionArgs, ActionFunctionArgs, redirect } from '@remix-run/node';
import { useLoaderData } from '@remix-run/react';
import { Profile } from "~/components/Profile";
import { MusicList } from "~/components/MusicList";
import { Follower, Music, UserInfo } from '~/types/types';
import { Followings } from '~/components/Followings';
import { Followers } from '~/components/Followers';
import { Header } from '~/components/Header';

export const action = async ({ request, params }: ActionFunctionArgs) => {
  const formData = await request.formData();
  const userId = params.user_id;
  console.log('Sending request to backend with user_id:', userId);
  const response = await fetch(`http://backend:8080/api/v1/music`, {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    body: JSON.stringify({
      user_id: userId,
    }),
    credentials: 'include',
  });

  console.log('Backend response status:', response.status);
  if (!response.ok) {
    const errorText = await response.text();
    console.error('Backend error:', errorText);
    throw new Error('Failed to update music list');
  }

  return redirect("./");
};
import { Button } from '~/components/ui/button';
import { useState } from 'react';

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

  // const myUserIdResponse = await fetch(`http://backend:8080/api/v1/myuserid`)
  // if (!myUserIdResponse.ok) {
  //   throw new Error (`Failed to fetch my user id: ${myUserIdResponse.statusText}`)
  // }
  // const myUserId = await myUserIdResponse.json();
  // console.log(myUserId)

  const userInfo: UserInfo = {
    displayName: userData.user.DisplayName,
    iconImage: userData.user.IconImage,
    xLink: userData.user.XLink,
    instagramLink: userData.user.InstagramLink,
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
      userId: follower.userId,
      displayName: follower.displayName,
      iconImage: follower.iconImage,
    };
  });

  const followings = followingsData.followees.map((following: Follower) => {
    return {
      userId: following.userId,
      displayName: following.displayName,
      iconImage: following.iconImage,
    };
  });

  return json({ userInfo, musicList, followers, followings });
};

export default function User() {
  const { userInfo, musicList, followers, followings } = useLoaderData<typeof loader>();
  const [myuserid, setMyuserid] = useState("")
  
  const getMyUserId = async () => {
    try {
      const response = await fetch(`http://localhost:8080/api/v1/myuserid`, { credentials: "include" });
      if (!response.ok) {
        throw new Error(`HTTP error! status: ${response.status}`);
      }
      const result = await response.json();
      console.log(result);
      setMyuserid(result);
    } catch (error) {
      console.error("There was a problem with the fetch operation:", error);
    }
  };

  return (
    <div>
      <Header />
      <div className="font-sans p-4 pt-20 flex flex-col gap-8">
      <Profile 
        displayName={userInfo.displayName}
        iconImage={userInfo.iconImage}
        xLink={userInfo.xLink}
        instagramLink={userInfo.instagramLink}
      />
      <MusicList musicList={musicList}/>
      <Followings followings={followings}/>
      <Followers followers={followers}/>
      <Button onClick={getMyUserId}>get my user id</Button>
    </div>
    </div>
  );
}


