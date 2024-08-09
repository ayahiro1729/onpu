import { Profile } from "~/components/Profile";
import { Friends } from "~/components/Friends";
import { MusicList } from "~/components/MusicList";

export default function UserPage() {
  return (
    <div className="font-sans p-4 pt-20 flex flex-col gap-8">
      <Profile />
      <MusicList />
      <Friends />
    </div>
  );
}