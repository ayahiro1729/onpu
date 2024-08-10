import { json, LoaderFunctionArgs } from "@remix-run/node";
import { useLoaderData } from "@remix-run/react";
import React from "react";
import { AnimatedTooltip } from "~/components/ui/animated-tooltip";
import { Follow, FollowData } from "~/types/types";

export const loader = async ({ params }: LoaderFunctionArgs) => {
  const response = await fetch(`https://localhost:8080/api/v1/followee/${params.userId}`);
  const data: FollowData[] = await response.json();

  const followings: Follow[] = data.map((item: FollowData) => {
    return {
      id: item.user_id,
      name: item.display_name,
      designation: "",
      image: item.icon_image,
    };
  });

  return json({ followings });
};

export const Followings = () => {
  const { followings } = useLoaderData<typeof loader>();

  if (followings.length === 0) {
    return <div>No followings found.</div>;
  }
  
  return (
    <div className="flex flex-col gap-4">
      <p className="text-2xl">Followings</p>
      <div className="flex flex-row items-center w-full">
        <AnimatedTooltip items={followings} />
      </div>
    </div>
  );
}
