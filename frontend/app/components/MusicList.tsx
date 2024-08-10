import React from "react";
import { Carousel, Card } from "~/components/ui/apple-cards-carousel";
import { Loader2 } from "lucide-react"
 
import { Button } from "~/components/ui/button"
import { ActionFunctionArgs, json, LoaderFunctionArgs } from '@remix-run/node';
import { Form, useLoaderData } from '@remix-run/react';

type Music = {
  music_id: number,
  name: string,
  image: string,
  artist_name: string,
  spotify_link: string
};

type MusicCard = {
  src: string;
  title: string;
  category: string;
  content: string;
};

export const action = async ({
  params,
}: ActionFunctionArgs) => {
  const response = await fetch(`https://localhost:8080/api/v1/music/${params.userId}`, {
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


export const loader = async ({ params }: LoaderFunctionArgs) => {
  const user_id = params.userId;
  const response = await fetch(`https://localhost:8080/api/v1/user/${user_id}`);
  const data = await response.json();

  const musicList = data.music_list;
  const musicCardList = musicList.map((music: Music) => {
    return {
      src: music.image,
      title: music.name,
      category: music.artist_name,
      content: music.spotify_link
    };
  });

  return json({ user_id, musicCardList });
};

export function MusicList() {
  const { user_id, musicCardList } = useLoaderData<typeof loader>();
  const musics = musicCardList.map((card: MusicCard, index: number) => (
    <a href={card.content}>
      <Card key={card.src} card={card} index={index} />
    </a>
  ));

  return (
    <div className="flex flex-col gap-4">
      <div className="flex justify-between items-center">
        <p className='flex items-center text-2xl'>Favorite Music</p>
        <Form key={user_id} id="contact-form" method="post">
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
