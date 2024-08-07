import type { MetaFunction } from "@remix-run/node";
import { Profile } from "~/components/Profile";

export const meta: MetaFunction = () => {
  return [
    { title: "New Remix App" },
    { name: "description", content: "Welcome to Remix!" },
  ];
};

export default function Index() {
  return (
    <div className="font-sans p-4">
      <Profile />
    </div>
  );
}