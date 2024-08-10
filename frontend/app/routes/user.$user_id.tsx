import { Profile } from "~/components/Profile";
import { Friends } from "~/components/Friends";
import { MusicList } from "~/components/MusicList";
import { json, LoaderFunction } from '@remix-run/node';
import { useLoaderData } from '@remix-run/react';
import type { LoaderFunctionArgs } from "@remix-run/node";


export const loader: LoaderFunction = async ({ params }: LoaderFunctionArgs) => {
  const user_id = params.userId;
  const response = await fetch(`http://localhost:8080/api/v1/user/${user_id}`);
  const data = await response.json();

  const displayName = data.display_name;
  const iconImage = data.icon_image;
  const xLink = data.x_link;
  const instagramLink = data.instagram_link;

  return json({ displayName, iconImage, xLink, instagramLink });
};

export default function UserPage() {
  const { displayName, iconImage, xLink, instagramLink } = useLoaderData<typeof loader>();

  return (
    <div className="font-sans p-4 pt-20 flex flex-col gap-8">
      <Profile displayName={displayName} iconImage={iconImage} xLink={xLink} instagramLink={instagramLink}/>
      <MusicList />
      <Friends />
    </div>
  );
}
