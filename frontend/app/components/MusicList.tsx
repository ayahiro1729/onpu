import React from "react";
import { Carousel, Card } from "~/components/ui/apple-cards-carousel";
import { Loader2 } from "lucide-react"
import { Button } from "~/components/ui/button"
import { Form, useNavigation, useParams } from '@remix-run/react';
import { MusicCard, MusicListProps} from "~/types/types";

export const MusicList: React.FC<MusicListProps> = ({ myUserId, musicList }) => {
  const navigation = useNavigation();
  const isUpdating = navigation.state === "submitting";

  const musics = musicList.map((card: MusicCard, index: number) => (
    <a href={card.content} target="_blank">
      <Card key={card.src} card={card} index={index} />
    </a>
  ));

  const params = useParams();
  const pageUserId = params.user_id ? Number(params.user_id) : undefined;

  return (
    <div className="flex flex-col gap-4">
      <div className="flex justify-between items-center">
        <p className='flex items-center text-2xl'>Favorite Music</p>
        {myUserId == pageUserId &&
          <Form method="post">
            <input type="hidden" name="user_id" value={myUserId} />
            <Button
              type="submit"
              name="_action"
              aria-label="post"
              value="post"
              className="px-2 py-1 bg-[#1ED760]"
              disabled={isUpdating}
            >
              {isUpdating && <Loader2 className="mr-2 h-4 w-4 animate-spin" />}
                <p>Update</p>
            </Button>
          </Form>
        }
      </div>
      <Carousel items={musics} />
    </div>
  );
}
