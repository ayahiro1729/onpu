import { json, LoaderFunctionArgs } from "@remix-run/node";
import { useLoaderData } from "@remix-run/react";
import React from "react";
import { AnimatedTooltip } from "~/components/ui/animated-tooltip";
import { Follow, FollowData } from "~/types/types";

export const loader = async ({ params }: LoaderFunctionArgs) => {
  const response = await fetch(`https://localhost:8080/api/v1/follower/${params.userId}`);
  const data: FollowData[] = await response.json();

  const followers: Follow[] = data.map((item: FollowData) => {
    return {
      id: item.user_id,
      name: item.display_name,
      designation: "",
      image: item.icon_image,
    };
  });

  return json({ followers });
};

export const Followers = () => {
  const { followers } = useLoaderData<typeof loader>();
  
  if (followers.length === 0) {
    return <div>No followers found.</div>;
  }

  return (
    <div className="flex flex-col gap-4">
      <p className="text-2xl">Followers</p>
      <div className="flex flex-row items-center w-full">
        <AnimatedTooltip items={followers} />
      </div>
    </div>
  );
}
