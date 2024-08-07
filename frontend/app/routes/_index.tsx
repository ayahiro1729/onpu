import type { MetaFunction } from "@remix-run/node";
import { Profile } from "~/components/Profile";
import { Friends } from "~/components/Friends";
import { MusicList } from "~/components/MusicList";

export const meta: MetaFunction = () => {
  return [
    { title: "New Remix App" },
    { name: "description", content: "Welcome to Remix!" },
  ];
};

export default function Index() {
  return (
    <div className="font-sans p-4 pt-20 flex flex-col gap-8">
      <Profile />
      <MusicList />
      <Friends />
    </div>
  );
}