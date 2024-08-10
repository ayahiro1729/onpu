import React from "react";
import { Carousel, Card } from "~/components/ui/apple-cards-carousel";
import { Loader2 } from "lucide-react"
 import { Button } from "~/components/ui/button"
import { ActionFunctionArgs, json, LoaderFunctionArgs } from '@remix-run/node';
import { Form, useLoaderData, useParams } from '@remix-run/react';
import { Music, MusicCard } from "~/types/types";

export const action = async ({
  params,
}: ActionFunctionArgs) => {
  const userId = params.user_id;
  const response = await fetch(`https://localhost:8080/api/v1/music/${userId}`, {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
  });

  if (!response.ok) {
    throw new Error('Failed to update music list');
  }

  return json({ success: true });
};

type MusicListProps = {
  musicList: MusicCard[];
};

export const MusicList: React.FC<MusicListProps> = ({ musicList }) => {
  const params = useParams();
  const userId = params.user_id;

  const musics = musicList.map((card: MusicCard, index: number) => (
    <a href={card.content} target="_blank">
      <Card key={card.src} card={card} index={index} />
    </a>
  ));

  return (
    <div className="flex flex-col gap-4">
      <div className="flex justify-between items-center">
        <p className='flex items-center text-2xl'>Favorite Music</p>
        <Form key={userId} id="contact-form" method="post">
          <Button className="px-2 py-1 bg-[#1ED760]">
            <Loader2 className="mr-2 h-4 w-4 animate-spin" />
              <p>Update</p>
          </Button>
        </Form>
      </div>
      <Carousel items={musics} />
    </div>
  );
}
