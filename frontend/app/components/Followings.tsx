import React from "react";
import { AnimatedTooltip } from "~/components/ui/animated-tooltip";
import { FollowingsProps } from "~/types/types";

export const Followings: React.FC<FollowingsProps> = ({followings}) => {
  return (
    <div className="flex flex-col gap-4">
      <p className="text-2xl">Followings</p>
      { followings.length === 0 && <div>No followings found.</div> }
      <div className="flex flex-row items-center w-full">
        <AnimatedTooltip items={followings} />
      </div>
    </div>
  );
}
