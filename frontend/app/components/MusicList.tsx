import React from "react";
import { Carousel, Card } from "~/components/ui/apple-cards-carousel";
import { Loader2 } from "lucide-react"
import { Button } from "~/components/ui/button"
import { Form, useNavigation } from '@remix-run/react';
import { Music, MusicCard } from "~/types/types";

type MusicListProps = {
  musicList: MusicCard[];
};

export const MusicList: React.FC<MusicListProps> = ({ musicList }) => {
  const navigation = useNavigation();
  const isUpdating = navigation.state === "submitting";
  // console.log(musicList);

  const musics = musicList.map((card: MusicCard, index: number) => (
    <a href={card.content} target="_blank">
      <Card key={card.src} card={card} index={index} />
    </a>
  ));

  return (
    <div className="flex flex-col gap-4">
      <div className="flex justify-between items-center">
        <p className='flex items-center text-2xl'>Favorite Music</p>
        <Form id="contact-form" method="post">
          <Button
            type="submit"
            className="px-2 py-1 bg-[#1ED760]"
            disabled={isUpdating}
          >
            {isUpdating && <Loader2 className="mr-2 h-4 w-4 animate-spin" />}
              <p>Update</p>
          </Button>
        </Form>
      </div>
      <Carousel items={musics} />
    </div>
  );
}
