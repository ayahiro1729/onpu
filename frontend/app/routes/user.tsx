import { Profile } from "~/components/Profile";
import { Friends } from "~/components/Friends";
import { MusicList } from "~/components/MusicList";
import { json, LoaderFunction } from '@remix-run/node';
import { useLoaderData } from '@remix-run/react';
import { useParams } from "@remix-run/react";


// export const loader: LoaderFunction = async () => {
//   const { user_id } = useParams();
//   const response = await fetch(`https://localhost:8080/api/v1/user/${user_id}`);
//   const data = await response.json();

//   // ビットコインの価格を取得
//   const displayName = data.display_name;
//   const iconImage = data.icon_image;
//   const xLink = data.x_link;
//   const instagramLink = data.instagram_link;

//   return json({ displayName, iconImage, xLink, instagramLink });
// };

export default function UserPage() {
  // const { displayName, iconImage, xLink, instagramLink } = useLoaderData<typeof loader>();

  const displayName = 'displayName';
  const iconImage = '/OnpuLogo.jpg';
  const xLink = 'https://x.com';
  const instagramLink = 'https://instagram.com';

  return (
    <div className="font-sans p-4 pt-20 flex flex-col gap-8">
      <Profile displayName={displayName} iconImage={iconImage} xLink={xLink} instagramLink={instagramLink}/>
      <MusicList />
      <Friends />
    </div>
  );
}