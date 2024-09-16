import { Link } from "@remix-run/react";
import React from "react";
import { FollowingsProps } from "~/types/types";

export const Followings: React.FC<FollowingsProps> = ({followings}) => {
  return (
    <div className="flex flex-col gap-4">
      <p className="text-2xl">Followings</p>
      { followings.length === 0 && <div>No followings found.</div> }
      { followings.map((following, index) => (
        <Link to={`/user/${following.user_id}`} key={index} className="flex items-center gap-2">
          <img src={following.icon_image} className="w-10 h-10 rounded-full" />
        </Link>
      ))}
    </div>
  );
}
