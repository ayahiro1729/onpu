import { json, LoaderFunctionArgs } from '@remix-run/node';
import { useLoaderData } from '@remix-run/react';
import { Profile } from "~/components/Profile";
import { MusicList } from "~/components/MusicList";
import { Music, UserInfo } from '~/types/types';

export const loader = async ({ params }: LoaderFunctionArgs) => {
  const userId = params.user_id;
  const userResponse = await fetch(`http://backend:8080/api/v1/user/${userId}`);
  const userData = await userResponse.json();

  const musicResponse = await fetch(`http://backend:8080/api/v1/music/${userId}`);
  const musicData = await musicResponse.json();

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

  return json({ userInfo, musicList });
};

export default function User() {
  const { userInfo, musicList } = useLoaderData<typeof loader>();

  return (
    <div className="font-sans p-4 pt-20 flex flex-col gap-8">
      <Profile 
        displayName={userInfo.displayName}
        iconImage={userInfo.iconImage}
        xLink={userInfo.xLink}
        instagramLink={userInfo.instagramLink}
      />
      <MusicList musicList={musicList}/>
      {/* <Followings />
      <Followers /> */}
    </div>
  );
}