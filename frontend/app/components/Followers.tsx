import { Link } from "@remix-run/react";
import React from "react";
import { FollowerProps } from "~/types/types";

export const Followers: React.FC<FollowerProps> = ({followers}) => {
  return (
    <div className="flex flex-col gap-4">
      <p className="text-2xl">Followers</p>
      { followers.length === 0 && <div>No followers found.</div> }
      <div className="flex gap-2">
      { followers.map((follower, index) => (
        <Link to={`/user/${follower.user_id}`} key={index} className="flex items-center gap-2">
          <img src={follower.icon_image} className="w-10 h-10 rounded-full" />
        </Link>
      ))}
      </div>
    </div>
  );
}
