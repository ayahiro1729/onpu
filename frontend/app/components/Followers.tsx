import React from "react";
import { AnimatedTooltip } from "~/components/ui/animated-tooltip";
import { FollowerProps } from "~/types/types";

export const Followers: React.FC<FollowerProps> = ({followers}) => {
  return (
    <div className="flex flex-col gap-4">
      <p className="text-2xl">Followers</p>
      { followers.length === 0 && <div>No followers found.</div> }
      <div className="flex flex-row items-center w-full">
        <AnimatedTooltip items={followers} />
      </div>
    </div>
  );
}
